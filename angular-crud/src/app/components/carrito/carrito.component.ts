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
  test: string
  constructor(private storeServices: TiendaService) { }
  carrito(){
    this.storeServices.pedido_Cart(this.Precio).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }
  print(){
    this.test="Hola"
  }
  offProd(file:any, num:any){
    this.storeServices.off_Cart(file,num).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
    console.log(file)
  }
  ngOnInit(){
    console.log(this.test)
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
