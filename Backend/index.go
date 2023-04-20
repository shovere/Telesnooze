package main

import (
	"bytes"
	"crypto/sha256"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"strconv"

	"time"

	"github.com/robfig/cron/v3"

	//"github.com/elgs/cron"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/relvacode/iso8601"
	"github.com/rs/cors"
)

type account struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type cronHolder struct {
	alarm_id string
	cron *cron.Cron
}

type alarm struct {
	User_ID string `json:"user_id"`
	Time string `json:"time"`
	Alarm_ID string 
	Week struct {
		Sunday    bool `json:"sunday"`
		Monday    bool `json:"monday"`
		Tuesday   bool `json:"tuesday"`
		Wednesday bool `json:"wednesday"`
		Thursday  bool `json:"thursday"`
		Friday    bool `json:"friday"`
		Saturday  bool `json:"saturday"`
	} `json:"days"`
}
type retAlarms struct {
	User_ID string `json:"user_id"`
	Alarms []alarm `json:"alarms"`
}

type App struct {
	router *mux.Router
	DB     *sql.DB
	cronDaemon []cronHolder
}

func hashPassword(password string) []byte {
	h := sha256.New()
	h.Write([]byte(password))
	return h.Sum(nil)
}

func (a *App) createUser(writer http.ResponseWriter, request *http.Request) {
	var account account
	decoder := json.NewDecoder(request.Body)
	errDecode := decoder.Decode(&account)
	id := uuid.New()
	if errDecode != nil {
		fmt.Println(errDecode)
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}

	allFieldsFilled := true
	containsNonAscii := false
	validPhoneNumber := true
	v := reflect.ValueOf(account)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			allFieldsFilled = false
		}
	}
	for _, char := range account.Password {
		if char > 127 {
			containsNonAscii = true
		}
	}
	if len(account.Phone) != 10 {
		validPhoneNumber = false
	}
	for _, char := range account.Phone {
		if char < 48 || char > 57 {
			validPhoneNumber = false
		}
	}

	if !allFieldsFilled {
		writer.Write([]byte("Problem: All fields must be filled"))
	} else if containsNonAscii {
		writer.Write([]byte("Problem: Password must only contain ASCII characters"))
	} else if !validPhoneNumber {
		writer.Write([]byte("Problem: Phone number is invalid - must be 10 digits and only contain numbers"))
	} else {
		hashedPassword := hashPassword(account.Password)
		_, err := a.DB.Exec("INSERT INTO users(id, email, username, password, phone) VALUES($1, $2, $3, $4, $5) RETURNING id",
			id, account.Email, account.Username, hashedPassword, account.Phone)

		if err != nil {
			fmt.Println(err)
		} else {
			writer.Write([]byte("Success"))
		}
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
	db, err := sql.Open("sqlite3", "telesnooze.db")
	if err != nil {
		fmt.Println("failed to connect database")
	}
	a.DB = db
	a.router = mux.NewRouter()
}

func (a *App) initCron(){
	rows, err := a.DB.Query(`SELECT *
				FROM alarms`)
	if err != nil {
		fmt.Println("failure: ", err)
	} else {
		var cronAlarms []alarm;
		defer rows.Close()
		for rows.Next() {
			var al alarm
			err := rows.Scan(&al.Alarm_ID, 
					&al.Time, 
					&al.Week.Sunday, 
					&al.Week.Monday, 
					&al.Week.Tuesday,
					&al.Week.Wednesday, 
					&al.Week.Thursday, 
					&al.Week.Friday, 
					&al.Week.Saturday, 
					&al.User_ID)
			if err != nil {
				fmt.Println("error in init", err)
			}
			cronAlarms = append(cronAlarms, al)
		}
		for i := 0; i < len(cronAlarms); i++ {
			a.createCronFromAlarm(cronAlarms[i])
		}
	}

}

func (a *App) createCronFromAlarm(al alarm){
	weekString := "";
	v := reflect.ValueOf(al.Week)

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == true {
			if(weekString == ""){
				weekString += strconv.Itoa(i)
			}else {
				 weekString +=  "," + strconv.Itoa(i) 
			}
		}
	}
	
	date, err := time.Parse(time.RFC3339, al.Time)
	
  
	if err != nil {
		fmt.Println("create cron alarm", err)
		return
	} else {
		updateCron := false;
		newCron := cron.New()
		var foundCron cronHolder
		for i:=0; i < len(a.cronDaemon); i++ {
			if(a.cronDaemon[i].alarm_id == al.Alarm_ID) {
				foundCron= a.cronDaemon[i]
				updateCron = true;
				break
			}
		}
	
		
		
		var cronString = fmt.Sprintf(`%d %d ? * %s`, date.Minute(), date.Hour(), weekString)
		fmt.Println(cronString)
		cronID,err := newCron.AddFunc(cronString, func() {
			fmt.Println("this is where we would call the user", al.Alarm_ID)
			//normally we would just query the user's phone number and put it here
			//but because we don't have a full version of twilio setup we just call my phone
			callNumber("+16035689902")
		})
		if(err != nil){
			fmt.Println(err)
		}
		if(updateCron){
			foundCron.cron.Stop()
			foundCron.cron = newCron
			foundCron.cron.Start()
		}else {
			foundCron.alarm_id = al.Alarm_ID
			foundCron.cron = newCron
			foundCron.cron.Start()
			a.cronDaemon = append(a.cronDaemon, foundCron)
		}
		fmt.Println(cronID, al.Alarm_ID)		
	}
}





func sayHello(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("hello new user")
	writer.Header().Set("hello", "there")
}

// this a *App  means it applies to the app struct type
func (a *App) createAlarm(writer http.ResponseWriter, request *http.Request) {

	//TODO
	//check that there is at least one true value for days of the week
	var alarm alarm
	decoder := json.NewDecoder(request.Body)

	errDecode := decoder.Decode(&alarm)
	fmt.Printf("%v: %v\n", alarm.Time, alarm.Week)
	if errDecode != nil {
		fmt.Println(errDecode)
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	id := uuid.New()
	_, tmErr := iso8601.ParseString(alarm.Time)
	v := reflect.ValueOf(alarm.Week)
	hasDaysOfWeek := false

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == true {
			hasDaysOfWeek = true
		}
	}

	if tmErr != nil {
		writer.Write([]byte("Timestamp is not in ISO format"))
	} else if !hasDaysOfWeek {
		writer.Write([]byte("Problem: Week needs at least one true value OR JSON be malformed"))

	} else {
		_, err := a.DB.Exec(
			`INSERT INTO alarms(id, time, sunday, monday, tuesday, wednesday, thursday,friday,saturday, user_id) 
			 VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`,
			id, alarm.Time, alarm.Week.Sunday, alarm.Week.Monday, alarm.Week.Tuesday, alarm.Week.Wednesday, alarm.Week.Thursday, alarm.Week.Friday, alarm.Week.Saturday, alarm.User_ID)
		if err != nil {
			fmt.Println("failure: ", err)
			writer.Write([]byte("Something went wrong in DB process"))
		} else {
			writer.Write([]byte("Success"))
		}
		a.createCronFromAlarm(alarm)
	
	}

	defer request.Body.Close()
}

func (a *App) retrieveAlarms(writer http.ResponseWriter, request *http.Request){
	
	//TODO
	//check that there is at least one true value for days of the week
	var tmpRetAlarm retAlarms
	decoder := json.NewDecoder(request.Body)
	errDecode := decoder.Decode(&tmpRetAlarm)
	fmt.Printf("%v", tmpRetAlarm.User_ID)
	if errDecode != nil {
		fmt.Println(errDecode)
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	
	rows, err := a.DB.Query(
		`SELECT * 
		FROM alarms
		WHERE user_id = $1`,
		tmpRetAlarm.User_ID)

	if err != nil {
		fmt.Println("failure: ", err)
		writer.Write([]byte("Something went wrong in DB process"))
	} else {
		defer rows.Close()
		var rowAlarm []alarm;
		for rows.Next() {
			var al alarm
			err := rows.Scan(&al.Alarm_ID, 
					&al.Time, 
					&al.Week.Sunday, 
					&al.Week.Monday, 
					&al.Week.Tuesday,
					&al.Week.Wednesday, 
					&al.Week.Thursday, 
					&al.Week.Friday, 
					&al.Week.Saturday, 
					&al.User_ID)
			if err != nil {
				log.Fatal(err)
			}
			rowAlarm = append(rowAlarm, al)
		}
		tmpRetAlarm.Alarms = rowAlarm;
		buf := new(bytes.Buffer)
		newErr := json.NewEncoder(buf).Encode(tmpRetAlarm)
		if newErr != nil {
			log.Fatal(newErr)
		}
		writer.Write(buf.Bytes())
	}
	defer request.Body.Close()
}
func (a *App) updateAlarm(writer http.ResponseWriter, request *http.Request){
	var alarm alarm
	decoder := json.NewDecoder(request.Body)

	errDecode := decoder.Decode(&alarm)
	fmt.Printf("%v: %v\n", alarm.Time, alarm.Week)
	if errDecode != nil {
		fmt.Println(errDecode)
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	_, tmErr := iso8601.ParseString(alarm.Time)
	v := reflect.ValueOf(alarm.Week)
	hasDaysOfWeek := false

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == true {
			hasDaysOfWeek = true
		}
	}

	if tmErr != nil {
		writer.Write([]byte("Timestamp is not in ISO format"))
	} else if !hasDaysOfWeek {
		writer.Write([]byte("Problem: Week needs at least one true value OR JSON be malformed"))

	} else {
		res, err := a.DB.Exec(
			`UPDATE alarms
			 SET time = $2, 
				sunday = $3, 
				monday = $4, 
				tuesday = $5, 
				wednesday = $6, 
				thursday = $7, 
				friday = $8, 
				saturday = $9
			 WHERE id = $1 AND user_id = $10`,
			alarm.Alarm_ID, 
			alarm.Time, 
			alarm.Week.Sunday, 
			alarm.Week.Monday, 
			alarm.Week.Tuesday, 
			alarm.Week.Wednesday, 
			alarm.Week.Thursday, 
			alarm.Week.Friday, 
			alarm.Week.Saturday, 
			alarm.User_ID)
		if err != nil {
			fmt.Println("failure: ", err)
			writer.Write([]byte("Something went wrong in DB process"))
		} else {
			fmt.Println("THIS IS AN ALARM", alarm.User_ID, alarm.Alarm_ID, 	alarm.Time,res)
			a.createCronFromAlarm(alarm)
		
			writer.Write([]byte("Success"))
		}
	}

	defer request.Body.Close()
}
func (a *App) deleteAlarm(writer http.ResponseWriter, request *http.Request){
	var alarm alarm
	decoder := json.NewDecoder(request.Body)

	errDecode := decoder.Decode(&alarm)
	fmt.Printf("%v: %v\n", alarm.Time, alarm.Week)
	if errDecode != nil {
		fmt.Println(errDecode)
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	
		_, err := a.DB.Exec(
			`DELETE
			 FROM alarms
			 WHERE id = $1 `,
			alarm.Alarm_ID)
		if err != nil {
			fmt.Println("failure: ", err)
			writer.Write([]byte("Something went wrong in DB process"))
		} else {
			writer.Write([]byte("Success"))
		}
	

	defer request.Body.Close()
}

func (a *App) authenticationEndpoint(writer http.ResponseWriter, request *http.Request) {
	var account account
	decoder := json.NewDecoder(request.Body)
	errDecode := decoder.Decode(&account)
	if errDecode != nil {
		fmt.Println(errDecode)
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	hashedPassword := hashPassword(account.Password)
	var id string
	err := a.DB.QueryRow("SELECT id FROM users WHERE username = $3 AND password = $4", account.Username, hashedPassword).Scan(&id)
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Problem: Username or password is incorrect")
		fmt.Println(err)
	} else {
		writer.Write([]byte("Successful find"))
	}
	defer request.Body.Close()
}

func (a *App) forgotPassword(writer http.ResponseWriter, request *http.Request) {
	var user account
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		respondWithError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if user.Username == "" || user.Phone == "" || user.Password == "" {
		respondWithError(writer, http.StatusBadRequest, "Missing required field(s)")
		return
	}
	for _, char := range user.Password {
		if char > 127 {
			respondWithError(writer, http.StatusNotFound, "Problem: Password must only contain ASCII characters")
			return
		}
	}
	var retrievedAccount account
	err = a.DB.QueryRow("SELECT email, password, phone FROM users WHERE username = ? AND phone = ?", user.Username, user.Phone).Scan(&retrievedAccount.Email, &retrievedAccount.Password, &retrievedAccount.Phone)
	if err != nil {
		respondWithError(writer, http.StatusNotFound, "Account not found")
		return
	}
	hashedPassword := hashPassword(user.Password)
	_, err = a.DB.Exec("UPDATE users SET password = ? WHERE username = ? AND phone = ?", hashedPassword, user.Username, user.Phone)
	if err != nil {
		log.Fatal(err)
	}
	writer.Write([]byte("Password successfully updated"))
}




func main() {
	app := &App{}
	app.initializeApp()
	app.initCron()
	app.router.HandleFunc("/api/v1/", sayHello).Methods("GET")
	app.router.HandleFunc("/api/v1/createAlarm", app.createAlarm).Methods("POST")
	app.router.HandleFunc("/api/v1/retrieveAlarms", app.retrieveAlarms).Methods("POST")
	app.router.HandleFunc("/api/v1/updateAlarm", app.updateAlarm).Methods("POST")
	app.router.HandleFunc("/api/v1/deleteAlarm", app.deleteAlarm).Methods("POST")
	app.router.HandleFunc("/api/v1/createUser", app.createUser).Methods("POST")
	app.router.HandleFunc("/api/v1/login", app.authenticationEndpoint).Methods("POST")
	app.router.HandleFunc("/api/v1/forgotPassword", app.forgotPassword).Methods("POST")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	handler := c.Handler(app.router)
	fmt.Println("Server at 8123")
	log.Fatal(http.ListenAndServe(":8123", handler))

}
