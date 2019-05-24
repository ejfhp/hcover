package hcover

import (
	"github.com/savardiego/geohash"
)

//ExtentHashes returns a list of hashes that covers the given extent with the given accuracy.
//Accuracy is the len of the expected hash (1-12)
func ExtentHashes(north, east, south, west float64, accuracy uint) []string {
	if accuracy == 0 {
		return []string{}
	}
	hashSE := geohash.EncodeWithPrecision(south, east, accuracy)
	hashNW := geohash.EncodeWithPrecision(north, west, accuracy)
	// north side from west to east
	NWE := make([]string, 0)
	NWE = append(NWE, hashNW)
	// east side from south to north
	ESN := make([]string, 0)
	ESN = append(ESN, hashSE)
	incESN := true
	//Growing NWE and ESN to find intersection
	crossed := indexIn(NWE, ESN[0])
	for crossed < 0 {
		if incESN {
			ESN = append(ESN, geohash.Neighbor(ESN[len(ESN)-1], geohash.North))
			crossed = indexIn(NWE, ESN[len(ESN)-1])
			if crossed >= 0 {
				NWE = NWE[:crossed+1]
			}
			incESN = false
		} else {
			NWE = append(NWE, geohash.Neighbor(NWE[len(NWE)-1], geohash.East))
			crossed = indexIn(ESN, NWE[len(NWE)-1])
			if crossed >= 0 {
				ESN = ESN[:crossed+1]
			}
			incESN = true
		}
	}
	//Find row per row from NWE (base) for the len of ESN (height)
	ins := make([]string, 0, len(NWE)*len(ESN))
	ins = append(ins, NWE...)
	for _, h := range NWE {
		pre := h
		for j := 0; j < len(ESN)-1; j++ {
			n := geohash.Neighbor(pre, geohash.South)
			ins = append(ins, n)
			pre = n
		}
	}
	return ins
}

func indexIn(slice []string, val string) int {
	for i, v := range slice {
		if v == val {
			return i
		}
	}
	return -1
}
