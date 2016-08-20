package main

import "golang.org/x/tour/pic"

func myPic(dx, dy int) [][]uint8 {
	nestedArray := make([][]uint8, dy)
	for i := range nestedArray {
		nestedArray[i] = make([]uint8, dx)
	}
	for x := 0; x < dx; x++ {
		for y := 0; y < dx; y++ {
			nestedArray[x][y] = uint8(x ^ y)
		}
	}
	return nestedArray
}

func main() {
	pic.Show(myPic)
}
