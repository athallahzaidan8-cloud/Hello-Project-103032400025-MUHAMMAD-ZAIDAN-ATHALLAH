// hasil nya nya

package main

import "fmt"

func menu() {
	fmt.Println("--- M E N U ---")
	fmt.Println("1. Hitung Penjumlahan")
	fmt.Println("2. Hitung Perkalian")
	fmt.Println("3. Hitung Pembagian")
	fmt.Println("4. Exit")
}

func hitungJumlah() {
	var a, b int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil penjumlahan:", a+b)
}

func hitungKali() {
	var a, b int
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&a, &b)
	fmt.Println("Hasil perkalian:", a*b)
}

func hitungBagi() {
	var a, b float64
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&a, &b)
	if b == 0 {
		fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
	} else {
		fmt.Println("Hasil pembagian:", a/b)
	}
}

func main() {
	var pilih int

	for {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			hitungJumlah()
		case 2:
			hitungKali()
		case 3:
			hitungBagi()
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}
