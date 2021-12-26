package go_concurrency_util

func PanicIfErrNotNil(err *error) {
	if *err != nil {
		panic(*err)
	}
}