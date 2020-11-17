package main

import (
	"fmt"
	"math"

	"github.com/jordanblakey/go_crash_course/03_packages/strutil"
)

func main() {
	// Import and used math package
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("Not an anagram!"))
}
