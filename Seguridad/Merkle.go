package Seguridad

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type Sub_arbol struct {
	nodo_izquierdo Nodo_markle
	nodo_derecho   Nodo_markle
	raiz           Nodo_markle
	sig            *Sub_arbol
}

type Arbol_Merkle struct {
	fondos     []*Nodo_markle
	insertados int
}

type Nodo_markle struct {
	hash      Data_hash
	Siguiente *Nodo_markle
	Estado    int //0 vacio, 1 ocupado, 2 eliminado
}

type Data_hash struct {
	Hash          string   //cabecera
	Data          string   //datos originales concatenado
	Data_original []string //arreglo de datos originales
}

func newnodo_m(indice Data_hash) *Nodo_markle {
	return &Nodo_markle{indice, nil, 1}
}
func newnodo_nil() *Nodo_markle {
	var b Data_hash
	hsh := sha256.New()
	hsh.Write([]byte("--"))
	y := hex.EncodeToString(hsh.Sum(nil))
	b.Hash = y
	b.Data = "--"
	b.Data_original = append(b.Data_original, "--")
	b.Data_original = append(b.Data_original, "--")
	return &Nodo_markle{b, nil, 0}
}

func raiz(der, izq *Nodo_markle) {
	var d []string
	d = append(d, der.hash.Hash)
	d = append(d, izq.hash.Hash)
	data := dev_hash(d)
	h := sha256.New()
	h.Write([]byte(data.Data))
	c := &Nodo_markle{data, nil, 1}
	der.Siguiente = c
	izq.Siguiente = c
}

func dev_raiz(der, izq *Nodo_markle) *Nodo_markle {
	var d []string
	j := der.hash.Hash + "-" + izq.hash.Hash
	d = append(d, der.hash.Hash)
	d = append(d, izq.hash.Hash)
	data := dev_hash(d)
	h := sha256.New()
	h.Write([]byte(j))
	z := hex.EncodeToString(h.Sum(nil))
	data.Hash = z
	data.Data_original = d
	c := &Nodo_markle{data, nil, 1}
	der.Siguiente = c
	izq.Siguiente = c
	return c
}

func (a *Arbol_Merkle) Insert(dato Data_hash) {
	/*fmt.Println("=======================")
	fmt.Println(dato.Data)*/
	if len(a.fondos) == 0 {
		//fmt.Println("len 0")
		nodo := newnodo_m(dato)
		a.fondos = append(a.fondos, nodo)
		nodo_vacio := newnodo_nil()
		a.fondos = append(a.fondos, nodo_vacio)
		a.insertados = 1
		a.fondos[0].Estado = 1
		completar_raices(a.fondos)
	} else {
		if len(a.fondos) == a.insertados {
			//fmt.Println("iguales", len(a.fondos))
			c := pow_merkle(len(a.fondos), a.insertados)
			//fmt.Println("Pow merkle es", c)
			d := len(a.fondos)
			for i := d; i < c; i++ {
				a.fondos = append(a.fondos, newnodo_nil())
			}
			//fmt.Println("Se inserto en el nulo de pos ", d)
			a.fondos[d] = newnodo_m(dato)
			a.fondos[d].Estado = 1
			a.insertados++
		} else {
			conf := 0
			for i := 0; i < len(a.fondos); i++ {
				if a.fondos[i].Estado != 1 && conf == 0 {
					conf = 1
					a.fondos[i].hash = dato
					a.fondos[i].Estado = 1
					a.insertados++
					break
				}
			}
		}
		completar_raices(a.fondos)
		rehacer_hash(a.fondos, len(a.fondos))
	}
	//fmt.Println(len(a.fondos), "-", a.insertados)
}

func pow_merkle(len, llenos int) int { //revisa que el fondo del arbol sea completo
	conf := 0
	a := 0
	i := 0.0
	for conf != 1 {
		if a > llenos {
			conf = 1

		} else {
			a = int(math.Pow(2, i))
		}
		i++
	}
	//fmt.Println("El pow es", a)
	return a
}

func (a *Arbol_Merkle) Fondo() {
	for i := 0; i < len(a.fondos); i++ {
		fmt.Println("==================")
		fmt.Println(a.fondos[i])
	}
}

func (a *Arbol_Merkle) Delete(dato Data_hash) {
	for i := 0; i < len(a.fondos); i++ {
		if a.fondos[i].hash.Data == dato.Data {
			fmt.Println("Se encontro a ", dato.Data)
			a.fondos[i].Estado = 2
			break
		}
	}
}

func dev_hash(datos []string) Data_hash {
	var a Data_hash
	b := datos[0] + "-"
	for i := 1; i < len(datos); i++ {
		b += datos[i] + "-"
	}
	a.Data_original = datos
	a.Data = b
	h := sha256.New()
	h.Write([]byte(b))
	z := hex.EncodeToString(h.Sum(nil))
	a.Hash = z
	return a
}

func rehacer_hash(nodo []*Nodo_markle, len_original int) {
	if len(nodo) > 1 {
		var nodos []*Nodo_markle
		for i := 0; i < len(nodo); i += 2 {
			root := nodo[i].Siguiente
			if nodo[i].Estado != 2 && nodo[i+1].Estado != 2 { //solo se modifico el if nodo estado
				var data Data_hash
				data.Data = nodo[i].hash.Hash + "-" + nodo[i+1].hash.Hash
				split := strings.Split(data.Data, "-")
				data.Data_original = split
				hsh := sha256.New()
				hsh.Write([]byte(data.Data))
				y := hex.EncodeToString(hsh.Sum(nil))
				data.Hash = y
				root.hash = data
			}
			nodos = append(nodos, root)
		}
		rehacer_hash(nodos, len_original)

	}
}

func (a *Arbol_Merkle) Fix() {
	Fix_hash(a.fondos, len(a.fondos))
}

func Fix_hash(nodo []*Nodo_markle, len_original int) {
	if len(nodo) > 1 {
		var nodos []*Nodo_markle
		for i := 0; i < len(nodo); i += 2 {
			if nodo[i].Estado == 2 {
				nodo[i] = newnodo_nil()

			}
			if nodo[i+1].Estado == 2 {
				nodo[i+1] = newnodo_nil()
			}
			root := nodo[i].Siguiente
			if nodo[i].Estado != 2 && nodo[i+1].Estado != 2 { //solo se modifico el if nodo estado
				var data Data_hash
				data.Data = nodo[i].hash.Hash + "-" + nodo[i+1].hash.Hash
				split := strings.Split(data.Data, "-")
				data.Data_original = split
				hsh := sha256.New()
				hsh.Write([]byte(data.Data))
				y := hex.EncodeToString(hsh.Sum(nil))
				data.Hash = y
				root.hash = data
			}
			nodos = append(nodos, root)
		}
		rehacer_hash(nodos, len_original)

	}
}

func completar_raices(nodo []*Nodo_markle) {
	//fmt.Println("Completar raices alv")
	if len(nodo) > 1 {
		var raices []*Nodo_markle
		for i := 0; i < len(nodo); i += 2 {
			if nodo[i].Siguiente == nil {
				new_root := dev_raiz(nodo[i], nodo[i+1])
				var t []string
				t = append(t, nodo[i].hash.Hash)
				t = append(t, nodo[i+1].hash.Hash)
				new_root.hash.Data_original = t
				raices = append(raices, new_root)
			} else {
				raices = append(raices, nodo[i].Siguiente)
			}
		}
		completar_raices(raices)
	}
}

func (a *Arbol_Merkle) Print(tipo string) {
	//fmt.Println("Vamo a graficar", len(a.fondos))
	graph = "digraph structs {\n"
	nodes = ""
	pointers = ""
	graph += "rankdir=BT;"
	graph += "node [shape = record style=filled];\n"
	_imprimir(a.fondos, 0, len(a.fondos))
	graph += nodes + "\n" + pointers
	graph += "\n}"
	data := []byte(graph)                            //Almacenar el codigo en el formato adecuado
	err := ioutil.WriteFile("graph.dot", data, 0644) //Crear el archivo .dot necesario para la imagen

	if err != nil {
		log.Fatal(err)
	}
	//Creación de la imagen
	//fmt.Println(graph)
	path, _ := exec.LookPath("dot") //Para que funcione bien solo asegurate de tener todas las herramientas para
	// Graphviz en tu compu, si no descargalas osea el Graphviz
	cmd, _ := exec.Command(path, "-Tpdf", "graph.dot").Output()     //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                               //Se mantiene igual
	ioutil.WriteFile("Merkle-"+tipo+".pdf", cmd, os.FileMode(mode)) //Creacion de la imagen
	pointers = ""
	nodes = ""
	graph = ""
}

func _imprimir(nodo []*Nodo_markle, cantidad, point int) {
	//fmt.Println("tamo imprimiendo ", len(nodo))
	if len(nodo) > 1 {
		var raices []*Nodo_markle
		for i := 0; i < len(nodo); i++ {
			nodes += "struct" + strconv.Itoa(cantidad) + " [label=\"{" + nodo[i].hash.Hash + "|{"
			for j := 0; j < len(nodo[i].hash.Data_original); j++ {
				//tam:=len(nodo[i].hash.data_original)
				nodes += nodo[i].hash.Data_original[j] + "\\n"
			}
			nodes += strconv.Itoa(cantidad) + "}}"
			if nodo[i].Estado != 2 {
				nodes += "\" fillcolor= olivedrab1]\n"
			} else {
				nodes += "\" fillcolor= firebrick3]\n"
			}
			//struct4 [shape=record,label="{ b |{c\na\nh}}"];

			if i%2 == 0 {
				raices = append(raices, nodo[i].Siguiente)
				pointers += "struct" + strconv.Itoa(cantidad) + " -> struct" + strconv.Itoa(point) + "\n"
			} else {
				pointers += "struct" + strconv.Itoa(cantidad) + " -> struct" + strconv.Itoa(point) + "\n"
				point++
			}
			cantidad++
		}

		_imprimir(raices, cantidad, point)

	} else {
		nodes += "struct" + strconv.Itoa(cantidad) + " [label=\"{" + nodo[0].hash.Hash + "|{"
		for j := 0; j < len(nodo[0].hash.Data_original); j++ {
			//tam:=len(nodo[i].hash.data_original)
			nodes += nodo[0].hash.Data_original[j] + "\\n"
		}
		nodes += strconv.Itoa(nodo[0].Estado) + "}}\" fillcolor= olivedrab1]\n"
	}
}
