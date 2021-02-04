package main

import (
	"encoding/json"
	"testing"
)

func TestMaximumReward(t *testing.T)  {
	var Items Knapsack
	jsonPayload:=`{"Items":[{"weight":4,"value":40},{"weight":6,"value":30}]}`
	_ = json.Unmarshal([]byte(jsonPayload), &Items)
	bag := Knapsack{knapsackCapacity: 10}
	a:=solution(Items.Items, bag)

	if a!=70 {
		t.Errorf("Failed expected %v but got %v", 70, a)
	}

}
