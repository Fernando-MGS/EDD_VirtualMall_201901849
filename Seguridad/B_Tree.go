package Seguridad

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/Fernando-MGS/TEST/Tipos"
)

var conf_b int
var nodes string
var graph string
var pointers string

type Pagina struct {
	altura int
	users  []*Nodo_B
}

type Nodo_B struct {
	altura    int
	izquierda *Pagina
	derecha   *Pagina
	User      Tipos.Usuario
}

func newnodo_b(indice Tipos.Usuario, altura int) *Nodo_B {
	return &Nodo_B{altura, nil, nil, indice}
}

type B_Tree struct {
	raiz *Pagina
}

func New_B() *B_Tree {
	return &B_Tree{nil}
}

//INSERCION
func (a *B_Tree) Insertar(user Tipos.Usuario) {
	user_b := newnodo_b(user, 0)
	//fmt.Println("Vamo a insertar")
	if a.raiz == nil {
		var n []*Nodo_B
		n = append(n, user_b)
		a.raiz = &Pagina{0, n}
	} else {
		Busqueda_Inser(a, user_b, &a.raiz, a.raiz, a.raiz.altura)
		rompimiento(a.raiz, &a.raiz)
	}
}

func rompimiento(actual *Pagina, anterior **Pagina) {
	if len(actual.users) == 5 {
		var n []*Nodo_B
		n = append(n, actual.users[0])
		n = append(n, actual.users[1])
		var m []*Nodo_B
		m = append(m, actual.users[3])
		m = append(m, actual.users[4])
		prim := &Pagina{actual.altura + 1, n}
		sec := &Pagina{actual.altura + 1, m}
		var l []*Nodo_B
		l = append(l, actual.users[2])
		uniq := &Pagina{actual.altura, l}
		uniq.users[0].izquierda = prim
		uniq.users[0].derecha = sec
		if actual == (*anterior) {
			(*anterior) = uniq
			fmt.Println("Reset Rompi")
		} else {
			ind := 0
			fmt.Println("Reset alae2")
			for ind < len(uniq.users) {
				(*anterior).users = append((*anterior).users, uniq.users[ind])
				slice := ordenar_slice((*anterior).users)
				(*anterior).users = slice
				ind++
			}
			b := len((*anterior).users) - 1
			for i := 0; i < len((*anterior).users); i++ {
				if i != b {
					(*anterior).users[i].derecha = nil
				}
			}
			impr((*anterior).users)
		}
	} else {
		for i := 0; i < len(actual.users); i++ {
			if actual.users[i].izquierda != nil {
				rompimiento(actual.users[i].izquierda, &actual)
			}
			if actual.users[i].derecha != nil {
				rompimiento(actual.users[i].derecha, &actual)
			}
		}
	}
}
func Busqueda_Inser(arbol *B_Tree, user *Nodo_B, anterior **Pagina, pag *Pagina, alt int) {
	conf := 0
	cont := 0
	index := 0
	for cont < len(pag.users) {
		//fmt.Println(len(pag.users))
		j, err := strconv.Atoi(pag.users[cont].User.DPI)
		i, err := strconv.Atoi(user.User.DPI)
		inutil(err)
		if i < j && conf != 2 {
			conf = 2
			index = cont
		} else {
			conf = 3
			index = cont
		}
		cont++
	}
	impr(pag.users)
	fmt.Println("--------")
	if conf == 2 {
		if pag.users[index].izquierda == nil {
			user.altura = alt
			pag.users = append(pag.users, user)
			slice := ordenar_slice(pag.users)
			pag.users = slice
		} else {
			Busqueda_Inser(arbol, user, &pag, pag.users[index].izquierda, alt+1)
			fmt.Println("B_izq")
		}
	} else if conf == 3 {
		if pag.users[index].derecha == nil {
			user.altura = alt
			pag.users = append(pag.users, user)
			slice := ordenar_slice(pag.users)
			pag.users = slice
			/*if len(pag.users) == 5 {
				impr(pag.users)
				fmt.Println("-/////")
				var n []*Nodo_B
				n = append(n, pag.users[0])
				n = append(n, pag.users[1])
				var m []*Nodo_B
				m = append(m, pag.users[3])
				m = append(m, pag.users[4])
				prim := &Pagina{alt + 1, n}
				sec := &Pagina{alt + 1, m}
				var l []*Nodo_B
				l = append(l, pag.users[2])
				uniq := &Pagina{alt, l}
				uniq.users[0].izquierda = prim
				uniq.users[0].derecha = sec
				if pag == (*anterior) {
					(*anterior) = uniq
					conf_b = 1
					fmt.Println("Resetear raíz 1")
				} else {
					ind := 0
					fmt.Println("Reseteo ale")
					for ind < len(uniq.users) {
						(*anterior).users = append((*anterior).users, uniq.users[ind])
						slice := ordenar_slice((*anterior).users)
						(*anterior).users = slice
						ind++
					}
				}
			}*/
		} else {
			Busqueda_Inser(arbol, user, &pag, pag.users[index].derecha, alt+1)
			fmt.Println("B_der")
		}
	}
}

func impr(a []*Nodo_B) {
	cont := 0
	for cont < len(a) {
		fmt.Print(a[cont].User.DPI, "-")
		cont++
	}
}

func (m *B_Tree) Print() {
	rompimiento(m.raiz, &m.raiz)
	fmt.Println("Vamo a graficar")
	graph = "digraph List {\n"
	graph += "rankdir=TB;"
	graph += "node [shape = record];"
	imprimir(m.raiz)
	graph += nodes + "\n" + pointers
	graph += "\n}"
	data := []byte(graph)                            //Almacenar el codigo en el formato adecuado
	err := ioutil.WriteFile("graph.dot", data, 0644) //Crear el archivo .dot necesario para la imagen

	if err != nil {
		log.Fatal(err)
	}
	//Creación de la imagen
	fmt.Println(graph)
	path, _ := exec.LookPath("dot") //Para que funcione bien solo asegurate de tener todas las herramientas para
	// Graphviz en tu compu, si no descargalas osea el Graphviz
	cmd, _ := exec.Command(path, "-Tpng", "graph.dot").Output() //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                           //Se mantiene igual
	ioutil.WriteFile("Usuarios.png", cmd, os.FileMode(mode))    //Creacion de la imagen
	pointers = ""
	nodes = ""
	graph = ""
	//imprimir(m.raiz)
}

func inutil(a error) {

}
func imprimir(pag *Pagina) {
	if pag != nil {
		cont := 0
		nodes += "Node" + pag.users[cont].User.DPI + "[label="
		for cont < len(pag.users) {
			if cont == 0 {
				if len(pag.users) > 1 {
					nodes += "\"<f" + strconv.Itoa(cont) + ">" + pag.users[cont].User.DPI + "|"
				} else {
					nodes += "\"<f" + strconv.Itoa(cont) + ">" + pag.users[cont].User.DPI + "\"]\n"
				}
			}
			if cont < len(pag.users)-1 && cont != 0 {
				nodes += "<f" + strconv.Itoa(cont) + ">" + pag.users[cont].User.DPI + "|"
			} else {
				if len(pag.users) > 1 && cont != 0 {
					nodes += "<f" + strconv.Itoa(cont) + ">" + pag.users[cont].User.DPI + "\"]\n"
				}
			}
			if pag.users[cont].izquierda != nil {
				pointers += "\"Node" + pag.users[0].User.DPI + "\":f" + strconv.Itoa(cont) + "->\"Node" + pag.users[cont].izquierda.users[0].User.DPI + "\":f0;\n"
			}
			if cont == len(pag.users)-1 {
				if pag.users[cont].derecha != nil {
					pointers += "\"Node" + pag.users[0].User.DPI + "\":f" + strconv.Itoa(cont) + "->\"Node" + pag.users[cont].derecha.users[0].User.DPI + "\":f0;\n"
				}
			}
			cont++
		}
		cont = 0
		for cont < len(pag.users) {
			imprimir(pag.users[cont].izquierda)
			imprimir(pag.users[cont].derecha)
			cont++
		}
	}
}
func ordenar_slice(array []*Nodo_B) []*Nodo_B {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array); j++ {
			m, err := strconv.Atoi(array[i].User.DPI)
			n, err := strconv.Atoi(array[j].User.DPI)
			inutil(err)
			if m < n {
				aux := array[i]
				array[i] = array[j]
				array[j] = aux
			}
		}
	}
	//var a []*Nodo_B
	/*for i := len(array) - 1; i >= 0; i-- {
		a = append(a, array[i])
	}*/
	return array
}

func busq_orden() {

}

/*func main() {
	fmt.Println("asa")
}*/
