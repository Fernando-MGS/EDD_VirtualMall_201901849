package list

import (
	"fmt"

	"github.com/Fernando-MGS/TEST/Tipos"
)

//a ver si se subio
type nodo_m struct { //lugar donde se almacena, tipo tienda del package empresa
	anterior  *nodo_m
	siguiente *nodo_m
	dato      Tipos.Tienda
}

type Lista struct { //apuntadores
	inicio *nodo_m
	ultimo *nodo_m
	tam    int
}

func NewLista() *Lista { //crea una lista
	return &Lista{nil, nil, 0} //el inicio es null, final null, y el tamaño es 0
}

func (m *Lista) Insertar(valor Tipos.Tienda) { //insertar un nodo_m
	nuevo := &nodo_m{nil, nil, valor}
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
		fmt.Println(aux.dato.Nombre)
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

func (m *Lista) GetItem(index int) Tipos.Tienda { //Devuelve un dato de la lista
	ind := 1
	aux := m.inicio
	if index <= m.tam {
		for ind < index {
			aux = aux.siguiente
			ind++
		}

	}
	return aux.dato
}

func (m *Lista) Comentar(nombre string, comment Tipos.Comentario) { //Devuelve una tienda en especifico
	aux := m.inicio
	for nombre != aux.dato.Nombre {
		aux = aux.siguiente
	}
	aux.dato.Comentarios.Insertar(comment.Contenido, comment.User)
}

func (m *Lista) Responder(nombre string, respuesta Tipos.Respuestas) { //Devuelve una tienda en especifico
	aux := m.inicio
	for nombre != aux.dato.Nombre {
		aux = aux.siguiente
	}
	aux.dato.Comentarios.Respuesta(respuesta)
}

func (m *Lista) Get(nombre string) Tipos.Tienda { //Devuelve una tienda en especifico
	aux := m.inicio
	for nombre != aux.dato.Nombre {
		aux = aux.siguiente
	}
	return aux.dato
}

func (m *Lista) Set_Inventario(tmp Tipos.Tienda) { //Devuelve un dato de la lista
	aux := m.inicio
	for tmp.Nombre != aux.dato.Nombre {
		aux = aux.siguiente
	}
	aux.dato = tmp
}
