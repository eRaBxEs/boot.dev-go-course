package main

import (
	"fmt"
)

/*
Assignment
We've been asked to add a feature that extracts costs for a given day.

Complete the getDayCosts() function using the append() function. It accepts a slice of cost structs and a day int, and it returns a float64 slice containing that day's costs.

If there are no costs for that day, return an empty slice.
*/

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {

	var daysCost []float64
	for i := 0; i < len(costs); i++ {
		if costs[i].day == day {
			daysCost = append(daysCost, costs[i].value)
		}
	}
	return daysCost
}

func test(costs []cost, day int) {
	defer fmt.Println("===============================================")
	costForDay := getDayCosts(costs, day)
	fmt.Printf("The costs for day: %v are as listed: %#v\n", day, costForDay)
}

func main() {
	test([]cost{
		{day: 1, value: 0.2},
		{day: 2, value: 0.2},
		{day: 2, value: 0.2},
		{day: 1, value: 0.2},
		{day: 3, value: 0.25},
		{day: 3, value: 0.42}}, 3)

	test([]cost{
		{day: 1, value: 0.2},
		{day: 2, value: 0.2},
		{day: 2, value: 0.2},
		{day: 1, value: 0.2},
		{day: 3, value: 0.25},
		{day: 3, value: 0.42}}, 1)

	test([]cost{
		{day: 1, value: 0.55},
		{day: 2, value: 0.12},
		{day: 2, value: 0.22},
		{day: 1, value: 0.20},
		{day: 3, value: 0.25},
		{day: 3, value: 0.42}}, 2)
}
