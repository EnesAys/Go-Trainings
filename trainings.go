package main

import (
	"fmt"
	"math"
	"time"
)

var global bool

func main() {

	var local int
	var x, y int = 1, 2
	var number float64 = 9

	fmt.Println("Hello World")
	fmt.Println("Result = ", math.Sqrt(number))
	fmt.Println("Result = ", math.Pi)

	fmt.Println(AddMessage("Merhaba", "Dünya"))
	fmt.Println(ChangeValue(5))
	fmt.Println(Swap(5, 10))
	fmt.Println(Split(8))

	fmt.Println(local, " ", global)
	fmt.Print(x, "\n", y)
	fmt.Print(&x)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	sum := 1
	for sum < 5 {
		sum += sum
	}
	fmt.Println(sum)

	if a := true; a == global {
		fmt.Println("if condition")
	} else {
		fmt.Println("else condition")
	}

	var Name string = "Enes"
	switch Name {
	case "Enes":
		fmt.Println("Case Enes, name is", Name)
		break
	case "Go Lang":
		fmt.Println("Case Go Lang, name is", Name)
		break
	default:
		fmt.Println("Case Default, name is", Name)
		break
	}

	day := time.Now().Hour()
	switch { //Like if then else
	case day < 12:
		fmt.Println("Good morning!")
	case day < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	defer fmt.Println("Defer function will start last end use Last in first out - 1")
	defer fmt.Println("Defer function will start last end use Last in first out - 2")
}

func AddMessage(a, b string) string {
	return a + " " + b
}

func ChangeValue(c int) int {
	number := c
	return number
}

func Swap(c, b int) (int, int) {
	return b, c
}

func Split(sum int) (x, y int) {
	x = sum * 3
	y = sum - x
	return
}
