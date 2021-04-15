package Tipos

import (
	"fmt"
)

var Listado []Producto

type nodo_m struct {
	indice   Producto
	altura   int
	izq, der *nodo_m
}

func newnodo_m(indice Producto) *nodo_m {
	return &nodo_m{indice, 0, nil, nil}
}

type AVL struct {
	raiz *nodo_m
}

func NewAVL() *AVL {
	return &AVL{nil}
}

func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func altura(temp *nodo_m) int {
	if temp != nil {
		return temp.altura
	}
	return -1
}

func rotacionIzquierda(temp **nodo_m) {
	aux := (*temp).izq
	(*temp).izq = aux.der
	aux.der = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).izq), (*temp).altura) + 1
	*temp = aux
}

func rotacionDerecha(temp **nodo_m) {
	aux := (*temp).der
	(*temp).der = aux.izq
	aux.izq = *temp
	(*temp).altura = max(altura((*temp).der), altura((*temp).izq)) + 1
	aux.altura = max(altura((*temp).der), (*temp).altura) + 1
	*temp = aux
}

func rotacionDobleIzquierda(temp **nodo_m) {
	rotacionDerecha(&(*temp).izq)
	rotacionIzquierda(temp)
}

func rotacionDobleDerecha(temp **nodo_m) {
	rotacionIzquierda(&(*temp).der)
	rotacionDerecha(temp)
}

func (avl *AVL) Buscar(indice int) int {
	fmt.Println("Llego a buscar ", indice)
	return _prob_exist(indice, &avl.raiz)
}

func (avl *AVL) Comp() int {
	if &avl.raiz != nil {
		return 1
	}
	return 0
}

func _prob_exist(indice int, root **nodo_m) int {
	if (*root) == nil {
		return 0
	}
	if indice < (*root).indice.Codigo {
		fmt.Println(indice, "--", (*root).indice.Codigo)
		prob_exist(indice, &(*root).izq)
	} else if indice > (*root).indice.Codigo {
		fmt.Println(indice, "--", (*root).indice.Codigo)
		prob_exist(indice, &(*root).der)
	} else {
		return 1
	}
	return 0
}

func insert(indice Producto, root **nodo_m) {
	if (*root) == nil {
		*root = newnodo_m(indice)
		return
	}
	if indice.Codigo < (*root).indice.Codigo {
		insert(indice, &(*root).izq)
		if (altura((*root).izq) - altura((*root).der)) == -2 {
			if indice.Codigo < (*root).izq.indice.Codigo {
				rotacionIzquierda(root)
			} else {
				rotacionDobleIzquierda(root)
			}
		}
	} else if indice.Codigo > (*root).indice.Codigo {
		insert(indice, &(*root).der)
		if (altura((*root).der) - altura((*root).izq)) == 2 {
			if indice.Codigo > (*root).der.indice.Codigo {
				rotacionDerecha(root)
			} else {
				rotacionDobleDerecha(root)
			}
		}
	} else {
		(*root).indice.Cantidad = (*root).indice.Cantidad + indice.Cantidad
	}

	(*root).altura = max(altura((*root).izq), altura((*root).der)) + 1
}

func (avl *AVL) Insertar(indice Producto) {
	if prob_exist(indice.Codigo, &avl.raiz) == 0 || prob_exist(indice.Codigo, &avl.raiz) == 2 {
		insert(indice, &avl.raiz)
	} else {
		agregar_cant(indice.Codigo, &avl.raiz, indice)
	}
}

func (avl *AVL) Print() {
	inOrden(avl.raiz)
}

func (avl *AVL) prob_nil() int {
	if avl.raiz == nil {
		return 0
	}
	return 1
}

func prob_exist(indice int, root **nodo_m) int { //0 no existe, 1 si existe
	if (*root) == nil {
		return 0
	}
	if indice < (*root).indice.Codigo {
		prob_exist(indice, &(*root).izq)
	} else if indice > (*root).indice.Codigo {
		prob_exist(indice, &(*root).der)
	} else {
		return 1
	}

	return 2
}

func agregar_cant(indice int, root **nodo_m, prod Producto) { //0 no existe, 1 si existe
	if indice < (*root).indice.Codigo {
		agregar_cant(indice, &(*root).izq, prod)
	} else if indice > (*root).indice.Codigo {
		agregar_cant(indice, &(*root).der, prod)
	} else {
		(*root).indice.Cantidad = (*root).indice.Cantidad + prod.Cantidad
	}
	return
}
func (avl *AVL) Quitar(cant int, prod Producto) {
	avl.Quitar_cant(cant, &avl.raiz, prod)
}

func (avl *AVL) Quitar_cant(indice int, root **nodo_m, prod Producto) { //0 no existe, 1 si existe
	if prod.Codigo < (*root).indice.Codigo {
		avl.Quitar_cant(indice, &(*root).izq, prod)
	} else if prod.Codigo > (*root).indice.Codigo {
		avl.Quitar_cant(indice, &(*root).der, prod)
	} else {
		(*root).indice.Cantidad = (*root).indice.Cantidad - indice
	}
	return
}

func (avl *AVL) Add(cant int, prod Producto) {
	avl.add_cant(cant, &avl.raiz, prod)
}

func (avl *AVL) add_cant(indice int, root **nodo_m, prod Producto) { //0 no existe, 1 si existe
	//fmt.Println("Llego al quitar", (*root).indice.Codigo)
	if prod.Codigo < (*root).indice.Codigo {
		//fmt.Println("Llego al quitar 1<", prod.Codigo, "--", (*root).indice.Codigo)
		avl.add_cant(indice, &(*root).izq, prod)
	} else if prod.Codigo > (*root).indice.Codigo {
		//fmt.Println("Llego al quitar 2>", prod.Codigo, "--", (*root).indice.Codigo)
		//fmt.Println("")
		avl.add_cant(indice, &(*root).der, prod)
	} else {
		//fmt.Println("Llego al AVL", (*root).indice.Cantidad)
		(*root).indice.Cantidad = (*root).indice.Cantidad + indice
	}
	return
}

func inOrden(temp *nodo_m) {
	if temp != nil {
		inOrden(temp.izq)
		fmt.Println("Index: ", temp.indice)
		inOrden(temp.der)
	}
}

func InOrden_prod(temp *nodo_m) {
	if temp != nil {
		InOrden_prod(temp.izq)
		var array []int
		sum := 1
		for sum <= temp.indice.Cantidad {
			array = append(array, sum)
			sum++
		}
		temp.indice.Cant = array
		Listado = append(Listado, temp.indice)
		InOrden_prod(temp.der)
	}
}

func (avl *AVL) Get_Inventario(id string) []Producto {
	var Nuevo []Producto
	Listado = Nuevo
	InOrden_prod(avl.raiz)
	sum := 0
	for sum < len(Listado) {
		Listado[sum].ID = id
		sum++
	}
	return Listado
}
