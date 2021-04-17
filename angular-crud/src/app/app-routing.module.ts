import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {ProductsComponent} from './components/products/products.component';
import {AppComponent} from './app.component'
import { BrowserModule } from '@angular/platform-browser'
import { CargaComponent } from './components/carga/carga.component';
import { TiendasComponent } from './components/tiendas/tiendas.component';
import { InventarioComponent } from './components/inventario/inventario.component';
import { CarritoComponent } from './components/carrito/carrito.component';
import { PedidosComponent } from './components/pedidos/pedidos.component';
import {LoginComponent} from './components/login/login.component';
import {UserComponent} from './components/user/user.component'
const routes: Routes = [
  
  { path:'', component: LoginComponent},
  { path:'', component: AppComponent},
  { path:'login', component: LoginComponent},
  { path:'profile', component: UserComponent},
  { path:'carga', component: CargaComponent},
  { path:'tiendas', component: TiendasComponent},
  {path:'inventario/:id', component: InventarioComponent},
  {path:'carrito', component: CarritoComponent},
  {path:'pedidos', component: PedidosComponent},
  {path:'pedido/:id', component: PedidosComponent}
  
];

@NgModule({
  declarations: [AppComponent,ProductsComponent],
  imports: [RouterModule.forRoot(routes),
            BrowserModule
  ],
})
export class AppRoutingModule { }
