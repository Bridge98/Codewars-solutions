package kata

import (
	"reflect"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {

	if array1 == nil || array2 == nil {
		return false
	} else {
		// put array1[i] * array1[i] in array1[i]:
		for i, value := range array1 {
			array1[i] *= value
		}
		// sort all the given arrays:
		sort.Ints(array1)
		sort.Ints(array2)
	}

	/***check if array1 is equal to array2***/
	return reflect.DeepEqual(array1, array2)
}