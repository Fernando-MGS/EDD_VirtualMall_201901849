import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import {TiendaService} from 'src/app/tienda.service';
import {Pedidos} from 'src/app/models/pedidos'
import {Prod} from 'src/app/models/producto'
import {M_front} from 'src/app/models/matriz'
import {POST} from 'src/app/models/post'
@Component({
  selector: 'app-pedidos',
  templateUrl: './pedidos.component.html',
  styleUrls: ['./pedidos.component.css'],
  providers: [TiendaService] 
})
export class PedidosComponent implements OnInit {
  id: string
  arr:Prod[]
  indice: any
  large: any
  url:string
  y1:string
  mes:string
  pedido: any
  year: any
  test: any
  testito:string="1"
  rest:M_front[]
  n_cif:POST={Tipo:"0",Par1:"0",Par2:"0",Par3:0,Par4:0}
  cif:POST={Tipo:"1",Par1:"0",Par2:"0",Par3:0,Par4:0}
  cif_s:POST={Tipo:"2",Par1:"0",Par2:"0",Par3:0,Par4:0}
  constructor(private rutaActiva: ActivatedRoute, private storeServices: TiendaService) { }

  tesl(file){
    var x=this.year.toString()
    this.url=x.concat("-")
    this.url=this.url.concat(file)
    this.storeServices.Dev_mes(this.url).subscribe((res)=>{
      this.rest=res
      
    })
    var post: POST={Tipo:"0",Par1:"0",Par2:file,Par3:this.year,Par4:0}
    
  }
  pedidos(chain:any){
    var split = chain.concat('-',this.year)
    
  }

  fix_user(){
    this.storeServices.fix_user().subscribe((res)=>{})
  }

  fix_store(){
    this.storeServices.fix_store().subscribe((res)=>{})
  }

  graf_alm(){
    this.storeServices.Graph_Alm().subscribe((res)=>{})
  }
  graf_store(){
    this.storeServices.Graph_store().subscribe((res)=>{})
  }
  graf_user(){
    this.storeServices.Graf_users(this.n_cif).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }
  graf_user__(){
    this.storeServices.Graf_users(this.cif).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }
  graf_user_(){
    this.storeServices.Graf_users(this.cif_s).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }
  graf_year(){
    this.storeServices.Graph_year().subscribe((res)=>{})
  }
  graf_mes(file:any){
    var post: POST={Tipo:"0",Par1:"0",Par2:file,Par3:this.year,Par4:0}
    console.log(post)
    this.storeServices.Graph_mes(post).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }
  m_store(){
    this.storeServices.m_store().subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }
  m_ped(){
    this.storeServices.m_ped().subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }
  m_user(){
    this.storeServices.m_user().subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
  }

  ngOnInit(): void {
    
    this.id=this.rutaActiva.snapshot.params.id
    this.storeServices.dev_pedido(this.id).subscribe((res)=>{
      this.pedido=res.Mes
      this.indice=res.Indice
      this.large=res.Large
      this.year=res.Año
      console.log(res)
      
    })
  }

}
