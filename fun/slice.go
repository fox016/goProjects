package main

import "fmt"

func buildMultTable(dx, dy int) [][]int {
	table := make([][]int, dy)
	for y:= 0; y < dy; y++ {
		table[y] = make([]int, dx)
		for x := 0; x < dx; x++ {
			table[y][x] = y*x
		}
	}
	return table
}

func main() {
	table := buildMultTable(13, 11)
	for _, v := range table {
		fmt.Println(v)
	}
}
