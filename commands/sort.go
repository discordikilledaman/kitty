package commands

import (
	"sort"
	"strconv"

	"github.com/acdenisSK/kitty"
)

// Sort f
type Sort struct{}

// IsOwnerOnly f
func (Sort) IsOwnerOnly() bool {
	return false
}

// Help f
func (Sort) Help() [2]string {
	return [2]string{"Sorts the provided numbers from lowest to highest", "<3> <2> <1>...."}
}

// Process f
func (Sort) Process(context kitty.Context) {
	nums := []int{}
	for _, num := range context.Args {
		result, err := strconv.Atoi(num)
		if err != nil {
			context.Error(err)
			return
		}
		nums = append(nums, result)
	}
	sort.Ints(nums)
	context.Say("Result: ", nums)
}
