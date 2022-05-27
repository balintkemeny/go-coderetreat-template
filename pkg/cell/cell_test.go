package cell_test

import (
	"github.com/balintkemeny/go-coderetreat-template/pkg/cell"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CellSuite struct {
	suite.Suite
	subject cell.Cell
}

func (suite *CellSuite) SetupTest() {
	suite.subject = cell.Cell{
		X: 2,
		Y: 3,
	}
}

func (suite *CellSuite) TestGetPositionX() {
	actual, _ := suite.subject.GetPosition()
	suite.Equal(actual, 2)
}

func TestCellSuite(t *testing.T) {
	suite.Run(t, new(CellSuite))
}
