import { Component } from '@angular/core';
import {Store} from './models/store';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  storearray: Store[]=[
    {Nombre:"Walmart",Descripcion:"Yes",Contacto:"4554545",Calificacion:5,Logo:"https://www.braindw.com/wp-content/uploads/2018/05/logo-walmart.jpg"},     
    {Nombre:"Gatorade",Descripcion:"Yes",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/05/Gatorade-Logo.png"},
    {Nombre:"PlayStation",Descripcion:"Yes",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/11/PlayStation-Logotipo1994-2009.jpg"}
  ]; 
}
