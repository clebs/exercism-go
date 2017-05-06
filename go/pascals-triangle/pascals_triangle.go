// Package pascal provides tools to calculate pascal's triangles
package pascal

const testVersion = 1

//Triangle calculates pascal's triangle for the given number of rows
func Triangle(n int) [][]int {
	pt := make([][]int, n)

	for row := 0; row < n; row++ {
		pt[row] = make([]int, row+1)
		for col := 0; col < row+1; col++ {
			if col == 0 || col == row {
				pt[row][col] = 1
			} else {
				pt[row][col] = pt[row-1][col-1] + pt[row-1][col]
			}
		}
	}
	return pt
}
