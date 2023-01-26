/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
// #region zero is close to 0
// package main

// import (
// 	"fmt"
// )

// type ID struct {
// 	row    int
// 	column int
// }

// type Vertex struct {
// 	id       ID
// 	value    int
// 	distance int
// 	points   []int
// 	wasUsed  bool
// }

// func minimumTotal(triangle [][]int) int {
// 	vertexes := make(map[ID]*Vertex)

// 	path := make([]int, 0)

// 	for rowNumber := range triangle {
// 		shortestVertex := vertexes[getShortest(vertexes)]
// 		if shortestVertex == nil {
// 			id := NewID(0, 0)
// 			shortestVertex = &Vertex{
// 				id:       NewID(0, 0),
// 				value:    triangle[0][0],
// 				distance: triangle[0][0],
// 				points:   []int{triangle[0][0]},
// 				wasUsed:  true,
// 			}

// 			vertexes[id] = shortestVertex
// 		}

// 		rowNumber = shortestVertex.id.row

// 		// if it is not the last row
// 		if rowNumber < len(triangle)-1 {
// 			columns := getNextStepsRange(shortestVertex.id.column, len(triangle[rowNumber]))
// 			for _, nextVertexColumn := range columns {
// 				nextVertexRow := rowNumber + 1
// 				newVertex := defineDistance(shortestVertex, triangle[nextVertexRow][nextVertexColumn], nextVertexRow, nextVertexColumn, vertexes)
// 				vertexes[newVertex.id] = newVertex
// 			}
// 		}
// 	}

// 	distance := 0
// 	wasChanged := false

// 	for _, vertex := range vertexes {
// 		// if it is the last row
// 		if vertex.id.row == len(triangle)-1 {
// 			if !wasChanged || (vertex.distance < distance) {
// 				distance = vertex.distance
// 				path = vertex.points

// 				wasChanged = true
// 			}
// 		}
// 	}

// 	fmt.Println(path)
// 	return distance
// }

// func getNextStepsRange(index, rowLen int) []int {
// 	switch index {
// 	case 0:
// 		return []int{0, 1}
// 	case rowLen - 1:
// 		return []int{index, rowLen}
// 	default:
// 		return []int{index, index + 1}

// 	}
// }

// func defineDistance(vertex *Vertex, point, row, column int, vertexes map[ID]*Vertex) *Vertex {
// 	defer func() {
// 		vertex.wasUsed = true
// 	}()

// 	id := NewID(row, column)
// 	if existingVertex, ok := vertexes[id]; ok {
// 		currentDistance := vertex.distance + point
// 		if existingVertex.distance < currentDistance {
// 			return existingVertex
// 		}
// 	}

// 	path := append([]int{}, vertex.points...)

// 	return &Vertex{
// 		id:       id,
// 		value:    point,
// 		distance: vertex.distance + point,
// 		points:   append(path, point),
// 		wasUsed:  false,
// 	}
// }

// func getShortest(vertexes map[ID]*Vertex) (id ID) {
// 	distance := 0
// 	isFirst := true
// 	id = ID{-1, -1}

// 	for _, vertex := range vertexes {
// 		if (distance > vertex.distance || isFirst) && !vertex.wasUsed {
// 			distance = vertex.distance
// 			id = vertex.id

// 			isFirst = false
// 		}
// 	}

// 	return id
// }

// func NewID(row, column int) ID {
// 	return ID{row, column}
// }
// #endregion

package main

import (
	"fmt"
)

type ID struct {
	row    int
	column int
}

type Vertex struct {
	id       ID
	value    int
	distance int
	points   []int
	wasUsed  bool
}

func minimumTotal(triangle [][]int) int {
	vertexes := make(map[ID]*Vertex)

	path := make([]int, 0)

	for {
		shortestVertex := vertexes[getShortest(vertexes)]

		// if all vertexes have been visited
		if len(vertexes) != 0 && shortestVertex == nil {
			break
		}

		// if it is the first vertex
		if len(vertexes) == 0 && shortestVertex == nil {
			id := NewID(0, 0)
			shortestVertex = &Vertex{
				id:       NewID(0, 0),
				value:    triangle[0][0],
				distance: triangle[0][0],
				points:   []int{triangle[0][0]},
				wasUsed:  true,
			}

			vertexes[id] = shortestVertex
		}

		rowNumber := shortestVertex.id.row

		// if it is not the last row
		if rowNumber < len(triangle)-1 {
			columns := getNextStepsRange(shortestVertex.id.column, len(triangle[rowNumber]))
			for _, nextVertexColumn := range columns {
				nextVertexRow := rowNumber + 1
				newVertex := defineDistance(shortestVertex, triangle[nextVertexRow][nextVertexColumn], nextVertexRow, nextVertexColumn, vertexes)
				vertexes[newVertex.id] = newVertex
			}
		}
	}

	distance := 0
	wasChanged := false

	for _, vertex := range vertexes {
		// if it is the last row
		if vertex.id.row == len(triangle)-1 {
			if !wasChanged || (vertex.distance < distance) {
				distance = vertex.distance
				path = vertex.points

				wasChanged = true
			}
		}
	}

	fmt.Println(path)
	return distance
}

func getNextStepsRange(index, rowLen int) []int {
	switch index {
	case 0:
		return []int{0, 1}
	case rowLen - 1:
		return []int{index, rowLen}
	default:
		return []int{index, index + 1}
	}
}

func defineDistance(vertex *Vertex, point, row, column int, vertexes map[ID]*Vertex) *Vertex {
	defer func() {
		vertex.wasUsed = true
	}()

	id := NewID(row, column)
	if existingVertex, ok := vertexes[id]; ok {
		currentDistance := vertex.distance + point
		if existingVertex.distance < currentDistance {
			return existingVertex
		}
	}

	path := append([]int{}, vertex.points...)

	return &Vertex{
		id:       id,
		value:    point,
		distance: vertex.distance + point,
		points:   append(path, point),
		wasUsed:  false,
	}
}

func getShortest(vertexes map[ID]*Vertex) (id ID) {
	for _, vertex := range vertexes {
		if !vertex.wasUsed {
			vertex.wasUsed = true
			return vertex.id
		}
	}

	return ID{-1, -1}

	// the code below does not work for this task coz
	// in this task we have to find the minimum value of the distance
	// rather than the closest to 0
	distance := 0
	isFirst := true
	id = ID{-1, -1}

	for _, vertex := range vertexes {
		if (distance > vertex.distance || isFirst) && !vertex.wasUsed {
			distance = vertex.distance
			id = vertex.id

			isFirst = false
		}
	}

	return id
}

func NewID(row, column int) ID {
	return ID{row, column}
}

// @lc code=end
