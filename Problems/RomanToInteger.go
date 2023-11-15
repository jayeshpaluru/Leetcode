package main

// Given a roman numeral, convert it to an integer

// Solution: Use a map to store the roman numerals and their corresponding integer values
// Iterate through the roman numeral string
// If the current roman numeral is less than the next roman numeral, then subtract the current roman numeral from the next roman numeral
// Otherwise, add the current roman numeral to the integer value
func romanToInt(s string) int {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	integer := 0
	lastNumeral := 0

	for _, currentNumeral := range s {
		currentValue := romanNumerals[currentNumeral]

		if currentValue > lastNumeral {
			integer += currentValue - 2*lastNumeral
		} else {
			integer += currentValue
		}

		lastNumeral = currentValue
	}

	return integer
}
