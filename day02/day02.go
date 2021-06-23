package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func findsum(s string) (counter3 int, counter2 int) {
	count3 := false
	count2 := false
	for _, t := range s {
		if strings.Count(s, string(t)) == 3 {
			count3 = true

		} else if strings.Count(s, string(t)) == 2 {
			count2 = true

		}
	}
	if count3 == true {
		counter3 = 1
	}
	if count2 == true {
		counter2 = 1
	}
	return counter3, counter2

}

func main() {
	num1 := 0
	num2 := 0
	sum2 := 0
	sum3 := 0
	b, err := ioutil.ReadFile("day02.txt")
	if err != nil {
		fmt.Println(err)
	}
	st := string(b)
	sliceData := strings.Split(string(st), "\n")
	for _, line := range sliceData {
		num1, num2 = findsum(string(line))
		sum2 += num2
		sum3 += num1
	}
	fmt.Println(sum2 * sum3)

}
