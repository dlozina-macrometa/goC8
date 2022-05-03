package graph

import (
	"github.com/marvin-hansen/goC8"
	"github.com/marvin-hansen/goC8/src/utils"
	config "github.com/marvin-hansen/goC8/tests/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllVertices(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())

	res, err := c.Graph.GetAllVertices(fabric, graphName)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}

func TestGetVertex(t *testing.T) {
	c := goC8.NewClient(config.GetDefaultConfig())
	collectionID := "teachers"
	edgeID := "Jean"

	res, err := c.Graph.GetVertex(fabric, graphName, collectionID, edgeID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	utils.PrintRes(res, verbose)
}
