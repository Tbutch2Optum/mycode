package main

import "fmt"

type Astro struct {
	name string
	age int
	lastmission string
	isneeded bool

}

func main() {
	ast1 := Astro{"Dan Casebeer", 29, "Titan",  true}
	ast2 := Astro{ "Tracy Singleton",  35, "Artimus", false}
	ast3 := Astro{ name:"Bill Fishmouth", age: 42, lastmission: "Titan", isneeded: true}

	fmt.Println(ast1)
	fmt.Println(ast2)
	fmt.Println(ast3)
}
