package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	arg := os.Args
	//Picks values from the commands line
	if len(arg) > 2 {
		x, err := strconv.Atoi(arg[1])
		y, err1 := strconv.Atoi(arg[2])
		if err == nil && err1 == nil {
			score := score(x, y)
			fmt.Printf("The score is %d", score)
		} else {
			fmt.Println("The values supplied are incorrect")
		}

	} else {
		fmt.Println("Not enough arguments were supplied")
	}

}
func score(x int, y int) int {
	// square the cordinates to also gather for negative values
	// then find squareroot of the positive to get point in
	// cartesion plane distance from origin where dart landed
	distance := math.Sqrt(float64(x*x + y*y))
	switch {
	case distance > 10:
		return 0
	case distance > 5:
		return 1
	case distance > 1:
		return 5
	default:
		return 10
	}
}
