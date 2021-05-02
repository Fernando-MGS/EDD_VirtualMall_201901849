package Seguridad

type Hash struct {
	Tamaño      float64
	Comentarios []Nod
}

type Nod struct {
	Comentario Comentario
	Indice     int
	Estado     int //0 sin insertar, 1 ocupado, -1 borrado
}

type Comentario struct {
	Contenido  string
	Respuestas Hash
}

func (h *Hash) rehash(tam float64) {
	if tam >= h.Tamaño*0.6 {

	}
}

func Insertar() {

}

func indice(tam, key int) int {
	index := 0
	index = tam * (key)
	return index
}
