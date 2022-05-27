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
}

func (suite *CellSuite) TestCellCanReturnItsPosition() {
	suite.subject = cell.Cell{
		X: 5,
		Y: 6,
	}

	actualX, actualY := suite.subject.GetPosition()

	suite.Equal(5, actualX)
	suite.Equal(6, actualY)
}

func TestCellSuite(t *testing.T) {
	suite.Run(t, new(CellSuite))
}
