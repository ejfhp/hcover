package hcover_test

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/trealtamira/dynatile"
)

func TestExtentHashes(t *testing.T) {
	//input is: north, east, south, west, accuracy
	input := [][]float64{
		{47.203194, 9.294485, 45.689868, 7.366385, 3},
		{45.161760, 11.277115, 44.969711, 11.172745, 4},
		{22.504744, 45.009896, 22.495070, 44.989769, 6},
		{-44.970504, -67.473327, -45.024886, -67.531001, 5},
		{-44.996500, -67.491879, -45.009366, -67.507672, 6},
		{83.606445, -34.544769, 83.602773, -34.559361, 6},
		{0.018667, 0.030479, -0.013604, -0.025139, 5},
	}
	hashes := [][]string{
		{"u0j", "u0m", "u0n", "u0q"},
		{"spzz", "srbp", "u0pb", "u200"},
		{"sgzzzz", "supbpb", "t5bpbp", "th0000"},
		{"4rzzz", "4xbpb", "62pbp", "68000"},
		{"4rzzzy", "4rzzzz", "4xbpbn", "4xbpbp", "62pbpb", "680000"},
		{"gnz7te", "gnz7tg", "gnz7ts", "gnz7tu"},
		{"7zzzz", "ebpbp", "kpbpb", "s0000"},
	}
	for i, in := range input {
		fmt.Printf("Loop %d  %v\n", i, in)
		cover := dynatile.HashCoverage(in[0], in[1], in[2], in[3], uint(in[4]))
		sort.Strings(cover)
		if !reflect.DeepEqual(cover, hashes[i]) {
			t.Errorf("Hashes are different: \n got:      %v \n expected: %v", cover, hashes[i])
		}
	}
}
