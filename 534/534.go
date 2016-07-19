// UVa 534 - Frogger

package main

import (
	"fmt"
	"math"
	"os"
)

type stone struct {
	x, y float64
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 { return a + b - min(a, b) }

func floydWarshall(stones []stone) float64 {
	n := len(stones)
	minMax := make([][]float64, n)
	for i := range minMax {
		minMax[i] = make([]float64, n)
		for j := range minMax[i] {
			minMax[i][j] = math.Sqrt((stones[i].x-stones[j].x)*(stones[i].x-stones[j].x) + (stones[i].y-stones[j].y)*(stones[i].y-stones[j].y))
		}
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				minMax[i][j] = min(minMax[i][j], max(minMax[i][k], minMax[k][j]))
			}
		}
	}
	return minMax[0][1]
}

func main() {
	in, _ := os.Open("534.in")
	defer in.Close()
	out, _ := os.Create("534.out")
	defer out.Close()

	var n, kase int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		stones := make([]stone, n)
		for i := range stones {
			fmt.Fscanf(in, "%f%f", &stones[i].x, &stones[i].y)
		}
		fmt.Fscanln(in)
		kase++
		fmt.Fprintf(out, "Scenario #%d\n", kase)
		fmt.Fprintf(out, "Frog Distance = %.3f\n\n", floydWarshall(stones))
	}
}