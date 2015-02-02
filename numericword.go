package numericword

/*
	numericword converts a number into equivalent words.

	ex: 1024 -> one thousand twenty four
	    -534 -> negative five hundred thirty four

	Author: Tony Grosinger <tony@grosinger.net>
*/

import (
	"fmt"
)

var (
	bigkeys = []int64{
		1000000000000000000,
		1000000000000000,
		1000000000000,
		1000000000,
		1000000,
		1000,
		100}

	bignames = map[int64]string{
		1000000000000000000: "quintillion",
		1000000000000000:    "quadrillion",
		1000000000000:       "trillion",
		1000000000:          "billion",
		1000000:             "million",
		1000:                "thousand",
		100:                 "hundred"}

	tens = []string{
		"",
		"",
		"twenty",
		"thirty",
		"forty",
		"fifty",
		"sixty",
		"seventy",
		"eighty",
		"ninety"}

	ones = []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"ten",
		"eleven",
		"twelve",
		"thirteen",
		"fourteen",
		"fifteen",
		"sixteen",
		"seventeen",
		"eighteen",
		"nineteen"}
)

func convert(i int64) string {
	if i < 20 {
		return ones[i]
	} else if i < 100 && i%10 == 0 {
		return tens[i/10]
	} else if i < 100 {
		return fmt.Sprintf("%s %s", tens[i/10], convert(i%10))
	}

	for j := 0; j < len(bigkeys); j++ {
		place := bigkeys[j]
		if i >= place {
			if i%place == 0 {
				// Skip the extra zero that would occur
				return fmt.Sprintf("%s %s", convert(i/place), bignames[place])
			} else {
				return fmt.Sprintf("%s %s %s",
					convert(i/place),
					bignames[place],
					convert(i%place))
			}
		}
	}

	return "oops"
}

// ToWords converts a number to words up to max or min size of int64
func ToWords(i int64) string {
	if i < 0 {
		return fmt.Sprintf("negative %s", convert(-1*i))
	}
	return convert(i)
}
