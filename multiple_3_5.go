package main

import "fmt"

// 1_000 equals to 1000. underscore(_) replaces comma(,) in real number(1,000)
func main() {
	sum := 0
	for i := 1; i < 1_000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Printf("Sum: %d\n", sum)
}
