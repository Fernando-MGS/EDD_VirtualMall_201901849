package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Fernando-MGS/TEST/AV"
	"github.com/Fernando-MGS/TEST/estructura"
	"github.com/Fernando-MGS/TEST/list"
	"github.com/gorilla/mux"
	//"github.com/Fernando-MGS/TEST/list"
)

type entrada struct {
	Datos []estructura.Data `json:"Datos,omitempty"`
}
type Stores struct {
	Array []list.Tienda
}

type Inventarios struct {
	Inventario []estructura.Inventario `json:"Inventarios,omitempty"`
}

type especifica []list.Tienda

type busqueda struct {
	Departamento string `json:"Departamento,omitempty"`
	Nombre       string `json:"Nombre,omitempty"`
	Calificacion int    `json:"Calificacion,omitempty"`
}

var rowmajor []list.Lista
var depto []string
var index []string
var e entrada

func graficar() {

	b := []byte("Hola mundo!\n")
	err := ioutil.WriteFile("personal.pdf", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func tienda_especifica(info entrada) {

}

func readBody(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e entrada
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

func llenar_depto(e entrada) {
	var inf = e.Datos
	var i = inf[0]
	var j = i.Departamentos
	sum := 0
	for sum < len(j) {
		depto = append(depto, j[sum].Nombre)
		sum++
	}
}

func llenar_index(e entrada) {
	var inf = e.Datos
	sum := 0
	for sum < len(inf) {
		index = append(index, inf[sum].Indice)
		sum++
	}
	fmt.Println(index)
}

func llenar_matriz(info entrada) {
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
		fmt.Println("tOTAL", len(inf))
		var suma = 0
		for suma <= len(j)-1 { //sa
			var k = j[suma].Tiendas
			var calif = 1
			for calif <= 5 {
				lis := list.NewLista()
				var cont = 0
				for cont <= len(k)-1 {
					if k[cont].Calificacion == calif {
						var store list.Tienda
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
	llenar_lista(prelista)
}

func llenar_lista(array []list.Lista) { //toma el arreglo line
	sum := 0
	for sum < len(array) {
		rowmajor = append(rowmajor, array[sum])
		sum++
	}
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	conv := params["numero"]
	i, err := strconv.Atoi(conv)
	tam := rowmajor[i].Tamaño()
	var salida []list.Tienda
	j := 1
	for j <= tam {
		salida = append(salida, rowmajor[i].GetItem(j))
		j++
	}
	fmt.Println(tam)
	fmt.Println(salida)
	salida[0].Inventario.Print()
	var exit especifica
	exit = salida
	json.NewEncoder(w).Encode(exit)
	fmt.Println(err)
}
func DeleteStore(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e busqueda
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
	borrar(e)
	fmt.Println("Final")
	return
}

func borrar(store busqueda) {
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
	fmt.Println("coor ", coordenada, " indice ", indice, " dept ", no_dept, " sum ", sum)
	fmt.Println("Final 1")
	rowmajor[coordenada].Borrar(sum)
	fmt.Println("Final 2")
}

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

func GetStore(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e busqueda
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

func give_tienda(store busqueda) list.Tienda {
	indice := dev_indice(store.Nombre)
	no_dept := dev_depto(store.Departamento)
	coordenada := store.Calificacion + 5*(no_dept+len(depto)*indice) - 1
	fmt.Println(store.Nombre, "nombre-coor", coordenada)
	fmt.Println(depto)
	tes := rowmajor[coordenada]
	sum := 0
	for sum < tes.Tamaño() {
		if tes.GetItem(sum).Nombre == store.Nombre {
			fmt.Println(tes.GetItem(sum).Nombre, "SUM ES", sum)
			break
		}
		sum++
	}
	fmt.Println(tes.GetItem(sum).Nombre, "SUM ES", sum)
	fmt.Println(tes.GetItem(sum - 1).Nombre)
	return (tes.GetItem(sum))
}

func Save(w http.ResponseWriter, r *http.Request) {
	var file estructura.Archivo
	arch := file.Datos
	sum := 0
	row := 0
	for sum < len(index) {
		var dato estructura.Data
		dato.Indice = index[sum]
		var depts estructura.Depto
		cont := 0
		for cont < len(depto) {
			var store []list.Tienda
			depts.Nombre = depto[cont]
			calif := 0
			for calif < 5 {
				long := 0
				for long < rowmajor[row].Tamaño() {
					store = append(store, rowmajor[row].GetItem(long))
				}
				calif++
				row++
			}
			depts.Tiendas = store

			cont++
		}
		sum++
		arch = append(arch, dato)
	}
	file.Datos = arch
	json.NewEncoder(w).Encode(file)
	return
}
func l_inventario(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e Inventarios
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

func llenar_avl(e Inventarios) {
	fmt.Println("llego")
	cont := 0
	for cont < len(e.Inventario) {
		index_dep := dev_depto(e.Inventario[cont].Departamento)
		index_ind := dev_indice(e.Inventario[cont].Tienda)
		fmt.Println("llEGO2", e.Inventario[cont].Departamento, index_dep, "---", e.Inventario[cont].Tienda, index_ind)
		_index := (e.Inventario[cont].Calificacion) + 5*(index_dep+len(depto)*index_ind) - 1
		fmt.Println(_index)
		tmp := rowmajor[_index].Get(e.Inventario[cont].Tienda)
		sum := 0
		fmt.Println("LLego 3")
		for sum < len(e.Inventario[cont].Productos) {
			var prod AV.Producto
			prod.Nombre = e.Inventario[cont].Productos[sum].Nombre
			prod.Cantidad = e.Inventario[cont].Productos[sum].Cantidad
			prod.Codigo = e.Inventario[cont].Productos[sum].Codigo
			prod.Descripcion = e.Inventario[cont].Productos[sum].Descripcion
			prod.Precio = e.Inventario[cont].Productos[sum].Precio
			prod.Imagen = e.Inventario[cont].Productos[sum].Imagen
			tmp.Inventario.Insertar(prod)
			fmt.Println("Llego 4")
			sum++
			rowmajor[_index].Set_Inventario(tmp)
		}
		fmt.Println("Llego 5")
		cont++
	}
}

func give_tiendas(w http.ResponseWriter, r *http.Request) {
	sum := 0
	var i Stores
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

func main() {
	graficar()
	router := mux.NewRouter()
	//endpoint-rutas

	router.HandleFunc("/TiendaEspecifica", GetStore).Methods("POST") //LISTO
	router.HandleFunc("/id/{numero}", GetList).Methods("GET")        //LISTO
	router.HandleFunc("/cargartienda", readBody).Methods("POST")     //LISTO
	router.HandleFunc("/Eliminar", DeleteStore).Methods("POST")
	router.HandleFunc("/Inventarios", l_inventario).Methods("POST")
	router.HandleFunc("/guardar", Save).Methods("GET")
	router.HandleFunc("/Tiendas", give_tiendas).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))

}
