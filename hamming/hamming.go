package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strands must be of equal length")
	}

	distance := 0
	for index := range a {
		if a[index] != b[index] {
			distance++
		}
	}
	return distance, nil
}
