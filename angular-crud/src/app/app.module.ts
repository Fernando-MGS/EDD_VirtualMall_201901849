import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from "@angular/common/http";
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

//servicios
import {  TiendaService} from './tienda.service';
import { ProductsComponent } from './components/products/products.component';
import { TiendasComponent } from './components/tiendas/tiendas.component';
import { CargaComponent } from './components/carga/carga.component'

@NgModule({
  declarations: [
  TiendasComponent,
  CargaComponent],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [
    TiendaService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
