package main

import (
	"fmt"
	"time"
)

func displayBucketInfo(bucket []int) {
	fmt.Printf("Bucket length: %2d, Bucket capacity: %2d\n", len(bucket), cap(bucket))
}

func main() {
	numberbucket := make([]int, 0)

	for i := 0; i < 20; i++ {
		numberbucket = append(numberbucket, i)
		displayBucketInfo(numberbucket)
		time.Sleep(time.Second)
	}
}