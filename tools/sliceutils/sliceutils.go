package sliceutils

import (
	"sort"
	"strings"
)

func ContainsString(arr []string, str string) bool {

	for _, s := range arr {

		if s == str {
			return true
		}
	}

	return false

}

func RemoveString(arr []string, str string) ([]string, bool) {

	for i, s := range arr {

		if s == str {

			arr = append(arr[:i], arr[i+1:]...)

			return arr, true
		}
	}

	return arr, false

}

func RemoveDuplicate(list []string) []string {
	sort.Strings(list)
	i := 0
	var newlist = []string{""}
	for j := 0; j < len(list); j++ {
		if strings.Compare(newlist[i], list[j]) == -1 {
			newlist = append(newlist, list[j])
			i++
		}
	}
	return newlist
}
