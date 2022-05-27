package game_test

import (
	"github.com/balintkemeny/go-coderetreat-template/pkg/game"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GameSuite struct {
	suite.Suite
	subject game.Game
}

func (suite *GameSuite) TestEmptyWorldRemainsEmptyAfterInit() {
	suite.subject = game.Game{
		World: game.World{},
	}

	suite.Equal(0, len(suite.subject.World))
}

func TestGameSuite(t *testing.T) {
	suite.Run(t, new(GameSuite))
}
