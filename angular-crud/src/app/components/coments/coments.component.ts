import { Component, OnInit } from '@angular/core';
import {TiendaService} from 'src/app/tienda.service';
import { ActivatedRoute, Params } from '@angular/router';
import {Comentario} from 'src/app/models/comentario';
import {Prod} from 'src/app/models/producto';
import {Store} from 'src/app/models/store';
import {Comentarios} from 'src/app/models/comentarios';
@Component({
  selector: 'app-coments',
  templateUrl: './coments.component.html',
  styleUrls: ['./coments.component.css']
})
export class ComentsComponent implements OnInit {
  id: string;
  user:any;
  articulo:any;  
  imagen:string;
  producto: Store;
  t:number[]=[];
  comentarios:any
  Comments:Comentarios[]
  r_activ:number
  constructor(private rutaActiva: ActivatedRoute, private storeServices: TiendaService) { 

  }
  Respondiendo:string
  test(user:any, index:any){
    this.r_activ=0
    console.log(index)
    this.Respondiendo="Respondiendo a "+user
  }
  comentar(text:string){
    var comentario: Comentario={Contenido:text,User:this.user}
    if (this.r_activ!=0){
    this.storeServices.Comentar(comentario,this.id).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
    }else{
      console.log("comentararara")
    }
  }

  ngOnInit(): void {
    this.id=this.rutaActiva.snapshot.params.id
    this.storeServices.Tipo_User().subscribe((resP) =>{
      this.user=resP;
    });
    var split = this.id.split("-",3);
this.storeServices.Articulo(this.id).subscribe((res_) =>{
  this.producto=res_;
});
this.storeServices.Comentarios(this.id).subscribe((res) =>{
  this.Comments=res.Comentarios;
  console.log(this.Comments,"Hola 2")
}); 
    console.log("Hola 3 ",this.r_activ)
  }

}
