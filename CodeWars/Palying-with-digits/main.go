package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(DigPow(46288, 3))
}

func DigPow(n, p int) int {
	var sum int
	sl := make([]int, 20)
	for i := 20; i > 0; i-- {
		if n/int(math.Pow(10, float64(i))) >= 1 {
			sl[20-i] = int(n / int(math.Pow(10, float64(i))))
		}
	}
	a := 0
	for j := 0; j < 20; j++ {
		if sl[j] != 0 {
			a = j
			break
		}

	}
	b := sl[a:20]
	for i := 0; i < 20-a; i++ {
		sum += int(math.Pow(float64(b[i]), float64(p+i)))
	}
	if sum%n == 0 {
		return sum / n
	}
	return -1
}
