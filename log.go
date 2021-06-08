package log

import "sync"

// Log is the abstraction of the whole commit log.
// It satisfies Logger interface.
type Log struct {
	// Right now going to take a lock over the complete log.
	// We might have a round robin segments in future and therefore only
	// benefit from taking lock over that segment
	mu sync.Mutex

	// activeSegment is the active segment to which the log is going to be written
	activeSegment *segment

	// segments is the slice of pointers to segments ( old segments )
	segments []*segment
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

func loadConfs() *Config {
	c := &Config{}
	c.Segment.MaxStoreBytes = 10000
	c.Segment.BaseDir = "/Users/vishal/work/src/github.com/vishal1132/commitlog"
	c.Segment.MaxIndexBytes = 10000
	return c
}

// Write writes the value to the log.
// It is concurrent safe.
func (l *Log) Write(value []byte) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	return l.activeSegment.write(value)
}

// New returns a new log instance
func New() *Log {
	c := loadConfs()
	return &Log{
		activeSegment: c.newSegment("testfile"),
	}
}

// Read reads the value set by the Write()
func (l *Log) Read(offset uint64) ([]byte, error) {
	return l.activeSegment.read(offset)
}
