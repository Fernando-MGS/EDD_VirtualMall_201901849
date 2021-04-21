package pedidos

import (
	"github.com/Fernando-MGS/TEST/Tipos"
	//"github.com/Fernando-MGS/TEST/AV"
)

type Product struct {
	Codigo      int    //Se cambio a primer letra mayuscula para poder acceder
	Nombre      string //Se cambio a primer letra mayuscula para poder acceder
	Descripcion string //Se cambio a primer letra mayuscula para poder acceder
	Cantidad    int
}

type nodo struct {
	//Estos atributos son especificos para la matriz
	x, y                              int              //Saber en que cabecera estoy
	producto                          []Tipos.Producto //tipo de objeto
	izquierda, derecha, arriba, abajo *nodo            //nodos con los que nos desplazamos dentro de la matriz
	//Estos atributos son especificos para la lista
	header              int   //tipo interno de la cabecera
	siguiente, anterior *nodo // nodos con los que nos vamos a desplazar dentro de las listas
}

type lista struct {
	first, last *nodo
}

type matriz struct {
	lst_h, lst_v *lista
}

func nodoMatriz(x int, y int, producto Tipos.Producto) *nodo {
	var array []Tipos.Producto
	array = append(array, producto)
	return &nodo{x, y, array, nil, nil, nil, nil, 0, nil, nil}

}

func nodoLista(header int) *nodo {
	return &nodo{0, 0, nil, nil, nil, nil, nil, header, nil, nil}
}

func newLista() *lista {
	return &lista{nil, nil}
}

//Se cambio a primer letra mayuscula para poder acceder
func NewMatriz() *matriz {
	return &matriz{newLista(), newLista()}
}

func (n *nodo) headerX() int { return n.x }
func (n *nodo) headerY() int { return n.y }

/*func (n *nodo) toString() string {
	return "Nombre: " + n.producto.Nombre + "\nDescripcion: " + n.producto.Descripcion
}*/

func (l *lista) ordenar(nuevo *nodo) {
	//fmt.Println("A ordenar: ", nuevo)
	aux := l.first
	for aux != nil {
		if nuevo.header > aux.header {
			aux = aux.siguiente
		} else {
			if aux == l.first {
				nuevo.siguiente = aux
				aux.anterior = nuevo
				l.first = nuevo
			} else {
				nuevo.anterior = aux.anterior
				aux.anterior.siguiente = nuevo
				nuevo.siguiente = aux
				aux.anterior = nuevo
			}
			return
		}
	}
	l.last.siguiente = nuevo
	nuevo.anterior = l.last
	l.last = nuevo
}

func (l *lista) insert(header int) {
	//fmt.Println("En insert, ", header)
	nuevo := nodoLista(header)
	if l.first == nil {
		//fmt.Println("first is nil ", header)
		l.first = nuevo
		l.last = nuevo
	} else {
		l.ordenar(nuevo)
	}
}

func (l *lista) search(header int) *nodo {
	//fmt.Println("Buscando a header :", header)
	temp := l.first
	for temp != nil {
		//fmt.Println(temp.header, "vs", header)
		if temp.header == header {
			return temp
		}
		temp = temp.siguiente
	}
	return nil
}

func (l *lista) Print() {
	temp := l.first
	for temp != nil {
		t := temp
		//fmt.Print("Cabecera:", temp.y, temp.header, ":   ")

		for t != nil {
			//	fmt.Print(t.producto, "  ", t.x, "-", t.y)
			t = t.derecha
		}
		//fmt.Println()
		temp = temp.siguiente
	}
}

func (l *lista) Dev() []Tipos.Matrices {
	var arr []Tipos.Matrices

	temp := l.first
	for temp != nil {
		t := temp
		//	fmt.Print("Cabecera:", temp.y, temp.header, ":   ")
		for t != nil {
			var matriz Tipos.Matrices
			matriz.Producto = t.producto
			arr = append(arr, matriz)
			//fmt.Print(t.producto, "  ", t.x, "-", t.y)
			t = t.derecha
		}
		//fmt.Println()
		temp = temp.siguiente
	}
	return arr
}

func (l *lista) buscar(x, y int, producto Tipos.Producto) int {
	temp := l.first
	find := 0
	//fmt.Println(producto)
	for temp != nil {
		t := temp

		for t != nil {
			if t.x == x && t.y == y {
				//fmt.Println("Producto ya existe")
				t.producto = append(t.producto, producto)
				find = 1
			}

			t = t.derecha
		}
		//fmt.Println("-")
		temp = temp.siguiente
	}
	temp = l.first
	for temp != nil {
		t := temp
		for t != nil {
			if t.x == x && t.y == y {
				//fmt.Println("Producto ya existe")
				t.producto = append(t.producto, producto)
				find = 1
			}
			t = t.abajo
		}
		//fmt.Println()
		temp = temp.siguiente
	}

	//fmt.Println("Salió del For de buscar matrix")
	return find
}

func (l *lista) _rec_head(t *nodo) {
	//	fmt.Println("----")
	for t.abajo != nil {
		//fmt.Println(t.x)
	}
	for t.derecha != nil {
		//fmt.Println(t.y)
	}
}

func (l *lista) _buscar() []Tipos.Matrices {
	var b []Tipos.Matrices

	temp := l.first
	//find := 0
	for temp != nil {
		t := temp
		for t != nil {
			if t.x != 0 {
				var l Tipos.Matrices
				l.Producto = t.producto
				l.X = t.x
				l.Y = t.y
				b = append(b, l)
			}
			t = t.derecha
		}
		//fmt.Println("-")
		temp = temp.siguiente
	}
	temp = l.first
	for temp != nil {
		t := temp
		for t != nil {
			if t.y != 0 {
				var l Tipos.Matrices
				l.Producto = t.producto
				l.X = t.x
				l.Y = t.y
				b = append(b, l)
			}
			t = t.abajo
		}
		//fmt.Println()
		temp = temp.siguiente
	}
	return b
}

func (l *lista) Print_h() {
	//fmt.Println("Print de las columnas")
	temp := l.first
	for temp != nil {
		t := temp
		//fmt.Print("Cabecera:", temp.x, temp.header, ":   ")
		for t != nil {
			//fmt.Print(t.x, "-", t.y, "v: ")
			//	fmt.Println(t.producto)
			t = t.abajo

		}
		//fmt.Println("________________")
		temp = temp.siguiente
	}
}

func (m *matriz) Insert(producto Tipos.Producto, x int, y int) {
	/*fmt.Println("________________________________")
	fmt.Println("Preparandose para insertar ", producto, " en", x, ",", y)
	fmt.Println("bUSCANDO H")*/
	h := m.lst_h.search(x)

	//fmt.Println("h es ", h.x, "--", h.y, "--", h.producto[0].Nombre)
	//fmt.Println("bUSCANDO V")
	v := m.lst_v.search(y)
	/*find := m.lst_v.buscar(x, y, producto)
	find_1 := m.lst_h.buscar(x, y, producto)*/
	//fmt.Println("Entro al find 0")
	if h == nil && v == nil {
		//fmt.Println("h y  v nill")
		m.noExisten(producto, x, y)
	} else if h == nil && v != nil {
		//	fmt.Println("h nil, v no")
		m.existeVertical(producto, x, y)
	} else if h != nil && v == nil {
		//	fmt.Println("h, v nill")
		m.existeHorizontal(producto, x, y)
	} else {
		m.existen(producto, x, y)
	}
	//fmt.Println("Fin del buscar del insert")
}

func (m *matriz) noExisten(producto Tipos.Producto, x int, y int) {
	//fmt.Println("no Existen")
	m.lst_h.insert(x) //insertamos en la lista que emula la cabecera horizontal
	m.lst_v.insert(y) //insertamos en la lista que emula la cabecera vertical

	h := m.lst_h.search(x) //vamos a buscar el n7odo que acabos de insertar para poder enlazarlo
	v := m.lst_v.search(y) //vamos a buscar el nodo que acabos de insertar para poder enlazarlo

	nuevo := nodoMatriz(x, y, producto) //creamos nuevo nodo tipo matriz

	h.abajo = nuevo  //enlazamos el nodo horizontal hacia abajo
	nuevo.arriba = h //enlazmos el nuevo nodo hacia arriba

	v.derecha = nuevo   //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierda = v //enlazamos el nuevo nodo hacia la izquierda
}

func (m *matriz) existeVertical(producto Tipos.Producto, x int, y int) {
	//fmt.Println("existe Vertical")
	m.lst_h.insert(x) //insertamos en la lista que emula la cabecera horizontal

	h := m.lst_h.search(x) //vamos a buscar el n7odo que acabos de insertar para poder enlazarlo
	v := m.lst_v.search(y) //vamos a buscar el nodo que acabos de insertar para poder enlazarlo
	//	fmt.Println("h.x es ", h.x, "  vy es:", v.y)
	nuevo := nodoMatriz(x, y, producto) //creamos nuevo nodo tipo matriz

	h.abajo = nuevo  //enlazamos el nodo horizontal hacia abajo
	nuevo.arriba = h //enlazmos el nuevo nodo hacia arriba

	v.derecha = nuevo   //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierda = v //enlazamos el nuevo nodo hacia la izquierda
}

func (m *matriz) existeHorizontal(producto Tipos.Producto, x int, y int) {
	//fmt.Println("existe Horizontal")
	m.lst_v.insert(y) //insertamos en la lista que emula la cabecera vertical

	h := m.lst_h.search(x) //vamos a buscar el n7odo que acabos de insertar para poder enlazarlo
	v := m.lst_v.search(y) //vamos a buscar el nodo que acabos de insertar para poder enlazarlo
	//fmt.Println("h.x es ", h.x, "  vy es:", v.y)
	nuevo := nodoMatriz(x, y, producto) //creamos nuevo nodo tipo matriz

	h.abajo = nuevo  //enlazamos el nodo horizontal hacia abajo
	nuevo.arriba = h //enlazmos el nuevo nodo hacia arriba

	v.derecha = nuevo   //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierda = v //enlazamos el nuevo nodo hacia la izquierda
}

func (m *matriz) existen(producto Tipos.Producto, x int, y int) {
	//fmt.Println("Si Existen")
	h := m.lst_h.search(x) //vamos a buscar el n7odo que acabos de insertar para poder enlazarlo
	v := m.lst_v.search(y) //vamos a buscar el nodo que acabos de insertar para poder enlazarlo
	//	fmt.Println("h.x es ", h.x, "  vy es:", v.y)
	nuevo := nodoMatriz(x, y, producto) //creamos nuevo nodo tipo matriz
	h.abajo = nuevo                     //enlazamos el nodo horizontal hacia abajo
	nuevo.arriba = h                    //enlazmos el nuevo nodo hacia arriba
	v.derecha = nuevo                   //enlazamos el nodo vertical hacia la derecha
	nuevo.izquierda = v
	nuevo.producto = append(nuevo.producto, producto) //enlazamos el nuevo nodo hacia la izquierda
}

/*func main() {
	m := NewMatriz()
	p1 := Producto{12, "Test", "Soy descripcion", 1}
	m.Insert(p1, 1, 1)
	p2 := Producto{13, "Test2", "Soy descripcion-1", 1}
	m.Insert(p2, 2, 2)
	p3 := Producto{13, "Test3", "Soy descripcion-1", 2}
	m.Insert(p3, 2, 2)
	m.lst_h.print_h()
	fmt.Println("--------------")
	m.lst_v.print()
}*/
