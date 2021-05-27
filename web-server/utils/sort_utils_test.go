package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortWorstCase(t *testing.T) {
	// initialization
	els := []int{10, 8, 7, 6, 5}

	// execution
	BubbleSort(els)

	// validation
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 10, els[4])

}

func TestBubbleSortBestCase(t *testing.T) {
	// initialization
	els := []int{5, 6, 7, 8, 11}

	// execution
	BubbleSort(els)

	// validation
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 11, els[4])

}

func TestBubbleSortNilSlice(t *testing.T) {
	BubbleSort(nil)
}
