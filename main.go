// Jevgenijs Kovaļonoks 251RTU125

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var n int
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			factor := 2
			n, err := strconv.Atoi(arg)
			if err != nil || n < 2 {
				continue
			} else {
				fmt.Printf("%d: ", n)
				for factor <= n {
					if n%factor == 0 {
						fmt.Printf("%d ", factor)
						n = n / factor
					} else {
						factor += 1
					}
				}

			}
			fmt.Println()
		}
	} else {
		_, err := fmt.Scan(&n)
		factor := 2
		fmt.Print("Input number: ")

		if err != nil {
			fmt.Println("input-output error")
			return
		}
		if n < 2 {
			fmt.Println("input-output error")
			return
		}
		fmt.Printf("%d ", n)
		fmt.Println()
		fmt.Printf("%d: ", n)
		for factor <= n {
			if n%factor == 0 {
				fmt.Printf("%d ", factor)
				n = n / factor
			} else {
				factor += 1
			}
		}
	}

}
