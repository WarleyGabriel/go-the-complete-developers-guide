package main

import (
	"fmt"
	"reflect"
)

type st struct {
	email string
}

func main() {

	//
	// Value types
	//
	var i int = 1

	var f float32 = 1.11

	var s string = "Original"

	var b bool = true

	var stv st
	stv.email = "test@mail.com"

	ai := [3]int{1, 2, 3}

	updateValueTypes(i, f, s, b, stv, ai)

	fmt.Println("Printing value types variables")
	fmt.Println(reflect.TypeOf(i), i)
	fmt.Println(reflect.TypeOf(f), f)
	fmt.Println(reflect.TypeOf(b), b)
	fmt.Println(reflect.TypeOf(s), s)
	fmt.Println(reflect.TypeOf(stv), stv)
	fmt.Println(reflect.TypeOf(ai), ai)
	fmt.Printf("*** Finished\n\n\n")

	//
	// Reference types
	//
	si := []int{1, 2, 3, 4}

	m := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}

	// channels

	// pointers

	// functions

	updateReferenceTypes(si, m)

	fmt.Println("Printing reference types variables")
	fmt.Println(reflect.TypeOf(si), si)
	fmt.Println(reflect.TypeOf(m), m)
	fmt.Println("*** Finished")
}

func updateValueTypes(i int, f float32, s string, b bool, stv st, ai [3]int) {
	i = 9
	f = 9.999
	s = "Updated"
	b = false
	stv.email = "lalallala@mail.com"
	ai[0] = 99
}

func updateReferenceTypes(si []int, m map[string]string) {
	si[0] = 99
	m["red"] = "RED"
}
