package main

import (
	"fmt"
)

//type emptyInterface interface{}

func whatIsThis(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("%s is a string\n", v)
	case int:
		fmt.Printf("%d is an int\n", v)
	default:
		fmt.Printf("type: %T\n", v)
	}
}

func scope() func() int {
	outer_var := 2
	foo := func() int { return outer_var }
	return foo
}

func main() {
	whatIsThis(42)
	whatIsThis("holy cow!!")
	whatIsThis([]int{1, 2, 3})
	whatIsThis([...]string{"hi", "how", "are", "you"})
}
