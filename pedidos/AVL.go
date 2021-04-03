package pedidos

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	//"github.com/Fernando-MGS/TEST/AV"
	"github.com/Fernando-MGS/TEST/Tipos"
)

type Year struct {
	Año  int
	List Lista
}

var años []string

type Producto struct {
	Nombre      string  `json:"Nombre,omitempty"`
	Codigo      int     `json:"Codigo,omitempty"`
	Descripcion string  `json:"Descripcion,omitempty"`
	Precio      float64 `json:"Precio,omitempty"`
	Cantidad    int     `json:"Cantidad,omitempty"`
	Imagen      string  `json:"Imagen,omitempty"`
	Cant        []int
}

type Años struct {
	Datos  []Meses
	Indice int
	Large  int
}

var mes []Meses

type Meses struct {
	Año int
	Mes []string
}

var Listado []Producto

type nodo_m struct {
	indice   Year
	altura   int
	izq, der *nodo_m
}

var nodes string
var graph string
var pointers string

func newnodo_m(indice Year) *nodo_m {
	return &nodo_m{indice, 0, nil, nil}
}

type AVL struct {
	raiz *nodo_m
}

func NewAVL() *AVL {
	return &AVL{nil}
}

func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func altura(temp *nodo_m) int {
	if temp != nil {
		return temp.altura
	}
	return -1
}

func rotacionIzquierda(temp **nodo_m) {
	aux := (*temp).izq
	(*temp).izq = aux.der
	aux.der = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).izq), (*temp).altura) + 1
	*temp = aux
}

func rotacionDerecha(temp **nodo_m) {
	aux := (*temp).der
	(*temp).der = aux.izq
	aux.izq = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).der), (*temp).altura) + 1
	*temp = aux
}

func rotacionDobleIzquierda(temp **nodo_m) {
	rotacionDerecha(&(*temp).izq)
	rotacionIzquierda(temp)
}

func rotacionDobleDerecha(temp **nodo_m) {
	rotacionIzquierda(&(*temp).der)
	rotacionDerecha(temp)
}

func prob_exist(indice int, root **nodo_m) int { //0 no existe, 1 si existe
	if (*root) == nil {
		return 0
	}
	if indice < (*root).indice.Año {
		prob_exist(indice, &(*root).izq)
	} else if indice > (*root).indice.Año {
		prob_exist(indice, &(*root).der)
	} else {
		return 1
	}
	return 2
}

func (avl *AVL) Insertar(indice Year, l_prod []Tipos.Producto, depto, mes, dia int) {
	fmt.Println("Llego a insertar")
	/*if prob_exist(indice.Año, &avl.raiz) == 0 || prob_exist(indice.Año, &avl.raiz) == 2 {
		fmt.Println()
		fmt.Println("---------")
		fmt.Println("El año ", indice.Año, " no existe")
		insert(indice, &avl.raiz)
	} else {
		fmt.Println("El año ", indice.Año, " ya existe")
		agregar_toList(&avl.raiz, l_prod, indice.Año, depto, mes, dia)
	}*/
	_insert(indice, &avl.raiz, l_prod, depto, mes, dia)
}

func _insert(indice Year, root **nodo_m, l_prod []Tipos.Producto, depto, mes, dia int) {
	if (*root) == nil {
		fmt.Println("Insertando como root ", indice.Año)
		*root = newnodo_m(indice)
		return
	}
	if indice.Año < (*root).indice.Año {
		_insert(indice, &(*root).izq, l_prod, depto, mes, dia)
		fmt.Println("Insertando izq ", indice.Año)
		if (altura((*root).izq) - altura((*root).der)) == -2 {
			if indice.Año < (*root).izq.indice.Año {
				rotacionIzquierda(root)
			} else {
				rotacionDobleIzquierda(root)
			}
		}
	} else if indice.Año > (*root).indice.Año {
		_insert(indice, &(*root).der, l_prod, depto, mes, dia)
		fmt.Println("Insertando der ", indice.Año)
		if (altura((*root).der) - altura((*root).izq)) == 2 {
			if indice.Año > (*root).der.indice.Año {
				rotacionDerecha(root)
			} else {
				rotacionDobleDerecha(root)
			}
		}
	} else {
		(*root).indice.List.Insercion(l_prod, depto, mes, dia)
	}

	(*root).altura = max(altura((*root).izq), altura((*root).der)) + 1
}
func insert(indice Year, root **nodo_m) {
	if (*root) == nil {
		fmt.Println("Insertando como root ", indice.Año)
		*root = newnodo_m(indice)
		return
	}
	if indice.Año < (*root).indice.Año {
		insert(indice, &(*root).izq)
		fmt.Println("Insertando izq ", indice.Año)
		if (altura((*root).izq) - altura((*root).der)) == -2 {
			if indice.Año < (*root).izq.indice.Año {
				rotacionIzquierda(root)
			} else {
				rotacionDobleIzquierda(root)
			}
		}
	} else if indice.Año > (*root).indice.Año {
		insert(indice, &(*root).der)
		fmt.Println("Insertando der ", indice.Año)
		if (altura((*root).der) - altura((*root).izq)) == 2 {
			if indice.Año > (*root).der.indice.Año {
				rotacionDerecha(root)
			} else {
				rotacionDobleDerecha(root)
			}
		}
	} else {

		fmt.Println("Solo para ver si llego")
	}

	(*root).altura = max(altura((*root).izq), altura((*root).der)) + 1
}

func agregar_toList(root **nodo_m, l_prod []Tipos.Producto, año, depto, mes, dia int) { //0 no existe, 1 si existe
	if año < (*root).indice.Año {
		agregar_toList(&(*root).izq, l_prod, año, depto, mes, dia)
	} else if año > (*root).indice.Año {
		agregar_toList(&(*root).der, l_prod, año, depto, mes, dia)
	} else {
		(*root).indice.List.Insercion(l_prod, depto, mes, dia)
	}
	return
}

func (avl *AVL) Print() {
	inOrden(avl.raiz)
}

func (avl *AVL) prob_nil() int {
	if avl.raiz == nil {
		return 0
	}
	return 1
}

/*func (avl *AVL) Buscar(indice int) int {
	return _prob_exist(indice, &avl.raiz)
}*/

/*func _prob_exist(indice int, root **nodo_m) int {
	if (*root) == nil {
		return 0
	}
	if indice < (*root).indice.Codigo {
		prob_exist(indice, &(*root).izq)
	} else if indice > (*root).indice.Codigo {
		prob_exist(indice, &(*root).der)
	} else {
		return 1
	}
	return 0
}*/

func graph_inOrden(temp *nodo_m) {
	if temp != nil {
		graph_inOrden(temp.izq)
		nodes += "Node" + strconv.Itoa(temp.indice.Año) + "[label=\"" + strconv.Itoa(temp.indice.Año) + "\"]\n"
		if temp.izq != nil {
			pointers += "Node" + strconv.Itoa(temp.indice.Año) + "->Node" + strconv.Itoa(temp.izq.indice.Año) + ";\n"
		}
		if temp.der != nil {
			pointers += "Node" + strconv.Itoa(temp.indice.Año) + "->Node" + strconv.Itoa(temp.der.indice.Año) + ";\n"
		}
		graph_inOrden(temp.der)
	}
}

func (avl *AVL) Grap() {
	graph = "digraph List {\n"
	graph += "rankdir=TB;"
	graph += "node [shape = circle, color=greenyellow , style=filled, fillcolor=darkgreen];"
	graph_inOrden(avl.raiz)
	fmt.Println("El root es ", avl.raiz.indice.Año)
	graph += nodes + "\n" + pointers
	graph += "\n}"
	data := []byte(graph)                            //Almacenar el codigo en el formato adecuado
	err := ioutil.WriteFile("graph.dot", data, 0644) //Crear el archivo .dot necesario para la imagen
	if err != nil {
		log.Fatal(err)
	}
	//Creación de la imagen
	path, _ := exec.LookPath("dot") //Para que funcione bien solo asegurate de tener todas las herramientas para
	// Graphviz en tu compu, si no descargalas osea el Graphviz
	cmd, _ := exec.Command(path, "-Tpng", "graph.dot").Output()  //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                            //Se mantiene igual
	ioutil.WriteFile("Pedidos-años.png", cmd, os.FileMode(mode)) //Creacion de la imagen
	pointers = ""
	nodes = ""
	graph = ""
}

func inOrden(temp *nodo_m) {
	if temp != nil {
		inOrden(temp.izq)
		fmt.Println("Index: ", temp.indice)
		fmt.Print("Lista: ")
		temp.indice.List.Imprimir()
		inOrden(temp.der)
	}
}

func (avl *AVL) Dev_year(ind string) Años {
	in_Orden(avl.raiz)
	var year Años
	year.Datos = mes
	//year.Indice = 0
	year.Large = len(mes)
	var niu []Meses
	mes = niu
	return year
}

func in_Orden(t *nodo_m) {
	if t != nil {
		in_Orden(t.izq)
		var month Meses
		month.Año = t.indice.Año
		month.Mes = t.indice.List.Dev_array()
		mes = append(mes, month)
		in_Orden(t.der)
	}
}

func (avl *AVL) Graph_lista(año, mes int) {
	temp := avl.raiz
	for temp != nil {
		if temp.indice.Año == año {
			y := strconv.Itoa(año)
			temp.indice.List.GraphData(y)
			break
		}
	}
}
