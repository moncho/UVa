// UVa 10205 - Stack 'em Up

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const total = 52

var (
	suits  = []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	values = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King", "Ace"}
)

func do(shuffles [][]int, order []int) []int {
	cards := make([]int, total)
	for i := range cards {
		cards[i] = i
	}
	for _, vi := range order {
		newCards := make([]int, total)
		for j, vj := range shuffles[vi] {
			newCards[j] = cards[vj-1]
		}
		copy(cards, newCards)
	}
	return cards
}

func main() {
	in, _ := os.Open("10205.in")
	defer in.Close()
	out, _ := os.Create("10205.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	s.Scan()
	s.Text()
	var tmp int
	var line string
	for kase > 0 && s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		var shuffles [][]int
		for n > 0 {
			var shuffle []int
			for s.Scan() {
				r := strings.NewReader(s.Text())
				for {
					if _, err := fmt.Fscanf(r, "%d", &tmp); err != nil {
						break
					}
					shuffle = append(shuffle, tmp)
				}
				if len(shuffle) == total {
					break
				}
			}
			shuffles = append(shuffles, shuffle)
			n--
		}
		var order []int
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			tmp, _ = strconv.Atoi(line)
			order = append(order, tmp-1)
		}
		for _, vi := range do(shuffles, order) {
			fmt.Fprintf(out, "%s of %s\n", values[vi%13], suits[vi/13])
		}
		kase--
		if kase > 0 {
			fmt.Fprintln(out)
		}
	}
}
