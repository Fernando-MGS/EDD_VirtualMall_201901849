package Tipos

import (
	//"math"
	"fmt"
	"strconv"
	"strings"
)

var resp []Respuesta

type Hash struct {
	Tamaño      float64 // es el numero de comentarios insertados
	index       int     //es el indice del arreglo de numeros primos
	Comentarios []*Nod  //Arreglo de nodos
}

type Nod struct {
	Comentario  Comentario_
	Respuesta   Respuesta
	Indice_next int
	ind_nod     string
	Estado      int //0 sin insertar, 1 ocupado
	//index       []*string
}

type Comentario_ struct {
	Creador    int
	Texto      string
	Contenido  string
	Nombre     string
	ind        string //el indice de la respuesta 1-2-3
	Respuestas Hash
}

func inutil(e error) {

}

func (h *Hash) Respuesta(res Respuestas) {
	indexs := strings.Split(res.Index, "-")
	index, er := strconv.Atoi(indexs[0])
	inutil(er)
	hash := h.Comentarios[index]
	//fmt.Println("llego a respuesta", res.Index)
	responder(&hash, res, dev_indexs(indexs), h.Comentarios[index].Comentario.Nombre)
}

func responder(h **Nod, res Respuestas, indexs []string, rec string) {
	/*fmt.Println("llego a responders", len(indexs))
	fmt.Println(res.Index)*/
	if len(indexs) == 0 {
		(*h).Comentario.Respuestas.Insert_r(res, rec)
		//(*h).Comentario.Respuestas.Print_res()
		/*h.Insert_r(res, rec)
		h.Print_res()*/
	} else {
		/*indexs := strings.Split(res.Index, "-")*/
		index, er := strconv.Atoi(indexs[0])
		inutil(er)
		hash := (*h).Comentario.Respuestas.Comentarios[index]
		responder(&hash, res, dev_indexs(indexs), rec)
	}
}

func dev_indexs(ind []string) []string {
	var niu []string
	for i := 1; i < len(ind); i++ {
		niu = append(niu, ind[i])
	}
	return niu
}

func (h *Hash) Insert_r(res Respuestas, rec string) {
	text := res.User.Nombre + " respondió a " + res.Receptor
	//fmt.Println(text)
	var comentario Comentario_
	comentario.Contenido = res.Respuesta
	comentario.Creador = res.User.Dpi_
	comentario.Nombre = res.User.Nombre
	comentario.Texto = text

	//fmt.Println("Marca 0.5")
	if len(h.Comentarios) == 0 {
		//fmt.Println("Marca 0.6")
		h.Comentarios = New_Hash_r().Comentarios
	}
	//fmt.Println("Marca 1")
	m := len(h.Comentarios)
	//fmt.Println("marca 1.1")
	form := res.User.Dpi_ % m
	if h.Comentarios[form].Estado == 1 {
		colision_r(h, form, 1, comentario, res)
	} else {
		var ind string
		ind = res.Index + "-" + (strconv.Itoa(form))
		comentario.ind = ind
		h.Comentarios[form].Comentario.ind = ind
		fmt.Println("el ind es ", ind)
		//h.Comentarios[form].index = append(h.Comentarios[form].index, &ind)
		h.Comentarios[form].Estado = 1
		h.Comentarios[form].Comentario = comentario
	}
	//fmt.Println("Marca 2")
	h.Tamaño++
	if h.rehash(h.Tamaño) == 1 {
		h.rehashing()
		h.index++
	}

}

func (h *Hash) Insertar(content string, user Usuario) {
	var comentario Comentario_
	comentario.Contenido = content
	comentario.Creador = user.Dpi_
	comentario.Nombre = user.Nombre
	if len(h.Comentarios) == 0 {
		h.Comentarios = New_Hash().Comentarios
	}
	m := len(h.Comentarios)
	form := user.Dpi_ % m
	if h.Comentarios[form].Estado == 1 {
		colision(h, form, 1, comentario)
	} else {
		h.Comentarios[form].Estado = 1
		h.Comentarios[form].Comentario = comentario
		fix_index(&h.Comentarios[form], strconv.Itoa(form))
	}
	h.Tamaño++
	if h.rehash(h.Tamaño) == 1 {
		h.rehashing()
		h.index++
	}
}

func new_index(ind, array string) string {
	niu := ind
	ar := strings.Split(array, "-")
	for i := 1; i < len(array); i++ {
		niu += "-" + ar[i]
	}
	fmt.Println("El nuevo ind es ", niu)
	return niu
}

func (t *Hash) rehashing() {
	var new []*Nod
	temp := t.Comentarios
	t.Tamaño = 0
	for i := 0; i < dev_tam(t.index); i++ {
		var vacio *Nod
		vacio.Estado = 0
		vacio.Indice_next = -1
		new = append(new, vacio)
	}
	t.Comentarios = new
	for i := 0; i < len(temp); i++ {
		if temp[i].Estado == 1 {
			var user Usuario
			user.Nombre = temp[i].Comentario.Nombre
			user.Dpi_ = temp[i].Comentario.Creador
			t.Insertar(temp[i].Comentario.Contenido, user)
		}
	}

}

func fix_index(h **Nod, ind string) {
	if len((*h).Comentario.Respuestas.Comentarios) > 0 {
		for i := 0; i < len((*h).Comentario.Respuestas.Comentarios); i++ {
			if (*h).Comentario.Respuestas.Comentarios[i].Estado == 1 {
				a := (*h).Comentario.Respuestas.Comentarios[i].Comentario.ind
				(*h).Comentario.Respuestas.Comentarios[i].Comentario.ind = new_index(ind, a)
				fix_index(&(*h).Comentario.Respuestas.Comentarios[i], ind)
			}
		}
	}
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

func New_Hash() Hash {
	var niu Hash

	for i := 0; i < 7; i++ {

		var a Comentario_
		var b Respuesta
		vacio := &Nod{Comentario: a, Respuesta: b, Indice_next: -1, Estado: 0}

		niu.Comentarios = append(niu.Comentarios, vacio)

	}

	return niu
}

func New_Hash_r() Hash {
	var niu Hash
	//fmt.Println("Marca 1.2")
	for i := 0; i < 30; i++ {
		//	fmt.Println("Marca 1.21")
		var a Comentario_
		var b Respuesta
		vacio := &Nod{Comentario: a, Respuesta: b, Indice_next: -1, Estado: 0}
		niu.Comentarios = append(niu.Comentarios, vacio)
		//fmt.Println("Marca 1.23")
	}
	//fmt.Println("Marca 1.3")
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
		//fmt.Println("ind es ", ind, " Y EL LEN ES ", len(h.Comentarios))
		ind = (ind + anterior) % len(h.Comentarios)
		//fmt.Println("Ahora ind es ", ind)
	}
	if h.Comentarios[ind].Estado == 1 {
		colision(h, ind, i, com)
	} else {
		h.Comentarios[ind].ind_nod = strconv.Itoa(ind)
		h.Comentarios[ind].Comentario = com
		h.Comentarios[ind].Estado = 1
		h.Comentarios[anterior].Indice_next = ind
		fix_index(&h.Comentarios[ind], strconv.Itoa(ind))
	}
}

func colision_r(h *Hash, index, i int, com Comentario_, res Respuestas) {
	anterior := index
	ind := index + i*i
	i++
	if ind >= len(h.Comentarios) {
		//fmt.Println("ind es ", ind, " Y EL LEN ES ", len(h.Comentarios))
		ind = (ind + anterior) % len(h.Comentarios)
		//fmt.Println("Ahora ind es ", ind)
	}
	if h.Comentarios[ind].Estado == 1 {
		colision_r(h, ind, i, com, res)
	} else {
		com.ind = res.Index + "-" + strconv.Itoa(ind)
		h.Comentarios[ind].ind_nod = strconv.Itoa(ind)
		h.Comentarios[ind].Comentario = com
		h.Comentarios[ind].Estado = 1
		h.Comentarios[anterior].Indice_next = ind
		fmt.Println("Colision El ind es ", h.Comentarios[ind].Comentario.ind)
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

func (h *Hash) Print_res() {
	fmt.Println()
	fmt.Println("////////////////////////////")
	fmt.Println("El len es ", len(h.Comentarios))
	fmt.Println("Existen ", h.Tamaño)
	for i := 0; i < len(h.Comentarios); i++ {
		fmt.Println("===============")
		fmt.Println("El index es ", h.Comentarios[i].Comentario.ind)
		fmt.Println("El texto es ", h.Comentarios[i].Comentario.Texto)
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
			var reset []Respuesta
			resp = reset
			coment.Index = i
			Dev_respuestas(&h.Comentarios[i])
			coment.Respuestas = resp
			//fmt.Println("El array de resp es")
			//fmt.Println(resp)
			coment.Contenido = h.Comentarios[i].Comentario.Contenido
			coment.User = h.Comentarios[i].Comentario.Nombre
			comments.Comentarios = append(comments.Comentarios, coment)
			/*fmt.Println("El array de resp desúeses")
			fmt.Println(resp)*/
		}
	}
	return comments
}
func Dev_respuestas(n **Nod) {
	//temp:=(*n).Comentario.Respuestas.Comentarios
	//fmt.Println("Dev_resp")
	if len((*n).Comentario.Respuestas.Comentarios) > 0 {
		for i := 0; i < len((*n).Comentario.Respuestas.Comentarios); i++ {
			var res Respuesta
			if (*n).Comentario.Respuestas.Comentarios[i].Estado == 1 {
				a := (*n).Comentario.Respuestas.Comentarios[i]
				res.Index = a.Comentario.ind
				res.Respuesta = a.Comentario.Contenido
				res.User = a.Comentario.Texto
				res.Autor = a.Comentario.Nombre
				resp = append(resp, res)
				//fmt.Println((*n).Comentario.Respuestas.Comentarios[i].Comentario.Contenido)
				//dev_resp(&(*n).Comentario.Respuestas.Comentarios[i])
				Dev_respuestas(&(*n).Comentario.Respuestas.Comentarios[i])
			}
		}
	}
}

func dev_resp(n **Nod) {
	fmt.Println("Estoy en dev_resp")
	var res Respuesta
	for i := 0; i < len((*n).Comentario.Respuestas.Comentarios); i++ {
		if (*n).Comentario.Respuestas.Comentarios[i].Estado == 1 {
			a := (*n).Comentario.Respuestas.Comentarios[i]
			res.Index = a.Comentario.ind
			res.Respuesta = a.Comentario.Contenido
			res.User = a.Comentario.Texto
			fmt.Println(res)
			resp = append(resp, res)
		}
	}
}

func dev_tam(index int) int {
	primos := []int{11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 105, 107, 109, 113}
	return primos[index]
}
