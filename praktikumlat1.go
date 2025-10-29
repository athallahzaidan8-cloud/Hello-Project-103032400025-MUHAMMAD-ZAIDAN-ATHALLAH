package main

import "fmt"

const NMAX int = 20
type tabInt [NMAX] int 

func main(){
	var dataAwal, dataHasil tabInt
	var nAwal, nHasil int
	var a, b float64
	
	baca (&dataAwal, &nAwal)
	fmt.Scan (&a, &b)
	jumlahNilai(&dataAwal, &dataHasil, &nAwal, &nHasil, a b)
	
	fmt.Println ("Statistika Tim B")
	cetak (dataHasil, nHasil)
}

func baca (A* tabInt, n* int){
	fmt.Scan (n)
	
	
}

func jumlahNilai (A tabInt, n int){
	
	
} 

func cetak (B tabInt, n int){
	
	
}