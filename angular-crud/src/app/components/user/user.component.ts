import { Component, OnInit } from '@angular/core';
import {User} from 'src/app/models/user'
@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  a:User={Dpi:"klj",Nombre:"Stonks",Correo:"hola@gmail.com",Password:""}
  constructor() { }

  ngOnInit(): void {
  }

}
