package forloop

import "fmt" 

func ForFive() {
	loop := [][]int{
		{	1, 3, 5 },
		{	4, 5, 6 },
		{	2, 4, 6 },
	}

	for i, j := range loop {
	fmt.Printf("The index of  %v is %v \n", j, i)
	for p, inside := range j {
		fmt.Printf("The index of  %v is %v \n", inside, p)
	}
	}
}