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
var cluster int

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
	_insert(indice, &avl.raiz, l_prod, depto, mes, dia)
}

func _insert(indice Year, root **nodo_m, l_prod []Tipos.Producto, depto, mes, dia int) {
	if (*root) == nil {
		*root = newnodo_m(indice)
		return
	}
	if indice.Año < (*root).indice.Año {
		_insert(indice, &(*root).izq, l_prod, depto, mes, dia)
		if (altura((*root).izq) - altura((*root).der)) == -2 {
			if indice.Año < (*root).izq.indice.Año {
				rotacionIzquierda(root)
			} else {
				rotacionDobleIzquierda(root)
			}
		}
	} else if indice.Año > (*root).indice.Año {
		_insert(indice, &(*root).der, l_prod, depto, mes, dia)
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
	//fmt.Println("El root es ", avl.raiz.indice.Año)
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
		//fmt.Println("Index: ", temp.indice)
		//fmt.Print("Lista: ")
		temp.indice.List.Imprimir()
		inOrden(temp.der)
	}
}
func (avl *AVL) Dev(mes string, año int, nombres []string) {
	b := a(mes, año, avl.raiz)
	//fmt.Println("Va mo a dev", len(b), "-")
	//fmt.Println(len(b[0].Producto))
	for i := 0; i < len(b); i++ {
		c := b[i].Producto
		//fmt.Println(b[i].X, "....", b[i].Y)
		for j := 0; j < len(c); j++ {
			//fmt.Print(c[j].Codigo, "--", c[j].ID, "--", c[j].Departamento)
		}
		/*fmt.Println()
		fmt.Println("______________________")*/
	}
	//fmt.Println("//////////////////////////")

	t := strconv.Itoa(año)
	graph_matriz(mes, t, ordenar_matriz(b, nombres), nombres)
}

func (avl *AVL) _Dev(mes string, año int, nombres []string) {
	b := a(mes, año, avl.raiz)
	//t := strconv.Itoa(año)
	e := ordenar_matriz(b, nombres)
	fmt.Println(e.Cabeceras_x)
	fmt.Println("Cabeceras_y", e.Cabeceras_y)
	fmt.Println(e.Cabeceras_x)
}

func a(mes string, año int, t *nodo_m) []Tipos.Matrices {
	if t != nil {
		if t.indice.Año == año {
			//fmt.Println("Encontrado")
			return t.indice.List._GetItem(dev_mes(mes))
		}
	}
	return nil
}

func dev_mes(mes string) int {
	Meses := []string{"ENERO", "FEBRERO", "MARZO", "ABRIL", "MAYO", "JUNIO", "JULIO", "AGOSTO", "SEPTIEMBRE", "OCTUBRE", "NOVIEMBRE", "DICIEMBRE"}
	for i := 0; i < len(Meses); i++ {
		if mes == Meses[i] {
			return i + 1
		}
	}
	return 0
}
func (avl *AVL) dev_matriz() {
	in_Orden(avl.raiz)
	var year Años
	year.Datos = mes
	//year.Indice = 0
	year.Large = len(mes)
	var niu []Meses
	mes = niu
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

func dev_nombre_depto(index int, nombres []string) string {
	return nombres[index]
}

func ordenar_matriz(e []Tipos.Matrices, nombres []string) Tipos.Pedidos_mes {

	var orden Tipos.Pedidos_mes
	for i := 0; i < len(e); i++ {
		orden.Cabeceras_x = append(orden.Cabeceras_x, e[i].X)
	}

	for i := 0; i < len(e); i++ {
		find := 0
		for j := 0; j < len(orden.Cabeceras_y); j++ {
			if e[i].Y == orden.Cabeceras_y[j] {
				find = 1
			}
		}
		if find == 0 {
			orden.Cabeceras_y = append(orden.Cabeceras_y, e[i].Y)
		}
	}

	//juntar los productos por departamento y luego compararlos
	for i := 0; i < len(e); i++ {
		conf := -1
		var a Tipos.Cabecera_y
		a.Dia = e[i].X
		a.Productos = e[i].Producto
		for j := 0; j < len(orden.Cabeceras_y); j++ {
			//fmt.Println(orden.Cabeceras_y[j], "-", dev_ind_dep(e[i].Producto[0].Departamento, nombres))
			if orden.Cabeceras_y[j] == dev_ind_dep(e[i].Producto[0].Departamento, nombres) {
				a.Conf_exis = append(a.Conf_exis, j)
				d := conf_dia(a.Dia, orden)
				conf = d
				if d >= 0 {
					orden.Pedidos_D[d].Conf_exis = append(orden.Pedidos_D[d].Conf_exis, j)
					for k := 0; k < len(e[i].Producto); k++ {
						orden.Pedidos_D[d].Productos = append(orden.Pedidos_D[d].Productos, e[i].Producto[k])
					}
				}
			}
		}
		if conf == -1 {
			orden.Pedidos_D = append(orden.Pedidos_D, a)
		}
	}

	for j := 0; j < len(orden.Pedidos_D); j++ {
		/*fmt.Print("El dia es", orden.Pedidos_D[j].Dia, "-", len(orden.Pedidos_D[j].Productos))
		fmt.Print(orden.Pedidos_D[j].Conf_exis)
		fmt.Println("===================")*/
	}
	//fmt.Println(orden.Pedidos_D)
	return orden
}

func conf_dia(dia int, e Tipos.Pedidos_mes) int {
	conf := -1
	for i := 0; i < len(e.Pedidos_D); i++ {
		if dia == e.Pedidos_D[i].Dia {
			conf = i
			//fmt.Println("El día que coincide es", conf)
		}
	}
	return conf
}

func dev_ind_dep(name string, nombres []string) int {
	cont := 0
	for i := 0; i < len(nombres); i++ {
		if nombres[i] == name {

			cont = i
		}
	}
	return cont + 1
}

func graph_matriz(mes string, año string, month Tipos.Pedidos_mes, nombres []string) {
	fmt.Println(month.Cabeceras_y)
	fmt.Println(nombres)
	graph = "digraph G{\n"
	cont_n := 0
	cluster = 1
	/*graph += "rankdir=TB;"
	graph += "node [shape = circle, color=greenyellow , style=filled, fillcolor=darkgreen];"*/
	nodes += "subgraph cluster" + strconv.Itoa(cluster) + "{\n"
	nodes += "node [style=filled,color =lightgrey,shape=Mrecord];\n"
	nodes += "style=filled;\n"
	nodes += "color=white;\n"
	nodes += "node" + strconv.Itoa(cont_n) + "[label=\"\" style=filled, color=white]\n"
	nodes += "node0 -> node1 [arrowhead=none, color=white]\n"
	cont_n = 1
	for i := 0; i < len(month.Cabeceras_y); i++ {
		//name:=dev_nombre_depto(month.Cabeceras_y[i])
		nodes += "node" + strconv.Itoa(cont_n) + "[label=\"" + nombres[month.Cabeceras_y[i]-1] + "\" style=filled]\n"
		if i <= len(month.Cabeceras_y)-2 {
			pointers += "node" + strconv.Itoa(cont_n) + " ->"
		} else if i == len(month.Cabeceras_y)-1 {
			pointers += "node" + strconv.Itoa(cont_n) + "[arrowhead=none, color=white]; \n"
			pointers += mes + " [shape=Mdiamond,color=white];"
			pointers += mes + " -> node0 [arrowhead=none, color=white]\n"
		}
		cont_n++
	}
	for i := 0; i < len(month.Pedidos_D); i++ {
		nodes += "subgraph cluster" + strconv.Itoa(cluster) + "{\n"
		nodes += "node [style=filled,color=lightgrey,shape=Mrecord];\n"
		nodes += "style=filled\n"
		nodes += "color=white\n"
		index := 0
		for j := 0; j < len(month.Cabeceras_y); j++ {
			if j == 0 {
				nodes += "node" + strconv.Itoa(cont_n) + "[label=\"" + strconv.Itoa(month.Pedidos_D[i].Dia) + "\"]\n"
				pointers += mes + " -> node" + strconv.Itoa(cont_n) + "[arrowhead=none, color=white]\n"
				pointers += "node" + strconv.Itoa(cont_n) + " -> "
				cont_n++
			}
			if month.Pedidos_D[i].Conf_exis[index] == j {
				nodes += "node" + strconv.Itoa(cont_n) + "[label=\"\"]\n"
				if j == len(month.Cabeceras_y)-1 {
					pointers += "node" + strconv.Itoa(cont_n) + "[arrowhead=none, color=white];\n"
				} else {
					pointers += "node" + strconv.Itoa(cont_n) + " ->"
				}
				if index < len(month.Pedidos_D[i].Conf_exis)-1 {
					index++
				}
			} else {
				nodes += "node" + strconv.Itoa(cont_n) + "[label=\"\" color=white]\n"
				if j == len(month.Cabeceras_y)-1 {
					pointers += "node" + strconv.Itoa(cont_n) + "[arrowhead=none, color=white];\n"
				} else {
					pointers += "node" + strconv.Itoa(cont_n) + " ->"
				}
			}
			cont_n++
		}
		nodes += "}\n"
		cluster++
	}
	nodes += "}"
	graph += nodes + "\n" + pointers
	graph += "\n}"
	data := []byte(graph)                            //Almacenar el codigo en el formato adecuado
	err := ioutil.WriteFile("graph.dot", data, 0644) //Crear el archivo .dot necesario para la imagen
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(graph)
	//Creación de la imagen
	path, _ := exec.LookPath("dot") //Para que funcione bien solo asegurate de tener todas las herramientas para
	// Graphviz en tu compu, si no descargalas osea el Graphviz
	cmd, _ := exec.Command(path, "-Tpng", "graph.dot").Output()            //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                                      //Se mantiene igual
	ioutil.WriteFile("Pedidos"+mes+"-"+año+".png", cmd, os.FileMode(mode)) //Creacion de la imagen
	pointers = ""
	nodes = ""
	graph = ""
}
