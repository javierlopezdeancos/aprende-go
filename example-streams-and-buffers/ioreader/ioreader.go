package ioreader

import "io"

// MyStringData - simple struct to hold string data
type MyStringData struct {
	Str       string
	readIndex int // default: 0
}

// add `Read` method (to impletement io.Reader)
func (myStringData *MyStringData) Read(p []byte) (n int, err error) {
	// convert `str` string to a slice of bytes
	strBytes := []byte(myStringData.Str)

	// if `readIndex` is GTE source length, return `EOF` error
	if myStringData.readIndex >= len(strBytes) {
		return 0, io.EOF // `0` bytes read
	}

	// get next readable limit (exclusive)
	nextReadLimit := myStringData.readIndex + len(p)

	// if `nextReadLimit` is GTE source length
	// set `nextReadLimit` to source length and `err` to `EOF`
	if nextReadLimit >= len(strBytes) {
		nextReadLimit = len(strBytes)
		err = io.EOF
	}

	// get next bytes to copy and set `n` to its length
	nextBytes := strBytes[myStringData.readIndex:nextReadLimit]
	n = len(nextBytes)

	// copy all bytes of `nextBytes` into `p` slice
	copy(p, nextBytes)

	// increment `readIndex` to `nextReadLimit`
	myStringData.readIndex = nextReadLimit

	// return values
	return
}

// SampleStore - sample store type
type SampleStore struct {
	Data []byte
}

// implement `io.Writer` interface
func (ss *SampleStore) Write(p []byte) (n int, err error) {

	// check if `18` bytes has been written
	if len(ss.Data) == 21 {
		return 0, io.EOF // end of limit error
	}

	// get remaining capacity of the `ss.data`
	remainingCap := 21 - len(ss.Data)

	// get length of data to write
	writeLength := len(p)
	if remainingCap <= writeLength {
		writeLength = remainingCap
		err = io.EOF
	}

	// append `writeLength` of data from `p` to `ss.data`
	ss.Data = append(ss.Data, p[:writeLength]...)

	// set number of bytes written and return
	n = writeLength
	return
}
