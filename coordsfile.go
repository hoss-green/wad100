package main

// import "wad/collision"


type CoordsFile struct {
	Width  int
	Height int
	Points [][]interface{} `json:"coords"`
}
