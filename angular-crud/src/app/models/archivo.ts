import {Store} from './store';
export class Data{
    Indice: string
    Departamentos: Depto[] 
}
class Depto{
    Nombre: string
    Tiendas: Store []
}

