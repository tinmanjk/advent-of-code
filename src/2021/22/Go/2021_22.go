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
	cuboids := parseInput(lines)

	// part 1
	result := part1(cuboids)
	fmt.Println("Part 1:", result)
	// part 2
	bigResult := part2(cuboids)
	fmt.Println("Part 2:", bigResult)

}

func parseInput(lines []string) (instrCuboids []InstructionCuboid) {
	instrCuboids = []InstructionCuboid{}

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lineSplitted := strings.Split(line, ",")
		instructionAndx := strings.Split(lineSplitted[0], " ")
		onOff := instructionAndx[0]
		x1x2 := strings.Split(strings.Split(instructionAndx[1], "x=")[1], "..")
		x1, _ := strconv.Atoi(x1x2[0])
		x2, _ := strconv.Atoi(x1x2[1])

		y1y2 := strings.Split(strings.Split(lineSplitted[1], "y=")[1], "..")
		y1, _ := strconv.Atoi(y1y2[0])
		y2, _ := strconv.Atoi(y1y2[1])

		z1z2 := strings.Split(strings.Split(lineSplitted[2], "z=")[1], "..")
		z1, _ := strconv.Atoi(z1z2[0])
		z2, _ := strconv.Atoi(z1z2[1])

		// TODO FIX

		cuboid := Cuboid{x1, x2, y1, y2, z1, z2}
		var on bool
		if onOff == "on" {
			on = true
		} else {
			on = false
		}
		instrCuboid := InstructionCuboid{on, cuboid}
		instrCuboids = append(instrCuboids, instrCuboid)
	}
	return
}

const inputPath = "../input.txt"

type InstructionCuboid struct {
	on     bool
	cuboid Cuboid
}

type Cube struct {
	x int
	y int
	z int
}

type Cuboid struct {
	x1 int
	x2 int
	y1 int
	y2 int
	z1 int
	z2 int
}

func part1(instrCuboids []InstructionCuboid) (result int) {

	cubes := map[Cube]Cube{}

	for i := 0; i < len(instrCuboids); i++ {
		cuboid := instrCuboids[i].cuboid
		var x1, x2, y1, y2, z1, z2 int
		// x
		if cuboid.x1 > 50 {
			continue
		}
		if cuboid.x1 < -50 {
			x1 = -50
		} else {
			x1 = cuboid.x1
		}

		if cuboid.x2 > 50 {
			x2 = 50
		} else {
			x2 = cuboid.x2
		}
		// y
		if cuboid.y1 > 50 {
			continue
		}
		if cuboid.y1 < -50 {
			y1 = -50
		} else {
			y1 = cuboid.y1
		}

		if cuboid.y2 > 50 {
			y2 = 50
		} else {
			y2 = cuboid.y2
		}

		// z
		if cuboid.z1 > 50 {
			continue
		}
		if cuboid.z1 < -50 {
			z1 = -50
		} else {
			z1 = cuboid.z1
		}

		if cuboid.z2 > 50 {
			z2 = 50
		} else {
			z2 = cuboid.z2
		}

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				for z := z1; z <= z2; z++ {
					cube := Cube{x, y, z}
					if instrCuboids[i].on {
						cubes[cube] = cube
					} else {
						delete(cubes, cube)
					}
				}
			}
		}
	}

	return len(cubes)
}

func part2(instrCuboids []InstructionCuboid) (result uint64) {

	// start with on cuboid as starting off cuboids do nothing
	foundOnCuboid := false
	for i := 0; i < len(instrCuboids); i++ {
		if instrCuboids[i].on {
			instrCuboids = instrCuboids[i:]
			foundOnCuboid = true
			break
		}
	}
	if len(instrCuboids) == 0 || !foundOnCuboid {
		return 0
	}

	// just one on
	if len(instrCuboids) == 1 {
		return calculateVolume(instrCuboids[0].cuboid)
	}

	// guaranteed first on and at least two
	nonOverlappingCuboids := map[Cuboid]Cuboid{}
	firstCuboid := instrCuboids[0].cuboid
	nonOverlappingCuboids[firstCuboid] = firstCuboid

	for i := 1; i < len(instrCuboids); i++ {
		instrCuboid := instrCuboids[i]
		if instrCuboid.on {
			rollingRightDiffWithNoOverlaps := map[Cuboid]Cuboid{}
			rollingRightDiffWithNoOverlaps[instrCuboid.cuboid] = instrCuboid.cuboid

			for _, nonOverLap := range nonOverlappingCuboids {
				tempRolling := map[Cuboid]Cuboid{}
				for right := range rollingRightDiffWithNoOverlaps {
					diffRights := diffWithRight(nonOverLap, right)
					for _, diffRight := range diffRights {
						tempRolling[diffRight] = diffRight
					}
				}
				rollingRightDiffWithNoOverlaps = tempRolling
			}

			for diffRight := range rollingRightDiffWithNoOverlaps {
				nonOverlappingCuboids[diffRight] = diffRight
			}

		} else { // off
			newNonOverlappingCuboids := map[Cuboid]Cuboid{}
			for _, nonOverLap := range nonOverlappingCuboids {
				diffRights := diffWithRight(instrCuboid.cuboid, nonOverLap)
				for _, diffRight := range diffRights {
					newNonOverlappingCuboids[diffRight] = diffRight
				}
			}
			nonOverlappingCuboids = newNonOverlappingCuboids
		}
	}

	for cuboid := range nonOverlappingCuboids {
		volume := calculateVolume(cuboid)
		result += volume
	}

	return
}

func diffWithRight(left Cuboid, right Cuboid) (diffRights []Cuboid) {

	ol := createOverlapCuboid(left, right)
	if ol == nil {
		return []Cuboid{right}
	}

	// diffing with the overlap
	// working to exclude evyrthing from right (or the input) that is not in the overlap
	in := right // in = input
	xDiffBottom := in.x1 < ol.x1 && ol.x2 == in.x2
	yDiffBottom := in.y1 < ol.y1 && ol.y2 == in.y2
	zDiffBottom := in.z1 < ol.z1 && ol.z2 == in.z2

	xNoDiff := in.x1 == ol.x1 && ol.x2 == in.x2
	yNoDiff := in.y1 == ol.y1 && ol.y2 == in.y2
	zNoDiff := in.z1 == ol.z1 && ol.z2 == in.z2

	xDiffTop := in.x1 == ol.x1 && ol.x2 < in.x2
	yDiffTop := in.y1 == ol.y1 && ol.y2 < in.y2
	zDiffTop := in.z1 == ol.z1 && ol.z2 < in.z2

	xDiffTopBottom := in.x1 < ol.x1 && ol.x2 < in.x2
	yDiffTopBottom := in.y1 < ol.y1 && ol.y2 < in.y2
	zDiffTopBottom := in.z1 < ol.z1 && ol.z2 < in.z2

	dBx1 := in.x1
	dBx2 := ol.x1 - 1
	dTx1 := ol.x2 + 1
	dTx2 := in.x2

	dBy1 := in.y1
	dBy2 := ol.y1 - 1
	dTy1 := ol.y2 + 1
	dTy2 := in.y2

	dBz1 := in.z1
	dBz2 := ol.z1 - 1
	dTz1 := ol.z2 + 1
	dTz2 := in.z2

	// 0 / 1 / 1 / 2
	// 1 / 3 / 3 / 5
	// 1 / 3 / 3 / 5
	// 2 / 5 / 5 / 8
	result := []Cuboid{}
	if xNoDiff {
		// 1. equal - whole line - DONE - 4/64
		// 0 / 1 / 1 / 2
		if yNoDiff {

			// 0 diffs
			if zNoDiff {
				return result
				// no differences from right - nothing to add
				// right is entirely overlapped see above already handled
			}
			// 1 diff
			if zDiffTop {

				diff1 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}
				result = append(result, diff1)
			}
			// 1 diff
			if zDiffBottom {
				diff1 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				result = append(result, diff1)
			}
			// 2 diff
			if zDiffTopBottom {
				diff1 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff2 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}
				result = append(result, diff1, diff2)
			}

			return result
		}

		// 2. leftbound to the middle somewhere - DONE - 8/64
		// 1 / 3 / 3 / 5
		if yDiffTop {

			// 1 diff
			if zNoDiff {

				diff1 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}
				return []Cuboid{diff1}
			}

			// 3 diffs
			if zDiffTop {

				// by y
				diff1 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, dTz2}

				// by z
				diff3 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff3}
			}

			// 3 diffs
			// 3. rightbound to middle somewhere
			if zDiffBottom {

				// 2 by Y
				diff1 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, ol.z2}

				// 1 by z
				diff3 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff1, diff3}
			}

			// 5 diffs
			// 4. center somehow
			if zDiffTopBottom {

				// 3 by Y
				diff1 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, dTz2}

				// 2 by z
				diff4 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff5 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff4, diff5}
			}
		}

		// 3. rightbound to middle somewhere DONE - 12 / 64
		// 1 / 3 / 3 / 5
		if yDiffBottom {

			// 1 diff
			if zNoDiff {

				diff1 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, ol.z2}
				return []Cuboid{diff1}
			}

			// 3 diffs
			// 2. leftbound to the middle somewhere
			if zDiffTop {

				// 2 combinations for yNoOverlap - outer if
				diff1 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, dTz2}

				// 1 combination for zNoOverlap s y1
				diff3 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff3}
			}

			// 3 diffs
			// 3. rightbound to middle somewhere
			if zDiffBottom {

				// 2 combinations for yNoOverlap - outer if
				diff1 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, ol.z2}

				// 1 combination for zNoOverlap s y1
				diff3 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff1, diff3}
			}

			// 5 diffs
			// 4. center somehow
			if zDiffTopBottom {

				// 3 combination for y1noLT
				diff1 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, dTz2}

				// 2 combination for overlapped.y1
				diff4 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff5 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff4, diff5}
			}
		}

		// 4. center somehow DONE 16/64
		// 2 / 5 / 5 / 8
		if yDiffTopBottom {

			// 2 diffs
			if zNoDiff {
				// 2 diffs
				diff1 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, ol.z2}
				diff2 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}

				return []Cuboid{diff1, diff2}
			}

			// 5 diffs
			// 2. leftbound to the middle somewhere
			if zDiffTop {

				// should be 5 diffs
				// all combinations for y1nolT and y2nolB with zol and no t
				diff1 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, dTz2}

				diff3 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, dTz2}

				// last combination

				diff5 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff3, diff5}
			}

			// 5 diffs
			// 3. rightbound to middle somewhere
			if zDiffBottom {
				// again should be 5 diffs

				// should be 5 diffs
				// all combinations for y1nolT and y2nolB with zol and no t
				diff2 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, ol.z2}
				diff4 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, ol.z2}

				// last combination

				diff5 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff2, diff4, diff5}
			}

			// 8 diffs
			// 4. center somehow
			if zDiffTopBottom {
				// we have 2 centers // 4 total different
				// 8 combinations basically

				// pochvame s y1nolb
				diff1 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, dTz2}
				// sled tova y1nolT
				diff4 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, dTz2}

				//sega e s dvete sredni na z1 pri overlap

				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff8 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff4, diff7, diff8}
			}
		}
	}

	// 1 / 3 / 3 / 5
	// 3 / 7 / 7 / 11
	// 3 / 7 / 7 / 11
	// 5 / 11 / 11 / 17
	if xDiffTop {
		// 1 / 3 / 3 / 5 DONE 20/64
		if yNoDiff {
			// 1 diff
			if zNoDiff {

				diff1 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, ol.z1, ol.z2}
				return []Cuboid{diff1}
			}
			// 3 diff
			if zDiffTop {

				// 2 combinations for yNoOverlap - outer if
				diff1 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, ol.z1, ol.z2}
				diff2 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, dTz1, dTz2}
				// 1 combination for zNoOverlap s y1

				diff3 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff2, diff3}
				// ot x.a pochvame
			}
			// 3 diff.a
			if zDiffBottom {

				// 2 combinations for yNoOverlap - outer if
				diff1 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, ol.z1, ol.z2}
				diff2 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, dBz1, dBz2}
				// 1 combination for zNoOverlap s y1

				diff3 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff1, diff2, diff3}
			}
			// 5 diff.a
			if zDiffTopBottom {

				// 3 combination for y1noLT
				diff1 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, dBz1, dBz2}
				diff2 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, ol.z1, ol.z2}
				diff3 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, dTz1, dTz2}
				// 2 combination for overlapped.y1

				diff4 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff5 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff2, diff3, diff4, diff5}
			}
		}

		// 3 / 7 / 7 / 11 DONE 24/64
		if yDiffTop {
			// 3
			if zNoDiff {

				// 2 combinations for yNoOverlap - outer if
				diff1 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, ol.z1, ol.z2}
				diff2 := Cuboid{dTx1, dTx2, dTy1, dTy2, ol.z1, ol.z2}
				// 1 combination for zNoOverlap s y1

				diff3 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}

				return []Cuboid{diff1, diff2, diff3}

			}

			// 7
			if zDiffTop {

				// 1 po x
				diff1 := Cuboid{dTx1, dTx2, ol.y1, dTy2, ol.z1, dTz2}

				// 2. po y
				diff5 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}
				diff6 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dTz1, dTz2}

				// 3. po z posledno
				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff5, diff6, diff7}

			}

			// 7
			if zDiffBottom {

				// 1 po x
				diff2 := Cuboid{dTx1, dTx2, ol.y1, dTy2, dBz1, ol.z2}

				// 2. po y
				diff6 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, dBz2}
				diff5 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}

				// 3. po z posledno
				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff2, diff5, diff6, diff7}

			}

			// 11 cases
			// 4. center somehow the MOST
			if zDiffTopBottom {

				// 1 po x -> 6 x1nolT is fixed
				diff1 := Cuboid{dTx1, dTx2, ol.y1, dTy2, dBz1, dTz2}

				// 2. po y -> 3 y1nolT is fixed
				diff7 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, dTz2}

				// 3. po z -> imame 2
				diff10 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff7, diff10, diff11}

			}
		}

		// 3 / 7 / 7 / 11 -> Done 28 / 64
		// 3. rightbound to middle somewhere
		if yDiffBottom {
			// 3
			if zNoDiff {

				// 2 combinations for yNoOverlap - outer if
				diff2 := Cuboid{dTx1, dTx2, dBy1, ol.y2, ol.z1, ol.z2}
				// 1 combination for zNoOverlap s y1

				diff3 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, ol.z2}

				return []Cuboid{diff2, diff3}

			}

			// 7
			if zDiffTop {
				// 1 po x
				diff3 := Cuboid{dTx1, dTx2, dBy1, ol.y2, ol.z1, dTz2}

				// 2. po y
				diff5 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, dTz2}

				// 3. po z posledno
				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff3, diff5, diff7}

			}

			// 7
			if zDiffBottom {
				// the same as before

				// 1 po x
				diff1 := Cuboid{dTx1, dTx2, dBy1, ol.y2, dBz1, ol.z2}

				// 2. po y
				diff5 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, ol.z2}

				// 3. po z posledno
				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff1, diff5, diff7}

			}

			// 11 cases
			if zDiffTopBottom {

				// 1 po x -> 6 x1nolT is fixed
				diff1 := Cuboid{dTx1, dTx2, dBy1, ol.y2, dBz1, dTz2}

				// 2. po y -> 3 y1nolT is fixed
				diff7 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, dTz2}

				// 3. po z -> imame 2
				diff10 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff7, diff10, diff11}

			}
		}

		// 5 / 11 / 11 / 17  Done  32/64
		// 4. center somehow
		if yDiffTopBottom {
			// 5
			if zNoDiff {

				// po x - 3
				diff1 := Cuboid{dTx1, dTx2, dBy1, dTy2, ol.z1, ol.z2}
				// po y - 2
				diff4 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, ol.z2}
				diff5 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}
				// z -> nqma zashtoto e sashtoto

				return []Cuboid{diff1, diff4, diff5}
			}

			// 11
			if zDiffTop {

				// po x - 6
				diff1 := Cuboid{dTx1, dTx2, dBy1, dTy2, ol.z1, dTz2}
				// po y - 4 - only  overlapped.x1 available
				diff7 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, dTz2}
				diff8 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, dTz2}
				// z -> 1 -> only x and y overlaps
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff7, diff8, diff11}

			}

			// 11
			if zDiffBottom {
				// po x - 6
				diff1 := Cuboid{dTx1, dTx2, dBy1, dTy2, dBz1, ol.z2}
				// po y - 4 - only  overlapped.x1 available
				diff7 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, ol.z2}
				diff8 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, ol.z2}

				// z -> 1 -> only x and y overlaps
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff1, diff7, diff8, diff11}
			}

			// 17
			if zDiffTopBottom {
				// po x - 9 - all y and all z
				diff1 := Cuboid{dTx1, dTx2, dBy1, dTy2, dBz1, dTz2}

				// po y - 6 - only  overlapped.x1 available
				diff10 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, dTz2}
				diff13 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, dTz2}

				// z -> 2 -> only x and y overlaps
				diff16 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff17 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff10, diff13, diff16, diff17}
			}
		}
	}

	// 1 / 3 / 3 / 5
	// 3 / 7 / 7 / 11
	// 3 / 7 / 7 / 11
	// 5 / 11 / 11 / 17
	if xDiffBottom {
		// 1 / 3 / 3 / 5 DONE 20/64
		if yNoDiff {
			// 1 diff
			if zNoDiff {

				diff1 := Cuboid{dBx1, dBx2, ol.y1, ol.y2, ol.z1, ol.z2}
				return []Cuboid{diff1}
			}
			// 3 diff.a
			if zDiffTop {

				// 2 combinations for yNoOverlap - outer if
				diff1 := Cuboid{dBx1, dBx2, ol.y1, ol.y2, ol.z1, dTz2}
				// 1 combination for zNoOverlap s y1

				diff3 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff3}
			}

			// 3 diff.a
			if zDiffBottom {

				// 2 combinations for yNoOverlap - outer if
				diff2 := Cuboid{dBx1, dBx2, ol.y1, ol.y2, dBz1, ol.z2}
				// 1 combination for zNoOverlap s y1

				diff3 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff2, diff3}
			}

			// 5 diff.a
			if zDiffTopBottom {

				// 3 combination for y1noLT
				diff1 := Cuboid{dBx1, dBx2, ol.y1, ol.y2, dBz1, dTz2}
				// 2 combination for overlapped.y1

				diff4 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff5 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff4, diff5}
			}
		}

		if yDiffTop {

			// 3
			if zNoDiff {

				// 2 combinations for yNoOverlap - outer if
				diff1 := Cuboid{dBx1, dBx2, ol.y1, dTy2, ol.z1, ol.z2}
				// 1 combination for zNoOverlap s y1

				diff3 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}

				return []Cuboid{diff1, diff3}

			}

			// 7
			if zDiffTop {
				// should be A LOT

				// 1 po x
				diff1 := Cuboid{dBx1, dBx2, ol.y1, dTy2, ol.z1, dTz2}

				// 2. po y
				diff5 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, dTz2}

				// 3. po z posledno
				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff5, diff7}

			}

			// 7
			if zDiffBottom {
				// 1 po x
				diff2 := Cuboid{dBx1, dBx2, ol.y1, dTy2, dBz1, ol.z2}

				// 2. po y
				diff6 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, ol.z2}

				// 3. po z posledno
				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff2, diff6, diff7}

			}

			// 11 cases
			if zDiffTopBottom {
				// 1 po x -> 6 x1nolT is fixed
				diff1 := Cuboid{dBx1, dBx2, ol.y1, dTy2, dBz1, dTz2}

				// 2. po y -> 3 y1nolT is fixed
				diff7 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, dTz2}

				// 3. po z -> imame 2
				diff10 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff7, diff10, diff11}

			}
		}

		// 3 / 7 / 7 / 11 -> Done 28 / 64
		if yDiffBottom {
			// 3
			if zNoDiff {
				// 2 combinations for yNoOverlap - outer if
				diff2 := Cuboid{dBx1, dBx2, dBy1, ol.y2, ol.z1, ol.z2}
				// 1 combination for zNoOverlap s y1

				diff3 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, ol.z2}

				return []Cuboid{diff2, diff3}

			}

			// 7
			if zDiffTop {
				// 1 po x
				diff1 := Cuboid{dBx1, dBx2, dBy1, ol.y2, ol.z1, dTz2}

				// 2. po y
				diff5 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, dTz2}

				// 3. po z posledno
				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff5, diff7}

			}

			// 7
			if zDiffBottom {
				// 1 po x
				diff1 := Cuboid{dBx1, dBx2, dBy1, ol.y2, dBz1, ol.z2}

				// 2. po y
				diff6 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, ol.z2}

				// 3. po z posledno
				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff1, diff6, diff7}

			}

			// 11 cases
			if zDiffTopBottom {

				// 1 po x -> 6 x1nolT is fixed
				diff4 := Cuboid{dBx1, dBx2, dBy1, ol.y2, dBz1, dTz2}

				// 2. po y -> 3 y1nolT is fixed
				diff7 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, dTz2}

				// 3. po z -> imame 2
				diff10 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff4, diff7, diff10, diff11}

			}
		}

		// 5 / 11 / 11 / 17  Done  32/64
		if yDiffTopBottom {

			// 5
			if zNoDiff {
				// po x - 3
				diff1 := Cuboid{dBx1, dBx2, dBy1, dTy2, ol.z1, ol.z2}
				// po y - 2
				diff4 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, ol.z2}
				diff5 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}
				// z -> nqma zashtoto e sashtoto

				return []Cuboid{diff1, diff4, diff5}
			}

			// 11
			if zDiffTop {
				// po x - 6
				diff1 := Cuboid{dBx1, dBx2, dBy1, dTy2, ol.z1, dTz2}

				// po y - 4 - only  overlapped.x1 available
				diff7 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, dTz2}

				diff8 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, dTz2}
				// z -> 1 -> only x and y overlaps
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff7, diff8, diff11}

			}

			// 11
			if zDiffBottom {

				// po x - 6
				diff1 := Cuboid{dBx1, dBx2, dBy1, dTy2, dBz1, ol.z2}
				// po y - 4 - only  overlapped.x1 available
				diff7 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, ol.z2}

				diff8 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, ol.z2}

				// z -> 1 -> only x and y overlaps
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff1, diff7, diff8, diff11}
			}

			// 17
			if zDiffTopBottom {

				// po x - 9 - all y and all z
				diff1 := Cuboid{dBx1, dBx2, dBy1, dTy2, dBz1, dTz2}

				// po y - 6 - only  overlapped.x1 available
				diff10 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, dTz2}

				diff13 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, dTz2}

				// z -> 2 -> only x and y overlaps
				diff16 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff17 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff10, diff13, diff16, diff17}
			}
		}
	}

	// 2 / 5 / 5 / 8
	// 5 / 11 / 11 / 17
	// 5 / 11 / 11 / 17
	// 8 / 17 / 17 / 26
	if xDiffTopBottom {

		// 2 / 5 / 5 / 8 -> DONE 52/64
		if yNoDiff {
			// 2
			if zNoDiff {

				diff1 := Cuboid{dBx1, dBx2, ol.y1, ol.y2, ol.z1, ol.z2}
				diff2 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, ol.z1, ol.z2}

				return []Cuboid{diff1, diff2}
			}

			// 5
			if zDiffTop {
				// parvo po x -> 4
				diff1 := Cuboid{dBx1, dBx2, ol.y1, ol.y2, ol.z1, dTz2}
				diff2 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, ol.z1, dTz2}

				// posle po z because y is stationary
				diff5 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff2, diff5}
			}

			// 5
			if zDiffBottom {

				// parvo po x -> 4
				diff1 := Cuboid{dBx1, dBx2, ol.y1, ol.y2, dBz1, ol.z2}
				diff2 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, dBz1, ol.z2}

				// posle po z because y is stationary
				diff5 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff1, diff2, diff5}
			}

			// 8
			if zDiffTopBottom {

				// parvo po x -> 6
				diff1 := Cuboid{dBx1, dBx2, ol.y1, ol.y2, dBz1, dTz2}
				diff2 := Cuboid{dTx1, dTx2, ol.y1, ol.y2, dBz1, dTz2}

				// posle po z -> 2 because y is stationary
				diff7 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff8 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff2, diff7, diff8}
			}
		}

		// 5 / 11 / 11 / 17 -> DONE 56/64
		if yDiffTop {

			// 5
			if zNoDiff {
				// parvo po x -> 4
				diff1 := Cuboid{dBx1, dBx2, ol.y1, dTy2, ol.z1, ol.z2}
				diff2 := Cuboid{dTx1, dTx2, ol.y1, dTy2, ol.z1, ol.z2}

				// posle po y -> 1
				diff5 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}

				return []Cuboid{diff1, diff2, diff5}
			}

			// 11
			if zDiffTop {

				// parvo po x -> 8
				diff1 := Cuboid{dBx1, dBx2, ol.y1, dTy2, ol.z1, dTz2}
				diff5 := Cuboid{dTx1, dTx2, ol.y1, dTy2, ol.z1, dTz2}

				// posle po y -> 2
				diff9 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, dTz2}

				// posle po z -> 1
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff5, diff9, diff11}
			}

			// 11
			if zDiffBottom {

				// parvo po x -> 8
				diff2 := Cuboid{dBx1, dBx2, ol.y1, dTy2, dBz1, ol.z2}
				diff5 := Cuboid{dTx1, dTx2, ol.y1, dTy2, dBz1, ol.z2}

				// posle po y -> 2
				diff10 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, ol.z2}

				// posle po z -> 1
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff2, diff5, diff10, diff11}
			}

			// 17
			if zDiffTopBottom {

				// parvo po x -> 12
				diff1 := Cuboid{dBx1, dBx2, ol.y1, dTy2, dBz1, dTz2}
				diff7 := Cuboid{dTx1, dTx2, ol.y1, dTy2, dBz1, dTz2}

				// posle po y - 3
				diff13 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, dTz2}

				// nakraq po z - 2
				diff16 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff17 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff7, diff13, diff16, diff17}
			}
		}

		// 5 / 11 / 11 / 17 - DONE 60/64
		if yDiffBottom {

			// 5
			if zNoDiff {
				// parvo po x -> 4
				diff3 := Cuboid{dBx1, dBx2, dBy1, ol.y2, ol.z1, ol.z2}
				diff4 := Cuboid{dTx1, dTx2, dBy1, ol.y2, ol.z1, ol.z2}

				// posle po y -> 1
				diff5 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, ol.z2}

				return []Cuboid{diff3, diff4, diff5}
			}

			// 11
			if zDiffTop {

				// parvo po x -> 8
				diff3 := Cuboid{dBx1, dBx2, dBy1, ol.y2, ol.z1, dTz2}
				diff5 := Cuboid{dTx1, dTx2, dBy1, ol.y2, ol.z1, dTz2}

				// posle po y -> 2
				diff9 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, dTz2}

				// posle po z -> 1
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff3, diff5, diff9, diff11}
			}

			// 11
			if zDiffBottom {

				// parvo po x -> 8
				diff3 := Cuboid{dBx1, dBx2, dBy1, ol.y2, dBz1, ol.z2}
				diff7 := Cuboid{dTx1, dTx2, dBy1, ol.y2, dBz1, ol.z2}

				// posle po y -> 2
				diff10 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, ol.z2}

				// posle po z -> 1
				diff11 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff3, diff7, diff10, diff11}
			}

			// 17
			if zDiffTopBottom {

				// parvo po x -> 12
				diff4 := Cuboid{dBx1, dBx2, dBy1, ol.y2, dBz1, dTz2}
				diff7 := Cuboid{dTx1, dTx2, dBy1, ol.y2, dBz1, dTz2}

				// posle po y - 3
				diff13 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, dTz2}
				// nakraq po z - 2
				diff16 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff17 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff4, diff7, diff13, diff16, diff17}
			}
		}

		// 8 / 17 / 17 / 26 DONE - 64/64
		if yDiffTopBottom {

			// 8
			if zNoDiff {
				// po x - 6
				diff1 := Cuboid{dBx1, dBx2, dBy1, dTy2, ol.z1, ol.z2}
				diff4 := Cuboid{dTx1, dTx2, dBy1, dTy2, ol.z1, ol.z2}

				// po y - 2
				diff7 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, ol.z2}
				diff8 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, ol.z2}

				return []Cuboid{diff1, diff4, diff7, diff8}
			}

			// 17
			if zDiffTop {

				// po x - 12
				diff1 := Cuboid{dBx1, dBx2, dBy1, dTy2, ol.z1, dTz2}
				diff7 := Cuboid{dTx1, dTx2, dBy1, dTy2, ol.z1, dTz2}

				// po y - 4
				diff13 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, ol.z1, dTz2}
				diff15 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, ol.z1, dTz2}
				// po z 1
				diff17 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff7, diff13, diff15, diff17}
			}

			// 17
			if zDiffBottom {

				// po x - 12
				diff4 := Cuboid{dBx1, dBx2, dBy1, dTy2, dBz1, ol.z2}
				diff7 := Cuboid{dTx1, dTx2, dBy1, dTy2, dBz1, ol.z2}

				// po y - 4

				diff14 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, dBz1, ol.z2}
				diff16 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, dBz1, ol.z2}
				// po z 1
				diff17 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}

				return []Cuboid{diff4, diff7, diff14, diff16, diff17}
			}

			// 26
			if zDiffTopBottom {

				// po x - 18
				diff1 := Cuboid{dBx1, dBx2, in.y1, in.y2, in.z1, in.z2}
				diff10 := Cuboid{dTx1, dTx2, in.y1, in.y2, in.z1, in.z2}

				// po y - 6

				diff19 := Cuboid{ol.x1, ol.x2, dBy1, dBy2, in.z1, in.z2}
				diff22 := Cuboid{ol.x1, ol.x2, dTy1, dTy2, in.z1, in.z2}

				// po y - 2
				diff25 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dBz1, dBz2}
				diff26 := Cuboid{ol.x1, ol.x2, ol.y1, ol.y2, dTz1, dTz2}

				return []Cuboid{diff1, diff10, diff19, diff22, diff25, diff26}
			}
		}
	}

	// should actually return here
	return
}

func createOverlapCuboid(first Cuboid, second Cuboid) (overlapped *Cuboid) {

	var x1, x2, y1, y2, z1, z2 int
	// X
	if first.x1 > second.x2 {
		return // nil
	}
	if first.x2 < second.x1 {
		return // nil
	}

	switch {
	case first.x1 <= second.x1 && second.x1 <= first.x2 && first.x2 <= second.x2:
		x1 = second.x1
		x2 = first.x2
	case first.x1 <= second.x1 && second.x2 <= first.x2: // second is the overlap entirely
		x1 = second.x1
		x2 = second.x2
	case second.x1 <= first.x1 && first.x2 <= second.x2: // first is the overlap entirely
		x1 = first.x1
		x2 = first.x2
	case second.x1 <= first.x1 && first.x1 <= second.x2 && second.x2 <= first.x2:
		x1 = first.x1
		x2 = second.x2
	}

	// Y
	if first.y1 > second.y2 {
		return // nil
	}
	if first.y2 < second.y1 {
		return // nil
	}
	switch {
	case first.y1 <= second.y1 && second.y1 <= first.y2 && first.y2 <= second.y2:
		y1 = second.y1
		y2 = first.y2
	case first.y1 <= second.y1 && second.y2 <= first.y2: // second is the overlap entirely
		y1 = second.y1
		y2 = second.y2
	case second.y1 <= first.y1 && first.y2 <= second.y2: // first is the overlap entirely
		y1 = first.y1
		y2 = first.y2
	case second.y1 <= first.y1 && first.y1 <= second.y2 && second.y2 <= first.y2:
		y1 = first.y1
		y2 = second.y2
	}

	// Z
	if first.z1 > second.z2 {
		return // nil
	}
	if first.z2 < second.z1 {
		return // nil
	}
	switch {
	case first.z1 <= second.z1 && second.z1 <= first.z2 && first.z2 <= second.z2:
		z1 = second.z1
		z2 = first.z2
	case first.z1 <= second.z1 && second.z2 <= first.z2: // second is the overlap entirely
		z1 = second.z1
		z2 = second.z2
	case second.z1 <= first.z1 && first.z2 <= second.z2: // first is the overlap entirely
		z1 = first.z1
		z2 = first.z2
	case second.z1 <= first.z1 && first.z1 <= second.z2 && second.z2 <= first.z2:
		z1 = first.z1
		z2 = second.z2
	}

	overlapped = &Cuboid{x1, x2, y1, y2, z1, z2}
	return overlapped
}

func calculateVolume(cuboid Cuboid) (volume uint64) {
	x1x2 := uint64(math.Abs(float64(cuboid.x2)-float64(cuboid.x1))) + 1
	y1y2 := uint64(math.Abs(float64(cuboid.y2)-float64(cuboid.y1))) + 1
	z1z2 := uint64(math.Abs(float64(cuboid.z2)-float64(cuboid.z1))) + 1

	volume = x1x2 * y1y2 * z1z2
	return
}
