package rounding

import (
	"math"
	"strconv"
	"strings"
)

// RoundNumber rounds a number based on its characteristics.
func RoundNumber(num float64) float64 {
	// Convert the number to a string
	strNum := strconv.FormatFloat(num, 'e', -1, 64)

	// Check if the string representation contains "e" (scientific notation)
	if strings.Contains(strNum, "e") {
		return RoundToScientificNotation(num)
	}

	// Check if the number has leading zeros
	if strings.HasPrefix(strNum, "0.") {
		return RoundToScientificNotation(num)
	}

	// Default to rounding to decimal places
	return RoundToDecimalPlaces(num)
}

// RoundToDecimalPlaces rounds a number to the appropriate number of decimal places.
func RoundToDecimalPlaces(num float64) float64 {
	if num == 0 {
		return 0
	}
	// Round the number to 2 decimal places
	rounded := math.Round(num*100) / 100
	return rounded
}

// RoundToScientificNotation rounds a number to significant figures and represents it in scientific notation.
func RoundToScientificNotation(num float64) float64 {
	if num == 0 {
		return 0
	}
	// Round the number to significant figures
	rounded := RoundToSignificantFigures(num)
	// Convert to scientific notation
	return rounded
}

// RoundToSignificantFigures rounds a number to significant figures.
func RoundToSignificantFigures(num float64) float64 {
	if num == 0 {
		return 0
	}
	// Calculate the magnitude of the number
	magnitude := math.Floor(math.Log10(math.Abs(num))) + 1
	// Round the number to significant figures
	rounded := math.Round(num*math.Pow(10, 3-magnitude)) / math.Pow(10, 3-magnitude)
	return rounded
}
