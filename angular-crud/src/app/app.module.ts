import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from "@angular/common/http";
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import {ReactiveFormsModule} from '@angular/forms';
//servicios
import {  TiendaService} from './tienda.service';
import { ProductsComponent } from './components/products/products.component';
import { TiendasComponent } from './components/tiendas/tiendas.component';
import { CargaComponent } from './components/carga/carga.component';
import { InventarioComponent } from './components/inventario/inventario.component';
import { CarritoComponent } from './components/carrito/carrito.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { PedidosComponent } from './components/pedidos/pedidos.component';
import { LoginComponent } from './components/login/login.component';
import { UserComponent } from './components/user/user.component';
import { ComentarioComponent } from './components/comentario/comentario.component'

@NgModule({
  declarations: [
  TiendasComponent,
  CargaComponent,
  InventarioComponent,
  CarritoComponent,
  NavbarComponent,
  PedidosComponent,
  LoginComponent,
  UserComponent,
  ComentarioComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [
    TiendaService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
