package utils

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestBubbleSortBestCase(t *testing.T) {
	//Initialization
	els := []int{5, 6, 7, 8, 9}

	//Execution
	els = BubbleSort(els)

	//Validation
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func TestBubbleSortWorstCase(t *testing.T) {
	//Initialization
	els := []int{9, 8, 7, 6, 5}

	//Execution
	els = BubbleSort(els)

	//Validation
	assert.NotNil(t, els)
	assert.EqualValues(t, 5, len(els))
	assert.EqualValues(t, 5, els[0])
	assert.EqualValues(t, 6, els[1])
	assert.EqualValues(t, 7, els[2])
	assert.EqualValues(t, 8, els[3])
	assert.EqualValues(t, 9, els[4])
}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func BenchmarkBubbleSort100(b *testing.B) {
	els := getElements(100)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort100(b *testing.B) {
	els := getElements(100)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort1000(b *testing.B) {
	els := getElements(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	els := getElements(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}

func BenchmarkSort10000(b *testing.B) {
	els := getElements(10000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}
