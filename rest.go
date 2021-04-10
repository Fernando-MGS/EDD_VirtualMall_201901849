package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Fernando-MGS/TEST/Tipos"
	"github.com/Fernando-MGS/TEST/list"
	"github.com/Fernando-MGS/TEST/lista"
	"github.com/Fernando-MGS/TEST/pedidos"
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
	fmt.Println(0, err)
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
	fmt.Println(err)
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
	fmt.Println(er)
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
	find := 0
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
			find = prob_exist_avl(envio[cont].Departamento, envio[cont].Tienda, envio[cont].Calificacion, elemento[cont2].Codigo)
			if find == 1 {
				prod_real = append(prod_real, elemento[cont2])
				fmt.Println(err)
				//matriz.Insert(elemento[cont2],dia,index_dep)
			}
			cont2++
		}
		if len(prod_real) > 0 {
			meses.Insercion(prod_real, index_dep, mes, dia)
			var year pedidos.Year
			year.Año = año
			year.List = *meses
			fmt.Println("Preparandose para año ", año)
			AVL_Pedidos.Insertar(year, prod_real, index_dep, mes, dia)
		}
		cont++
	}
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
	} else { //baja otro
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

func pedido_carrito(w http.ResponseWriter, r *http.Request) {
	sum := 0
	t := time.Now()
	año := t.Year()
	mes := t.String()
	num_mes := strings.Split(mes, "-")
	_mes, err := strconv.Atoi(num_mes[1])
	dia := t.Day()
	fmt.Println(err)
	meses := pedidos.NewLista()

	fmt.Println("El tamaño del carrito es ", carrito.Cantidad)
	for sum < carrito.Cantidad {
		prod := carrito.GetItem(sum)
		fmt.Println(prod.ID)
		index := strings.Split(prod.ID, "-")
		ID, err := strconv.Atoi(index[0])
		fmt.Println("Entro al segundo for", err)
		t := rowmajor[ID].Get(index[1])
		fmt.Println(t.Nombre)
		index_dep := dev_depto(t.Departamento)
		/*tmp := rowmajor[ID].Get(index[1]).Inventario
		tmp.Print()
		fmt.Println("El arbol")*/
		var prod_real []Tipos.Producto
		prod_real = append(prod_real, prod)
		meses.Insercion(prod_real, index_dep, _mes, dia)
		var year pedidos.Year
		year.Año = año
		year.List = *meses
		fmt.Println("/---------------------------------/")
		fmt.Println("Preparandose para insertar en avl pedidos", prod.Nombre)
		AVL_Pedidos.Insertar(year, prod_real, index_dep, _mes, dia)
		sum++
	}
	var new lista.List
	carrito = new
	fmt.Println("Vacío el carrito")
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
	fmt.Println("Año es ", i)
	AVL_Pedidos.Graph_lista(i, j)
	fmt.Println(err)
}

//SESIONES

func devolver_t_user(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(user_tipo)
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
	llenar_users(e)
	return
	//json.NewEncoder(w).Encode(user_tipo)
}

func llenar_users(e Tipos.Cuentas) {
	array_users := e.Usuarios

	sum := 0
	for sum < len(array_users) {
		if array_users[sum].Cuenta == "Admin" {
			array_users[sum].Tipo = 1
		} else {
			array_users[sum].Tipo = 2
		}
		sum++
	}
	fmt.Println(array_users)
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
	sum := sha256.Sum256([]byte(e.Password))
	str2 := string(e.SHA_pass[:])
	fmt.Println("-------")
	fmt.Println(str2)
	//str3 := bytes.NewBuffer(e.SHA_pass[]).String()
	e.SHA_pass = sum
	fmt.Println(e)
	fmt.Printf("%x", e.SHA_pass)
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
	fmt.Println(e)
}

//CODIFICACION FERNET
func Key(w http.ResponseWriter, r *http.Request) {
	m_key = ""
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
	crear_grafo(e)
}

func crear_grafo(e Tipos.File_grafo) {
	//long := len(e.Nodos)
	//var lista_grafo []*Tipos.Nodo_G
	storage.Pos_Robot = e.Pos_init
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
	fmt.Println(len(e.Nodos), "-", len(storage.Estructura))
	fmt.Println(len(e.Nodos[0].Enlaces), "-", len(storage.Estructura[0].Enlaces), "-", len(storage.Estructura[0].Arista))
	c := 0
	for c < len(storage.Estructura) {
		cont := 0
		for cont < len(storage.Estructura[c].Arista) {
			sum := 0
			for sum < len(storage.Estructura) {
				if storage.Estructura[sum].Nombre == storage.Estructura[c].Arista[cont].Destino.Nombre {
					if storage.Prob_exist(sum, storage.Estructura[c].Nombre) == 0 {

						fmt.Println(storage.Estructura[sum].Nombre, "-", storage.Estructura[c].Arista[cont].Destino.Nombre)
						var arista Tipos.Arista
						arista.Peso = storage.Estructura[c].Arista[cont].Peso
						arista.Destino = storage.Estructura[c]
						storage.Estructura[sum].Arista = append(storage.Estructura[sum].Arista, arista)
						fmt.Println(storage.Estructura[sum].Nombre, "se le agrego como arista ", arista.Destino.Nombre)
					}
				}
				sum++
			}
			cont++
		}
		c++
	}
	fmt.Println("--------")
	//storage.Aristas()
	storage.Graficar()
	storage.Grafos()
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

func main() {
	router := mux.NewRouter()
	//endpoint-rutas
	router.HandleFunc("/TiendaEspecifica", GetStore).Methods("POST") //LISTO
	router.HandleFunc("/id/{numero}", GetList).Methods("GET")        //LISTO
	router.HandleFunc("/cargartienda", readBody).Methods("POST")     //LISTO
	router.HandleFunc("/Inventarios", l_inventario).Methods("POST")
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
	router.HandleFunc("/CartSize", CartSize).Methods("GET")
	router.HandleFunc("/user", devolver_t_user).Methods("GET")
	router.HandleFunc("/LoadUsers", cargar_users).Methods("POST")
	router.HandleFunc("/regisUser", regisUser).Methods("POST")
	router.HandleFunc("/loginUser", loginUser).Methods("POST")
	router.HandleFunc("/masterKey", Key).Methods("POST")
	router.HandleFunc("/Loadgrafo", loadGrafo).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", router))
}
