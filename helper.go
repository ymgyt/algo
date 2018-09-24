package main

import "testing"

func error(t *testing.T, diff string) {
	t.Helper()
	t.Errorf("(-got +want)\n%s", diff)
}
