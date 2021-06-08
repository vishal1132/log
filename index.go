package log

type index struct {
	// index is the map of offset to byteoffset in the file.
	// this is just going an in memory map right now, will write to a file.
	// and use mmap to memory map the index file for faster use.
	index      map[uint64]uint64
	currOffset uint64
}

func newIndex() *index {
	return &index{
		index:      make(map[uint64]uint64),
		currOffset: 0,
	}
}

// write will map the offset to the position
func (i *index) write(offset uint64, position uint64) error {
	i.index[offset] = position
	return nil
}

// read will return the position for the offset
func (i *index) read(offset uint64) uint64 {
	// return the position for the offset
	return i.index[offset]
}
