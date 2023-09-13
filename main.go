package main

import "fmt"

func main() {
	// create a list of any type
	list := []interface{}{}
	for i := 1; i <= 100; i++ {
		msg := ""
		if i%3 == 0 {
			msg += "Fizz"
		}
		if i%5 == 0 {
			msg += "Buzz"
		}
		if msg == "" {
			msg = fmt.Sprintf("%d", i)
		}
		list = append(list, msg)
	}
	// print the list
	fmt.Println(list)
}
