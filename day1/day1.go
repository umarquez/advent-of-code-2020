/*
 * Copyright (c) 2020. Uriel Márquez All Rights Reserved
 * https://umarquez.c0d3.mx
 */

package day1

import (
	"github.com/pkg/errors"
)

const ErrNumsLessThanTwoo = "nums param must be >= 2"

func FindNumbersAndMultiply(sum, nums int, values ...int) (int, error) {
	if nums < 2 {
		return -1, errors.New(ErrNumsLessThanTwoo)
	}
	table := make(map[int][]int)

	for _, v := range values {
		table[sum-v] = []int{v}
	}

	total := 0

	for i := 1; i < nums; i++ {
		for left, ints := range table {
			for _, v := range values {
				id := left - v
				table[id] = append(ints, v)
				if id == 0 && len(table[id]) == nums {
					total = 1
					for _, f := range table[id] {
						total *= f
					}
				}
			}
		}
	}

	return total, nil
}
