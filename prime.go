package main

import "math"

// PrimeFactors return factors of n. int slice is sorted (ascendent)
func PrimeFactors(n int) []int {
	if n <= 1 {
		return nil
	}
	if n == 2 {
		return []int{2}
	}

	var ns []int

	for {
		if n%2 == 0 {
			ns = append(ns, 2)
			n = n / 2
		} else {
			break
		}
	}

	// n is odd
	for i := 3; i <= int(math.Sqrt(float64(n))); i = i + 2 {
		for {
			if n%i == 0 {
				ns = append(ns, i)
				n = n / i
			} else {
				break
			}
		}
	}

	if n > 2 {
		ns = append(ns, n)
	}

	return ns
}
