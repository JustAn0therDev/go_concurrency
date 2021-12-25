package concurrency_tests

import (
	"fmt"
	"time"
)

func IncrementAsync(value *int) {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println(*value)
		*value++
	}
}
