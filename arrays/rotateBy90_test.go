package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Only applicable to Square matrices
func Test_RotateMatrixby90(t *testing.T) {
	assert := assert.New(t)

	input := make(TwoDArray, 4)
	input[0] = Rows{1, 2, 3, 4}
	input[1] = Rows{5, 6, 7, 8}
	input[2] = Rows{9, 10, 11, 12}
	input[3] = Rows{13, 14, 15, 16}

	output := make(TwoDArray, 4)
	output[0] = Rows{13, 9, 5, 1}
	output[1] = Rows{14, 10, 6, 2}
	output[2] = Rows{15, 11, 7, 3}
	output[3] = Rows{16, 12, 8, 4}

	arr := NewTwoDArray(input)
	arr.PrintArr()

	arr.RotateByNinety()

	arr.PrintArr()
	assert.Equal(input, output, "Array clockwise rotation should be equal")
}

func Test_RotateMatrixby903x3(t *testing.T) {
	assert := assert.New(t)

	input := make(TwoDArray, 3)
	input[0] = Rows{1, 2, 3}
	input[1] = Rows{4, 5, 6}
	input[2] = Rows{7, 8, 9}

	output := make(TwoDArray, 3)
	output[0] = Rows{7, 4, 1}
	output[1] = Rows{8, 5, 2}
	output[2] = Rows{9, 6, 3}

	arr := NewTwoDArray(input)
	arr.PrintArr()

	arr.RotateByNinety()

	arr.PrintArr()
	assert.Equal(input, output, "Array clockwise rotation should be equal for 3x3 as well")
}
