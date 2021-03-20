import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import {Observable} from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class TiendaService {

  lista:any[]=[
    {
      Nombre:"Walmart",Descripcion:"Soy una descripción yei",Contacto:"4554545",Calificacion:5,Logo:"https://www.braindw.com/wp-content/uploads/2018/05/logo-walmart.jpg"},
      {Nombre:"Gatorade",Descripcion:"Soy otra descripción no yei",Contacto:"4554545",Calificacion:5, Logo:"https://logos-marcas.com/wp-content/uploads/2020/05/Gatorade-Logo.png"
    }

  ]

  constructor(private http: HttpClient) { 
    
  }
  obtenerLista(){
   return this.lista;
  }
  getStore():Observable <any>{
    console.log("hOLA1")
    return this.http.get<any>('/Tiendas')
  }
  putStores(stores:any):Observable<any>{
    console.log("hOLA2");
    return this.http.post<any>('/cargartienda',stores,)
  }
  putInvts(stores:any):Observable<any>{
    console.log("hOLA3");
    return this.http.post<any>('/Inventarios',stores,)
  }
  getProd(prod:any):Observable<any>{
    console.log("hoLA4");
    return this.http.get<any>('/products/'+prod)
  }
}
