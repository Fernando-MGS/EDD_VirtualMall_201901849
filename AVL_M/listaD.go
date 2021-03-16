package AVL_M

import (
	"fmt"
)

//a ver si se subio
type nodo_l struct { //lugar donde se almacena, tipo tienda del package empresa
	anterior  *nodo_l
	siguiente *nodo_l
	mes       int
}

type Lista struct { //apuntadores
	inicio *nodo_l
	ultimo *nodo_l
	tam    int
	año    int
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
}

func (m *Lista) Insertar(valor Tienda) { //insertar un nodo_l
	nuevo := &nodo_l{nil, nil, valor}
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
		fmt.Println(aux.dato)
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

func (m *Lista) GetItem(index int) Tienda { //Devuelve un dato de la lista
	ind := 1
	aux := m.inicio
	if index <= m.tam {
		for ind < index {
			aux = aux.siguiente
			ind++
		}

	}
	return aux
}

/*func (m *Lista) Get(nombre string) Tienda { //Devuelve una tienda en especifico
	aux := m.inicio
	for nombre != aux.dato.Nombre {
		aux = aux.siguiente
	}
	return aux.dato
}*/
