import { Component, OnInit } from '@angular/core';
import {Store} from './models/store';
import {TiendaService} from './tienda.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [TiendaService]
})
export class AppComponent implements OnInit{
  tiendas: any[]=[];

  storearray: Store[]=[
    /*{Nombre:"Walmart",Descripcion:"Soy una descripción yei",Contacto:"4554545",Calificacion:5,Logo:"https://www.braindw.com/wp-content/uploads/2018/05/logo-walmart.jpg"},     
    {Nombre:"Gatorade",Descripcion:"Soy otra descripción no yei",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/05/Gatorade-Logo.png"},
    {Nombre:"PlayStation",Descripcion:"Meh, que ves",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/11/PlayStation-Logotipo1994-2009.jpg"}*/
  ];
  
  Prueba: string="Hola string";
  Stores: Store[]=[];
  constructor(public storeServices: TiendaService){ }
  
  ngOnInit(){
    this.storeServices.getStore().subscribe((res) =>{
      this.storearray=res.Array;
      console.log(res);
        console.log(this.storearray);
    });
  }
}
