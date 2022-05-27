package game

import "github.com/balintkemeny/go-coderetreat-template/pkg/cell"

type World map[cell.Cell]struct{}

type Game struct {
	World World
}
