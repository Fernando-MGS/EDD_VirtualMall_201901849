import { Component, OnInit } from '@angular/core';
import {Store} from 'src/app/models/store';
import {Data} from 'src/app/models/archivo'; 
import {POST} from 'src/app/models/post';
import {Master} from 'src/app/models/master';
import { FormControl,FormGroup,FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
//import {Archivo} from '.models/store';
import {TiendaService} from 'src/app/tienda.service';
@Component({
  selector: 'app-carga',
  templateUrl: './carga.component.html',
  styleUrls: ['./carga.component.css'],
  providers: [TiendaService
]
})
export class CargaComponent implements OnInit {
  tiendas: any[]=[];
  master: Master
  test: any;
  inv:any;
  post:POST
  p:string
  storearray: Data[]=[
    /*{Nombre:"Walmart",Descripcion:"Soy una descripción yei",Contacto:"4554545",Calificacion:5,Logo:"https://www.braindw.com/wp-content/uploads/2018/05/logo-walmart.jpg"},     
    {Nombre:"Gatorade",Descripcion:"Soy otra descripción no yei",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/05/Gatorade-Logo.png"},
    {Nombre:"PlayStation",Descripcion:"Meh, que ves",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/11/PlayStation-Logotipo1994-2009.jpg"}*/
  ];
  Register= new FormGroup({
    Key:new FormControl('',Validators.required),
  })
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
  postPedido(file:any){
    this.inv=JSON.parse(file);
    this.storeServices.addPedido(this.inv).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
    console.log(this.inv);
  }
  postUser(file:any){
    this.inv=JSON.parse(file)
    this.storeServices.LoadUser(this.inv).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
    console.log(this.inv);
  }
  Grafo(file:any){
    
    
  }
  PostGrafo(file:any){
    this.test=JSON.parse(file)
    this.storeServices.Post_grafo(this.test).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }
  Master(){
    console.log(this.master)
    console.log("-??????????????????")
    this.storeServices.Master_Key(this.Register.value).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
    console.log("/")
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
