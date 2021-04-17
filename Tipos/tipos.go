package Tipos

//"github.com/Fernando-MGS/EDD_MALL/Back/AV"

//"github.com/Fernando-MGS/TEST/list"

type Archivo struct {
	Datos []Data `json:"Datos,omitempty"`
}

type Data struct {
	Indice        string  `json:"Indice,omitempty"`
	Departamentos []Depto `json:"Departamentos,omitempty"`
}
type Depto struct {
	Nombre  string   `json:"Nombre,omitempty"`
	Tiendas []Tienda `json:"Tiendas,omitempty"`
}

type Tienda struct {
	Nombre       string `json:"Nombre,omitempty"`
	Descripcion  string `json:"Descripcion,omitempty"`
	Contacto     string `json:"Contacto,omitempty"`
	Calificacion int    `json:"Calificacion,omitempty"`
	Logo         string `json:"Logo,omitempty"`
	Departamento string
	ID           string
	Inventario   AVL
}

type _Inventario struct {
	Tienda       string     `json:"Tienda,omitempty"`
	Departamento string     `json:"Departamento,omitempty"`
	Calificacion int        `json:"Calificacion,omitempty"`
	Productos    []Producto `json:"Productos,omitempty"`
}

type Producto struct {
	Nombre         string  `json:"Nombre,omitempty"`
	Codigo         int     `json:"Codigo,omitempty"`
	Descripcion    string  `json:"Descripcion,omitempty"`
	Precio         float64 `json:"Precio,omitempty"`
	Cantidad       int     `json:"Cantidad,omitempty"`
	Imagen         string  `json:"Imagen,omitempty"`
	Almacenamiento string  `json:"Almacenamiento,omitempty"`
	ID             string
	Cant           []int
}
type Pedidos struct {
	Pedidos []Pedido `json:"Pedidos,omitempty"`
}

type Pedido struct {
	Fecha        string     `json:"Fecha,omitempty"`
	Tienda       string     `json:"Tienda,omitempty"`
	Departamento string     `json:"Departamento,omitempty"`
	Calificacion int        `json:"Calificacion,omitempty"`
	Producto     []Producto `json:"Productos,omitempty"`
}

type T_especifica []Tienda

type Busqueda struct {
	Departamento string `json:"Departamento,omitempty"`
	Nombre       string `json:"Nombre,omitempty"`
	Calificacion int    `json:"Calificacion,omitempty"`
}

type Stores struct {
	Array []Tienda
}

type Products struct {
	Array  []Producto
	Tamaño int
	Precio float64
}

type Inventarios struct {
	Inventario []_Inventario `json:"Inventarios,omitempty"`
}

type Months struct {
	Año    int
	Large  int
	Indice int
	Mes    []string
}

type Cuentas struct {
	Usuarios []Usuario `json:"Usuarios,omitempty"`
}

type Usuario struct {
	DPI      string `json:"DPI,omitempty"`
	Dpi_     int    `json:"Dpi,omitempty"`
	Nombre   string `json:"Nombre,omitempty"`
	Correo   string `json:"Correo,omitempty"`
	Password string `json:"Password,omitempty"`
	Pass     string
	Cuenta   string `json:"Cuenta,omitempty"`
	SHA_pass [32]byte
	Tipo     int //1 admin, 2 user
}

type Consulta struct {
	Nombre   string `json:"Nombre,omitempty"`
	Password string `json:"Password,omitempty"`
	DPI      string `json:"dpi,omitempty"`
}

type Clave struct {
	Clave string `json:"Clave,omitempty"`
}
type File_grafo struct {
	Nodos    []_NodoG `json:"Nodos,omitempty"`
	Pos_init string   `json:"PosicionInicialRobot,omitempty"`
	Entrega  string   `json:"Entrega,omitempty"`
}

type _NodoG struct {
	Nombre  string   `json:"Nombre,omitempty"`
	Enlaces []Enlace `json:"Enlaces,omitempty"`
}
type Arista struct {
	Destino *Nodo_G
	Peso    int
}
type Almacen struct {
	Estructura []*Nodo_G
	Pos_Robot  string
	Entrega    string
}

type Enlace struct {
	Nombre    string `json:"Nombre,omitempty"`
	Distancia int    `json:"Distancia,omitempty"`
}

/*type Años struct {
	Datos  []pedidos.Meses
	indice int
	large  int
}*/
