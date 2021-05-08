package Tipos

import (
	//"math"
	"fmt"
)

type Hash struct {
	Tamaño      float64 // es el numero de comentarios insertados
	index       int     //es el indice del arreglo de numeros primos
	Comentarios []Nod   //Arreglo de nodos
}

type Nod struct {
	Comentario  Comentario_
	Indice_next int
	Estado      int //0 sin insertar, 1 ocupado

}

type Comentario_ struct {
	Creador    int
	Contenido  string
	Nombre     string
	Respuestas Hash
}

func (h *Hash) Insertar(content string, user Usuario) {
	/*fmt.Println("INSERTANDO", content, "-", user)
	fmt.Println("eL TAMAÑO ES", h.Tamaño, "/", content, "-", user.Nombre)
	fmt.Println("El len es", len(h.Comentarios))*/
	var comentario Comentario_
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
		//fmt.Println("eL TAMAÑO ES", h.Tamaño)
		h.rehashing()
		/*for i := len(h.Comentarios); i < dev_tam(h.index); i++ {
			var vacio Nod
			vacio.Estado = 0
			h.Comentarios = append(h.Comentarios, vacio)
		}*/
		h.index++
	}
}

func (t *Hash) rehashing() {
	//fmt.Println("Rehashiando")
	var new []Nod
	temp := t.Comentarios
	t.Tamaño = 0
	for i := 0; i < dev_tam(t.index); i++ {
		var vacio Nod
		vacio.Estado = 0
		vacio.Indice_next = -1
		new = append(new, vacio)
	}
	t.Comentarios = new
	/*fmt.Println("EL LEN DE REHASH ES", len(t.Comentarios))
	fmt.Println("EL LEN DE temp ES", len(temp))*/
	for i := 0; i < len(temp); i++ {
		if temp[i].Estado == 1 {
			var user Usuario
			user.Nombre = temp[i].Comentario.Nombre
			user.Dpi_ = temp[i].Comentario.Creador
			t.Insertar(temp[i].Comentario.Contenido, user)
			//fmt.Println("i es ", i)
		}
	}
	/*
		m := len(new)
		for i := 0; i < len(h.Comentarios); i++ {
			if h.Comentarios[i].Estado == 1 {
				k := h.Comentarios[i].Comentario.Creador
				form := k % m
				if new[form].Estado != 1 {
					new[form] = h.Comentarios[i]
				} else {
					form = dev_colision(h.Comentarios, 1, form)
					new[form] = h.Comentarios[i]
				}
			}
		}
		fmt.Println("vamo a hacer rehashs, el len es", len(new))*/
}

func dev_colision(h []Nod, i, form int) int {
	anterior := form
	ind := form + i*i
	i++
	if ind >= len(h) {
		ind = (ind + anterior) % len(h)
	}
	if h[ind].Estado == 1 {
		return dev_colision(h, i, form)
	} else {
		return ind
	}
}

func (h *Hash) rehash(tam float64) int {
	if tam >= float64(len(h.Comentarios))*0.6 {
		return 1
	} else {
		return 0
	}
}

func (h Hash) Respuesta(dpi int) {

}

func New_Hash() Hash {
	var niu Hash
	for i := 0; i < 7; i++ {
		var vacio Nod
		vacio.Estado = 0
		vacio.Indice_next = -1
		niu.Comentarios = append(niu.Comentarios, vacio)
	}
	return niu
}

func indice(tam, key int) int {
	index := 0
	index = tam * (key)
	return index
}

func colision(h *Hash, index, i int, com Comentario_) {
	anterior := index
	ind := index + i*i
	i++
	if ind >= len(h.Comentarios) {
		fmt.Println("ind es ", ind, " Y EL LEN ES ", len(h.Comentarios))
		ind = (ind + anterior) % len(h.Comentarios)
		fmt.Println("Ahora ind es ", ind)
	}
	if h.Comentarios[ind].Estado == 1 {
		colision(h, ind, i, com)
	} else {
		h.Comentarios[ind].Comentario = com
		h.Comentarios[ind].Estado = 1
		h.Comentarios[anterior].Indice_next = ind
	}
}

func (h *Hash) Print_com() {
	fmt.Println()
	fmt.Println("////////////////////////////")
	fmt.Println("El len es ", len(h.Comentarios))
	fmt.Println("Existen ", h.Tamaño)
	for i := 0; i < len(h.Comentarios); i++ {
		fmt.Println("===============")
		fmt.Println("El index es ", i)
		fmt.Println("El estado es ", h.Comentarios[i].Estado)
		fmt.Println("El siguiente es ", h.Comentarios[i].Indice_next)
		fmt.Println("Comentario es ", h.Comentarios[i].Comentario.Contenido)
		fmt.Println("DPI es ", h.Comentarios[i].Comentario.Creador)
		fmt.Println("Name es ", h.Comentarios[i].Comentario.Nombre)
	}
}

func (h *Hash) Dev_Comentarios() Comments {
	var comments Comments
	var coment Comentarios
	for i := 0; i < len(h.Comentarios); i++ {
		if h.Comentarios[i].Estado == 1 {
			coment.Index = i
			coment.Contenido = h.Comentarios[i].Comentario.Contenido
			coment.User = h.Comentarios[i].Comentario.Nombre
			comments.Comentarios = append(comments.Comentarios, coment)
		}
	}
	return comments
}
func (h *Hash) Dev_respuestas() {

}

func dev_tam(index int) int {
	primos := []int{11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 105, 107, 109, 113}
	return primos[index]
}
