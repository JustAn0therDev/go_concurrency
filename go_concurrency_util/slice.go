package go_concurrency_util

func AnyTrue(slice []bool) bool {
	for _, val := range slice {
		if val {
			return true
		}
	}
	return false
}