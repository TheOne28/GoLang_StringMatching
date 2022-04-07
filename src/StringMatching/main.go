package main

import (
	"StringMatching/matchingMethod"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var choice int64

	scanner := bufio.NewScanner(os.Stdin)

	// reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Selamat datang di Pattern Matching\n")
	fmt.Printf("Ada 2 algoritma yang tersedia, silahkan pilih\n")
	fmt.Printf("1. Knuth-Morris-Pratt Algoritm\n")
	fmt.Printf("2. Boyer-Moore Algorithm\n")
	fmt.Printf("3. Exit\n")

	for cond := true; cond; cond = (choice < 1 || choice > 3) {
		fmt.Printf("Silahkan masukkan angka pilihan anda (1-3): ")
		scanner.Scan()
		choice, _ = strconv.ParseInt(scanner.Text(), 10, 64)

		if choice == 3 {
			fmt.Printf("Terima kasih sudah menggunakan\n")
		} else if choice == 1 || choice == 2 {

			fmt.Printf("Silahkan masukkan text yang ingin dicek: ")
			var text string
			scanner.Scan()
			text = scanner.Text()

			fmt.Printf("Silahkan masukkan pattern yang ingin digunakan: ")
			var pattern string
			scanner.Scan()
			pattern = scanner.Text()

			var res int

			if choice == 1 {
				res = matchingMethod.KMPMatch(text, pattern)
			} else {
				res = matchingMethod.BMMatch(text, pattern)
			}

			if res == -1 {
				fmt.Println("Pattern tidak ditemukan")
			} else {
				fmt.Printf("Pattern dimulai pada posisi %d\n", res)
			}
		} else {
			fmt.Println("Input tidak valid")
		}
	}
}
