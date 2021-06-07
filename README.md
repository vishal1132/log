A commit log written in go.

### Idea
Write logs to append only file encoded into binary, Read the logs via offset.
Appending to a single file is going to be very costly. Therefore a log will be breaken into multiple segments, and each segment will have it's index file and the log file, in which the log is actually stored.

Reading a log from the offset is therefore going to be a 2 step process, first it will read from the segments index file about the byte offset of the log in the logfile, then it will read the corresponding log. 

We are going to maintain a slice of pointer to segments, which will tell the base offset that segment is going to deal with. 

Segment is an abstraction over index and logfile, while log is an abstraction over segments.

Smallest unit of storage of log is logfile, and it's index file is index, to lookup logs faster.

For now there is just going to be an in memory map for index file. But later we will use mmap to load this file and write in to it. Also try to make an algorithm to work with sparse indices.