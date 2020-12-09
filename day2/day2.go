/*
 * Copyright (c) 2020. Uriel Márquez All Rights Reserved
 * https://umarquez.c0d3.mx
 */

package day2

import (
	"github.com/pkg/errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var ErrIvalidPattern = "patter format is incorrect, it must match with regexp: '[0-9]*-[0-9]* [a-z]'"

func ValidateCharsRepetitions(values [][]string) (valids int) {
	for _, pair := range values {
		pattern := pair[0]
		password := pair[1]
		if !regexp.MustCompile("[0-9]*-[0-9]* [a-z]").MatchString(pattern) {
			log.Printf("%v", errors.New(ErrIvalidPattern))
			continue
		}

		strMin, strMax, char := "", "", ""
		ix := 0

		for string(pattern[ix]) != "-" {
			strMin += string(pattern[ix])
			ix++
		}
		ix++

		for string(pattern[ix]) != " " {
			strMax += string(pattern[ix])
			ix++
		}
		ix++

		char = string(pattern[ix])
		min, _ := strconv.Atoi(strMin)
		max, _ := strconv.Atoi(strMax)
		counter := strings.Count(password, char)

		if counter >= min && counter <= max {
			valids++
		}
	}

	return
}

func ValidateCharsLocations(values [][]string) (valids int) {
	for _, pair := range values {
		pattern := pair[0]
		password := pair[1]
		if !regexp.MustCompile("[0-9]*-[0-9]* [a-z]").MatchString(pattern) {
			log.Printf("%v", errors.New(ErrIvalidPattern))
			continue
		}

		strA, strB, char := "", "", ""
		ix := 0

		for string(pattern[ix]) != "-" {
			strA += string(pattern[ix])
			ix++
		}
		ix++

		for string(pattern[ix]) != " " {
			strB += string(pattern[ix])
			ix++
		}
		ix++

		char = string(pattern[ix])
		aPos, _ := strconv.Atoi(strA)
		bPos, _ := strconv.Atoi(strB)
		a := string(password[aPos-1]) == char
		b := string(password[bPos-1]) == char
		result := (a || b) && !(a && b) // XORing results (only one positions is allowed to match)

		if result {
			valids++
		}
	}

	return
}
