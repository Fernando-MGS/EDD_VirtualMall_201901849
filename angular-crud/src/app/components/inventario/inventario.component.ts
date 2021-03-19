import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import {TiendaService} from 'src/app/tienda.service';
import {Prod} from 'src/app/models/producto';
@Component({
  selector: 'app-inventario',
  templateUrl: './inventario.component.html',
  styleUrls: ['./inventario.component.css'],
  providers: [TiendaService] 
})
export class InventarioComponent implements OnInit {
  coche: {marca: string, modelo: string};
  id: string
  constructor(private rutaActiva: ActivatedRoute, private storeServices: TiendaService) { }
  Prodc: Prod[]
  ngOnInit(): void {
    this.id=this.rutaActiva.snapshot.paramMap.get('')
    
  }

}
