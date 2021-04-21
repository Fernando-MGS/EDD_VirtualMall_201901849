package lista

import (
	//	"fmt"

	"github.com/Fernando-MGS/TEST/Tipos"
)

/*type Producto struct {
	Codigo   int
	Cantidad int
}*/
type Node struct {
	data Tipos.Producto
	next *Node
}

type List struct {
	head     *Node
	tamaño   int
	precio   float64
	Cantidad int
}

var carrito []Tipos.Producto

func (l *List) Tamaño() int {
	return l.tamaño
}

func (l *List) Precio() float64 {
	return float64(l.precio)
}

func (l *List) Add(data Tipos.Producto) {
	if l.prob_exist(data.Codigo) == 0 {
		if l.head == nil {
			tmp := &Node{data: data, next: l.head}
			l.tamaño = l.tamaño + data.Cantidad
			l.precio = l.precio + float64(data.Cantidad)*data.Precio
			l.head = tmp
			l.Cantidad = l.Cantidad + 1
		} else {
			tmp := l.head
			for tmp.next != nil {
				tmp = tmp.next
			}
			tmp.next = &Node{data: data, next: nil}
			l.tamaño = l.tamaño + data.Cantidad
			l.precio = l.precio + float64(data.Cantidad)*data.Precio
			l.Cantidad = l.Cantidad + 1
		}
	} else {
		l.add_prod(data.Codigo, data.Cantidad)
		l.precio = l.precio + float64(data.Cantidad)*data.Precio
	}
}

func (l *List) add_prod(codigo int, cantidad int) {
	tmp := l.head
	for tmp.next != nil {
		if tmp.data.Codigo == codigo {
			tmp.data.Cantidad = tmp.data.Cantidad + cantidad
			l.tamaño = l.tamaño + cantidad
		}
		tmp = tmp.next
	}
	if tmp.data.Codigo == codigo {
		tmp.data.Cantidad = tmp.data.Cantidad + cantidad
		l.tamaño = l.tamaño + cantidad
	}
}

func (l *List) Putoff_car(prod Tipos.Producto, cantidad int) {
	tmp := l.head
	for tmp.next != nil {
		if tmp.data.Codigo == prod.Codigo {
			tmp.data.Cantidad = tmp.data.Cantidad - cantidad
			if tmp.data.Cantidad == 0 {
				//				fmt.Println("Ya no hay :v")
				l.Cantidad = l.Cantidad - 1
				l.delete(tmp.data.Codigo)
				l.tamaño = l.tamaño - cantidad
				l.precio = l.precio - float64(cantidad)*prod.Precio
			}
		}
		tmp = tmp.next
	}
	if tmp.data.Codigo == prod.Codigo {
		tmp.data.Cantidad = tmp.data.Cantidad - cantidad
		if tmp.data.Cantidad == 0 {
			l.delete(tmp.data.Codigo)
			l.tamaño = l.tamaño - cantidad
			l.precio = l.precio - float64(cantidad)*prod.Precio
		}
	}
}

func (l *List) prob_exist(codigo int) int { // revisa si un producto ya fue ingresado
	tmp := l.head
	conf := 0
	if l.head == nil {
		return 0
	}
	for tmp.next != nil {
		if tmp.data.Codigo == codigo {
			conf = 1
		}
		tmp = tmp.next
	}
	if tmp.data.Codigo == codigo {
		conf = 1
	}
	return conf
}

func (l *List) delete(data int) {
	tmp := l.head
	if l.head != nil {
		for tmp.next != nil {
			if l.head.data.Codigo == data {
				if l.head.next != nil {
					l.head = l.head.next
				}
			} else if tmp.next.data.Codigo == data {
				tmp.next = tmp.next.next
			} else {
				tmp = tmp.next
			}
		}
	}
}

func (l *List) GetProducts() []Tipos.Producto {
	var carr []Tipos.Producto
	tmp := l.head
	for tmp.next != nil {
		carr = append(carr, tmp.data)
		tmp = tmp.next
	}
	carr = append(carr, tmp.data)
	carrito = carr
	return carrito
}

func (l *List) GetItem(index int) Tipos.Producto {
	sum := 1
	temp := l.head
	for sum <= index {
		temp = temp.next
		sum++
	}
	return temp.data
}

func (l *List) Show() {
	tmp := l.head
	for tmp != nil {
		//fmt.Print(tmp.data, " ")
		tmp = tmp.next
	}
}
