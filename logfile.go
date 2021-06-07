// logfile.go is the actual file write imoplementation of the log.
package log

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"path"
	"sync"
)

type logfile struct {
	mu        sync.Mutex
	file      *os.File
	bufWriter *bufio.Writer
}

// newLogFile creates a new log file for the segment
func newLogFile(filename string, conf Config) (*logfile, error) {

	f, err := os.OpenFile(path.Join(conf.Segment.BaseDir, fmt.Sprintf("%s.log", filename)), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &logfile{
		file:      f,
		bufWriter: bufio.NewWriterSize(f, conf.Segment.Log.InMemorySize),
	}, nil
}

func (l *logfile) close() {
	l.file.Close() // error deliberately ignored
}

// write actually writes to the file. In the first 8 bytes it is going to write
// the length of the log that we are going to write, encoded using bigendian,
// therefore decode with bigendian as well.
// // Then without any offset, we write the log
func (l *logfile) write(value []byte) error {
	if err := binary.Write(l.bufWriter, binary.BigEndian, uint64(len(value))); err != nil {
		return err
	}
	if _, err := l.bufWriter.Write(value); err != nil {
		return err
	}
	return nil
}

func (l *logfile) read(offset uint64) ([]byte, error) {
	// implement the read method here
	return nil, nil
}
