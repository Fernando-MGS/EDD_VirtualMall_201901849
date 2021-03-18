import { Component, OnInit } from '@angular/core';
import {Store} from './models/store';
import {Data} from './models/archivo'; 
//import {Archivo} from '.models/store';
import {TiendaService} from './tienda.service';
import {CargaService} from './carga.service';
@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [TiendaService,
            CargaService
    ]
})
export class AppComponent implements OnInit{
  tiendas: any[]=[];
  test: any;
  inv:any;
  storearray: Data[]=[
    /*{Nombre:"Walmart",Descripcion:"Soy una descripción yei",Contacto:"4554545",Calificacion:5,Logo:"https://www.braindw.com/wp-content/uploads/2018/05/logo-walmart.jpg"},     
    {Nombre:"Gatorade",Descripcion:"Soy otra descripción no yei",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/05/Gatorade-Logo.png"},
    {Nombre:"PlayStation",Descripcion:"Meh, que ves",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/11/PlayStation-Logotipo1994-2009.jpg"}*/
  ];
  
  Prueba: string="Hola string";
  Stores: Store[]=[];
  constructor(private storeServices: TiendaService){ }
  
getStores(file:any){
  this.test=JSON.parse(file)
  this.storearray=this.test
  this.storeServices.putStores(this.test).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  console.log(this.test);
}
getInvent(file:any){
  this.inv=JSON.parse(file);
  this.storeServices.putInvts(this.inv).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  console.log(this.inv);
}
  ngOnInit(){
    console.log("Jola");
   /* this.storeServices.getStore().subscribe((res) =>{
      this.storearray=res.Array;
      console.log(res);
        console.log(this.storearray);
    });*/
  }
}
