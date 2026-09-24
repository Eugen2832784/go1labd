// Jevgēnijs Kovaļonoks 125RDB251

package main

import "fmt"

func main() {
	fmt.Println("First sequence of numbers")
	sequence1 := readSequence()
	fmt.Println("Second sequence of numbers")
	sequence2 := readSequence()
	result := findUniqueElements(sequence1, sequence2)
	printResult(result)
}

func readSequence() []int {
	var n int
	newSlice := []int{}
	for {
		fmt.Print("Input number: ")
		_, err := fmt.Scan(&n)
		if err != nil {
			fmt.Println("input-output error")
			var dump string
			fmt.Scanln(&dump)
			continue
		}
		if n == 0 {
			break
		}
		newSlice = append(newSlice, n)
	}
	return newSlice
}

func findUniqueElements(seq1, seq2 []int) []int {
	resSeq := []int{}
	for _, valSeq1 := range seq1 {
		found := false
		for _, valSeq2 := range seq2 {
			if valSeq1 == valSeq2 {
				found = true
				break
			}
		}
		if !found {
			resSeq = append(resSeq, valSeq1)
		}
	}
	return resSeq
}

func printResult(result []int) {
	fmt.Print("Result: ")
	for _, val := range result {
		fmt.Print(val, " ")
	}
}
