import { Component, OnInit } from '@angular/core';
import { TiendaService } from 'src/app/tienda.service';
import { ActivatedRoute, Params } from '@angular/router';
import { Comentario } from 'src/app/models/comentario';
import { Prod } from 'src/app/models/producto';
import {Comentarios} from 'src/app/models/comentarios';
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
  t: number[] = [];
  constructor(private rutaActiva: ActivatedRoute, private storeServices: TiendaService) { }
  Respondiendo: string
  test() {
    console.log("Hola si funciono uwu")
    this.Respondiendo = "Respondiendo a"
  }
  comentar(text: string) {
    var comentario: Comentario = { Contenido: text, User: this.user }
    this.storeServices.Comentar(comentario, this.id).subscribe(data => console.log(data), err => console.log(err), () => console.log("Finish"));
    console.log(text)
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
      this.Comments = res;
      console.log(this.Comments, "Hola2")
    });
  }


}
