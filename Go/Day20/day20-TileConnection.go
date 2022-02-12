package main

import "fmt"

type TileConnection struct {
	IdFrom        int
	FromSide      int
	FromDirection int
	IdTo          int
	ToSide        int
	ToDirection   int
}

func (tc TileConnection) String() string {
	return fmt.Sprintf("Conn(%d[%d/%d] -> %d[%d/%d])", tc.IdFrom, tc.FromSide, tc.FromDirection, tc.IdTo, tc.ToSide, tc.ToDirection)
}
