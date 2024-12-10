package input

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadLines(t *testing.T) {
	t.Run("Take a filename read in and return a string array of the lines", func(t *testing.T) {
		expected := []string{"test", "1", "", "2"}
		actual := ReadLines("test-file.txt")

		assert.Equal(t, expected, actual)
	})
}

func TestReadParagraphs(t *testing.T) {
	t.Run("Read in a file and return a string array of the lines, each empty line a new paragraph", func(t *testing.T) {
		expected := [][]string{{"test", "1"}, {"2"}}
		actual := ReadParagraphs("test-file.txt")

		assert.Equal(t, expected, actual)
	})
}
