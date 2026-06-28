package raindrops

import "fmt"

func Convert(number int) string {
	output := ""

	if number%3 == 0 {
		output += "Pling"
	}
	if number%5 == 0 {
		output += "Plang"
	}
	if number%7 == 0 {
		output += "Plong"
	}

	if output == "" {
		return fmt.Sprintf("%d", number)
	}
	return output
}
