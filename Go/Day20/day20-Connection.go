package main

type SideDescr struct {
	Side      Border
	Direction Direction
}

type Connection struct {
	TargetTileId int
	SourceSide   SideDescr
	TargetSide   SideDescr
}
