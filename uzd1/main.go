// Jevgenijs Kovaļonoks 251RDB125

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
		fmt.Print("Input number: ")
		_, err := fmt.Scan(&n)

		if err != nil || n < 2 {
			fmt.Println("input-output error")
			return
		}

		factor := 2

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
