package sliceutils

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
