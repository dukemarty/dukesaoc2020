package main

type TileColor int

const (
	BLACK TileColor = 1
	WHITE TileColor = 0
)

type Position struct {
	X int
	Y int
}

type Tile struct {
	NavInstructions string
	Pos             Position
}

func MakeTile(rawTile string) Tile {
	res := Tile{
		NavInstructions: rawTile,
	}

	for i := 0; i < len(rawTile); i++ {
		switch rawTile[i] {
		case 'n':
			res.Pos.Y++
			i++
			if rawTile[i] == 'w' {
				res.Pos.X--
			}
		case 'e':
			res.Pos.X++
		case 's':
			res.Pos.Y--
			i++
			if rawTile[i] == 'e' {
				res.Pos.X++
			}
		case 'w':
			res.Pos.X--
		}
	}

	return res
}
