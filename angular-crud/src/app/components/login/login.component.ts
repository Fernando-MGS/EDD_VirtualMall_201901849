import { Component, OnInit } from '@angular/core';
import { FormControl,FormGroup,FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import {TiendaService} from 'src/app/tienda.service';
import {User} from 'src/app/models/user'
import {Consulta} from 'src/app/models/consulta'
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})

export class LoginComponent implements OnInit {
  
  constructor(private storeServices:TiendaService) { }
  user: User;
  consulta: Consulta
  Register= new FormGroup({
    Dpi:new FormControl('',Validators.required),
    Nombre:new FormControl('',Validators.required),
    Correo:new FormControl('',Validators.required),
    Password:new FormControl('',Validators.required)
    
  })
  Login_form= new FormGroup({
    Nombre:new FormControl('',Validators.required),
    Password:new FormControl('',Validators.required),
    dpi:new FormControl('',Validators.required)
  })

  log(){
    console.log("vamo a ver")
    console.log(this.Register.value)
    console.log(this.Login_form.value.l_userName)
  }
  registrar(){
    this.user=this.Register.value
    this.storeServices.RegisUser(this.user).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
    console.log(this.user)
  }
  logear(){
    this.consulta=this.Login_form.value
    this.storeServices.LoginUser(this.consulta).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
    console.log(this.consulta)
  }
  ngOnInit(){
    
  }

}
