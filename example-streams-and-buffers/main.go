package main

// A data stream as the name suggests is a stream of data.
// We can visualize this as the stream of water flowing from one end to another.
// Each droplet of water is the packet of data flowing in the stream and we can resize this data packet
// as per our needs.

// Buffer, on the other hand, is a temporary region of volatile memory (RAM)
// where data can be stored before consuming it.
// We can read from or write to a data buffer using streams.

func main() {
	/****************************************/
	/*      Reading from a Data Source      */
	/****************************************/

	// A data-source is a data container (in memory or file storage) that can stream data.
	// This container implements basic interfaces provided by the io package and exposes some methods to read data from.

	/****************************************/
	/*            io.Reader                 */
	/****************************************/

	// To be able to read data from a source as a stream, the source must implement io.Reader interface.
	// The io.Reader interface declares the basic Read method. Any type that implements the Read method is a
	// type of Reader

	/*
		type Reader interface {
			Read(p []byte) (n int, err error)
		}
	*/
}
