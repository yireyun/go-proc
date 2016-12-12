// proc_test
package proc

import (
	"testing"
)

func TestProc(t *testing.T) {
	pid := ProcPin()
	t.Logf("PID: %v", pid)
	ProcUnpin()
}

func BenchmarkProc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProcPin()
		ProcUnpin()
	}
}

func BenchmarkProcParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ProcPin()
			ProcUnpin()
		}
	})
}
