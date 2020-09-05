//https://leetcode.com/problems/flood-fill/
package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}
	rl := len(image)
	rc := len(image[0])
	color := image[sr][sc]
	image[sr][sc] = newColor
	r, c := sr-1, sc
	if r >= 0 && image[r][c] == color {
		floodFill(image, r, c, newColor)
	}
	r, c = sr, sc-1
	if c >= 0 && image[r][c] == color {
		floodFill(image, r, c, newColor)
	}
	r, c = sr, sc+1
	if c < rc && image[r][c] == color {
		floodFill(image, r, c, newColor)
	}
	r, c = sr+1, sc
	if r < rl && image[r][c] == color {
		floodFill(image, r, c, newColor)
	}
	return image
}

func floodFillBFS(image [][]int, sr, sc, newColor int) [][]int {

	i := 1
	for i < 3 {
		//  loop through directional

		// loop through diagonal
		i++
	}
	image[sr][sc] = newColor
	return image
}

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(image, 1, 1, 2))
}
