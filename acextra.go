package main

import (
	"fmt"
)

type No struct {
	dado interface{}
	proximo *No
}

type Fila struct {
	inicio *No
	fim *No
	tamanho int
}

func novaFila() *Fila {
	return &Fila{}
}

func (f *Fila) vazia() bool {
	return f.tamanho == 0
}

func (f *Fila) Tamanho() int {
	return f.tamanho
}

func (f *Fila) enfileira (dado interface{}) {
	novoNo := &No{dado:nil}
	if f.fim == nil {
		f.inicio = novoNo
		f.fim = novoNo
	}
	f.tamanho++
}

func (f *Fila) desenfileira () interface{} {
	if f.vazia() {
		return nil
	}
	dado := f.inicio.dado
	f.inicio = f.inicio.proximo
	if f.inicio == nil {
		f.fim = nil
	}
	f.tamanho--
	return dado
}

func (f *Fila) percorre() {
	atual := f.inicio
	for atual != nil {
		fmt.Println("v ", atual.dado)
		atual = atual.proximo
	}
	fmt.Println()
}




func main() {
		fila := novaFila()

		fila.enfileira(1)
		fila.enfileira(2)
		fila.enfileira(3)
		fila.enfileira(4)

		fmt.Println("Elementos da fila após a inserção:")
		fila.percorre()

		fila.desenfileira()
		fila.desenfileira()

		fmt.Println("Elementos da fila após a remoção:")
		fila.percorre()
	}
