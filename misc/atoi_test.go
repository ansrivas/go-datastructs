package misc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Atoi(t *testing.T) {
	assert := assert.New(t)
	actual, _ := Atoi("23")
	expected := 23
	assert.Equal(expected, actual, "23 string representation should be converted to int")

	actual, _ = Atoi("-23")
	expected = -23
	assert.Equal(expected, actual, "-23 string representation should be converted to int -23")

	_, err := Atoi("2-3")
	assert.NotNil(err, "Ill formed strings should throw error.")

	_, err = Atoi("")
	assert.NotNil(err, "Empty string should throw error.")
}
