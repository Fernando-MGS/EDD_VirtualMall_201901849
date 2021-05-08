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
  addProd(inv:any, url:any):Observable<any>{
    console.log("hOLA5");
    return this.http.post<any>('/addProducto/'+url,inv)
  }
  GetCart():Observable <any>{
    console.log("hOLA6")
    return this.http.get<any>('/getCart')
  }
  addPedido(pedido:any):Observable<any>{
    console.log("hOLA7")
    return this.http.post<any>('/Pedido',pedido)
  }
  pedido_Cart(pedido:any):Observable<any>{
    console.log("hOLA8")
    return this.http.post<any>('/PedidoCart',pedido)
  }
  off_Cart(pedido:any, url:any):Observable<any>{
    console.log("hOLA9")
    return this.http.post<any>('/offProduct/'+url,pedido)
  }
  dev_pedido(index:any):Observable<any>{
    console.log("hOLA10")
    return this.http.get<any>('/pedidos/'+index)
  }
  Cart_Size():Observable<any>{
    console.log("Hola 10")
    return this.http.get<any>('/CartSize')
  }
  Tipo_User():Observable<any>{
    console.log("Hola 11")
    return this.http.get<any>('/user')
  }
  LoadUser(users:any):Observable<any>{
    console.log("hOLA12");
    return this.http.post<any>('/LoadUsers',users,)
  }
  RegisUser(user:any):Observable<any>{
    console.log("hOLA13");
    return this.http.post<any>('/regisUser',user,)
  }
  LoginUser(user:any):Observable<any>{
    console.log("hOLA14");
    return this.http.post<any>('/loginUser',user,)
  }
  Logout():Observable<any>{
    return this.http.get<any>('/Logout')
  }
  Graph_year():Observable<any>{
    return this.http.get<any>('/year')
  }
  Graph_store():Observable<any>{
    return this.http.get<any>('/graf_stores')
  }
  Graph_Alm():Observable<any>{
    console.log('hOLA 15')
    return this.http.get<any>('/graf_grafo')
  }
  Graph_stores():Observable<any>{
    console.log('hOLA 15')
    return this.http.get<any>('/graf_stores')
  }
  Graph_mes(file:any): Observable<any>{
    return this.http.post<any>('/graf_mes',file)
  }
  Master_Key(file:any): Observable<any>{
    return this.http.post<any>('/masterKey',file)
  }
  Post_grafo(file:any):Observable<any>{
    return this.http.post<any>('/Loadgrafo',file)
  }
  Graf_users(file:any):Observable<any>{
    return this.http.post<any>('/graf_users',file)
  }
  Dev_mes(file:any):Observable<any>{
    return this.http.get<any>('/pedidos_mes/'+file)
  }
  Comentar(file:any, dir:any):Observable<any>{
    return this.http.post<any>('/comentario/'+dir,file)
  }
  Comentarios(dir:any):Observable<any>{
      return this.http.get<any>('/comentarios/'+dir)
  }
  Articulo(dir:any):Observable<any>{
    return this.http.get<any>('/articulo/'+dir)
  }
}
