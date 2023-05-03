package main

import "fmt"

func removeFromSlice(a []int, b []int) []int {
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(a); j++ {
			if a[j] == b[i] {
				a = append(a[:j], a[j+1:]...)
				j--
			}
		}
	}
	return a
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2}
	fmt.Println(removeFromSlice(a, b))

}
