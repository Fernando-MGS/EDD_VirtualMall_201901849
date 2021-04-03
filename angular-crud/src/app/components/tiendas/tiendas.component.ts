import { Component, OnInit } from '@angular/core';
import {Store} from 'src/app/models/store'
import {TiendaService} from 'src/app/tienda.service';
@Component({
  selector: 'app-tiendas',
  templateUrl: './tiendas.component.html',
  styleUrls: ['./tiendas.component.css'],
  providers: [TiendaService]
})
export class TiendasComponent implements OnInit {
  tiendas: any[]=[];

  storearray: Store[]=[
    /*{Nombre:"Walmart",Descripcion:"Soy una descripción yei",Contacto:"4554545",Calificacion:5,Logo:"https://www.braindw.com/wp-content/uploads/2018/05/logo-walmart.jpg"},     
    {Nombre:"Gatorade",Descripcion:"Soy otra descripción no yei",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/05/Gatorade-Logo.png"},
    {Nombre:"PlayStation",Descripcion:"Meh, que ves",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/11/PlayStation-Logotipo1994-2009.jpg"}*/
  ];
  
  Prueba: string="Hola string";
  Stores: Store[]=[];
  constructor(public storeServices: TiendaService){ }

  ngOnInit() {
    let rootVar = window['hola'];
      rootVar += 1;
      window['hola'] = rootVar;
    this.storeServices.getStore().subscribe((res) =>{
      this.storearray=res.Array;
      console.log(res);
        console.log(this.storearray);
    });
  }

}
