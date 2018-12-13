package day03

import (
	"regexp"
	"strconv"
	"strings"
)

func CalculateOverlappingAreas(input string) int {
	var x = strings.Split(input, "\n")

	board := make(map[string]int)

	overlapping := 0

	for _, i := range x {
		coor := getCoordinates(i)

		for left := coor.position.fromLeft + 1; left <= coor.position.fromLeft+coor.width; left++ {
			for top := coor.position.fromTop + 1; top <= coor.position.fromTop+coor.height; top++ {

				posLeft := strconv.Itoa(left)
				posTop := strconv.Itoa(top)

				index := posLeft + ":" + posTop

				if board[index] == 0 {
					board[index] += 1
				} else if board[index] >= 1 {
					overlapping++
				}
			}
		}
	}

	//fmt.Println(board)

	return overlapping
}

type pos struct {
	fromLeft int
	fromTop  int
}

type coordinates struct {
	id       int
	position pos
	width    int
	height   int
}

func getCoordinates(c string) coordinates {
	re := regexp.MustCompile(`(\d)`)
	single := re.FindAllString(c, -1)

	id, _ := strconv.Atoi(single[0])
	fromLeft, _ := strconv.Atoi(single[1])
	fromTop, _ := strconv.Atoi(single[2])
	width, _ := strconv.Atoi(single[3])
	height, _ := strconv.Atoi(single[4])

	var coordinates = coordinates{
		id: id,
		position: pos{
			fromLeft: fromLeft,
			fromTop:  fromTop,
		},
		width:  width,
		height: height,
	}

	return coordinates
}
