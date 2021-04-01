package models

import (
	"fmt"
)

// Item ...
type Item struct {
	ID          int
	Name        string
	Description string
	Created     int64
	Updated     int64
	Deleted     int64
}

func main() {
	m := map[int]Item{
		0: Item{0, "Alejandro", "desc0", 0, 0, 0},
	}
	fmt.Println(m[0])

}
