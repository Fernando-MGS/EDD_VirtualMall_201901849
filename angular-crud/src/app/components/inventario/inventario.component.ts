import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import {TiendaService} from 'src/app/tienda.service';
import {Prod} from 'src/app/models/producto';
import { FormControl, FormGroup, Validators } from '@angular/forms';
@Component({
  selector: 'app-inventario',
  templateUrl: './inventario.component.html',
  styleUrls: ['./inventario.component.css'],
  providers: [TiendaService] 
})
export class InventarioComponent implements OnInit {
  coche: {marca: string, modelo: string};
  id: string
  url: string
  conc:string
  constructor(private rutaActiva: ActivatedRoute, private storeServices: TiendaService) {}
  Prodc: Prod[]
  Cantidad : FormGroup;
  opcionSeleccionado: string  = '0';
  verSeleccion: string        = '';
  mensaje: string
  capturar(file:any) {
    // Pasamos el valor seleccionado a la variable verSeleccion
    console.log(file)

}

  addProd(file:any, num:any, stock:any){
    var cantidad: number=+num
    var _stock: number=+stock
    if (_stock>=cantidad){
      this.mensaje="Si hay stock"
      this.url=this.id.concat("-")
      this.url=this.url.concat(num)
      console.log("Hola add")
      this.storeServices.addProd(file,this.url).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
      console.log(file);
    }else{
      this.mensaje="No hay stock"
    }
    
    
   
  }
  ngOnInit(): void {
    this.id=this.rutaActiva.snapshot.params.id
    this.storeServices.getProd(this.id).subscribe((res)=>{
      this.Prodc=res.Array
      console.log(this.Prodc[0])
    })
  }

}
