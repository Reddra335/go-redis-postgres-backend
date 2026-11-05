package main

import (
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
)

type Person struct {
	Name string
}
type Asal struct {
	Dari string
}

type Gabung struct {
	Person
	Asal
}

func TestDisplay(t *testing.T) {
	var name string
	var g *errgroup.Group
	g.Go(func() error {
		name = "Reddra"
		return nil
	})

	g.Wait()

	fmt.Println(name)
}
