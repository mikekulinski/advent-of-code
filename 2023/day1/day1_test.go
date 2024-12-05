package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseNumbers(t *testing.T) {
	input := []string{
		"eightwothree",
		"xtwone3four",
		"oneeight",
	}
	expected := []string{
		"823",
		"2134",
		"18",
	}
	for i := range input {
		actual := parseNumbers(input[i])
		assert.Equal(t, expected[i], actual)
	}
}
