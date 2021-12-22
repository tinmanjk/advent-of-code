package main

import (
	"aoc/libs/go/inputParse"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := inputParse.ReturnSliceOfLinesFromFile(inputPath)
	cuboids := parseInput(lines)

	var result int
	// part 1
	result = findResult(cuboids)
	fmt.Println(result)
	// part 2

}

func parseInput(lines []string) (instrCuboids []*InstructionCuboid) {
	instrCuboids = []*InstructionCuboid{}

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

		cuboid := Cuboid{x1, x2, y1, y2, z1, z2}
		var on bool
		if onOff == "on" {
			on = true
		} else {
			on = false
		}
		instrCuboid := InstructionCuboid{on, &cuboid}
		instrCuboids = append(instrCuboids, &instrCuboid)
	}
	return
}

const inputPath = "../input.txt"

type InstructionCuboid struct {
	on     bool
	cuboid *Cuboid
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

func findResult(instrCuboids []*InstructionCuboid) (result int) {

	cubes := map[Cube]Cube{}

	for i := 0; i < len(instrCuboids); i++ {
		cuboid := instrCuboids[i].cuboid
		// -60 33, -60 33, -60 33
		// does it affect us
		// any bottom is higher than max
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
