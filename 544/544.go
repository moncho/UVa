// UVa 544 - Heavy Cargo

package main

import (
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int { return a + b - min(a, b) }

func floydWarshall(matrix map[string]map[string]int, c1, c2 string) int {
	for k, _ := range matrix {
		for i, _ := range matrix {
			for j, _ := range matrix {
				matrix[i][j] = max(matrix[i][j], min(matrix[i][k], matrix[k][j]))
			}
		}
	}
	return matrix[c1][c2]
}

func main() {
	in, _ := os.Open("544.in")
	defer in.Close()
	out, _ := os.Create("544.out")
	defer out.Close()

	var kase, n, r, l int
	var c1, c2 string
	first := true
	for {
		if fmt.Fscanf(in, "%d%d", &n, &r); n == 0 && r == 0 {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		matrix := make(map[string]map[string]int)
		for r > 0 {
			fmt.Fscanf(in, "%s%s%d", &c1, &c2, &l)
			if _, ok := matrix[c1]; !ok {
				matrix[c1] = make(map[string]int)
			}
			if _, ok := matrix[c2]; !ok {
				matrix[c2] = make(map[string]int)
			}
			matrix[c1][c2], matrix[c2][c1] = l, l
			r--
		}
		fmt.Fscanf(in, "%s%s", &c1, &c2)
		kase++
		fmt.Fprintf(out, "Scenario #%d\n%d tons\n", kase, floydWarshall(matrix, c1, c2))
	}
}
