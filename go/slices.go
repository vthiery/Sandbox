package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var output [][]uint8
	for i := 0; i < dy; i++ {
		var in []uint8
		for j := 0; j < dx; j++ {
			in = append(in, uint8(i * j))
		}
		output = append(output, in)
	}
	return output
}

func main() {
	pic.Show(Pic)
}