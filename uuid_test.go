package uuid

import "testing"

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
			t.Fatalf(err.Error())
		}
		s[n] = u
	}

	if len(dedup(s)) != len(s) {
		t.Errorf("duplicates generated : Set-%d Seq-%d", len(dedup(s)), len(s))
	}
}

func TestVersion(t *testing.T) {
	u, err := Generate()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if u[14] != '4' {
		t.Errorf("invalid version expected 4 got %c", u[14])
	}
}

func TestVariant(t *testing.T) {
	u, err := Generate()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if u[19] != '8' {
		t.Errorf("invalid variant expected 8 got %c", u[19])
	}
}

func TestCapacity(t *testing.T) {
	u, err := Generate()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if cap(u) != 36 {
		t.Errorf("expected capacity 36 actual %d", cap(u))
	}
}


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
