package uuid

import "testing"

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Generate()
	}
}

func BenchmarkGenerateParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Generate()
		}
	})
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

func TestGenerate(t *testing.T) {
	s := make([][]byte, 1000)

	for n := 0; n < 1000; n++ {
		u, err := Generate()
		if err != nil {
			t.Errorf("expected nil error but got %+v", err)
			return
		}
		s[n] = u
	}

	if len(dedup(s)) != len(s) {
		t.Errorf("duplicates generated : Set-%d Seq-%d", len(dedup(s)), len(s))
	}
}
