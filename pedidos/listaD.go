package pedidos

import (
	//"fmt"

	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	//"github.com/Fernando-MGS/TEST/AV"
	"github.com/Fernando-MGS/TEST/Tipos"
)

//a ver si se subio
type nodo_l struct { //lugar donde se almacena, tipo tienda del package empresa
	anterior  *nodo_l
	siguiente *nodo_l
	Mes       int
	pedidos   matriz
}

type Lista struct { //apuntadores
	inicio *nodo_l
	ultimo *nodo_l
	tam    int
}

func NewLista() *Lista { //crea una lista
	return &Lista{nil, nil, 0} //el inicio es null, final null, y el tamaño es 0
}

/*type Tienda struct {
	Nombre       string `json:"Nombre,omitempty"`
	Descripcion  string `json:"Descripcion,omitempty"`
	Contacto     string `json:"Contacto,omitempty"`
	Calificacion int    `json:"Calificacion,omitempty"`
	Logo         string `json:"Logo,omitempty"`
	ID           string
	Inventario   AV.AVL
}*/

func (m *Lista) buscar(mes, dia, depto int, l_prod []Tipos.Producto) int {
	aux := m.inicio
	ind := 1
	find := 0
	//fmt.Print("Buscando ", mes)
	for ind <= m.tam {
		if aux.Mes == mes {
			find = 1
			cont := 0
			//fmt.Print(", ha sido encontrado", aux.Mes)
			for cont < len(l_prod) {
				aux.pedidos.Insert(l_prod[cont], dia, depto)
				cont++
			}
			break
		}
		aux = aux.siguiente
		ind++
	}
	return find
}

func (m *Lista) Insercion(l_prod []Tipos.Producto, depto, mes, dia int) {
	find := m.buscar(mes, dia, depto, l_prod)
	if find == 0 {
		m.Insertar(mes, depto, dia, l_prod) //si es un mes nuevo
		//fmt.Println("Insertando nuevo mes")
	}
}

func (m *Lista) Insertar(mes, depto, dia int, l_prod []Tipos.Producto) { //insertar un nodo_m
	matriz := NewMatriz()
	cont := 0
	//fmt.Println("List insertar---------------")
	for cont < len(l_prod) {
		matriz.Insert(l_prod[cont], dia, depto)
		//fmt.Println("For de list-------------")
		cont++
	}
	nuevo := &nodo_l{nil, nil, mes, *matriz}
	if m.inicio == nil {
		//fmt.Println("!!!!!!!")
		m.inicio = nuevo
		m.ultimo = nuevo
	} else {
		//fmt.Println("????????")
		m.ultimo.siguiente = nuevo
		nuevo.anterior = m.ultimo
		m.ultimo = nuevo
	}
	m.tam++
}

func (m *Lista) Imprimir() { //IMPRIMIR
	aux := m.inicio
	for aux != nil {
		//fmt.Println(aux.Mes)
		/*aux.pedidos.lst_h.print_h()
		aux.pedidos.lst_v.print()*/
		aux = aux.siguiente
	}
	//fmt.Print("El tamaño es", m.tam)
}

func (m *Lista) Tamaño() int { //IMPRIMIR TAMAÑO DE LISTA
	return m.tam
}

func (m *Lista) _GetItem(index int) []Tipos.Matrices {
	ind := 1
	aux := m.inicio
	var b []Tipos.Matrices
	var c []Tipos.Matrices
	for ind <= m.tam {
		if index == aux.Mes {

			b = aux.pedidos.lst_h._buscar()
			c = aux.pedidos.lst_v._buscar()
			//aux.pedidos.lst_h._rec_head(aux.pedidos.lst_h.first)
			//aux.pedidos.lst_v._rec_head(aux.pedidos.lst_v.first)
		}
		aux = aux.siguiente
		ind++
	}
	for i := 0; i < len(c); i++ {
		//var d Tipos.Matrices
		//conf := 0
		for j := 0; j < len(b); j++ {
			//fmt.Println(c[i].X, "x-y", c[i].Y, "-", b[j].X, "x-b-y", b[j].Y)
			if c[i].X == b[j].X {
				//conf = 1
			}

		}
		b = append(b, c[i])
		//

	}

	return b
}

func (m *Lista) GetItem(index int) matriz { //Devuelve un dato de la lista
	ind := 1
	aux := m.inicio
	if index <= m.tam {
		for ind < index {
			aux = aux.siguiente
			ind++
		}
	}
	return aux.pedidos
}

func (m *Lista) Dev_array() []string {
	var month []string
	Meses := []string{"ENERO", "FEBRERO", "MARZO", "ABRIL", "MAYO", "JUNIO", "JULIO", "AGOSTO", "SEPTIEMBRE", "OCTUBRE", "NOVIEMBRE", "DICIEMBRE"}
	aux := m.inicio
	for aux != nil {
		month = append(month, Meses[aux.Mes-1])
		aux = aux.siguiente
	}
	return month
}

func (elist *Lista) GraphData(año string) {
	auxiliar := elist.inicio
	Meses := []string{"ENERO", "FEBRERO", "MARZO", "ABRIL", "MAYO", "JUNIO", "JULIO", "AGOSTO", "SEPTIEMBRE", "OCTUBRE", "NOVIEMBRE", "DICIEMBRE"}
	//la variable graph me ayudara a guardar toda el codigo del grafico.
	var graph string = "digraph List {\n" //Este es el encabezado no debe cambiar nada solo se puede cambiar el nombre List por el de
	// de su preferencia lo demás se queda así.
	graph += "rankdir=TB;" //Esto es solo para que la grafica se ordene en modo horizontal, puede cambiar si es necesario si se quiere
	//vertical se cambia LR por TB.
	//Esta linea es para modificar como se ve el nodo tanto el color interno como los bordes.
	graph += "node [shape = circle, color=greenyellow , style=filled, fillcolor=darkgreen];"
	var nodes string = ""    //Manejara todos los nodos la declaracion
	var pointers string = "" //Manejara todos los punteros y conexiones, es mejor separarlo para que no haya conflicto luego.
	for auxiliar != nil {
		//Como los nodos deben tener un nombre unico entonces le concatene su valor, entonces si un nodo tiene dentro un 5 entonces
		//se llamaria node5 y ahora bien el label es lo que tendra dentro del nodo aqui puede ir el nombre de la tienda.
		nodes += "Node" + strconv.Itoa(auxiliar.Mes) + "[label=\"" + Meses[auxiliar.Mes-1] + "\"]\n"
		if auxiliar.siguiente != nil {
			//Aqui se almacenan los punteros permite apuntar del actual al siguiente
			pointers += "Node" + strconv.Itoa(auxiliar.Mes) + "->Node" + strconv.Itoa(auxiliar.siguiente.Mes) + ";\n"
		}
		auxiliar = auxiliar.siguiente
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
	cmd, _ := exec.Command(path, "-Tpng", "graph.dot").Output()           //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                                     //Se mantiene igual
	ioutil.WriteFile("Pedidos-meses-"+año+".png", cmd, os.FileMode(mode)) //Creacion de la imagen
}
