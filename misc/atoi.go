package misc

import "fmt"

var NotAnInt = -999

func Atoi(input string) (int, error) {
	//Check empty
	if input == "" {
		return NotAnInt, fmt.Errorf("Can not convert empty string")
	}

	sign := 1
	if input[0] == '-' {
		sign = -1
		input = input[1:]
	}

	result := 0

	for _, val := range input {
		diff := int(val - '0')

		if diff > 9 || diff < 0 {
			return NotAnInt, fmt.Errorf("Can not convert ill formatted string: %q", val)
		}
		result = result*10 + diff
	}
	return sign * result, nil

}
