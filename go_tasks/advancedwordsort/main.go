package main

import (
	"fmt"

)

func AdvancedSortWordArr(a []string, f func(a, b string) int) {

	for i := 0; i < len(a)-1; i++ {
		for j :=0; j < len(a)-i-1; j++ {
			fmt.Printf("Compare %v and %v → %v\n", a[j], a[j+1], f(a[j], a[j+1]))
			if f(a[j], a[j+1]) < 0 {
				fmt.Printf("Swap %v ↔ %v\n", a[j], a[j+1])
				a[j],a[j+1] = a[j+1], a[j]
			}
			fmt.Println("Current:", a)

		}
	}

}
func Compare(a,b string) int {
	if a == b {
		return 0
	} else if a > b {
		return -1
	} else {
		return 1
	}
    
}

func main() {


	result := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)

	fmt.Println(result)
}