package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/relvacode/iso8601"
	"github.com/rs/cors"
)


type account struct {
	Email string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone string `json:"phone"`
}
type alarm struct {
	Time string `json:"time"`
	Week struct {
		Sunday bool `json:"sunday"`
		Monday bool `json:"monday"`
		Tuesday bool `json:"tuesday"`
		Wednesday bool `json:"wednesday"`
		Thursday bool `json:"thursday"`
		Friday bool `json:"friday"`
		Saturday bool `json:"saturday"`
	} `json:"days"`
}

type App struct {
	router *mux.Router
	DB *sql.DB
}

func hashPassword(password string) []byte {
	h := sha256.New()
	h.Write([]byte(password))
	return h.Sum(nil)
}

func (a *App) createUser(writer http.ResponseWriter, request *http.Request) {
	var account account
	decoder := json.NewDecoder(request.Body);
    errDecode := decoder.Decode(&account);
	id := uuid.New()
	if (errDecode != nil) {
		fmt.Println(errDecode)
        respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
        return
    }
	hashedPassword := hashPassword(account.Password)
	_, err := a.DB.Exec("INSERT INTO users(id, email, username, password, phone) VALUES($1, $2, $3, $4, $5) RETURNING id",
							id, account.Email, account.Username, hashedPassword, account.Phone)

	if err != nil {
		fmt.Println(err)
	}

	defer request.Body.Close()
}

func respondWithError(w http.ResponseWriter, code int, message string) {
    respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

func (a *App) initializeApp() {
	db, err := sql.Open("sqlite3","telesnooze.db");
	if err != nil {
		panic("failed to connect database")
	  }
	a.DB = db
	a.router = mux.NewRouter();
}

func sayHello(writer http.ResponseWriter, request *http.Request){
	fmt.Println("hello new user")
	writer.Header().Set("hello", "there")
}

//this a *App  means it applies to the app struct type
func (a *App) setAlarm(writer http.ResponseWriter, request *http.Request){

	//TODO 
	//check that there is at least one true value for days of the week
	var alarm alarm
	decoder := json.NewDecoder(request.Body);
	
	errDecode := decoder.Decode(&alarm);
	fmt.Printf("%v: %v\n", alarm.Time, alarm.Week)
	if (errDecode != nil) {
		fmt.Println(errDecode)
        respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
        return
    }
	id := uuid.New();
	_,tmErr := iso8601.ParseString(alarm.Time);
	v := reflect.ValueOf(alarm.Week)
	hasDaysOfWeek := false;

	for i := 0; i< v.NumField(); i++ {
		if(v.Field(i).Interface() == true){
			hasDaysOfWeek = true
		}
    }
	
	if(tmErr != nil){
   		writer.Write([]byte("Timestamp is not in ISO format"))
	} else if(!hasDaysOfWeek){
   		writer.Write([]byte("Problem: Week needs at least one true value OR JSON be malformed"))

	}else  {
		_, err := a.DB.Exec(
			`INSERT INTO alarms(id, time, sunday, monday, tuesday, wednesday, thursday,friday,saturday) 
			 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
			 id, alarm.Time, alarm.Week.Sunday, alarm.Week.Monday, alarm.Week.Tuesday, alarm.Week.Wednesday, alarm.Week.Thursday, alarm.Week.Friday, alarm.Week.Saturday); 
		if(err != nil){
			fmt.Println("failure: ", err);
			writer.Write([]byte("Something went wrong in DB process"))
		} else {
			writer.Write([]byte("Success"))
		}
	}

	defer request.Body.Close()
}

func main(){
	app := &App{}
	app.initializeApp()
	app.router.HandleFunc("/api/v1/", sayHello).Methods("GET");
	app.router.HandleFunc("/api/v1/setAlarm", app.setAlarm).Methods("POST")
	app.router.HandleFunc("/api/v1/createUser", app.createUser).Methods("POST")
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowCredentials: true,
    })
	handler := c.Handler(app.router)
	fmt.Println("Server at 8123")
    log.Fatal(http.ListenAndServe(":8123",handler))
}
