package main

import (
	"StringMatching/matchingMethod"
	"fmt"
)

func main() {
	var choice int

	// reader := bufio.NewReader(os.Stdin)

	fmt.Println("Selamat datang di Pattern Matching")
	fmt.Println("Ada 2 algoritma yang tersedia, silahkan pilih")
	fmt.Println("1. Knuth-Morris-Pratt Algoritm")
	fmt.Println("2. Boyer-Moore Algorithm")
	fmt.Println("3. Exit")

	fmt.Println("Silahkan masukkan angka pilihan anda (1-3)")

	for cond := true; cond; cond = (choice < 1 || choice > 3) {
		fmt.Scanf("%d", &choice)

		if choice == 3 {
			fmt.Println("Terima kasih sudah menggunakan")
		} else if choice == 1 || choice == 2 {

			fmt.Print("Silahkan masukkan text yang ingin dicek: ")
			var text string
			fmt.Scan(&text)

			fmt.Print("Silahkan masukkan pattern yang ingin digunakan: ")
			var pattern string
			fmt.Scan(&pattern)

			var res int

			if choice == 1 {
				res = matchingMethod.KMPMatch(text, pattern)
			} else {
				res = matchingMethod.KMPMatch(text, pattern)
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
