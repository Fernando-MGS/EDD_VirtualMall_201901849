import { Component, OnInit } from '@angular/core';
import { TiendaService } from 'src/app/tienda.service';
import { ActivatedRoute, Params } from '@angular/router';
import { Comentario } from 'src/app/models/comentario';
import { Prod } from 'src/app/models/producto';
import {Comentarios} from 'src/app/models/comentarios';
import {Respuestas} from 'src/app/models/respuestas';
@Component({
  selector: 'app-comments',
  templateUrl: './comments.component.html',
  styleUrls: ['./comments.component.css']
})
export class CommentsComponent implements OnInit {
  id: string;
  user: any;
  articulo: any;
  imagen: string;
  producto: Prod;
  comentario:any
  Comments:Comentarios[]
  indice:any
  receptor: string
  t: number[] = [];
  constructor(private rutaActiva: ActivatedRoute, private storeServices: TiendaService) { }
  Respondiendo: string
  r_activ=1
  r_index:any
  test(user:any,index:any) {
    this.r_activ=0
    console.log(index)
    this.indice=index
    console.log("Hola si funciono uwu")
    this.Respondiendo = "Respondiendo a "+user
    this.receptor=user
  }
  comentar(text: string) {
    this.r_index=this.indice+""
    var comentario: Comentario={Contenido:text,User:this.user}
    if (this.r_activ!=0){
    this.storeServices.Comentar(comentario,this.id).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
    }else{
      var Respuestas:Respuestas={Index:this.r_index,Respuesta:text,Receptor:this.receptor,User:this.user}
      this.storeServices.Responder(Respuestas,this.id).subscribe(data=>console.log(data),err=>console.log(err),()=>console.log("Finish"));
    }
  }
  ngOnInit(): void {
    this.id = this.rutaActiva.snapshot.params.id
    this.storeServices.Tipo_User().subscribe((resP) => {
      this.user = resP;
    });
    console.log(this.id)
    var split = this.id.split("-", 3);
    this.storeServices.Articulo(this.id).subscribe((res_) => {
      this.producto = res_;
      console.log(this.producto, "Hola")
    });
    this.storeServices.Comentarios(this.id).subscribe((res) => {
      this.Comments = res.Comentarios;
      console.log(this.Comments, "Hola2")
    });
    console.log("Hola 3 ",this.r_activ)
  }


}
