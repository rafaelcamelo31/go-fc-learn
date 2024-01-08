package main

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y := range picture {
		picture[y] = make([]uint8, dx)
		for x := range picture[y] {
			picture[y][x] = uint8((x + y) / 2)
		}
	}

	return picture
}
