package AVL_M

/*import (
	"fmt"
)*/
/*
type nodo_m struct {
	indice   Lista
	altura   int
	izq, der *nodo_m
}

func newnodo_m(indice Lista) *nodo_m {
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

func insert(indice Lista, root **nodo_m) {
	if (*root) == nil {
		*root = newnodo_m(indice)
		return
	}
	if indice.año < (*root).indice.año {
		insert(indice, &(*root).izq)
		if (altura((*root).izq) - altura((*root).der)) == -2 {
			if indice.año < (*root).izq.indice.año {
				rotacionIzquierda(root)
			} else {
				rotacionDobleIzquierda(root)
			}
		}
	} else if indice.año > (*root).indice.año {
		insert(indice, &(*root).der)
		if (altura((*root).der) - altura((*root).izq)) == 2 {
			if indice.año > (*root).der.indice.año {
				rotacionDerecha(root)
			} else {
				rotacionDobleDerecha(root)
			}
		}
	} else {
		fmt.Println("Año ya ha sido ingresado")
	}

	(*root).altura = max(altura((*root).izq), altura((*root).der)) + 1
}

func (avl *AVL) Insertar(indice Lista) {
	insert(indice, &avl.raiz)
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

func inOrden(temp *nodo_m) {
	if temp != nil {
		inOrden(temp.izq)
		fmt.Println("Index: ", temp.indice)
		inOrden(temp.der)
	}
}
*/
