package main

import (
	"io/ioutil"
)

func main() {
	contenido := []byte("Esto es un dato almacenado desde GO")

	datos := ioutil.WriteFile("archivoEscrito.txt", contenido, 0755)
	mostrarError(datos)
}

func mostrarError(e error) {
	if e != nil {
		panic(e)
	}
}
