package main

import (
	"fmt"
	"math"
)

func hitungLuasKelilingLingkaran(r float64, l, k *float64) {
	*l = math.Pi * r * r
	*k = 2 * math.Pi * r
}

func hitungLuasKelilingPersegi(s float64, l, k *float64) {
	*l = s * s
	*k = 4 * s
}

func hitungTotal(ll, lp, kl, kp float64, toluas, totKel *float64) {
	*toluas = ll + lp
	*totKel = kl + kp
}

func main() {
	var r, s, ll, lp, kl, kp, toluas, totKel float64


	for {
		fmt.Scan(&r, &s)
		
		if r == 0 && s == 0 {
			break
		}
		fmt.Println("R      S      LL      LP      KL      KP      TL      TP")
		hitungLuasKelilingLingkaran(r, &ll, &kl)
		hitungLuasKelilingPersegi(s, &lp, &kp)
		hitungTotal(ll, lp, kl, kp, &toluas, &totKel)
		fmt.Printf("%.2f  %.2f  %.2f  %.2f  %.2f  %.2f  %.2f  %.2f\n", r, s, ll, lp, kl, kp, toluas, totKel)
	}
}