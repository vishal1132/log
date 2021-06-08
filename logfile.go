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
	size      uint64
}

var encoding = binary.BigEndian

// newLogFile creates a new log file for the segment
func newLogFile(filename string, conf Config) (*logfile, error) {
	f, err := os.OpenFile(path.Join(conf.Segment.BaseDir, fmt.Sprintf("%s.log", filename)), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}
	return &logfile{
		file:      f,
		bufWriter: bufio.NewWriterSize(f, conf.Segment.Log.InMemorySize),
		size:      0,
	}, nil
}

func (l *logfile) close() {
	l.file.Close() // error deliberately ignored
}

// write actually writes to the file. In the first 8 bytes it is going to write
// the length of the log that we are going to write, encoded using bigendian,
// therefore decode with bigendian as well.
// Then without any width, we write the log
// Returns offset and error if the record can not be inserted
func (l *logfile) write(value []byte) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	offset := l.size
	if err := binary.Write(l.bufWriter, encoding, uint64(len(value))); err != nil {
		return 0, err
	}
	l.size += 8

	/*
		what if the number of bytes is written but not the entire log.
		should we remove?
		not if we adjust the size of the log here and not after writing the complete log, except from worthless size point of view.
		this might be addressed later.
	*/

	w, err := l.bufWriter.Write(value)
	if err != nil {
		return 0, err
	}

	l.size += uint64(w)
	return offset, nil
}

func (l *logfile) read(offset uint64) ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Flush anything already in the buffer on the file
	l.bufWriter.Flush()

	// 2 step read process.
	// first read 64 bits from the file- this is the length of the bytes that you have to read
	// now read those bytes and return

	logSizeRead := make([]byte, 8)
	if _, err := l.file.ReadAt(logSizeRead, int64(offset)); err != nil {
		return nil, err
	}

	// encoding.Uint64(logSizeRead) is the size of the log decoded by encoding
	logRead := make([]byte, encoding.Uint64(logSizeRead))

	if _, err := l.file.ReadAt(logRead, int64(offset)+8); err != nil {
		return nil, err
	}

	return logRead, nil
}

func (l *logfile) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.bufWriter.Flush() // error deliberately ignored
	l.file.Close()      // error deliberately ignored
}
