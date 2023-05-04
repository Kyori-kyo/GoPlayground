package main

import (
	"fmt"
	"log"

	"gorgonia.org/gorgonia"
)

func main() {

	g := gorgonia.NewGraph()

	var x, y, z *gorgonia.Node
	var err error

	// definir a expressão
	x = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("x"))
	y = gorgonia.NewScalar(g, gorgonia.Float64, gorgonia.WithName("y"))
	if z, err = gorgonia.Add(x, y); err != nil {
		log.Fatal(err)
	}

	// criando uma VM pra rodar o programa
	machine := gorgonia.NewTapeMachine(g)
	defer machine.Close()

	// valores iniciais
	gorgonia.Let(x, 2.0)
	gorgonia.Let(y, 2.5)
	if err = machine.RunAll(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", z.Value())
}
