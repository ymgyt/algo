package main

// Comb return combination patterns from n to r
func Comb(n, r int) int {
	return Permutation(n, r) / Factorical(r)
}

// Factorical return n!
func Factorical(n int) int {
	return Permutation(n, n-1)
}

// Permutation return nPr
func Permutation(n, r int) int {
	if n < r {
		return 0
	}
	if n == 0 {
		return 1
	}

	v := 1
	for i := 0; i < r; i++ {
		v *= (n - i)
	}
	return v
}
