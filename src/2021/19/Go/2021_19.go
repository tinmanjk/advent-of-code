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
	first              DimOrient
	second             DimOrient
	third              DimOrient
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
	zeroScanner.first.dimension = "x"
	zeroScanner.first.orientation = "pos"
	zeroScanner.second.dimension = "y"
	zeroScanner.second.orientation = "pos"
	zeroScanner.third.dimension = "z"
	zeroScanner.third.orientation = "pos"

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
	checkAllPairs bool, toCheckIndeces []pairIndeces) (mapDiffSliceIndeces map[int][]pairIndeces) {

	// map ot razlikata i count
	mapDiffSliceIndeces = map[int][]pairIndeces{}
	if checkAllPairs {
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

	mapFiltered := map[int][]pairIndeces{}
	for diff, indeces := range mapDiffSliceIndeces {
		if len(indeces) >= 12 {
			mapFiltered[diff] = indeces
		}
	}

	return mapFiltered
}

func checkTwoScannersOverlap(zeroBased *Scanner, other *Scanner) bool {

	compZeroOther := map[string]map[string]map[int][]pairIndeces{}
	dimensionNames := []string{
		"first", "second", "third",
	}

	compZeroOtherFiltered := map[string]map[string]map[int][]pairIndeces{}

	for z := 0; z < len(dimensionNames); z++ {
		// we dont'have matches so we compare ALL
		// at 0 -> we need to optimizer for second and third
		zeroDim := dimensionNames[z]
		compZeroOther[zeroDim] = map[string]map[int][]pairIndeces{}
		firstDimension := (z == 0)
		checkAll := firstDimension
		if firstDimension {
			for o := 0; o < len(dimensionNames); o++ {
				otherDim := dimensionNames[o]

				flipOther := false
				compZeroOther[zeroDim][otherDim] = getDiffZeroToOther12(zeroBased, other, zeroDim, otherDim,
					flipOther, checkAll, nil)

				flipOther = true
				compZeroOther[zeroDim][otherDim+"Flipped"] = getDiffZeroToOther12(zeroBased, other, zeroDim, otherDim,
					flipOther, checkAll, nil)

			}
		} else {
			for p := 0; p < z; p++ {
				previous := dimensionNames[p]
				previousDimensionFiltered := compZeroOtherFiltered[previous]

				otherDimensionUsed := map[string]bool{}
				for otherDimensionTaken, diffIndecesMap := range previousDimensionFiltered {
					for o := 0; o < len(dimensionNames); o++ {
						flipOther := false
						otherDim := dimensionNames[o]
						if otherDim == otherDimensionTaken {
							continue
						}
						for _, indecesPairs := range diffIndecesMap {
							compZeroOther[zeroDim][otherDim] = getDiffZeroToOther12(zeroBased, other, zeroDim, otherDim,
								flipOther, checkAll, indecesPairs)
							if len(compZeroOther[zeroDim][otherDim]) > 0 {
								otherDimensionUsed[otherDimensionTaken] = true
							}
						}

						flipOther = true
						if otherDim+"Flipped" == otherDimensionTaken {
							continue
						}
						for _, indecesPairs := range diffIndecesMap {
							compZeroOther[zeroDim][otherDim+"Flipped"] = getDiffZeroToOther12(zeroBased, other, zeroDim, otherDim,
								flipOther, checkAll, indecesPairs)
							if len(compZeroOther[zeroDim][otherDim+"Flipped"]) > 0 {
								otherDimensionUsed[otherDimensionTaken] = true
							}
						}
					}
					// tuka ako nqma compZeroOther[zeroDim] trqbva da go triem otherDimensionTaken

					// eto tuka
				}

				// TODO FIX
				if len(otherDimensionUsed) == 0 {
					// need to return because no match on this level
					return false
				}
				// filter out the
				// make a copy
				// previousCopied
				previousCopied := map[string]map[int][]pairIndeces{}
				for k, v := range compZeroOtherFiltered[previous] {
					if _, ok := otherDimensionUsed[k]; ok {
						previousCopied[k] = v
					}
				}

				// TODO
				if len(previousCopied) == 0 {
					return false
				}

				compZeroOtherFiltered[previous] = previousCopied
			}
		}

		// SET CcompZEROOTHERFILTERED copy of compZeroOther...s removed
		compZeroOtherFiltered[zeroDim] = map[string]map[int][]pairIndeces{}
		for otherDimension, mapDifferencesToPairs := range compZeroOther[zeroDim] {
			if len(mapDifferencesToPairs) == 0 {
				continue
			}
			compZeroOtherFiltered[zeroDim][otherDimension] = compZeroOther[zeroDim][otherDimension]
			// "firstFlipped" -> 68 -> pairIndexes
		}
		// check after first if nothing matches then bye bye
		if len(compZeroOtherFiltered[zeroDim]) == 0 {
			return false
		}
	}

	// MAKE CHECKS

	// should be just ONE
	// diff no use at the moment
	for dimension, mapDiffIndeces := range compZeroOtherFiltered["first"] {
		if strings.Contains(dimension, "Flipped") {
			other.first.orientation = "neg"
			other.first.dimension = strings.Split(dimension, "Flipped")[0]
		} else {
			other.first.orientation = "pos"
			other.first.dimension = dimension
		}
		for diff := range mapDiffIndeces {
			other.x = diff
		}

	}

	for dimension, mapDiffIndeces := range compZeroOtherFiltered["second"] {
		if strings.Contains(dimension, "Flipped") {
			other.second.orientation = "neg"
			other.second.dimension = strings.Split(dimension, "Flipped")[0]
		} else {
			other.second.orientation = "pos"
			other.second.dimension = dimension
		}

		for diff := range mapDiffIndeces {
			other.y = diff
		}
	}

	for dimension, mapDiffIndeces := range compZeroOtherFiltered["third"] {
		if strings.Contains(dimension, "Flipped") {
			other.third.orientation = "neg"
			other.third.dimension = strings.Split(dimension, "Flipped")[0]
		} else {
			other.third.orientation = "pos"
			other.third.dimension = dimension
		}

		for diff := range mapDiffIndeces {
			other.z = diff
		}
	}

	other.hasZeroCoordinates = true

	for _, beacon := range other.beacons {
		saveFirst := beacon.dimVal["first"]
		saveSecond := beacon.dimVal["second"]
		saveThird := beacon.dimVal["third"]

		switch other.first.dimension {
		case "first":
			if other.first.orientation == "neg" {
				beacon.dimVal["first"] = -saveFirst
			} else {
				beacon.dimVal["first"] = saveFirst
			}
		case "second":
			if other.first.orientation == "neg" {
				beacon.dimVal["first"] = -saveSecond
			} else {
				beacon.dimVal["first"] = saveSecond
			}
		case "third":
			if other.first.orientation == "neg" {
				beacon.dimVal["first"] = -saveThird
			} else {
				beacon.dimVal["first"] = saveThird
			}
		}

		switch other.second.dimension {
		case "first":
			if other.second.orientation == "neg" {
				beacon.dimVal["second"] = -saveFirst
			} else {
				beacon.dimVal["second"] = saveFirst
			}
		case "second":
			if other.second.orientation == "neg" {
				beacon.dimVal["second"] = -saveSecond
			} else {
				beacon.dimVal["second"] = saveSecond
			}
		case "third":
			if other.second.orientation == "neg" {
				beacon.dimVal["second"] = -saveThird
			} else {
				beacon.dimVal["second"] = saveThird
			}
		}

		switch other.third.dimension {
		case "first":
			if other.third.orientation == "neg" {
				beacon.dimVal["third"] = -saveFirst
			} else {
				beacon.dimVal["third"] = saveFirst
			}
		case "second":
			if other.third.orientation == "neg" {
				beacon.dimVal["third"] = -saveSecond
			} else {
				beacon.dimVal["third"] = saveSecond
			}
		case "third":
			if other.third.orientation == "neg" {
				beacon.dimVal["third"] = -saveThird
			} else {
				beacon.dimVal["third"] = saveThird
			}
		}

		beacon.x = other.x + beacon.dimVal["first"]
		beacon.y = other.y + beacon.dimVal["second"]
		beacon.z = other.z + beacon.dimVal["third"]

		beacon.dimVal["first"] = beacon.x
		beacon.dimVal["second"] = beacon.y
		beacon.dimVal["third"] = beacon.z

		beacon.hasZeroCoordinates = true
	}
	// rewrite beacons

	return true
}
