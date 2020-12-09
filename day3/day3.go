/*
 * Copyright (c) 2020. Uriel Márquez All Rights Reserved
 * https://umarquez.c0d3.mx
 */

package day3

import (
	"strings"
)

type Vector [2]int

func (v Vector) Add(vec Vector) Vector {
	return Vector{v[0] + vec[0], v[1] + vec[1]}
}

func CountTreesUntilTheBottom(pattern string, start Vector, steps Vector) int {
	position := start
	rows := strings.Split(pattern, "\n")
	trees := 0

	for position[1] < len(rows) {
		row := rows[position[1]]
		cell := string(row[position[0]])

		if cell == "#" {
			trees++
		}

		position = position.Add(steps)
		position[0] = (position[0] + len(row)) % len(row)
	}

	return trees
}

func MultiplyTreesOf(pattern string, start Vector, steps []Vector) int {
	var total = 1
	for _, step := range steps {
		res := CountTreesUntilTheBottom(pattern, start, step)
		total *= res
		//log.Printf("%v = %v", step, res)
	}

	return total
}
