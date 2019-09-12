package failing

import "testing"

func BenchmarkSolveProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SolveProblem()
	}
}
