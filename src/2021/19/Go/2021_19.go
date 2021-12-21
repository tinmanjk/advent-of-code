package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	scanners := parseInput(lines)
	var result int
	// part 1
	result = findResult(scanners, false)
	fmt.Println(result)

	// part 2
	scanners = parseInput(lines)
	result = findResult(scanners, true)
	fmt.Println(result)
}

func parseInput(lines []string) (scanners []*Scanner) {

	var currentScanner *Scanner
	scanners = []*Scanner{}
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			continue
		}
		if line[0:3] == "---" {
			currentScanner = new(Scanner)
			scanners = append(scanners, currentScanner)
			continue
		}
		splitted := strings.Split(line, ",")
		firstDim, _ := strconv.Atoi(splitted[0])
		secondDim, _ := strconv.Atoi(splitted[1])
		thirdDim, _ := strconv.Atoi(splitted[2])

		beacon := Beacon{}
		beacon.dimVal = map[string]int{}
		beacon.dimVal["first"] = firstDim
		beacon.dimVal["second"] = secondDim
		beacon.dimVal["third"] = thirdDim

		currentScanner.beacons = append(currentScanner.beacons, &beacon)
	}
	return
}

type Beacon struct {
	dimVal             map[string]int
	x                  int
	y                  int
	z                  int
	hasZeroCoordinates bool
}

type DimOrient struct {
	dimension   string // x y z
	orientation string // positive/negative
}

type Scanner struct {
	beacons            []*Beacon
	dimOrientation     map[string]DimOrient
	x                  int
	y                  int
	z                  int
	hasZeroCoordinates bool
}

const inputPath = "../input.txt"

func findResult(scanners []*Scanner, partTwo bool) (result int) {

	zeroScanner := scanners[0]
	zeroScanner.x = 0
	zeroScanner.y = 0
	zeroScanner.z = 0
	zeroScanner.hasZeroCoordinates = true
	zeroScanner.dimOrientation = map[string]DimOrient{}
	zeroScanner.dimOrientation["first"] = DimOrient{"x", "pos"}
	zeroScanner.dimOrientation["second"] = DimOrient{"y", "pos"}
	zeroScanner.dimOrientation["third"] = DimOrient{"z", "pos"}

	zeroScanners := []*Scanner{zeroScanner}
	for len(zeroScanners) != len(scanners) {
	allScannerLoop:
		for i := 0; i < len(scanners); i++ {
			for _, z := range zeroScanners {
				if scanners[i] == z {
					continue allScannerLoop
				}
			}
			for j := 0; j < len(zeroScanners); j++ {
				zeroScanner = zeroScanners[j]
				otherScanner := scanners[i]
				overlapped := checkTwoScannersOverlap(zeroScanner, otherScanner)
				if overlapped {
					zeroScanners = append(zeroScanners, otherScanner)
					continue allScannerLoop
				}
			}
		}
	}

	if partTwo {
		// maxdistance
		manhattanDistance := 0
		for i := 0; i < len(zeroScanners); i++ {
			for j := 0; j < len(zeroScanners); j++ {
				if zeroScanners[i] == zeroScanners[j] {
					continue
				}
				// |x1 - x2| + |y1 - y2|.
				s1 := zeroScanners[i]
				s2 := zeroScanners[j]
				x1x2 := int(math.Abs(float64(s1.x - s2.x)))
				y1y2 := int(math.Abs(float64(s1.y - s2.y)))
				z1z2 := int(math.Abs(float64(s1.z - s2.z)))
				distance := x1x2 + y1y2 + z1z2
				if distance >= manhattanDistance {
					manhattanDistance = distance
				}
			}
		}

		return manhattanDistance
	}

	mapBeacons := map[coord]coord{}
	for _, zs := range zeroScanners {
		for _, beacon := range zs.beacons {
			beaconCoord := coord{}
			beaconCoord.x = beacon.dimVal["first"]
			beaconCoord.y = beacon.dimVal["second"]
			beaconCoord.z = beacon.dimVal["third"]
			mapBeacons[beaconCoord] = beaconCoord
		}
	}

	return len(mapBeacons)
}

type coord struct {
	x int
	y int
	z int
}

// what index corresponds to other other : zero
type pairIndeces struct {
	zero  int
	other int
}

func getDiffZeroToOther12(zeroBased *Scanner, other *Scanner,
	zDim string, oDim string, flipOther bool,
	checkAllBeaconPairs bool, toCheckIndeces []pairIndeces) (mapDiffSliceIndeces map[int][]pairIndeces) {

	// differences between two beacons on a certain dimension
	// either all beacons with all beacons permutations
	// or only check pairs of beacons
	mapDiffSliceIndeces = map[int][]pairIndeces{}
	if checkAllBeaconPairs {
		for z := 0; z < len(zeroBased.beacons); z++ {
			for o := 0; o < len(other.beacons); o++ {
				zeroBeacon := zeroBased.beacons[z]
				otherBeacon := other.beacons[o]
				pairInd := pairIndeces{z, o}
				var diff int
				if flipOther {
					diff = zeroBeacon.dimVal[zDim] - (-otherBeacon.dimVal[oDim])
				} else {
					diff = zeroBeacon.dimVal[zDim] - otherBeacon.dimVal[oDim]
				}
				if _, ok := mapDiffSliceIndeces[diff]; !ok {
					mapDiffSliceIndeces[diff] = []pairIndeces{pairInd}
				} else {
					mapDiffSliceIndeces[diff] = append(mapDiffSliceIndeces[diff], pairInd)
				}
			}
		}
	} else {
		for _, indexPair := range toCheckIndeces {
			zeroIndex := indexPair.zero
			otherIndex := indexPair.other
			zeroBeacon := zeroBased.beacons[zeroIndex]
			otherBeacon := other.beacons[otherIndex]
			var diff int
			if flipOther {
				diff = zeroBeacon.dimVal[zDim] - (-otherBeacon.dimVal[oDim])
			} else {
				diff = zeroBeacon.dimVal[zDim] - otherBeacon.dimVal[oDim]
			}
			if _, ok := mapDiffSliceIndeces[diff]; !ok {
				mapDiffSliceIndeces[diff] = []pairIndeces{indexPair}
			} else {
				mapDiffSliceIndeces[diff] = append(mapDiffSliceIndeces[diff], indexPair)
			}
		}
	}

	//https://stackoverflow.com/questions/23229975/is-it-safe-to-remove-selected-keys-from-map-within-a-range-loop
	for diff, indeces := range mapDiffSliceIndeces {
		if len(indeces) < 12 {
			delete(mapDiffSliceIndeces, diff)
		}
	}

	return
}

func checkTwoScannersOverlap(zeroBased *Scanner, other *Scanner) bool {

	compZeroOtherDimDiffs := map[string]map[string]map[int][]pairIndeces{}
	dimensionNames := []string{
		"first", "second", "third",
	}

	for z := 0; z < len(dimensionNames); z++ {
		// we dont'have matches so we compare ALL
		// at 0 -> we need to optimizer for second and third
		zeroDim := dimensionNames[z]
		compZeroOtherDimDiffs[zeroDim] = map[string]map[int][]pairIndeces{}
		firstDimension := (z == 0)
		checkAllBeacons := firstDimension
		if firstDimension {
			for o := 0; o < len(dimensionNames); o++ {
				otherDim := dimensionNames[o]

				flipOther := false
				diffToIndexPairs := getDiffZeroToOther12(zeroBased, other,
					zeroDim, otherDim,
					flipOther, checkAllBeacons, nil)
				if len(diffToIndexPairs) != 0 {
					compZeroOtherDimDiffs[zeroDim][otherDim] = diffToIndexPairs
				}

				flipOther = true
				diffToIndexPairs = getDiffZeroToOther12(zeroBased, other,
					zeroDim, otherDim,
					flipOther, checkAllBeacons, nil)
				if len(diffToIndexPairs) != 0 {
					compZeroOtherDimDiffs[zeroDim][otherDim+"Flipped"] = diffToIndexPairs
				}
			}
			goto CheckFound
		}

		for p := 0; p < z; p++ {
			previous := dimensionNames[p]
			// We need to check the candidate dimensions found for the previous dimension
			// e.g. "first" and "second flipped"
			// each of those can have 1 or more diffs that have 12 indeces of beacon pairs
			// unlikely but possible

			// if the pairs of a prev candidate work for this dimension
			// and produce a differently named dimension i.e. "third flipped"
			// we mark the previous dimensions as Used (i.e. "second flipped")

			prevOtherDimensionUsed := map[string]bool{} // see check below loop
			for otherDimTakenByPrev, diffIndecesMap := range compZeroOtherDimDiffs[previous] {
				for o := 0; o < len(dimensionNames); o++ {
					otherDim := dimensionNames[o]
					if otherDim == otherDimTakenByPrev {
						continue
					}

					flipOther := false
					// TODO mark Diff used or NOT
					for _, indecesPairs := range diffIndecesMap {
						diffToIndexPairs := getDiffZeroToOther12(zeroBased, other,
							zeroDim, otherDim,
							flipOther, checkAllBeacons, indecesPairs)
						if len(diffToIndexPairs) != 0 {
							compZeroOtherDimDiffs[zeroDim][otherDim] = diffToIndexPairs
							prevOtherDimensionUsed[otherDimTakenByPrev] = true
						}
					}

					if otherDim+"Flipped" == otherDimTakenByPrev {
						continue
					}

					flipOther = true
					for _, indecesPairs := range diffIndecesMap {
						diffToIndexPairs := getDiffZeroToOther12(zeroBased, other,
							zeroDim, otherDim,
							flipOther, checkAllBeacons, indecesPairs)
						if len(diffToIndexPairs) != 0 {
							compZeroOtherDimDiffs[zeroDim][otherDim+"Flipped"] = diffToIndexPairs
							prevOtherDimensionUsed[otherDimTakenByPrev] = true
						}
					}
				}
			}

			// none of the pairs from the previous dimension worked for the current one
			if len(prevOtherDimensionUsed) == 0 {
				return false
			}
			// filter out the
			// non-used dimensions for previous as they cannot be chained to the current
			for otherDimensionCandidate := range compZeroOtherDimDiffs[previous] {
				if _, ok := prevOtherDimensionUsed[otherDimensionCandidate]; !ok {
					delete(compZeroOtherDimDiffs[previous], otherDimensionCandidate)
				}
			}
		}

	CheckFound:
		// no other dimension found for the current zeroDim
		if len(compZeroOtherDimDiffs[zeroDim]) == 0 {
			return false
		}
	}

	// dimensionNames -> "first, second, third"
	for _, dimName := range dimensionNames {
		// only one dimension candidate left
		if len(compZeroOtherDimDiffs[dimName]) != 1 {
			return false // cannot determine a single other dimension for each zero dimension
		}

		// within a dimension there is only ONE diff
		// only one diff left TODO
		// counter := 0
		// for _, _ := range compZeroOtherDimDiffs[dimName] {
		// 	counter++
		// }

	}

	// set zerobased coordinations of the other scanner
	other.dimOrientation = map[string]DimOrient{}
	for _, dimName := range dimensionNames {
		// should be just one dimension -> no way of accessing it
		// TODO - find out a better way for First or Default then iterating here
		for dimension, mapDiffIndeces := range compZeroOtherDimDiffs[dimName] {
			if strings.Contains(dimension, "Flipped") {
				dimensionOnly := strings.Split(dimension, "Flipped")[0]
				other.dimOrientation[dimName] = DimOrient{dimensionOnly, "neg"}
			} else {
				other.dimOrientation[dimName] = DimOrient{dimension, "pos"}
			}

			// see above that only ONE diff
			diffFound := math.MaxInt32
			// len of mapDifferences is 1
			for diff := range mapDiffIndeces {
				diffFound = diff
			}

			switch dimName {
			case "first":
				other.x = diffFound
			case "second":
				other.y = diffFound
			case "third":
				other.z = diffFound
			}
		}
	}
	other.hasZeroCoordinates = true

	// other scanner beacons zero coordinates
	for _, beacon := range other.beacons {
		saveFirst := beacon.dimVal["first"]
		saveSecond := beacon.dimVal["second"]
		saveThird := beacon.dimVal["third"]

		for _, dimName := range dimensionNames {
			switch other.dimOrientation[dimName].dimension {
			case "first":
				if other.dimOrientation[dimName].orientation == "neg" {
					beacon.dimVal[dimName] = -saveFirst
				} else {
					beacon.dimVal[dimName] = saveFirst
				}
			case "second":
				if other.dimOrientation[dimName].orientation == "neg" {
					beacon.dimVal[dimName] = -saveSecond
				} else {
					beacon.dimVal[dimName] = saveSecond
				}
			case "third":
				if other.dimOrientation[dimName].orientation == "neg" {
					beacon.dimVal[dimName] = -saveThird
				} else {
					beacon.dimVal[dimName] = saveThird
				}
			}
		}

		// zero-based coordinates for beacons
		beacon.dimVal["first"] += other.x
		beacon.dimVal["second"] += other.y
		beacon.dimVal["third"] += other.z
		beacon.x = beacon.dimVal["first"]
		beacon.y = beacon.dimVal["second"]
		beacon.z = beacon.dimVal["third"]
		beacon.hasZeroCoordinates = true
	}

	return true
}
