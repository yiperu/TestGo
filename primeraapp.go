package main

import "fmt"

func main() {

	saludo()
}

func saludo() {

	var frutas = []string{"Manzanas", "Uvas", "Cerezas"}
	frutas = append(frutas, "Peras")

	fmt.Println(frutas[1:2])
	fmt.Println(len(frutas))
}

// type Cafe struct {
// 	nombre string
// 	precio float32
// 	azucar bool
// 	leche  int
// }
