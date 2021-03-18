import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import {Observable} from 'rxjs';
@Injectable({
  providedIn: 'root'
})
export class CargaService {

  constructor(private http: HttpClient) { }

  putStores(stores:any):Observable<any>{
    return this.http.post('/cargartienda',stores)
  }
}
