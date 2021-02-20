package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Fernando-MGS/TEST/estructura"
	"github.com/Fernando-MGS/TEST/list"
	"github.com/gorilla/mux"
	//"github.com/Fernando-MGS/TEST/list"
)

type Person struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json: "firstname,omitempty"`
	LastName  string `json: "lastname,omitempty"`
}

type entrada struct {
	Datos []estructura.Data `json: "Datos,omitempty"`
}

type especifica struct{}

var rowmajor []list.Lista
var people []Person
var e entrada
var largo = 3 * 2
var ti list.Lista

func graficar() {
	b := []byte("Hola mundo!\n")
	err := ioutil.WriteFile("personal.pdf", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
	fmt.Print("Hola")
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	fmt.Print(params)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})

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

func llenar_matriz(info entrada) {
	contador := 0
	var inf = info.Datos
	var i = inf[contador]
	var j = i.Departamentos
	long := len(inf) * len(j) * 5 //inf es filas, j es columnas y las 5 calificaciones
	var prelista = make([]list.Lista, long)
	for contador <= len(inf)-1 {
		fmt.Println("tOTAL", long)
		fmt.Println("Filas", len(inf))
		fmt.Println("Columnas,", len(j))
		var suma = 0
		//fmt.Println(i.Indice)
		for suma <= len(j)-1 { //sa
			//fmt.Println(j[suma].Nombre, suma)
			//fmt.Println("{")
			var k = j[suma].Tiendas
			var calif = 1

			for calif <= 5 {
				lis := list.NewLista()
				var cont = 0
				for cont <= len(k)-1 {
					if k[cont].Calificacion == calif {
						//fmt.Println(k[cont].Nombre, "--", k[cont].Calificacion)
						var store list.Tienda
						store.Nombre = k[cont].Nombre
						store.Contacto = k[cont].Contacto
						store.Descripcion = k[cont].Descripcion
						store.Calificacion = k[cont].Calificacion
						lis.Insertar(store)
					}
					cont++
				}
				var a = (calif - 1) + 5*(suma+len(j)*contador)
				prelista[a] = *lis
				//fmt.Println("Indice", a)
				//terminar de arreglar la posicion y hacerpruebas para guardar la lista
				calif++
			}
			suma++
			//fmt.Println("}")
		}
		contador = contador + 1
	}
	fmt.Println("La long de row es ", len(rowmajor))
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

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func CreateGraphicArray(w http.ResponseWriter, req *http.Request) {

}
func GetStore(w http.ResponseWriter, r *http.Request) {
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
	tienda_especifica(e)
	return
}
func DeleteStore(w http.ResponseWriter, req *http.Request) {

}

func main() {
	graficar()
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Ryan", LastName: "Rey"})
	people = append(people, Person{ID: "2", FirstName: "Joe", LastName: "Yei"})
	//endpoint-rutas
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET") //buscar por id
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/getArreglo", CreateGraphicArray).Methods("GET")
	router.HandleFunc("/TiendaEspecifica", GetStore).Methods("GET")
	router.HandleFunc("/cargartienda", readBody).Methods("POST")
	router.HandleFunc("/id/{numero}", readBody).Methods("GET")
	router.HandleFunc("/Eliminar", DeleteStore).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))

}
