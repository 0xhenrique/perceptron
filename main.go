package main

import (
	"fmt"
)

const WIDTH = 20
const HEIGHT = 20

func fillRect(inputs [HEIGHT][WIDTH]float32, x int, y int, w int, h int, value float32) {
	if w <= 0 { panic("width wrong!") }
	if h <= 0 { panic("height wrong!") }

	x0 := Clamp(x, 0, WIDTH - 1)
	y0 := Clamp(y, 0, HEIGHT - 1)
	x1 := Clamp(x0 + w - 1, 0, WIDTH - 1)
	y1 := Clamp(y0 + y - 1, 0, HEIGHT - 1)

	for y := y0; y <= y1; y++ {
		for x := x0; x <= x1; x++ {
			inputs[x][y] = value
		}
		
	}
}

// @TODO: Maybe not necessary
func Clamp(value, min, max int) int {
	if value < min { return min }
	if value > max { return max }

	return value
}

func forwardPropagation(inputs [HEIGHT][WIDTH]float32, weights [HEIGHT][WIDTH]float32) float32 {
	var output float32 = 0

    for y := 0; y < HEIGHT; y++ {
        for x := 0; x < WIDTH; x++ {
            output += inputs[y][x] * weights[y][x]
        }
    }

	return output
}

func main() {
    var inputs [HEIGHT][WIDTH]float32
    var weights [HEIGHT][WIDTH]float32

	inputs[0][0] = 38.0
	weights[0][0] = 382.0

	fillRect(inputs, 0, 0, WIDTH, HEIGHT, 0.00)
    result := forwardPropagation(inputs, weights)
	fmt.Println("Result: ", result)
}


