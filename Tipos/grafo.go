package Tipos

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	//"github.com/Fernando-MGS/TEST/AV"
	//"strconv"
)

type Nodo_G struct {
	Nombre   string
	Enlaces  []*Nodo_G
	Arista   []Arista
	Visitado int //0 no, 1 si
}

type search struct {
	name1 string
	name2 string
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

func (m *Almacen) Camino_corto(origen, destino string) {
	/*index:=m.find_index(origen)
	n:=m*/

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
	//Creación de la imagen
	path, _ := exec.LookPath("dot") //Para que funcione bien solo asegurate de tener todas las herramientas para
	// Graphviz en tu compu, si no descargalas osea el Graphviz
	cmd, _ := exec.Command(path, "-Tpng", "graph.dot").Output() //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                           //Se mantiene igual
	ioutil.WriteFile("Almacen2"+".png", cmd, os.FileMode(mode)) //Creacion de la imagen
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
