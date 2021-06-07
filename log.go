package log

import "sync"

// Log is the abstraction of the whole commit log.
// It satisfies Logger interface.
type Log struct {
	// Right now going to take a lock over the complete log.
	// We might have a round robin segments in future and therefore only
	// benefit from taking lock over that segment
	mu sync.Mutex
}

// Logger interface is the set of methods that a standard log should provide.
// Writing the log, Reading the log.
// Deleting the log is not going to be implemented in near future.
// It can be implemented by marking a record as delete, and then compact the log
// in a separate goroutine.
type Logger interface {
	// Write writes the value in the log,
	Write(value []byte) (uint64, error)
	Read(offset uint64) ([]byte, error)
}

// New returns a new log instance
func New() {

}
