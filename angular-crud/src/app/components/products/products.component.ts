import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import {TiendaService} from 'src/app/tienda.service';
import {Prod} from 'src/app/models/producto';
@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css'],
  providers: [TiendaService] 
})
export class ProductsComponent implements OnInit {

  constructor(private rutaActiva: ActivatedRoute, private storeServices: TiendaService) { }
  //Products: Prod[]
  ID: string
  Nombre: string
  Productos: any
  ngOnInit() {
    
    /*this.Nombre=(this.rutaActiva.snapshot.params.Nombre)
    console.log(this.rutaActiva.snapshot.params.id);*/
  }

}
