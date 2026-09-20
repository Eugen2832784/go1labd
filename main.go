// Jevgenijs Kovaļonoks 251RTU125

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	factor := 2
	var n int
	fmt.Println("Enter a number: ")
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			n, err := strconv.Atoi(arg)
			if err != nil || n < 2 {
				continue
			} else {

			}

		}
		// 1. SCENĀRIJS: Ir komandrindas parametri
		// Ar parasto for ciklu ej cauri os.Args[1:]
		// Izmanto strconv.Atoi(), lai pārvērstu tekstu par skaitli
		// Ja ir kļūda vai skaitlis < 2 -> continue (IGNORĒT)
		// Ja skaitlis >= 2 -> palaid savu faktoru ciklu un izdrukā rezultātu
	} else {
		fmt.Println("Enter a number: ")
		_, err := fmt.Scan(&n)

		if err != nil {
			fmt.Println("input-output error")
			return
		}
		if n < 2 {
			fmt.Println("input-output error")
			return
		}
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
