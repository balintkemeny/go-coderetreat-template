package main

import (
	"fmt"
	"github.com/balintkemeny/go-coderetreat-template/pkg/cell"
)

func main() {
	world := make(map[cell.Cell]struct{})
	fmt.Println(world)
}
