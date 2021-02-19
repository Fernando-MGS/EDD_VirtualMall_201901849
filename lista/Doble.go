package lista

import (
	"fmt"
)

//lugar donde se almacena
type nodo struct {
	anterior  *nodo
	siguiente *nodo
	dato      int
}

type Lista struct {
	inicio *nodo
	ultimo *nodo
	tam    int
}


func NewLista() *Lista {//crea una lista
	return &Lista{nil, nil, 0} //el inicio es null, final null, y el tamaño es 0
}


func (m *Lista) Insertar(valor int) {//insertar un nodo
	nuevo := &nodo{nil, nil, valor}
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


func (m *Lista) Imprimir() {//IMPRIMIR
	aux := m.inicio
	for aux != nil {
		fmt.Println(aux.dato)
		aux = aux.siguiente
	}
	fmt.Print(m.tam)
}


func (m *Lista) Tamaño() {//IMPRIMIR TAMAÑO DE LISTA
	fmt.Print(m.tam)
}

func (m *Lista) GetItem(index int) {//Devuelve un dato de la lista
	ind := 1
	aux := m.inicio
	if index <= m.tam {
		for ind < index {
			aux = aux.siguiente
			ind++
		}
		fmt.Println(aux.dato)
	} else {
		fmt.Println("Dato inexistente")
	}

}
