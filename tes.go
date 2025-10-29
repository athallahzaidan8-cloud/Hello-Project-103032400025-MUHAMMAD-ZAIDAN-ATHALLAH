package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Team struct {
	Name    string
	Players []string
	Wins    int
	Losses  int
}

type Match struct {
	Team1  string
	Team2  string
	Score1 int
	Score2 int
	Date   string
}

var teams []Team
var matches []Match

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n=== APLIKASI PENGELOLAAN TURNAMEN ESPORT ===")
		fmt.Println("1. Tambah Data Tim")
		fmt.Println("2. Tambah Hasil Pertandingan")
		fmt.Println("3. Tampilkan Daftar Tim")
		fmt.Println("4. Tampilkan Hasil Pertandingan")
		fmt.Println("5. Cari Tim")
		fmt.Println("6. Tampilkan Klasemen")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		scanner.Scan()
		choice := scanner.Text()
		switch choice {
		case "1":
			tambahTim(scanner)
		case "2":
			tambahPertandingan(scanner)
		case "3":
			tampilkanTim()
		case "4":
			tampilkanPertandingan()
		case "5":
			cariTim(scanner)
		case "6":
			tampilkanKlasemen()
		case "0":
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahTim(scanner *bufio.Scanner) {
	fmt.Print("Nama tim: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("Jumlah pemain: ")
	scanner.Scan()
	jumlahStr := scanner.Text()
	jumlah, _ := strconv.Atoi(jumlahStr)

	var pemain []string
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Nama pemain %d: ", i+1)
		scanner.Scan()
		pemain = append(pemain, scanner.Text())
	}

	tim := Team{nama, pemain, 0, 0}
	teams = append(teams, tim)
	fmt.Println("Tim berhasil ditambahkan!")
}

func tambahPertandingan(scanner *bufio.Scanner) {
	fmt.Print("Nama tim 1: ")
	scanner.Scan()
	t1 := scanner.Text()
	fmt.Print("Skor tim 1: ")
	scanner.Scan()
	s1, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Nama tim 2: ")
	scanner.Scan()
	t2 := scanner.Text()
	fmt.Print("Skor tim 2: ")
	scanner.Scan()
	s2, _ := strconv.Atoi(scanner.Text())

	fmt.Print("Tanggal pertandingan (YYYY-MM-DD): ")
	scanner.Scan()
	tgl := scanner.Text()

	matches = append(matches, Match{t1, t2, s1, s2, tgl})

	for i := range teams {
		if teams[i].Name == t1 {
			if s1 > s2 {
				teams[i].Wins++
			} else {
				teams[i].Losses++
			}
		} else if teams[i].Name == t2 {
			if s2 > s1 {
				teams[i].Wins++
			} else {
				teams[i].Losses++
			}
		}
	}

	fmt.Println("Hasil pertandingan berhasil ditambahkan!")
}

func tampilkanTim() {
	fmt.Println("\n--- DAFTAR TIM ---")
	for _, t := range teams {
		fmt.Printf("Nama: %s\n", t.Name)
		fmt.Printf("Pemain: %s\n", strings.Join(t.Players, ", "))
		fmt.Printf("Menang: %d, Kalah: %d\n\n", t.Wins, t.Losses)
	}
}

func tampilkanPertandingan() {
	fmt.Println("\n--- HASIL PERTANDINGAN ---")
	for _, m := range matches {
		fmt.Printf("%s vs %s | %d - %d | %s\n", m.Team1, m.Team2, m.Score1, m.Score2, m.Date)
	}
}

func cariTim(scanner *bufio.Scanner) {
	insertionSortByName() // harus urut dulu untuk binary search

	fmt.Print("Masukkan nama tim yang dicari: ")
	scanner.Scan()
	nama := scanner.Text()

	idx := binarySearchTeam(strings.ToLower(nama))
	if idx != -1 {
		t := teams[idx]
		fmt.Printf("\nNama: %s\nPemain: %s\nMenang: %d, Kalah: %d\n", t.Name, strings.Join(t.Players, ", "), t.Wins, t.Losses)
	} else {
		fmt.Println("Tim tidak ditemukan.")
	}
}

func tampilkanKlasemen() {
	insertionSortByWins()
	fmt.Println("\n--- KLASEMEN ---")
	for i, t := range teams {
		fmt.Printf("%d. %s - Menang: %d | Kalah: %d\n", i+1, t.Name, t.Wins, t.Losses)
	}
}

// Insertion Sort untuk klasemen berdasarkan jumlah menang (desc)
func insertionSortByWins() {
	for i := 1; i < len(teams); i++ {
		key := teams[i]
		j := i - 1
		for j >= 0 && teams[j].Wins < key.Wins {
			teams[j+1] = teams[j]
			j--
		}
		teams[j+1] = key
	}
}

// Insertion Sort untuk pengurutan nama tim (asc)
func insertionSortByName() {
	for i := 1; i < len(teams); i++ {
		key := teams[i]
		j := i - 1
		for j >= 0 && strings.ToLower(teams[j].Name) > strings.ToLower(key.Name) {
			teams[j+1] = teams[j]
			j--
		}
		teams[j+1] = key
	}
}

// Binary Search untuk mencari nama tim
func binarySearchTeam(target string) int {
	low := 0
	high := len(teams) - 1

	for low <= high {
		mid := (low + high) / 2
		midName := strings.ToLower(teams[mid].Name)

		if midName == target {
			return mid
		} else if midName < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
