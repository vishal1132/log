package log

type segment struct {
	index   *index
	logfile *logfile
	offset  uint64
}

func (c *Config) newSegment(file string) *segment {
	log, err := newLogFile("somefile", *c)
	if err != nil {
		return nil
	}
	return &segment{
		logfile: log,
		index:   newIndex(),
	}
}

func (s *segment) write(value []byte) (uint64, error) {
	pos, err := s.logfile.write(value)
	if err != nil {
		return 0, err
	}
	s.index.write(s.offset, pos)
	s.offset++
	return pos, nil
}

func (s *segment) read(offset uint64) ([]byte, error) {
	position := s.index.read(offset)
	return s.logfile.read(position)
}
