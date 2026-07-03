package main

import (
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
