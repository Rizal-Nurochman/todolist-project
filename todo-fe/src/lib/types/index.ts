export interface FormProps{
    msg:string;
    action:string;
    defaultvalues:{
        email:string;
        password:string;
        confirmPassword?:string
    }
}
