package pedidos

import (
	"fmt"

	"github.com/Fernando-MGS/TEST/AV"
)

//a ver si se subio
type nodo_l struct { //lugar donde se almacena, tipo tienda del package empresa
	anterior  *nodo_l
	siguiente *nodo_l
	Mes       int
	pedidos   matriz
}

type Lista struct { //apuntadores
	inicio *nodo_l
	ultimo *nodo_l
	tam    int
}

func NewLista() *Lista { //crea una lista
	return &Lista{nil, nil, 0} //el inicio es null, final null, y el tamaño es 0
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

func (m *Lista) Insertar(mes, depto, dia int, l_prod []AV.Producto) { //insertar un nodo_m
	matriz := NewMatriz()
	cont := 0
	for cont < len(l_prod) {
		matriz.Insert(l_prod[cont], dia, depto)
		cont++
	}
	nuevo := &nodo_l{nil, nil, mes, *matriz}
	if m.inicio == nil {
		m.inicio = nuevo
		m.ultimo = nuevo
	} else {
		m.ultimo.siguiente = nuevo
		nuevo.anterior = m.ultimo
		m.ultimo = nuevo
	}
	m.tam++
}

func (m *Lista) Imprimir() { //IMPRIMIR
	aux := m.inicio
	for aux != nil {
		fmt.Println(aux.Mes)
		aux.pedidos.lst_h.print_h()
		aux.pedidos.lst_v.print()
		aux = aux.siguiente
	}
	fmt.Print(m.tam)
}

func (m *Lista) Tamaño() int { //IMPRIMIR TAMAÑO DE LISTA
	return m.tam
}

func (m *Lista) Borrar(pos int) {
	aux := m.inicio
	sum := 0
	for sum < pos {
		aux = aux.siguiente
		sum++
	}
	fmt.Println("Llegom", sum)
	if m.inicio == aux {
		fmt.Println("inicio", m.tam)
		//i := m.tam - 1
		m.inicio = m.inicio.siguiente
		m.inicio.anterior = nil
		/*if i > 0 {
			fmt.Println("Llego 1.", i)
			m.inicio = aux.siguiente
			aux.siguiente.anterior = nil
			aux.siguiente = nil
		} else {
			fmt.Println("Llego 2", aux.dato.Nombre, " ", i)
			aux = nil
			m.inicio = aux
			m.inicio.siguiente = nil
			m.inicio.anterior = nil
		}*/

	} else if m.ultimo == aux {
		fmt.Println("adios1")
		m.ultimo = aux.anterior
		aux.anterior.siguiente = nil
	} else {
		fmt.Println("adios2")
		aux.anterior.siguiente = aux.siguiente
		aux.siguiente.anterior = aux.anterior
		aux.anterior = nil
		aux.siguiente = nil
	}

}

func (m *Lista) GetItem(index int) matriz { //Devuelve un dato de la lista
	ind := 1
	aux := m.inicio
	if index <= m.tam {
		for ind < index {
			aux = aux.siguiente
			ind++
		}
	}
	return aux.pedidos
}
func (m *Lista) buscar(mes, dia, depto int, l_prod []AV.Producto) int {
	aux := m.inicio
	ind := 1
	find := 0
	for ind <= m.tam {
		if aux.Mes == mes {
			find = 1
			cont := 0
			for cont < len(l_prod) {
				aux.pedidos.Insert(l_prod[cont], dia, depto)
				cont++
			}
			break
		}
		ind++
	}
	return find
}

func (m *Lista) Insercion(l_prod []AV.Producto, depto, mes, dia int) {
	find := m.buscar(mes, dia, depto, l_prod)
	if find == 0 {
		m.Insertar(mes, depto, dia, l_prod)
	}
}

/*func (m *Lista) Get(nombre string) Tienda { //Devuelve una tienda en especifico
	aux := m.inicio
	for nombre != aux.dato.Nombre {
		aux = aux.siguiente
	}
	return aux.dato
}*/

/*func (m *Lista) Set_Inventario(tmp Tienda) { //Devuelve un dato de la lista
	aux := m.inicio
	for tmp.Nombre != aux.dato.Nombre {
		aux = aux.siguiente
	}
	aux.dato = tmp
}*/
