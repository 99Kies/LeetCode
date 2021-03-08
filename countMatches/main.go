package main

import "fmt"

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	index := 0
	if ruleKey == "color" {
		index = 1
	} else if ruleKey == "name" {
		index = 2
	}
	result := 0
	for i := range items {
		if ruleValue == items[i][index] {
			result++
		}
	}

	return result
}

func main() {
	items := [][]string{{"phone", "blue", "pixel"}, {"computer", "silver", "lenovo"}, {"phone", "gold", "phone"}}
	ruleKey := "color"
	ruleValue := "silver"
	a := countMatches(items, ruleKey, ruleValue)
	fmt.Println(a)

}
