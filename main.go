package main

import (
	"fmt"
	"github.com/ptflp/gomapper"
)

type Source struct {
	A string `mapper:"a"`
	B string `mapper:"b"`
}

type Dest struct {
	A string `mapper:"a"`
	B string `mapper:"b"`
}

func main() {
	source := Source{
		A: "teasasfasf",
		B: "asfasf",
	}

	dest := Dest{}

	gomapper.MapStructs(&dest, &source)

	fmt.Println(source, dest)
}
