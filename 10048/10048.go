// UVa 10048 - Audiophobia

package main

import (
	"fmt"
	"math"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int { return a + b - min(a, b) }

func floydWarshall(matrix [][]int) {
	n := len(matrix)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i][j] = min(matrix[i][j], max(matrix[i][k], matrix[k][j]))
			}
		}
	}
}

func main() {
	in, _ := os.Open("10048.in")
	defer in.Close()
	out, _ := os.Create("10048.out")
	defer out.Close()

	var kase, c, s, q, c1, c2, d int
	for {
		if fmt.Fscanf(in, "%d%d%d", &c, &s, &q); c == 0 && s == 0 && q == 0 {
			break
		}
		matrix := make([][]int, c)
		for i := range matrix {
			matrix[i] = make([]int, c)
			for j := range matrix[i] {
				matrix[i][j] = math.MaxInt32
			}
		}
		for i := 0; i < s; i++ {
			fmt.Fscanf(in, "%d%d%d", &c1, &c2, &d)
			matrix[c1-1][c2-1], matrix[c2-1][c1-1] = d, d
		}
		floydWarshall(matrix)
		kase++
		if kase > 1 {
			fmt.Fprintln(out)
		}
		fmt.Fprintf(out, "Case #%d\n", kase)
		for i := 0; i < q; i++ {
			fmt.Fscanf(in, "%d%d", &c1, &c2)
			min := matrix[c1-1][c2-1]
			if min == math.MaxInt32 {
				fmt.Fprintln(out, "no path")
			} else {
				fmt.Fprintln(out, matrix[c1-1][c2-1])
			}
		}
	}
}
