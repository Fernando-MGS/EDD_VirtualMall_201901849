import { Component, OnInit } from '@angular/core';
import {TiendaService} from 'src/app/tienda.service';
import {Prod} from 'src/app/models/producto';
@Component({
  selector: 'app-carrito',
  templateUrl: './carrito.component.html',
  styleUrls: ['./carrito.component.css'],
  providers: [TiendaService] 
})
export class CarritoComponent implements OnInit {
  Prodc: Prod[]   
  Large: number
  Precio: number
  constructor(private storeServices: TiendaService) { }
  ngOnInit(){
    this.storeServices.GetCart().subscribe((res) =>{
      this.Prodc=res.Array;
      this.Large=res.Tamaño;
      this.Precio=res.Precio;
      console.log(res);
        console.log(this.Large);
        console.log(this.Precio);
    });
  }

}
