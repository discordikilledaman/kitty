package commands

import (
	"fmt"
	"strconv"

	"sort"

	"github.com/acdenisSK/kitty"
)

// Sort f
type Sort struct{}

// Checks f
func (Sort) Checks() kitty.Checks {
	return kitty.Checks{}
}

// Process f
func (Sort) Process(context kitty.Context) {
	nums := []int{}
	if len(context.Args) > 30 {
		context.Say("Sort only allows 30 integers")
		return
	}
	for _, num := range context.Args {
		result, err := strconv.Atoi(num)
		if err != nil {
			context.Say(fmt.Sprint("Error converting to a number: ", err))
			return
		}
		nums = append(nums, result)
	}
	sort.Ints(nums)
	context.Say(fmt.Sprint("Result: ", nums))
}
