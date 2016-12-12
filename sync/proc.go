// proc.go
package sync

func ProcPin() int {
	return runtime_procPin()
}

func ProcUnpin() {
	runtime_procUnpin()
}
