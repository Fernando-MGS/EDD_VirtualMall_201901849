import { Component, OnInit } from '@angular/core';
import {User} from 'src/app/models/user'
import {TiendaService} from 'src/app/tienda.service';
@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  user:any
  tipo:any
  a:User={DPI:"klj",Nombre:"Stonks",Correo:"hola@gmail.com",Password:"",Tipo:0}
  constructor(private storeServices: TiendaService) { }

  log_out(){
    this.storeServices.Logout().subscribe()
  }

  ngOnInit(): void {
    this.storeServices.Tipo_User().subscribe((resP) =>{
      this.user=resP;
      this.tipo=this.user.Tipo
      console.log(this.user)
      console.log(this.tipo)
    });
  }

}
