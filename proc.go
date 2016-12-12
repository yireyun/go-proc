// proc.go
package proc

import (
	"sync"
)

func ProcPin() int {
	return sync.ProcPin()
}

func ProcUnpin() {
	sync.ProcUnpin()
}
