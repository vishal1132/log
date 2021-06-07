package log

type Config struct {
	Segment struct {
		MaxStoreBytes uint64
		MaxIndexBytes uint64
		InitialOffset uint64
		BaseDir       string
		Log           struct {
			InMemorySize int
		}
	}
}
