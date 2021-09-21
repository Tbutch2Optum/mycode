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
	ast3 := Astro{ "Bill Fishmouth", 42,n: "Titan", isneeded: true}

	fmt.Println(ast1)
	fmt.Println(ast2)
	fmt.Println(ast3)

	// slice named people made up of astro
	p := []astro{ast1, ast2, ast3}
	
	// display the slice
   	 fmt.Println(p)
	 fmt.Println(p[2].name)

}
