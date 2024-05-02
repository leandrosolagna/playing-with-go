package bhaskara

import (
	"fmt"
)

func GetDelta(a, b, c int) int {
	delta := int((b * b) - (4 * (a * c)))

	return delta
}

func GetRoot(delta int) int {
	/*
		This function will receive the value of GetDelta
		Initialize a variable to help with the calculation of the root value
		A forloop will run and will only stop when the value of variable "root"
		is the same as variable "i".
		This is the simplest way I could think to find the value of a square
		root number. If the values are not the same, it will continue dividing
		delta by "i", and then increment variable "root".
	*/
	fmt.Println("Delta:", delta)
	root := 0
	for i := 1; i < delta; i++ {
		root = (delta / i)
		if root == i {
			break
		}
		root += i
	}

	return root
}

func BhaskaraMethod(a, b, root int) (int, int) {
	x1 := ((-(b) + root) / (2 * a))
	x2 := ((-(b) - root) / (2 * a))

	return x1, x2
}
