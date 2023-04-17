import { FormGroup, ValidatorFn, ValidationErrors, AbstractControl } from '@angular/forms';



export const passwordMatchValidator: ValidatorFn = (formGroup: AbstractControl): ValidationErrors | null => {
    const passwordControl = formGroup.get('passwordControl');
    const confirmPasswordControl = formGroup.get('confirmPasswordControl');
    if(passwordControl == null || confirmPasswordControl == null){
        return null;
    }
      if(passwordControl.value === confirmPasswordControl.value){
        return null;
      }
      else{
        return {'passwordMismatch':true};
      }
};
