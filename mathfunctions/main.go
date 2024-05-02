package main

import (
	"fmt"

	"github.com/leandrosolagna/playing-with-go/mathfunctions/bhaskara"
)

func main() {
	var firstTerm, secondTerm, thirdTerm int

	fmt.Println("Type a number for first term:")
	fmt.Scan(&firstTerm)
	fmt.Println("Type a number for second term:")
	fmt.Scan(&secondTerm)
	fmt.Println("Type a number for third term:")
	fmt.Scan(&thirdTerm)

	deltaRes := bhaskara.GetDelta(firstTerm, secondTerm, thirdTerm)
	fmt.Println("The value of delta is:", deltaRes)

	rootRes := bhaskara.GetRoot(deltaRes)
	fmt.Println("The value of root is:", rootRes)

	x1, x2 := bhaskara.BhaskaraMethod(firstTerm, secondTerm, rootRes)
	fmt.Println("The result of the equation are", x1, "and", x2)
}
