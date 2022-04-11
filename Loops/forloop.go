package main

import "fmt"

//Go has single loop type: For
//There is no while or do-while as in c/c++

func main() {
	for i := 0; i < 4; i++ {
		fmt.Printf("Value of i is %d ", i)
		fmt.Println()
	}

	var j = 1
	for j <= 3 {
		fmt.Printf("This is new type for loop")
		j++
		fmt.Println()
	}

	for k := 0; k < 5; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Printf("K is odd, it's value is %d", k)
		fmt.Println()
	}

	//For loop without conditionals

	for {
		fmt.Printf("It is infinite loop if we don't break manually")
		break
	}
}
