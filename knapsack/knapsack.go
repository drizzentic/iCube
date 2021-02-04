package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

type Item struct {
	Weight int `json:"weight"`
	Value  int `json:"value"`
}

type Knapsack struct {
	knapsackItemsWeight, knapsackCapacity, totalLoot int
	Items                                            []Item
}

func main() {
	var Items Knapsack

	jsonPayload := `{"Items":[{"weight":5,"value":10},{"weight":4,"value":40},{"weight":6,"value":30},{"weight":4,"value":98}]}`

	err := json.Unmarshal([]byte(jsonPayload), &Items)
	if err != nil {
		fmt.Print(error.Error(err))
		return
	}
	// Create a knapsack to hold the totalLoot

	knapsack := Knapsack{knapsackCapacity: 10}

	g := solution(Items.Items, knapsack)

	// Print the final weight

	fmt.Print(g)
}

func solution(items []Item, b Knapsack) int {
	// Sort the totalLoot from the largest to the smallest
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})
	// Add the items to the array
	for i := range items {
		_ = b.addLoot(items[i])
	}

	// Loop through the bag to get the final worth of the totalLoot
	for _, v := range b.Items {
		b.totalLoot += v.Value
	}
	// return the totalLoot total worth

	return b.totalLoot
}
func (b *Knapsack) addLoot(i Item) error {
	// If the value of the bag has been exceeded continue to the next value that will fit in the bag
	if b.knapsackItemsWeight+i.Weight <= b.knapsackCapacity {
		b.knapsackItemsWeight += i.Weight
		b.Items = append(b.Items, i)
		return nil
	}
	return errors.New("The totalLoot is bigger than the size of the bag")
}
