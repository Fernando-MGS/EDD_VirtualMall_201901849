package estructura

type archivo struct {
	Datos []Data `json: "Dato,omitempty"`
}

type Data struct {
	Indice        string  `json: "Indice, omitempty"`
	Departamentos []Depto `json: "Departamento, omitempty"`
}
type Depto struct {
	Nombre  string   `json: "Nombre, omitempty"`
	Tiendas []Tienda `json: "Tiendas, omitempty"`
}

type Tienda struct {
	Nombre       string `json: "Nombre, omitempty"`
	Descripcion  string `json: "Descripcion, omitempty"`
	Contacto     string `json: "Contacto, omitempty"`
	Calificacion int    `json: "Calificacion, omitempty"`
}
