package main

import (
	"encoding/json"
	"log"
	"os"
	"wad/collision"
	// "wad/collision"
)

func main() {
	file, err := os.ReadFile("coords.json")
	if err != nil {
		log.Println("could not read file")
	}

	output := &CoordsFile{}
	json.Unmarshal(file, output)

  log.Printf("num points: %d", len(output.Points))

  count := 0
	for _, v := range output.Points {
		point := collision.Point{
			X: float64(v[0].(float64)),
			Y: float64(v[1].(float64)),
		}

    colliding := false
    colliding = collision.AabbCollision(point, 155, 150, 20, 150)
    if colliding {
      count +=1
      continue
    }

    colliding = collision.RadialCollision(point, 250, 150, 55, 20)
    if colliding {
      count +=1
      continue
    }

    colliding = collision.RadialCollision(point, 410, 150, 55, 20)
    if colliding {
      count +=1
    }

	}

  log.Printf("count: %d", count)
}
