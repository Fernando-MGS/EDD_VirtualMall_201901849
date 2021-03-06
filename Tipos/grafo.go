package Tipos

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
	//"github.com/Fernando-MGS/TEST/AV"
	//"strconv"
)

type Nodo_G struct {
	Nombre   string
	Enlaces  []*Nodo_G
	Arista   []Arista
	Visitado int //0 no, 1 si
}

var large1 int
var large2 int
var large3 int
var large4 int
var search1 Recorridos
var search2 Recorridos
var search3 Recorridos
var search4 Recorridos
var local []string
var final Recorridos
var dest string

type search struct {
	name1 string
	name2 string
	peso  int
}

var caminos []Recorridos

type Recorridos struct {
	Nodos []search
	Peso  int
}

func (m *Almacen) Aristas() {
	c := 0
	for c < len(m.Estructura) {
		sum := 0
		fmt.Println("-------")
		fmt.Println(m.Estructura[c].Nombre, "///", len(m.Estructura[c].Arista), "----", len(m.Estructura[c].Enlaces))
		for sum < len(m.Estructura[c].Arista) {
			fmt.Println(m.Estructura[c].Arista[sum].Destino.Nombre, "-", m.Estructura[c].Arista[sum].Peso)
			sum++
		}
		c++
	}
}

func (m *Almacen) Prob_exist(index int, nombre string) int { //0 no, 1 si existe
	conf := 0 //confirmacion
	cont := 0
	for cont < len(m.Estructura[index].Arista) {
		if nombre == m.Estructura[index].Arista[cont].Destino.Nombre {
			conf = 1
		}
		cont++
	}
	return conf
}

func (m *Almacen) Despacho(destinos []string) {
	inicio := destinos[0]
	fmt.Println("El incicio es", inicio, "///", len(destinos))
	destino := destinos[len(destinos)-1]
	//fmt.Println(destinos)
	//ind := 0
	dest = destino
	local = append(local, inicio)
	for i := 0; i < len(destinos); i++ {
		conf := 0
		for j := 0; j < len(local); j++ {
			//fmt.Println(local[j], "--", destinos[i])
			if local[j] == destinos[i] || destinos[i] == inicio || destinos[i] == destino {
				conf = 1
				//			ind = i
			}
		}
		if conf == 0 {
			local = append(local, destinos[i])
		}
	}
	local = append(local, destino)
	fmt.Println("/////////////////////")
	fmt.Println(local)
	fmt.Println("------")
	fmt.Println(local[0])
	//b := local[0]
	for len(local) > 1 {
		fmt.Println(local[0], "vs", local[1])
		t := m.Camino_corto(local[0], local[1])
		b := local[1]
		fix_destinos(t, b)
		llenar_final(t)
		fmt.Println("vuelta")
	}
	fmt.Println(final)
	Graficar_camino(final)
}

func fix_destinos(usos Recorridos, destino string) {
	var a []string
	var b []string
	var niu []string

	back := local

	for i := 0; i < len(usos.Nodos); i++ {
		a = append(a, usos.Nodos[i].name1)
		if i == len(usos.Nodos)-1 {
			a = append(a, usos.Nodos[i].name2)
		}
	}
	fmt.Println("El recorrido es ", a)
	for j := 0; j < len(a); j++ {
		for i := 0; i < len(local); i++ {
			if a[j] == local[i] {
				b = append(b, local[i])
			}
		}
	}
	fmt.Println("MARCA 2")
	niu = append(niu, destino)
	for i := 0; i < len(local); i++ {
		conf := 0
		for j := 0; j < len(b); j++ {
			if b[j] == local[i] {
				conf = 1
			}
		}
		if conf == 0 {
			niu = append(niu, local[i])
		}
	}
	fmt.Println("MARCA 1")
	if len(a) > 1 {
		local = niu
	} else {
		if len(back) > 0 {
			for i := 0; i < len(back); i++ {
				if i >= 1 {
					a = append(a, back[i])
				}
			}
			local = a
		}
	}

	fmt.Println("=========Local es")
	fmt.Println(local)
}

func llenar_final(a Recorridos) {
	for i := 0; i < len(a.Nodos); i++ {
		final.Nodos = append(final.Nodos, a.Nodos[i])
	}
}

func (m *Almacen) Camino_corto(origen, destino string) Recorridos {
	n := m
	index := n.find_index(origen)
	_index := n.find_index(destino)
	recorrer(n.Estructura[index], n.Estructura[_index], n.Estructura[index], destino, "")
	fmt.Println("1era")
	n.reset_visit()
	_recorrer(m.Estructura[index], m.Estructura[_index], m.Estructura[index], destino)
	fmt.Println("2da")
	n.reset_visit()
	__recorrer(m.Estructura[index], m.Estructura[_index], m.Estructura[index], destino)
	n.reset_visit()
	fmt.Println("3era")
	recorrer_(m.Estructura[index], m.Estructura[_index], m.Estructura[index], destino)
	fmt.Println("4ta")
	n.reset_visit()
	answer := camino_final()
	return answer
}

func (n *Almacen) reset_visit() {
	cont := 0
	for cont < len(n.Estructura) {
		n.Estructura[cont].Visitado = 0
		cont++
	}
}

func arista_menor(nodo, nodo2 *Nodo_G, destino, origen string) int {
	index := 0
	cont := 0
	conf := 0
	if nodo.Arista[cont].Destino.Visitado == 1 {
		index = 1
	}
	for cont < len(nodo.Arista) {
		/*fmt.Println("========")
		fmt.Println(nodo.Arista[cont].Peso, "vs-ind-", nodo.Arista[index].Peso)*/
		if nodo.Arista[cont].Destino.Nombre == destino {
			index = cont
			conf = 1
		}
		if arista_destino(nodo.Arista[cont].Destino.Nombre, nodo2) == 1 &&
			nodo.Arista[cont].Destino.Visitado != 1 && conf != 1 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				index = cont
				conf = 2
			}
		}
		if nodo.Arista[cont].Peso <= nodo.Arista[index].Peso &&
			nodo.Arista[cont].Destino.Visitado == 0 && conf < 1 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				fmt.Println("ENTRO")
				index = cont
			}
		}
		cont++
	}
	fmt.Print("Arista menor", nodo.Arista[index].Destino.Nombre, "conf es ", conf, "-")
	fmt.Println(nodo.Arista[index].Destino.Visitado)
	return index
}

func arista_destino(nombre string, destino *Nodo_G) int {
	conf := 0
	cont := 0
	//index := 0
	for cont < len(destino.Arista) {
		if nombre == destino.Arista[cont].Destino.Nombre {
			conf = 1
		}
		cont++
	}
	return conf
}

func (m *Almacen) find_index(nombre string) int {
	index := 0
	cont := 0
	for cont < len(m.Estructura) {
		if m.Estructura[cont].Nombre == nombre {
			index = cont
			break
		}
		cont++
	}
	return index
}

//PRIMER RECORRIDO
func recorrer(origen, destiny, actual *Nodo_G, destino, anterior string) {
	//fmt.Println("VAMO A RECORRER")
	if exit(anterior, actual.Nombre) == 0 {
		fmt.Println(origen.Nombre, "vs", destiny.Nombre)
		fmt.Println("Estamos en ", actual.Nombre, " y esta", actual.Visitado)
		fmt.Println("Destino en ", destiny.Nombre, " y esta", destiny.Visitado)
		if actual.Visitado == 0 {
			if actual.Nombre == destino {
				search1.Peso = large1
			} else {
				actual.Visitado = 1
				index := arista_menor(actual, destiny, destino, origen.Nombre)
				large1 += actual.Arista[index].Peso
				//actual.Arista[index].Destino.Visitado = 1
				fmt.Println(actual.Nombre, "--", actual.Arista[index].Destino.Nombre)
				time.Sleep(2 * time.Second)
				var s1 search
				s1.name1 = actual.Nombre
				s1.name2 = actual.Arista[index].Destino.Nombre
				s1.peso = actual.Arista[index].Peso
				search1.Nodos = append(search1.Nodos, s1)
				recorrer(origen, destiny, actual.Arista[index].Destino, destino, actual.Nombre)
			}
		} else {
			origen.Visitado = 0
			large1 = 0
			var a Recorridos
			search1 = a
			fmt.Println("vAMO A REINICIAR")
			recorrer(origen, destiny, origen, destino, "")
		}
	}
}

func exit(name1, name2 string) int {
	if name1 == name2 {
		return 1
	}
	return 0
}

func arista_menor_(nodo, nodo2 *Nodo_G, destino, origen string) int {
	index := 0
	cont := 0
	conf := 0
	for cont < len(nodo.Arista) {
		if nodo.Arista[cont].Destino.Nombre == destino {
			index = cont
			conf = 1

		} else if nodo.Arista[cont].Peso <= nodo.Arista[index].Peso &&
			nodo.Arista[cont].Destino.Visitado != 1 && conf != 1 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				index = cont
				conf = 2
			}
		} else if arista_destino(nodo.Arista[cont].Destino.Nombre, nodo2) == 1 &&
			nodo.Arista[cont].Destino.Visitado != 1 && conf > 3 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				index = cont
			}
		}
		cont++
	}
	fmt.Println("Arista menor", nodo.Arista[index].Destino.Nombre)
	return index
}

//SEGUNDA BUSQUEDA
func _recorrer(origen, destiny, actual *Nodo_G, destino string) {
	//fmt.Println("VAMO A RECORRER")
	if actual.Visitado == 0 {
		if actual.Nombre == destino {
			search2.Peso = large2
		} else {
			actual.Visitado = 1
			index := arista_menor(actual, destiny, destino, origen.Nombre)
			large2 += actual.Arista[index].Peso
			//actual.Arista[index].Destino.Visitado = 1
			var s1 search
			s1.name1 = actual.Nombre
			s1.name2 = actual.Arista[index].Destino.Nombre
			s1.peso = actual.Arista[index].Peso
			search2.Nodos = append(search2.Nodos, s1)
			fmt.Println(actual.Nombre, "--", actual.Arista[index].Destino.Nombre)
			_recorrer(origen, destiny, actual.Arista[index].Destino, destino)
		}
	} else {
		origen.Visitado = 0
		var a Recorridos
		search3 = a
		//fmt.Println("vAMO A REINICIAR")
		_recorrer(origen, destiny, origen, destino)
	}
}

func _arista_menor(nodo, nodo2 *Nodo_G, destino, origen string) int {
	index := 0
	cont := 0
	conf := 0
	for cont < len(nodo.Arista) {
		if nodo.Arista[cont].Destino.Nombre == destino {
			index = cont
			conf = 1

		} else if nodo.Arista[cont].Peso < nodo.Arista[index].Peso &&
			nodo.Arista[cont].Destino.Visitado != 1 && conf != 1 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				index = cont
				conf = 2
			}
		} else if arista_destino(nodo.Arista[cont].Destino.Nombre, nodo2) == 1 &&
			nodo.Arista[cont].Destino.Visitado != 1 && conf > 3 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				index = cont
			}
		}
		cont++
	}
	fmt.Println("Arista menor", nodo.Arista[index].Destino.Nombre)
	return index
}

//TERCERA BUSQUEDA
func __recorrer(origen, destiny, actual *Nodo_G, destino string) {
	//fmt.Println("VAMO A RECORRER")
	if actual.Visitado == 0 {
		if actual.Nombre == destino {
			search3.Peso = large3
		} else {
			actual.Visitado = 1
			index := __arista_menor(actual, destiny, destino, origen.Nombre)
			large3 += actual.Arista[index].Peso
			//actual.Arista[index].Destino.Visitado = 1
			var s1 search
			s1.name1 = actual.Nombre
			s1.name2 = actual.Arista[index].Destino.Nombre
			s1.peso = actual.Arista[index].Peso
			search3.Nodos = append(search3.Nodos, s1)
			fmt.Println(actual.Nombre, "--", actual.Arista[index].Destino.Nombre)
			__recorrer(origen, destiny, actual.Arista[index].Destino, destino)
		}
	} else {
		origen.Visitado = 0
		var a Recorridos
		search3 = a
		//fmt.Println("vAMO A REINICIAR")
		__recorrer(origen, destiny, origen, destino)
	}
}

func __arista_menor(nodo, nodo2 *Nodo_G, destino, origen string) int {
	index := 0
	cont := 0
	conf := 0
	for cont < len(nodo.Arista) {
		if nodo.Arista[cont].Peso <= nodo.Arista[index].Peso &&
			nodo.Arista[cont].Destino.Visitado != 1 && conf != 1 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				index = cont
				conf = 2
			}
		} else if nodo.Arista[cont].Destino.Nombre == destino && conf != 2 {
			index = cont
			conf = 1

		} else if arista_destino(nodo.Arista[cont].Destino.Nombre, nodo2) == 1 &&
			nodo.Arista[cont].Destino.Visitado != 1 && conf > 3 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				index = cont
			}
		}
		cont++
	}
	fmt.Println("Arista menor", nodo.Arista[index].Destino.Nombre)
	return index
}

//CUARTA BUSQUEDA
func recorrer_(origen, destiny, actual *Nodo_G, destino string) {
	//fmt.Println("VAMO A RECORRER")
	if actual.Visitado == 0 {
		if actual.Nombre == destino {
			search4.Peso = large4
		} else {
			actual.Visitado = 1
			index := arista__menor(actual, destiny, destino, origen.Nombre)
			large4 += actual.Arista[index].Peso
			//actual.Arista[index].Destino.Visitado = 1
			var s1 search
			s1.name1 = actual.Nombre
			s1.name2 = actual.Arista[index].Destino.Nombre
			s1.peso = actual.Arista[index].Peso
			search4.Nodos = append(search4.Nodos, s1)
			fmt.Println(actual.Nombre, "--", actual.Arista[index].Destino.Nombre)
			recorrer_(origen, destiny, actual.Arista[index].Destino, destino)
		}
	} else {
		origen.Visitado = 0
		var a Recorridos
		search4 = a
		//fmt.Println("vAMO A REINICIAR")
		recorrer_(origen, destiny, origen, destino)
	}
}

func arista__menor(nodo, nodo2 *Nodo_G, destino, origen string) int {
	index := 0
	cont := 0
	conf := 0
	fmt.Println("Arista menor")
	for cont < len(nodo.Arista) {
		if nodo.Arista[cont].Peso < nodo.Arista[index].Peso &&
			nodo.Arista[cont].Destino.Visitado != 1 && conf != 1 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				index = cont
				conf = 2
			}
		} else if nodo.Arista[cont].Destino.Nombre == destino && conf != 2 {
			index = cont
			conf = 1

		} else if arista_destino(nodo.Arista[cont].Destino.Nombre, nodo2) == 1 &&
			nodo.Arista[cont].Destino.Visitado != 1 && conf > 3 {
			if nodo.Arista[cont].Destino.Nombre != origen {
				index = cont
			}
		}
		cont++
	}
	fmt.Println("Arista menor", nodo.Arista[index].Destino.Nombre)
	return index
}

//ELECCION FINAL
func camino_final() Recorridos {
	cont := 0
	index := 0

	rutas := []Recorridos{search1, search2, search3, search4}
	peso := rutas[0].Peso
	for cont < 4 {
		if rutas[cont].Peso <= peso {
			index = cont
			peso = rutas[cont].Peso
		}
		cont++
	}
	var a Recorridos
	search1 = a
	search2 = a
	search3 = a
	search4 = a
	return rutas[index]
}

func (m *Almacen) Graficar() {
	var graph string = "digraph List {\n"
	graph += "rankdir=LR;"
	graph += "node [shape = circle, color=greenyellow , style=filled, fillcolor=darkgreen];"
	var nodes string = ""
	var pointers string = ""
	cont := 0
	for cont < len(m.Estructura) {

		nodes += "Node" + m.Estructura[cont].Nombre + "[label=\"" + m.Estructura[cont].Nombre + "\"]\n"
		cont++
	}
	cont = 0
	for cont < len(m.Estructura) {
		cont2 := 0
		for cont2 < len(m.Estructura[cont].Enlaces) {
			pointers += "Node" + m.Estructura[cont].Nombre + "->Node" + m.Estructura[cont].Enlaces[cont2].Nombre + " [arrowhead=none]" + ";\n"
			cont2++
		}
		cont++
	}
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
	cmd, _ := exec.Command(path, "-Tpng", "graph.dot").Output() //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                           //Se mantiene igual
	ioutil.WriteFile("Almacen"+".png", cmd, os.FileMode(mode))  //Creacion de la imagen
}

func Graficar_camino(ruta Recorridos) {
	var graph string = "digraph List {\n"
	graph += "rankdir=LR;"
	graph += "node [shape = circle, color=greenyellow , style=filled, fillcolor=darkgreen];"
	var nodes string = ""
	var pointers string = ""
	cont := 0
	for cont < len(ruta.Nodos) {
		nodes += "Node" + ruta.Nodos[cont].name1 + "[label=\"" + ruta.Nodos[cont].name1 + "\"]\n"
		cont++
	}
	a := len(ruta.Nodos) - 1
	nodes += "Node" + ruta.Nodos[a].name2 + "[label=\"" + ruta.Nodos[a].name2 + "\"]\n"
	cont = 0
	for cont < len(ruta.Nodos) {
		a := strconv.Itoa(ruta.Nodos[cont].peso)
		pointers += "Node" + ruta.Nodos[cont].name1 + "->Node" + ruta.Nodos[cont].name2 + " [arrowhead=none label=" + a + "]" + ";\n"
		cont++
	}
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
	cmd, _ := exec.Command(path, "-Tpng", "graph.dot").Output()                                  //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                                                            //Se mantiene igual
	ioutil.WriteFile(ruta.Nodos[0].name1+"-"+ruta.Nodos[3].name2+".png", cmd, os.FileMode(mode)) //Creacion de la imagen
}

func (m *Almacen) Grafos() {
	var registro []search
	var graph string = "digraph List {\n"
	graph += "rankdir=LR;"
	graph += "node [shape = circle, color=greenyellow , style=filled, fillcolor=darkgreen];"
	var nodes string = ""
	var pointers string = ""
	cont := 0
	for cont < len(m.Estructura) {
		nodes += "Node" + m.Estructura[cont].Nombre + "[label=\"" + m.Estructura[cont].Nombre + "\"]\n"
		cont++
	}
	cont = 0
	for cont < len(m.Estructura) {
		cont2 := 0
		for cont2 < len(m.Estructura[cont].Arista) {
			var a search
			a.name1 = m.Estructura[cont].Nombre
			a.name2 = m.Estructura[cont].Arista[cont2].Destino.Nombre
			if conf_graf(registro, a) == 0 {
				registro = append(registro, a)
				a := strconv.Itoa(m.Estructura[cont].Arista[cont2].Peso)
				pointers += "Node" + m.Estructura[cont].Nombre + "->Node" + m.Estructura[cont].Arista[cont2].Destino.Nombre + " [arrowhead=none label=" + a + "]" + ";\n"
			}
			cont2++
		}
		cont++
	}
	graph += nodes + "\n" + pointers
	graph += "\n}"
	data := []byte(graph)                            //Almacenar el codigo en el formato adecuado
	err := ioutil.WriteFile("graph.dot", data, 0644) //Crear el archivo .dot necesario para la imagen
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(graph)
	//Creación de la imagen
	path, _ := exec.LookPath("dot") //Para que funcione bien solo asegurate de tener todas las herramientas para
	// Graphviz en tu compu, si no descargalas osea el Graphviz
	cmd, _ := exec.Command(path, "-Tpdf", "graph.dot").Output() //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                           //Se mantiene igual
	ioutil.WriteFile("Almacen2"+".pdf", cmd, os.FileMode(mode)) //Creacion de la imagen
}

func conf_graf(array []search, busq search) int {
	conf := 0
	cont := 0
	for cont < len(array) {
		if array[cont].name1 == busq.name1 {
			if array[cont].name2 == busq.name2 {
				conf = 1
			}
		}
		if array[cont].name1 == busq.name2 {
			if array[cont].name2 == busq.name1 {
				conf = 1
			}
		}
		cont++
	}
	return conf
}
