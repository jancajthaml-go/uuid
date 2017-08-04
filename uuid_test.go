package main

import "testing"

func BenchmarkUUID(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UUID()
	}
}

func dedup(x [][]byte) [][]byte {
	visited := map[string]bool{}
	r := [][]byte{}

	for v := range x {
		w := string(x[v])
		if !visited[w] {
			visited[w] = true
			r = append(r, x[v])
		}
	}
	return r
}

func TestUUID(t *testing.T) {
	s := make([][]byte, 1000)

	for n := 0; n < 1000; n++ {
		s[n] = UUID()
	}

	if len(dedup(s)) != len(s) {
		t.Errorf("Duplicates generated : Set-", len(dedup(s)), "Seq-", len(s))
	}
}
