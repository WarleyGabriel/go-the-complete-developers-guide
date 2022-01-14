package main

import "fmt"

func main() {
	// var colorsA map[string]string
	// colorsA["blue"] = "#0000FF"
	// fmt.Println(colorsA)
	// https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/

	// colorsB := make(map[string]string)
	// colorsB["white"] = "#FFFFFF"
	// fmt.Println(colorsB)

	colorsC := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"white": "#FFFFFF",
	}
	delete(colorsC, "red")
	// fmt.Println(colorsC)
	printMap(colorsC)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for:", color, "is:", hex)
	}
}
