package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements_InvalidSize(t *testing.T) {

	testSize := []struct {
		name string
		size int
	}{
		{"1", 0},
		{"2", -8},
		{"3", -100_000_000},
	}

	for _, w := range testSize {
		t.Run(w.name, func(t *testing.T) {
			_, err := generateRandomElements(w.size)
			require.Error(t, err)
		})
	}
}

func TestGenerateRandomElements_DifferentSize(t *testing.T) {

	testSize := []struct {
		name string
		size int
	}{
		{"1", 8},
		{"2", 5_812},
		{"3", 184_444},
		{"4", 438_939_948},
		{"5", 999_999_999},
	}

	for _, w := range testSize {
		t.Run(w.name, func(t *testing.T) {
			q, err := generateRandomElements(w.size)
			require.NoError(t, err)
			require.Len(t, q, w.size)
		})
	}
}

func TestMaximum_InvalidSize(t *testing.T) {

	testSlice := make([]int, 0)

	t.Run("test", func(t *testing.T) {
		_, err := maximum(testSlice)
		require.Error(t, err)
	})
}

func TestMaximum_DifferentSlices(t *testing.T) {

	testSlices := []struct {
		name  string
		slice []int
	}{
		{"1", []int{243, 245, 21, 463, 12, 6, 1, 75, 256}},
		{"2", []int{134, 1346, 16, 754, 1345, 7, 1456, 75}},
		{"3", []int{623, 17, 743, 542, 134, 85, 4567, 136, 865, 134, 65}},
	}

	for _, w := range testSlices {
		t.Run(w.name, func(t *testing.T) {
			number, err := maximum(w.slice)
			require.NoError(t, err)
			require.Equal(t, number, slices.Max(w.slice))
		})
	}
}
