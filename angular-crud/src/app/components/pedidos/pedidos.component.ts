import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import {TiendaService} from 'src/app/tienda.service';
import {Pedidos} from 'src/app/models/pedidos'
@Component({
  selector: 'app-pedidos',
  templateUrl: './pedidos.component.html',
  styleUrls: ['./pedidos.component.css'],
  providers: [TiendaService] 
})
export class PedidosComponent implements OnInit {
  id: string
  indice: any
  large: any
  pedido: any
  year: any
  test: any
  constructor(private rutaActiva: ActivatedRoute, private storeServices: TiendaService) { }

  pedidos(chain:any){
    var split = chain.concat('-',this.year)
    console.log(split)
  }

  ngOnInit(): void {
    
    this.id=this.rutaActiva.snapshot.params.id
    this.storeServices.dev_pedido(this.id).subscribe((res)=>{
      this.pedido=res.Mes
      this.indice=res.Indice
      this.large=res.Large
      this.year=res.Año
      console.log(res)
      console.log(this.indice)
    })
  }

}
