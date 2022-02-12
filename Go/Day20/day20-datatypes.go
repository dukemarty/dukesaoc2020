package main

import (
	"math"
)


type TileMap struct {
	Size           int
	Map            [][]Tile
	processedTiles map[int]bool
}

func (tm TileMap) DeterminePositions(connectedTiles []Tile, corners []Tile) TileMap {
	upperLeftCorner := corners[0]

	tm.Map[0][0] = upperLeftCorner
	tm.processedTiles[upperLeftCorner.Id] = true

	tc := upperLeftCorner.Connections[0]
	tid := 0
	if tc.IdFrom == upperLeftCorner.Id {
		tid = tc.IdTo
	} else {
		tid = tc.IdFrom
	}
	tm.Map[1][0], _ = FindTileById(connectedTiles, tid)

	tc = upperLeftCorner.Connections[1]
	tid = 0
	if tc.IdFrom == upperLeftCorner.Id {
		tid = tc.IdTo
	} else {
		tid = tc.IdFrom
	}
	tm.Map[0][1], _ = FindTileById(connectedTiles, tid)

	return tm
}

func MakeTileMap(tiles []Tile) TileMap {
	res := TileMap{
		Size:           int(math.Sqrt(float64(len(tiles)))),
		processedTiles: make(map[int]bool),
	}

	res.Map = make([][]Tile, res.Size)
	for i := 0; i < res.Size; i++ {
		res.Map[i] = make([]Tile, res.Size)
	}

	return res
}
