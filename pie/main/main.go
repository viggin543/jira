package main

import (
	"fmt"
	"github.com/ahmetb/go-linq/v3"
	"github.com/viggin543/jira/pie/collections"
)

func main() {
	cars := collections.Pies {
		{"Bob", 1},
		{"Sally", 2},
		{"John", 3},
		{"Jane", 4},

	}

	redCars := cars.Filter(func(car collections.Pie) bool {
		return car.Id == "Bob"
	})

	var bigPies []collections.Pie
	linq.From(redCars).Where(func(x interface{}) bool {
		return x.(collections.Pie).Number > 2
	}).ToSlice(&bigPies)

	var numbers []int
	linq.From(cars).Select(func(i interface{}) interface{} {
		return i.(collections.Pie).Number
	}).ToSlice(&numbers)

	fmt.Println(numbers)

}
