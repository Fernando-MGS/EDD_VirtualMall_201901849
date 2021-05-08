package Seguridad

import (
	//"math"
	"fmt"

	"github.com/Fernando-MGS/TEST/Tipos"
)

type Hash struct {
	Tamaño      float64
	index       int
	Comentarios []Nod
}

type Nod struct {
	Comentario  Comentario
	Indice_next int
	Estado      int //0 sin insertar, 1 ocupado

}

type Comentario struct {
	Creador    int
	Contenido  string
	Nombre     string
	Respuestas Hash
}

func (h *Hash) rehash(tam float64) int {
	if tam >= h.Tamaño*0.6 {
		return 1
	} else {
		return 0
	}
}

func (h *Hash) Insertar(content string, user Tipos.Usuario) {
	var comentario Comentario
	comentario.Contenido = content
	comentario.Creador = user.Dpi_
	comentario.Nombre = user.Nombre
	/*k := float32(dpi)
	a := math.Sqrt(5) - 1
	a = a / 2*/
	if len(h.Comentarios) == 0 {
		h.Comentarios = New_Hash().Comentarios
	}
	m := len(h.Comentarios)
	/*kA := k * a
	int, mod := math.Modf(k * a)*/
	form := user.Dpi_ % m
	if h.Comentarios[form].Estado == 1 {
		colision(h, form, 1, comentario)
	} else {
		h.Comentarios[form].Estado = 1
		h.Comentarios[form].Comentario = comentario
	}
	h.Tamaño++
	if h.rehash(h.Tamaño) == 1 {
		h.index++
		for i := len(h.Comentarios); i < dev_tam(h.index); i++ {
			var vacio Nod
			vacio.Estado = 0
			h.Comentarios = append(h.Comentarios, vacio)
		}
	}
}

func (h Hash) Respuesta(dpi int) {

}

func New_Hash() Hash {
	var niu Hash
	for i := 0; i < 7; i++ {
		var vacio Nod
		vacio.Estado = 0
		niu.Comentarios = append(niu.Comentarios, vacio)
	}
	return niu
}

func indice(tam, key int) int {
	index := 0
	index = tam * (key)
	return index
}

func colision(h *Hash, index, i int, com Comentario) {
	anterior := index
	ind := index + i*i
	i++
	if h.Comentarios[ind].Estado == 1 {
		colision(h, ind, i, com)
	} else {
		h.Comentarios[ind].Comentario = com
		h.Comentarios[ind].Estado = 1
		h.Comentarios[anterior].Indice_next = ind
	}
}

func (h *Hash) Print_com() {
	fmt.Println("El len es ", len(h.Comentarios))
	fmt.Println("Existen ", h.Tamaño)
	for i := 0; i < len(h.Comentarios); i++ {
		fmt.Println("===============")
		fmt.Println("El estado es ", h.Comentarios[i].Estado)
		fmt.Println("El siguiente es ", h.Comentarios[i].Indice_next)
		fmt.Println("Comentario es ", h.Comentarios[i].Comentario.Contenido)
		fmt.Println("DPI es ", h.Comentarios[i].Comentario.Creador)
		fmt.Println("Name es ", h.Comentarios[i].Comentario.Nombre)
	}
}

func dev_tam(index int) int {
	primos := []int{11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 105, 107, 109, 113}
	return primos[index]
}
