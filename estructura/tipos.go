package estructura

import (
	"github.com/Fernando-MGS/TEST/AV"
	"github.com/Fernando-MGS/TEST/list"
)

type Archivo struct {
	Datos []Data `json:"Dato,omitempty"`
}

type Data struct {
	Indice        string  `json:"Indice,omitempty"`
	Departamentos []Depto `json:"Departamentos,omitempty"`
}
type Depto struct {
	Nombre  string        `json:"Nombre,omitempty"`
	Tiendas []list.Tienda `json:"Tiendas,omitempty"`
}

type Tienda struct {
	Nombre       string `json:"Nombre,omitempty"`
	Descripcion  string `json:"Descripcion,omitempty"`
	Contacto     string `json:"Contacto,omitempty"`
	Calificacion int    `json:"Calificacion,omitempty"`
	Logo         string `json:"Logo,omitempty"`
	ID           string
	Inventario   AV.AVL
}

type Inventario struct {
	Tienda       string     `json:"Tienda,omitempty"`
	Departamento string     `json:"Departamento,omitempty"`
	Calificacion int        `json:"Calificacion,omitempty"`
	Productos    []Producto `json:"Productos,omitempty"`
}

type Producto struct {
	Nombre      string  `json:"Nombre,omitempty"`
	Codigo      int     `json:"Codigo,omitempty"`
	Descripcion string  `json:"Descripcion,omitempty"`
	Precio      float64 `json:"Precio,omitempty"`
	Cantidad    int     `json:"Cantidad,omitempty"`
	Imagen      string  `json:"Imagen,omitempty"`
	Cant        []int
}
type Pedidos struct {
	Pedidos []Pedido `json:"Pedidos,omitempty"`
}

type Pedido struct {
	Fecha        string        `json:"Fecha,omitempty"`
	Tienda       string        `json:"Tienda,omitempty"`
	Departamento string        `json:"Departamento,omitempty"`
	Calificacion int           `json:"Calificacion,omitempty"`
	Producto     []AV.Producto `json:"Productos,omitempty"`
}
