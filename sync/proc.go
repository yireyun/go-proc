// proc.go
package sync

func ProcPin() int {
	return runtime_procPin()
}

func ProcUnpin() {
	runtime_procUnpin()
}

func Semacquire(s *uint32) {
	runtime_Semacquire(s)
}

func Semrelease(s *uint32) {
	runtime_Semrelease(s)
}
