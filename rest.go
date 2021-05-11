package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	//"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Fernando-MGS/TEST/Seguridad"
	"github.com/Fernando-MGS/TEST/Tipos"
	"github.com/Fernando-MGS/TEST/list"
	"github.com/Fernando-MGS/TEST/lista"
	"github.com/Fernando-MGS/TEST/pedidos"
	"github.com/fernet/fernet-go"
	"github.com/gorilla/mux"
)

//VARIABLES GLOBALES

var a Tipos.AVL
var rowmajor []list.Lista
var depto []string
var index []string
var AVL_Pedidos pedidos.AVL
var carrito lista.List
var index_pedido int
var user_tipo int //0 no hay sesión, 1 admin, 2 cliente
var m_key string
var storage Tipos.Almacen
var usuarios Seguridad.B_Tree
var admin_def Tipos.Usuario
var user_actual Tipos.Usuario
var pedido_merkle Seguridad.Arbol_Merkle
var user_merkle Seguridad.Arbol_Merkle
var tiendas_merkle Seguridad.Arbol_Merkle
var product_merkle Seguridad.Arbol_Merkle
var bloque int

//F U N C I O N E S

//LLENAR ARREGLO ROWMAJOR
func readBody(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Archivo
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	llenar_matriz(e)
	return
}

func llenar_matriz(info Tipos.Archivo) {
	contador := 0
	llenar_depto(info)
	llenar_index(info)
	var inf = info.Datos
	var a = inf[0]
	var b = a.Departamentos
	long := len(inf) * len(b) * 5 //inf es filas, b es columnas y las 5 calificaciones
	var prelista = make([]list.Lista, long)
	for contador <= len(inf)-1 {
		var i = inf[contador]
		var j = i.Departamentos

		var suma = 0
		for suma <= len(j)-1 { //sa
			var k = j[suma].Tiendas
			var calif = 1
			for calif <= 5 {
				lis := list.NewLista()
				var cont = 0
				for cont <= len(k)-1 {
					if k[cont].Calificacion == calif {
						var a = (calif - 1) + 5*(suma+len(j)*contador)
						var store Tipos.Tienda
						id := strconv.Itoa(a)
						k[cont].ID = id + "-" + k[cont].Nombre
						k[cont].Departamento = b[suma].Nombre
						store = k[cont]
						data := store.Nombre + "-" + store.Departamento + "-" + strconv.Itoa(store.Calificacion)
						var datos Seguridad.Data_hash
						datos.Data = data
						split := strings.Split(data, "-")
						datos.Data_original = split
						h := sha256.New()
						h.Write([]byte(data))
						z := hex.EncodeToString(h.Sum(nil))
						datos.Hash = z
						tiendas_merkle.Insert(datos)
						lis.Insertar(store)
					}
					cont++
				}
				var a = (calif - 1) + 5*(suma+len(j)*contador)
				prelista[a] = *lis
				calif++
			}
			suma++
		}
		contador++
	}
	fmt.Println("tOTAL", len(inf))
	llenar_lista(prelista)

	//pedido_merkle.Fondo()
	//os.Exit(0)

}

func llenar_lista(array []list.Lista) { //toma el arreglo line
	sum := 0
	for sum < len(array) {
		rowmajor = append(rowmajor, array[sum])
		sum++
	}
}

func llenar_depto(e Tipos.Archivo) {
	var inf = e.Datos
	var i = inf[0]
	var j = i.Departamentos
	sum := 0
	for sum < len(j) {
		depto = append(depto, j[sum].Nombre)
		sum++
	}
}

func llenar_index(e Tipos.Archivo) {
	var inf = e.Datos
	sum := 0
	for sum < len(inf) {
		index = append(index, inf[sum].Indice)
		sum++
	}
}

//FUNCIONES PARA ROWMAJOR
func dev_indice(indice string) int { //devuelve el indice del alfabeto
	sum := 0
	nombre := strings.Split(indice, "")
	//alfabeto := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for sum < len(index) {
		if index[sum] == nombre[0] {
			break
		}
		sum++
	}
	return sum
}

func dev_depto(nombre string) int {
	sum := 0
	for sum < len(depto) {
		if nombre == depto[sum] {
			break
		}
		sum++
	}
	return sum
}

//Acciones con tiendas

func GetList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	conv := params["numero"]
	i, err := strconv.Atoi(conv)
	tam := rowmajor[i].Tamaño()
	var salida []Tipos.Tienda
	j := 1
	for j <= tam {
		salida = append(salida, rowmajor[i].GetItem(j))
		j++
	}
	var exit Tipos.T_especifica
	exit = salida
	json.NewEncoder(w).Encode(exit)
	fmt.Println(0, err)
}

func GetStore(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Busqueda
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	exit := give_tienda(e)
	json.NewEncoder(w).Encode(exit)
	return
}

func give_tienda(store Tipos.Busqueda) Tipos.Tienda {
	indice := dev_indice(store.Nombre)
	no_dept := dev_depto(store.Departamento)
	coordenada := store.Calificacion + 5*(no_dept+len(depto)*indice) - 1
	tes := rowmajor[coordenada]
	sum := 0
	for sum < tes.Tamaño() {
		if tes.GetItem(sum).Nombre == store.Nombre {
			break
		}
		sum++
	}
	return (tes.GetItem(sum))
}

func give_tiendas(w http.ResponseWriter, r *http.Request) {
	sum := 0
	var i Tipos.Stores
	for sum < len(rowmajor) {
		if rowmajor[sum].Tamaño() > 0 {
			cont := 0
			tmp := rowmajor[sum]
			for cont < rowmajor[sum].Tamaño() {
				i.Array = append(i.Array, tmp.GetItem(cont))
				cont++
			}

		}
		sum++
	}
	json.NewEncoder(w).Encode(i)
	return
}

func graph_(w http.ResponseWriter, r *http.Request) {
	var graph string
	var pointers string
	var nodes string
	par := 0
	//aux := 0
	graph = "digraph List {\n"
	graph += "rankdir=LR;"
	graph += "node [shape = record];"
	for cont := 0; cont < len(rowmajor); cont++ {
		list := rowmajor[cont]
		if list.Tamaño() > 0 {
			nodes += "node" + strconv.Itoa(cont) + "[label="
		}
		for i := 0; i < list.Tamaño(); i++ {
			if i == 0 {
				if list.Tamaño() > 1 {
					nodes += "\"<f" + strconv.Itoa(i) + ">" + list.GetItem(i+1).Nombre + "|"
				} else {
					nodes += "\"<f" + strconv.Itoa(i) + ">" + list.GetItem(i+1).Nombre + "\"]\n"
				}
			}
			if i < list.Tamaño()-1 && i != 0 {
				nodes += "<f" + strconv.Itoa(i) + ">" + list.GetItem(i+1).Nombre + "|"
			} else {
				if list.Tamaño() > 1 && i != 0 {
					nodes += "<f" + strconv.Itoa(i) + ">" + list.GetItem(i+1).Nombre + "\"]\n"
				}
			}

			//list.GetItem(i+1).
		}

		if list.Tamaño() > 0 {
			if par == 1 {
				pointers += "->\"node" + strconv.Itoa(cont) + "\":f0"
				par = 0
			} else {
				pointers += "\"node" + strconv.Itoa(cont) + "\":f0"
			}
			par = 1
			//aux = cont
		}

	}

	pointers += ";\n"
	graph += nodes + "\n" + pointers
	graph += "\n}"
	//fmt.Println(graph)
	data := []byte(graph)                            //Almacenar el codigo en el formato adecuado
	err := ioutil.WriteFile("graph.dot", data, 0644) //Crear el archivo .dot necesario para la imagen

	if err != nil {
		log.Fatal(err)
	}
	//Creación de la imagen
	//fmt.Println(graph)
	path, _ := exec.LookPath("dot") //Para que funcione bien solo asegurate de tener todas las herramientas para
	// Graphviz en tu compu, si no descargalas osea el Graphviz
	cmd, _ := exec.Command(path, "-Tpdf", "graph.dot").Output() //En esta parte en lugar de graph va el nombre de tu grafica
	mode := int(0777)                                           //Se mantiene igual
	ioutil.WriteFile("Tiendas.pdf", cmd, os.FileMode(mode))     //Creacion de la imagen
}

//Funciones  de inventario

func l_inventario(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Inventarios
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	llenar_avl(e)
	return

}

func llenar_avl(e Tipos.Inventarios) {
	cont := 0
	for cont < len(e.Inventario) {
		index_dep := dev_depto(e.Inventario[cont].Departamento)
		index_ind := dev_indice(e.Inventario[cont].Tienda)
		_index := (e.Inventario[cont].Calificacion) + 5*(index_dep+len(depto)*index_ind) - 1
		tmp := rowmajor[_index].Get(e.Inventario[cont].Tienda)
		sum := 0
		for sum < len(e.Inventario[cont].Productos) {
			prod := e.Inventario[cont].Productos[sum]
			prod.Departamento = e.Inventario[cont].Departamento
			prod.Almacenamiento = nombre_corregid(prod.Almacenamiento)
			tmp.Inventario.Insertar(prod)
			sum++
			rowmajor[_index].Set_Inventario(tmp)
		}
		cont++
	}
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	conv := params["numero"]
	index := strings.Split(conv, "-") //El index[0] tiene el indice en el array y el [1] tiene el nombre
	i, err := strconv.Atoi(index[0])
	str := rowmajor[i].Get(index[1])
	var j Tipos.Products
	j.Array = str.Inventario.Get_Inventario(conv)
	json.NewEncoder(w).Encode(j)
	inutil(err)
}

//Funciones del CARRITO
func getCart(w http.ResponseWriter, r *http.Request) {
	var j Tipos.Products
	j.Array = carrito.GetProducts()
	j.Tamaño = carrito.Tamaño()
	j.Precio = carrito.Precio()
	json.NewEncoder(w).Encode(j)
	return
}

func addProd(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Producto
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	carrito.Add(e)
	return
}

func resetCart() {
	var reset lista.List
	carrito = reset
}

func offProduct(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Producto
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	params := mux.Vars(r)
	conv := params["num"]
	num, err := strconv.Atoi(conv)

	con := strings.Split(e.ID, "-")
	ID, err := strconv.Atoi(con[0])
	t := rowmajor[ID].Get(con[1]).Inventario
	carrito.Putoff_car(e, num)
	if carrito.Tamaño() == 0 {
		resetCart()
	}
	//fmt.Println(err)
	t.Add(num, e)
	return
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var unmarshalErr *json.UnmarshalTypeError
	var e Tipos.Producto
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	params := mux.Vars(r)
	conv := params["id"]
	index := strings.Split(conv, "-") //El index[0] tiene el indice en el array y el [1] tiene el nombre
	i, err := strconv.Atoi(index[0])
	str := rowmajor[i].Get(index[1])
	j, er := strconv.Atoi(index[2])
	str.Inventario.Quitar(j, e)
	inutil(er)
	//fmt.Println(er)
	e.Cantidad = j
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	carrito.Add(e)
	return
}

func CartSize(w http.ResponseWriter, r *http.Request) {
	tamaño := carrito.Tamaño()
	json.NewEncoder(w).Encode(tamaño)
	return
}

func cart_s() {
	fmt.Println(carrito.Tamaño())
}

//PEDIDOS
func addPedido(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Pedidos
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	pedido_json(e)
	return

}
func pedido_json(pedido Tipos.Pedidos) {
	envio := pedido.Pedidos
	var find Tipos.Consulta_prod
	cont := 0
	for cont < len(envio) {
		fecha := strings.Split(envio[cont].Fecha, "-") //dd-mm-aa
		dia, err := strconv.Atoi(fecha[0])
		mes, err := strconv.Atoi(fecha[1])
		año, err := strconv.Atoi(fecha[2])
		cont2 := 0
		meses := pedidos.NewLista()
		elemento := envio[cont].Producto
		index_dep := dev_depto(envio[cont].Departamento) + 1
		var prod_real []Tipos.Producto
		for cont2 < len(elemento) {
			find = _prob_exist_avl(envio[cont].Departamento, envio[cont].Tienda, envio[cont].Calificacion, elemento[cont2].Codigo)
			if find.Find == 1 {
				find.Prod.Dueño = strconv.Itoa(envio[cont].Cliente)
				prod_real = append(prod_real, find.Prod)
				inutil(err)
				var data Seguridad.Data_hash
				data.Data = strconv.Itoa(envio[cont].Cliente) + "," + envio[cont].Fecha + "," + strconv.Itoa(find.Prod.Codigo)
				data.Data_original = strings.Split(data.Data, ",")
				hsh := sha256.New()
				hsh.Write([]byte(data.Data))
				y := hex.EncodeToString(hsh.Sum(nil))
				data.Hash = y
				pedido_merkle.Insert(data)
				//matriz.Insert(elemento[cont2],dia,index_dep)
			}
			cont2++
		}
		if len(prod_real) > 0 {
			//fmt.Println(len(prod_real), dia, "-", mes, "-", año)
			//fmt.Println("El len prod", len(prod_real), index_dep, mes, dia)
			meses.Insercion(prod_real, index_dep, mes, dia)
			var year pedidos.Year
			year.Año = año
			year.List = *meses
			//fmt.Println("Preparandose para año ", año)
			AVL_Pedidos.Insertar(year, prod_real, index_dep, mes, dia)
		}
		cont++
	}

	//AVL_Pedidos.Dev("ENERO", 2013, depto)
	//AVL_Pedidos.Print()
}

func prob_exist_avl(Departamento, Nombre string, Calificacion, Codigo int) int {
	find := 0
	index_dep := dev_depto(Departamento)
	index_ind := dev_indice(Nombre)
	_index := (Calificacion) + 5*(index_dep+len(depto)*index_ind) - 1
	tmp := rowmajor[_index].Get(Nombre).Inventario
	find = tmp.Buscar(Codigo)
	return find
}

func _prob_exist_avl(Departamento, Nombre string, Calificacion, Codigo int) Tipos.Consulta_prod {
	var find Tipos.Consulta_prod
	index_dep := dev_depto(Departamento)
	index_ind := dev_indice(Nombre)
	_index := (Calificacion) + 5*(index_dep+len(depto)*index_ind) - 1
	tmp := rowmajor[_index].Get(Nombre).Inventario
	find = tmp.Buscar_(Codigo)
	return find
}

func dev_pedidos(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	conv := params["id"]
	var resp Tipos.Months
	y := AVL_Pedidos.Dev_year(conv)
	if conv == "0" {
		index_pedido = 0
	} else if conv == "1" { //sube uno
		index_pedido = index_pedido + 1
		if index_pedido == AVL_Pedidos.Dev_year(conv).Large {
			index_pedido = index_pedido - 1
		}
	} else if conv == "2" { //baja otro
		index_pedido = index_pedido - 1
		if index_pedido < 0 {
			index_pedido = index_pedido + 1
		}
	}

	y.Indice = index_pedido
	y.Large = AVL_Pedidos.Dev_year(conv).Large
	resp.Mes = y.Datos[index_pedido].Mes
	resp.Indice = index_pedido
	resp.Año = y.Datos[index_pedido].Año
	resp.Large = y.Large
	json.NewEncoder(w).Encode(resp)
	return
}

func dev_mes(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	conv := params["id"]
	num_mes := strings.Split(conv, "-")
	año, err := strconv.Atoi(num_mes[0])
	inutil(err)
	mes := num_mes[1]
	resp := AVL_Pedidos.Devo(mes, año, depto)
	//fmt.Println(resp)
	json.NewEncoder(w).Encode(resp)
	return
}

func graf_mes(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Si llego al graf_mes")
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.POST // par 1 tiene el año y par 2 el mes
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)

	inutil(err)
	AVL_Pedidos.Dev(e.Par2, e.Par3, depto)
	AVL_Pedidos.Dev_(e.Par2, e.Par3, depto)
	return
}

func pedido_carrito(w http.ResponseWriter, r *http.Request) {
	sum := 0
	t := time.Now()
	año := t.Year()
	mes := t.String()
	num_mes := strings.Split(mes, "-")
	_mes, err := strconv.Atoi(num_mes[1])
	dia := t.Day()
	inutil(err)
	//fmt.Println(err)
	meses := pedidos.NewLista()
	//carro := carrito.GetProducts()
	//fmt.Println("El tamaño del carrito es ", carrito.Cantidad)
	for sum < carrito.Cantidad {
		prod := carrito.GetItem(sum)
		//fmt.Println(prod)
		var data Seguridad.Data_hash
		fecha := strconv.Itoa(dia) + "-" + mes + "-" + strconv.Itoa(año)
		data.Data = strconv.Itoa(user_actual.Dpi_) + "," + fecha + "," + strconv.Itoa(prod.Codigo)
		data.Data_original = strings.Split(data.Data, ",")
		hsh := sha256.New()
		hsh.Write([]byte(data.Data))
		y := hex.EncodeToString(hsh.Sum(nil))
		data.Hash = y
		pedido_merkle.Insert(data)
		index := strings.Split(prod.ID, "-")
		ID, err := strconv.Atoi(index[0])
		inutil(err)
		//fmt.Println("Entro al segundo for", err)
		t := rowmajor[ID].Get(index[1])
		//fmt.Println(t.Nombre)
		index_dep := dev_depto(t.Departamento)
		/*tmp := rowmajor[ID].Get(index[1]).Inventario
		tmp.Print()
		fmt.Println("El arbol")*/
		prod.Dueño = user_actual.Nombre
		var prod_real []Tipos.Producto
		prod_real = append(prod_real, prod)
		meses.Insercion(prod_real, index_dep, _mes, dia)
		//fmt.Println(prod_real)
		var year pedidos.Year
		year.Año = año
		year.List = *meses
		//fmt.Println("/---------------------------------/")
		//fmt.Println("Preparandose para insertar en avl pedidos", prod.Nombre)
		AVL_Pedidos.Insertar(year, prod_real, index_dep, _mes, dia)
		sum++
	}
	var new lista.List
	carrito = new
	//Camino_corto(carro)
	//fmt.Println("Vacío el carrito")
}

// GRAFICAR
func graph_año(w http.ResponseWriter, r *http.Request) {
	//AVL_Pedidos.Print()
	AVL_Pedidos.Grap()
}

func graph_month(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	conv := params["id"]
	index := strings.Split(conv, "-") //El index[0] tiene el año y el [1] tiene el mes

	i, err := strconv.Atoi(index[0])
	j, err := strconv.Atoi(index[1])
	//fmt.Println("Año es ", i)
	AVL_Pedidos.Graph_lista(i, j)
	inutil(err)
	//fmt.Println(err)
}

//SESIONES

func devolver_t_user(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(user_actual)
}

func default_admin() {
	admin_def.DPI = "1234567890101"
	admin_def.Correo = "auxiliar@edd.com"
	admin_def.Password = "1234"
	admin_def.Nombre = "EDD2021"
	admin_def.Tipo = 1
}

func default_user() {
	var a Tipos.Usuario
	user_actual = a
	user_actual.Nombre = "N/A"
	user_actual.Correo = "N/A"
	user_actual.Tipo = 1
}

func cargar_users(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Cuentas
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	//fmt.Println("LLego 1")
	llenar_users(e)
	return
	//json.NewEncoder(w).Encode(user_tipo)
}

func _cargar_users(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Cuentas
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	//fmt.Println("LLego 1")
	_llenar_users(e)
	return
	//json.NewEncoder(w).Encode(user_tipo)
}

func _llenar_users(e Tipos.Cuentas) {
	//fmt.Println("LLego 2")
	array_users := e.Usuarios
	sum := 0
	c := len(array_users) - 1
	array := ordenar_users(array_users, 0, c)
	for sum < len(array) {
		if array_users[sum].Cuenta == "Admin" {
			array[sum].Tipo = 1
		} else {
			array[sum].Tipo = 2
		}
		k := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
		tok, err := fernet.EncryptAndSign([]byte(array[sum].Correo), k[0])
		if err != nil {
			panic(err)
		}
		b := string(tok)
		array[sum].DPI = strconv.Itoa(array[sum].Dpi_)
		array[sum].D_PI = b
		t, err := fernet.EncryptAndSign([]byte(array[sum].Correo), k[0])
		if err != nil {
			panic(err)
		}
		d := string(t)
		array[sum].Mail = d
		h := sha256.New()
		h.Write([]byte(array[sum].DPI))
		z := hex.EncodeToString(h.Sum(nil))
		array[sum].Pass = z
		fmt.Println(array[sum].D_PI, "--", a)
		usuarios.Insertar(array[sum], sum)
		sum++
	}
	usuarios.Print()
	//fmt.Println(array)
	//fmt.Println(usuarios.Buscar(array[0]))

}

func logout(w http.ResponseWriter, r *http.Request) {
	default_user()
}

func test_b() {
	cont := 160
	//users := []int{34, 16, 13, 21, 1, 25, 89, 14, 15, 23, 94, 67, 88, 90, 24, 91, 93, 95, 96, 97, 98, 99, 212, 26, 27, 20, 214, 215, 216, 217, 218, 219, 2}
	var b []Tipos.Usuario
	for cont > 0 {
		var a Tipos.Usuario
		//c := rand.Intn(10000)
		a.DPI = strconv.Itoa(cont)
		a.Dpi_ = cont
		a.Tipo = 2
		a.Password = "Hola"
		b = append(b, a)
		/*usuarios.Insertar(a, cont)
		usuarios.Print(cont)*/
		cont--
	}
	var a Tipos.Usuario
	a.DPI = strconv.Itoa(3031970130108)
	a.Dpi_ = 3031970130108
	a.Tipo = 1
	a.Password = "test"
	//usuarios.Insertar(a, cont)

	/*
		i := 0
		for i < len(b) {
			fmt.Println(b[i].DPI)
			i++
		}*/
	//fmt.Println(b)
	//fmt.Println("Va mo a imprimir") Solo quitar el 88 y ya correra
	final := len(b) - 1
	r := ordenar_users(b, 0, final)
	t := 0
	for t < len(r) {
		usuarios.Insertar(r[t], t)
		t++
	}
	usuarios.Insertar(a, 2)
	//usuarios.Print(t)
}

func ordenar_users(slice []Tipos.Usuario, left int, right int) []Tipos.Usuario {
	pivote := slice[left]
	i := left
	j := right
	var aux Tipos.Usuario
	for i < j {
		for slice[i].Dpi_ <= pivote.Dpi_ && i < j {
			i++
		}
		for slice[j].Dpi_ > pivote.Dpi_ {
			j--
		}
		if i < j {
			aux = slice[i]
			slice[i] = slice[j]
			slice[j] = aux
		}
	}
	slice[left] = slice[j]
	slice[j] = pivote
	if left < j-1 {
		ordenar_users(slice, left, j-1)
	}
	if j+1 < right {
		ordenar_users(slice, j+1, right)
	}
	return slice
}

func llenar_users(e Tipos.Cuentas) {
	//fmt.Println("LLego 2")
	array_users := e.Usuarios
	sum := 0
	c := len(array_users) - 1
	array := ordenar_users(array_users, 0, c)
	for sum < len(array) && sum < 150 {
		var data Seguridad.Data_hash
		if array_users[sum].Cuenta == "Admin" {
			array[sum].Tipo = 1
		} else {
			array[sum].Tipo = 2
		}
		array[sum].DPI = strconv.Itoa(array[sum].Dpi_)
		h := sha256.New()
		h.Write([]byte(array[sum].Password))
		z := hex.EncodeToString(h.Sum(nil))
		array[sum].Pass = z

		k := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
		tok, err := fernet.EncryptAndSign([]byte(array[sum].DPI), k[0])
		if err != nil {
			panic(err)
		}
		b := string(tok)
		array[sum].D_PI = b
		t, err := fernet.EncryptAndSign([]byte(array[sum].Correo), k[0])
		if err != nil {
			panic(err)
		}
		d := string(t)
		array[sum].Mail = d
		usuarios.Insertar(array[sum], sum)
		//insertar en el merkle de usuarios
		data.Data = array[sum].DPI + "-" + array[sum].Nombre + "-" + strconv.Itoa(sum)
		split := strings.Split(data.Data, "-")
		data.Data_original = split
		hsh := sha256.New()
		hsh.Write([]byte(data.Data))
		y := hex.EncodeToString(hsh.Sum(nil))
		data.Hash = y
		user_merkle.Insert(data)
		sum++
	}

	//usuarios.Print()
	//fmt.Println(array)
	//fmt.Println(usuarios.Buscar(array[0]))
}

func graf_users(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.POST
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	if e.Tipo == "0" {
		usuarios.Print()
	} else if e.Tipo == "1" {
		usuarios.Print_()
	} else if e.Tipo == "2" {
		usuarios.Print__()
	}
}

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}
func regisUser(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Usuario

	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)

	//sum := sha256.Sum256([]byte(e.Password))
	//str2 := string(e.SHA_pass[:])
	//fmt.Println("-------")
	//fmt.Println(str2, "str2")
	//str3 := bytes.NewBuffer(e.SHA_pass[]).String()
	h := sha256.New()
	h.Write([]byte(e.Password))
	z := hex.EncodeToString(h.Sum(nil))
	e.Dpi_, err = strconv.Atoi(e.DPI)
	e.Tipo = 2

	k := fernet.MustDecodeKeys("cw_0x689RpI-jtRR7oE8h_eQsKImvJapLeSbXpwF4e4=")
	tok, err := fernet.EncryptAndSign([]byte(e.Correo), k[0])
	if err != nil {
		panic(err)
	}
	b := string(tok)
	e.DPI = strconv.Itoa(e.Dpi_)
	e.D_PI = b
	t, err := fernet.EncryptAndSign([]byte(e.Correo), k[0])
	if err != nil {
		panic(err)
	}
	d := string(t)
	e.Mail = d
	e.Pass = z
	usuarios.Insertar(e, 1)
	//usuarios.Print()
	/*fmt.Println(e.Pass)
	fmt.Println("(())")
	fmt.Println(e.D_PI)
	fmt.Println("(())")
	fmt.Println(e.Mail)*/
	//fmt.Printf("%x", e.SHA_pass)
	user_actual = e
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Consulta

	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	dev_user(e)
}

func dev_user(e Tipos.Consulta) {
	if e.Nombre == admin_def.Nombre && e.Password == admin_def.Password {
		user_actual = admin_def
	} else {
		var buscar_user Tipos.Usuario
		buscar_user.DPI = e.DPI
		//i, err := strconv.Atoi(conv)
		r, err := strconv.Atoi(buscar_user.DPI)
		inutil(err)
		buscar_user.Dpi_ = r
		resultado := usuarios.Buscar(buscar_user)
		//fmt.Println(resultado.DPI, "-", resultado.Dpi_, "-", resultado.Tipo)
		if resultado.Tipo != 0 {
			if e.Password == resultado.Password {
				user_actual = resultado
			} else {
				default_user()
			}
		} else {
			default_user()
		}
	}
}

//CODIFICACION FERNET
func Key(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Llego al mk")
	setupCorsResponse(&w, r)
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Master
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	//fmt.Println(e)
	m_key = e.Key
	//fmt.Println(e)
	fmt.Println(m_key)
	//fmt.Println("___{")
}

//GRAFOS
func loadGrafo(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.File_grafo
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	//fmt.Println(e)
	crear_grafo(corregir_nombres(e))
}

func corregir_nombres(e Tipos.File_grafo) Tipos.File_grafo {
	for i := 0; i < len(e.Nodos); i++ {
		e.Nodos[i].Nombre = nombre_corregid(e.Nodos[i].Nombre)
		k := e.Nodos[i]
		for j := 0; j < len(k.Enlaces); j++ {
			e.Nodos[i].Enlaces[j].Nombre = nombre_corregid(e.Nodos[i].Enlaces[j].Nombre)
		}
	}
	e.Entrega = nombre_corregid(e.Entrega)
	e.Pos_init = nombre_corregid(e.Pos_init)
	return e
}

func nombre_corregid(name string) string {
	k := strings.Split(name, " ")
	p := 0
	niu := ""
	if len(k) > 1 {
		for p < len(k) {
			niu += k[p]
			p++
		}
		return niu
	} else {
		return name
	}

}

func crear_grafo(e Tipos.File_grafo) {
	//long := len(e.Nodos)
	//var lista_grafo []*Tipos.Nodo_G

	storage.Pos_Robot = e.Pos_init
	storage.Entrega = e.Entrega
	fmt.Println(e.Pos_init, "//", storage.Pos_Robot)
	fmt.Println(e.Entrega, "//", storage.Entrega)
	cont := 0
	for cont < len(e.Nodos) {
		var a Tipos.Nodo_G
		a.Nombre = e.Nodos[cont].Nombre
		storage.Estructura = append(storage.Estructura, &a)
		//	lista_grafo=append(lista_grafo, &a)
		cont++
	}
	cont = 0
	for cont < len(storage.Estructura) {
		//	lista_grafo=append(lista_grafo, &a)
		sum := 0
		for sum < len(e.Nodos[cont].Enlaces) {
			sum2 := 0
			for sum2 < len(storage.Estructura) {
				var arista Tipos.Arista
				if e.Nodos[cont].Enlaces[sum].Nombre == storage.Estructura[sum2].Nombre {
					arista.Destino = storage.Estructura[sum2]
					arista.Peso = e.Nodos[cont].Enlaces[sum].Distancia
					/*cont2:=0
					for cont2<len(e.Nodos[cont].Enlaces){
						if storage.Estructura[cont2].Nombre==e.Nodos[cont].Enlaces[sum].Nombre{
							storage.Estructura[cont2].Arista=append(storage.Estructura[cont2].Arista, storage.Estructura[sum2])
						}
						cont2++
					}*/
					storage.Estructura[cont].Enlaces = append(storage.Estructura[cont].Enlaces, storage.Estructura[sum2])
					storage.Estructura[cont].Arista = append(storage.Estructura[cont].Arista, arista)
					break
				}
				sum2++
			}
			sum++
		}
		cont++
	}
	//fmt.Println(len(e.Nodos), "-", len(storage.Estructura))
	//fmt.Println(len(e.Nodos[0].Enlaces), "-", len(storage.Estructura[0].Enlaces), "-", len(storage.Estructura[0].Arista))
	c := 0
	for c < len(storage.Estructura) {
		cont := 0
		for cont < len(storage.Estructura[c].Arista) {
			sum := 0
			for sum < len(storage.Estructura) {
				if storage.Estructura[sum].Nombre == storage.Estructura[c].Arista[cont].Destino.Nombre {
					if storage.Prob_exist(sum, storage.Estructura[c].Nombre) == 0 {

						//fmt.Println(storage.Estructura[sum].Nombre, "-", storage.Estructura[c].Arista[cont].Destino.Nombre)
						var arista Tipos.Arista
						arista.Peso = storage.Estructura[c].Arista[cont].Peso
						arista.Destino = storage.Estructura[c]
						storage.Estructura[sum].Arista = append(storage.Estructura[sum].Arista, arista)
						//fmt.Println(storage.Estructura[sum].Nombre, "se le agrego como arista ", arista.Destino.Nombre)
					}
				}
				sum++
			}
			cont++
		}
		c++
	}
	fmt.Println("se creo el grafo--------")
	//storage.Aristas()
	//storage.Graficar()
	//storage.Grafos()
	//storage.Camino_corto("a", "d")
}

func graf_alm(w http.ResponseWriter, r *http.Request) {
	storage.Graficar()
	storage.Grafos()
	return
}

func Camino_corto(prod []Tipos.Producto) {
	var destinos []string
	fmt.Println("El storage", storage.Pos_Robot)
	destinos = append(destinos, storage.Pos_Robot)
	for i := 0; i < len(prod); i++ {
		destinos = append(destinos, prod[i].Almacenamiento)
	}
	destinos = append(destinos, storage.Entrega)
	storage.Despacho(destinos)
}

func graf_corto(w http.ResponseWriter, r *http.Request) {
	setupCorsResponse(&w, r)
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.File_grafo
	var unmarshalErr *json.UnmarshalTypeError
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	//fmt.Println(e)
	crear_grafo(e)
}

//FASE 4
//Comentarios
func add_comentario(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Comentario
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	params := mux.Vars(r)
	conv := params["id"]
	index := strings.Split(conv, "-") //El index[0] tiene el indice en el array
	//y el [1] tiene el nombre y el [2] tiene el codigo de producto
	//tipo, err := strconv.Atoi(index[0])
	i, err := strconv.Atoi(index[1])
	str := rowmajor[i].Get(index[2])
	if len(index) > 3 {
		codigo, er := strconv.Atoi(index[3])
		inutil(er)
		str.Inventario.Comentar(codigo, e)
		/*prod := str.Inventario.Buscar_(codigo).Prod.Comentarios
		prod.Print_com()*/
	} else {
		rowmajor[i].Comentar(index[2], e)
		//st := rowmajor[i].Get(index[2])
		//str.Comentarios.Insertar(e.Contenido, e.User)
		//st.Comentarios.Print_com()
	}
	//inutil(er)
}

func add_respuesta(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.Respuestas
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
	/*fmt.Println("Estoy respondiendo")
	fmt.Println(e)*/
	params := mux.Vars(r)
	conv := params["id"]
	index := strings.Split(conv, "-") //El index[0] tiene el indice en el array
	//y el [1] tiene el nombre y el [2] tiene el codigo de producto
	//tipo, err := strconv.Atoi(index[0])
	i, err := strconv.Atoi(index[1])
	str := rowmajor[i].Get(index[2])
	if len(index) > 3 {
		codigo, er := strconv.Atoi(index[3])
		inutil(er)
		str.Inventario.Responder(codigo, e)
		/*prod := str.Inventario.Buscar_(codigo).Prod.Comentarios
		prod.Print_com()*/
	} else {
		rowmajor[i].Responder(index[2], e)
		//st := rowmajor[i].Get(index[2])
		//str.Comentarios.Insertar(e.Contenido, e.User)
		//st.Comentarios.Print_com()
	}
}

func ver_comments(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("Llego")
	params := mux.Vars(r)
	conv := params["id"]
	index := strings.Split(conv, "-") //El index[0] tiene el indice en el array
	//y el [1] tiene el nombre y el [2] tiene el codigo de producto
	//tipo, err := strconv.Atoi(index[0])
	i, err := strconv.Atoi(index[1])
	str := rowmajor[i].Get(index[2])
	inutil(err)
	if len(index) > 3 {
		codigo, er := strconv.Atoi(index[3])
		inutil(er)
		producto := str.Inventario.Buscar_(codigo)
		comment := producto.Prod.Comentarios.Dev_Comentarios()
		//fmt.Println(comment)
		json.NewEncoder(w).Encode(comment)
		return
	} else {
		json.NewEncoder(w).Encode(str.Comentarios.Dev_Comentarios())
		return
	}
}

func ver_articulo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	conv := params["id"]
	index := strings.Split(conv, "-") //El index[0] tiene el indice en el array
	//y el [1] tiene el nombre y el [2] tiene el codigo de producto
	//tipo, err := strconv.Atoi(index[0])
	i, err := strconv.Atoi(index[1])
	str := rowmajor[i].Get(index[2])
	inutil(err)
	if len(index) > 3 {
		codigo, er := strconv.Atoi(index[3])
		inutil(er)
		producto := str.Inventario.Buscar_(codigo)
		json.NewEncoder(w).Encode(producto.Prod)
		return
	} else {
		json.NewEncoder(w).Encode(str)
		return
	}
}

//BLOQUES
func add_carp() {
	if _, err := os.Stat("Hash"); os.IsNotExist(err) {
		err = os.Mkdir("Hash", 0755)
		if err != nil {
			// Aquí puedes manejar mejor el error, es un ejemplo
			fmt.Println(err, "-")
			panic(err)
		}
	}
}

//ERRORES DE RESPONSE
func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func inutil(a error) {

}

//MERKLES
func merkle_store(w http.ResponseWriter, r *http.Request) {
	tiendas_merkle.Print("Tiendas")
	return
}

func merkle_ped(w http.ResponseWriter, r *http.Request) {
	pedido_merkle.Print("Pedidos")
	return
}
func merkle_user(w http.ResponseWriter, r *http.Request) {
	user_merkle.Print("Users")
	return
}

func del_store(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	conv := params["id"]
	index := strings.Split(conv, "-") //El index[0] tiene el indice en el array y el [1] tiene el nombre
	i, err := strconv.Atoi(index[0])
	inutil(err)
	fmt.Println(index)
	store := rowmajor[i].Get(index[1])
	data := store.Nombre + "-" + store.Departamento + "-" + strconv.Itoa(store.Calificacion)
	var datos Seguridad.Data_hash
	datos.Data = data
	split := strings.Split(data, "-")
	datos.Data_original = split
	h := sha256.New()
	h.Write([]byte(data))
	z := hex.EncodeToString(h.Sum(nil))
	datos.Hash = z
	tiendas_merkle.Delete(datos)
	return
}

func del_ped(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Tipos.POST
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Archivo Recibido", http.StatusOK)
}

func del_user(w http.ResponseWriter, r *http.Request) {

}

func main() {
	router := mux.NewRouter()
	//test_b()

	cart_s()
	add_carp()
	default_user()
	//endpoint-rutas
	default_admin()
	router.HandleFunc("/TiendaEspecifica", GetStore).Methods("POST") //LISTO
	router.HandleFunc("/id/{numero}", GetList).Methods("GET")        //LISTO
	router.HandleFunc("/cargartienda", readBody).Methods("POST")     //LISTO
	router.HandleFunc("/Inventarios", l_inventario).Methods("POST")
	router.HandleFunc("/graf_stores", graph_).Methods("GET")
	router.HandleFunc("/Tiendas", give_tiendas).Methods("GET")
	router.HandleFunc("/products/{numero}", getProduct).Methods("GET")
	router.HandleFunc("/addProduct", addProd).Methods("POST")
	router.HandleFunc("/offProduct/{num}", offProduct).Methods("POST")
	router.HandleFunc("/addProducto/{id}", addProduct).Methods("POST")
	router.HandleFunc("/getCart", getCart).Methods("GET")
	router.HandleFunc("/Pedido", addPedido).Methods("POST")
	router.HandleFunc("/PedidoCart", pedido_carrito).Methods("POST")
	router.HandleFunc("/year", graph_año).Methods("GET")
	router.HandleFunc("/month/{id}", graph_month).Methods("GET")
	router.HandleFunc("/pedidos/{id}", dev_pedidos).Methods("GET")
	router.HandleFunc("/pedidos_mes/{id}", dev_mes).Methods("GET")
	router.HandleFunc("/CartSize", CartSize).Methods("GET")
	router.HandleFunc("/user", devolver_t_user).Methods("GET")
	router.HandleFunc("/Logout", logout).Methods("GET")
	router.HandleFunc("/LoadUsers", cargar_users).Methods("POST")
	router.HandleFunc("/Load_Users", _cargar_users).Methods("POST")
	router.HandleFunc("/regisUser", regisUser).Methods("POST")
	router.HandleFunc("/loginUser", loginUser).Methods("POST")
	router.HandleFunc("/masterKey", Key).Methods("POST")
	router.HandleFunc("/Loadgrafo", loadGrafo).Methods("POST")
	router.HandleFunc("/graf_grafo", graf_alm).Methods("GET")
	router.HandleFunc("/graf_users", graf_users).Methods("POST")
	//router.HandleFunc("/graf_Lista", graf_list).Methods("POST")
	router.HandleFunc("/graf_mes", graf_mes).Methods("POST")
	router.HandleFunc("/graf_corto", graf_corto).Methods("POST")
	router.HandleFunc("/comentario/{id}", add_comentario).Methods("POST")
	router.HandleFunc("/respuesta/{id}", add_respuesta).Methods("POST")
	router.HandleFunc("/comentarios/{id}", ver_comments).Methods("GET")
	router.HandleFunc("/articulo/{id}", ver_articulo).Methods("GET")
	router.HandleFunc("/merkle_store", merkle_store).Methods("GET")
	router.HandleFunc("/merkle_ped", merkle_ped).Methods("GET")
	router.HandleFunc("/merkle_user", merkle_user).Methods("GET")
	router.HandleFunc("/del_store/{id}", del_store).Methods("GET")
	router.HandleFunc("/del_ped/{id}", merkle_ped).Methods("GET")
	router.HandleFunc("/del_user/{id}", merkle_user).Methods("GET")
	//router.HandleFunc("/graf_b", graf_b).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
