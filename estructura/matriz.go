package estructura

import (
	"fmt"

	"github.com/Fernando-MGS/TEST/list"
	"github.com/Fernando-MGS/TEST/lista"
)

var prueba = 2 * 2 * 5
var deep = make([]list.Tienda, 5)
var filas = make([]int, 3)

var tiendas [2][2][5]int

var matriz = make([][][]list.Tienda, 5)
var rowmajor = make([]int, prueba)
var arr [40]list.Lista

func testa() {
	cont := 1
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 5; k++ {
				tiendas[i][j][k] = cont
				cont++
			}
		}
	}
	fmt.Println(tiendas)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 5; k++ {
				rowmajor[k+5*(j+2*i)] = tiendas[i][j][k] //fallo porque hay que revisar que el dos tal vez tenga que ver con las dimensiones
			} //5=PROFUNDIDAD, 2 = FILAS
		}
	}
	fmt.Println("Hola_", rowmajor)
	lis := lista.NewLista()
	lis.Insertar(2)
	lis.Insertar(4)
	lis.Insertar(3)
	lis.GetItem(2)
	var store list.Tienda
	store.Nombre = "JOJO YESU"
	store.Descripcion = "Aveerrr si funciona"
	store.Calificacion = 1
	test := list.NewLista()
	test.Insertar(store)
	test.GetItem(1)
	fmt.Print(store)
	//var vector[]
}
