package main

import (
	"aprende-go/streams-and-buffers/ioreader"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// A data stream as the name suggests is a stream of data.
// We can visualize this as the stream of water flowing from one end to another.
// Each droplet of water is the packet of data flowing in the stream and we can resize this data packet
// as per our needs.

// Buffer, on the other hand, is a temporary region of volatile memory (RAM)
// where data can be stored before consuming it.
// We can read from or write to a data buffer using streams.

func main() {
	// create a packet to stream
	p := make([]byte, 3) // slice of length `3`

	/****************************************/
	/*      Reading from a Data Source      */
	/****************************************/

	// A data-source is a data container (in memory or file storage) that can stream data.
	// This container implements basic interfaces provided by the io package and exposes some methods to read data from.

	fmt.Println()
	fmt.Println("Reading from a Data Source")

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

	fmt.Println()
	fmt.Println("io.Reader")

	// create data source
	customStringDataSource := ioreader.MyStringData{
		Str: "Hello Amazing World!",
	}

	// read `src` until an error is returned
	for {

		// read `p` bytes from `src`
		n, err := customStringDataSource.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		// handle error
		if err == io.EOF {
			fmt.Println("--end-of-source--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occured!", err)
			break
		}
	}

	/****************************************/
	/*         strings.NewReader            */
	/****************************************/

	// Go provides streaming capabilities over some built-in types.
	// To get an object from a string that implements the Read method, we can use strings.NewReader function.

	// func NewReader(s string) *strings.Reader

	// This function returns a pointer to strings.Reader struct that implements the Read method
	// defined by the io.Reader interface.

	fmt.Println()
	fmt.Println("strings.NewReader")

	// create data source
	stringDataSource := strings.NewReader("Hello Amazing World!")

	// read `src` until an error is returned
	for {

		// read `p` bytes from `src`
		n, err := stringDataSource.Read(p)
		fmt.Printf("%d bytes read, data: %s\n", n, p[:n])

		// handle error
		if err == io.EOF {
			fmt.Println("--end-of-source--")
			break
		} else if err != nil {
			fmt.Println("Oops! Some error occured!", err)
			break
		}
	}

	/****************************************/
	/*           ioutil.ReadAll             */
	/****************************************/

	// Read everything from a source which implements io.Reader interface,
	// then you can use ioutil.ReadAll method.

	/*
		func ReadAll(src io.Reader) ([]byte, error)
	*/

	// The ReadAll method will read from src until src returns an EOF error or some other uneventful error.
	// The ReadAll method will only return an error if the underlying source src returns this uneventful
	// error (other than EOF).

	fmt.Println()
	fmt.Println("ioutil.ReadAll")

	// create data source
	dataSrc := strings.NewReader("I want read all this message")

	// read all data from `stringDataSource`
	data, _ := ioutil.ReadAll(dataSrc)

	// print `data`
	fmt.Printf("Read data of length %d : %s\n", len(data), data)

	/****************************************/
	/*           ioutil.ReadFull            */
	/****************************************/

	// If you are interested in exactly x number of bytes from the source and somehow,
	// if x number of bytes could not be read, get an error instead, io.ReadFull is the function for you.

	/*
		func ReadFull(src Reader, buf []byte) (n int, err error)
	*/

	// he ReadFull function will read from the source src exactly len(buf) bytes.
	// In an ideal condition, number of bytes read n will be len(buf). If n is less than len(buf)
	// because src returned EOF, then ReadFull will return an io.ErrUnexpectedEOF error.

	// If ReadFull reads 0 bytes from src (when n=0), the error returned by the ReadFull is io.EOF.
	// In a special condition, if an error is returned by src after len(buf) bytes were read, that error
	// is dropped by ReadFull.

	fmt.Println()
	fmt.Println("ioutil.ReadFull")

	// create data source
	dataSrc2 := strings.NewReader("Hello Amazing World!") // 20 characters

	// create buffer of length 14
	buffer := make([]byte, 14)

	// call 1: read from `dataSrc2`
	bytesRead1, err1 := io.ReadFull(dataSrc2, buffer)
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead1, buffer[:bytesRead1], err1)

	// call 2: read from `dataSrc2`
	bytesRead2, err2 := io.ReadFull(dataSrc2, buffer)
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead2, buffer[:bytesRead2], err2)

	// call 3: read from `dataSrc2`
	bytesRead3, err3 := io.ReadFull(dataSrc2, buffer)
	fmt.Printf("Bytes read: %d, value: %s, err: %v\n", bytesRead3, buffer[:bytesRead3], err3)

	/****************************************/
	/*           io.LimitReader            */
	/****************************************/

	// Many at times, we need to set a cap on a data source. The io.LimitReader takes a Reader
	// object r and returns a Reader object with a cap of n bytes.

	/*
		func LimitReader(r Reader, n int64) Reader
	*/

	// The LimitReader function returns a Reader object that abstracts source r.
	// This returned Reader object will keep reading from r until n bytes are read
	// and it will return io.EOF error once n bytes are completely read.

	fmt.Println()
	fmt.Println("ioutil.LimitReader")

	// create a main data source
	mainSrc3 := strings.NewReader("Hello Amazing World!") // 20 characters

	// create data source from `mainSrc3` with cap of `10` bytes
	src3 := io.LimitReader(mainSrc3, 10)

	// create a packet
	buffer3 := make([]byte, 3) // slice of length `3`

	// read `src3` until an error is returned
	for {

		// read `buffer3` bytes from `src3`
		n3, err3 := src3.Read(buffer3)
		fmt.Printf("%d bytes read, data: %s\n", n3, buffer3[:n3])

		// handle error
		if err3 == io.EOF {
			fmt.Println("--end-of-file--")
			break
		} else if err3 != nil {
			fmt.Println("Oops! Some error occured!", err3)
			break
		}
	}

	/****************************************/
	/*      Writing to a Data Store         */
	/****************************************/

	// A data-store is a data container (in memory or file storage) that can store an incoming stream of data.
	// This container implements basic interfaces provided by the io package and exposes some methods to write data to.

	fmt.Println()
	fmt.Println("Writing to a Data Store")

	/****************************************/
	/*           io.Writter                 */
	/****************************************/

	// The io.Writer interface declares the basic Write method.
	// Any type that implements this interface is a type of Writer.

	/*
		type Writer interface {
			Write(p []byte) (n int, err error)
		}
	*/

	// The Write method accepts a data packet p which is a slice of byte and writes to the store.
	// It returns the number of bytes written to the store n and an error if the packet p could not be written.

	fmt.Println()
	fmt.Println("io.Writter ")

	ss1 := ioreader.SampleStore{}

	// write 1: "Hello!"
	packet1 := []byte("Hello! ")
	bytesWritten1, err1 := ss1.Write(packet1)
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten1, err1)
	fmt.Printf("Value of ss1.Data: %s\n\n", ss1.Data)

	// write 2: " Amazing"
	packet2 := []byte("Amazing ")
	bytesWritten2, err2 := ss1.Write(packet2)
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten2, err2)
	fmt.Printf("Value of ss1.Data: %s\n\n", ss1.Data)

	// write 3: " World!"
	packet3 := []byte("World!")
	bytesWritten3, err3 := ss1.Write(packet3)
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten3, err3)
	fmt.Printf("Value of ss1.Data: %s\n\n", ss1.Data)

	/**************************************************/
	/*               io.WritteString                  */
	/**************************************************/

	// To make things simple, io.WriteString function writes bytes from a string to the source by making
	// source.Write() call internally and return the result of that Write call.

	fmt.Println()
	fmt.Println("io.WritteString")

	// func WriteString(w Writer, s string) (n int, err error)

	// If the source implements io.StringWriter interface, WriteString function will call the WriteString
	// method instead and return the result of the WriteString call.

	/*
		type StringWriter interface {
		  WriteString(s string) (n int, err error)
		}
	*/

	ss2 := &ioreader.SampleStore{}

	// write 1: "Hello!"
	bytesWritten4, err4 := io.WriteString(ss2, "Hello!")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten4, err4)
	fmt.Printf("Value of ss.data: %s\n\n", ss2.Data)

	// write 2: " Amazing"
	bytesWritten5, err5 := io.WriteString(ss2, " Amazing")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten5, err5)
	fmt.Printf("Value of ss.data: %s\n\n", ss2.Data)

	// write 3: " World!"
	bytesWritten6, err6 := io.WriteString(ss2, " World!")
	fmt.Printf("Bytes written %d, error: %v\n", bytesWritten6, err6)
	fmt.Printf("Value of ss.data: %s\n\n", ss2.Data)

	/********************************************/
	/*          Standard I/O Streams            */
	/********************************************/

	fmt.Println()
	fmt.Println("Standard I/O Streams")

	// The standard I/O streams viz. os.Stdin, os.Stdout and os.Stderr implement the io.Writer and
	// io.StringWriter interfaces. Hence we can call their Write or WriteString methods to write some data.
	// The fmt package also provides some functions to write some data to io.Writer objects.

	// use `io.WriteString` to write to a `io.Writer`
	io.WriteString(os.Stdout, "Hello World!\n")

	// call `Write` method of of a `io.Writer`
	os.Stdout.Write([]byte("Hello World!\n"))

	// use `fmt` package function to write to a `io.Writer`
	fmt.Fprint(os.Stdout, "Hello World!\n")
	fmt.Fprintln(os.Stdout, "Hello World!") // adds new line
	fmt.Fprintf(os.Stdout, "%s, World!\n", "Hello")

	/********************************************/
	/*        Closing I/O Operations            */
	/********************************************/

	fmt.Println()
	fmt.Println("Closing I/O Operations")

	// When we have an object that implements the io.Reader or io.Writer interface, we can call Read() or Write()
	// method on it any time we want.

	// But once we are done with either read or write operations, that object is no longer needed for any I/O
	// operations. Since to provide I/O capabilities may take additional system resources, releasing such
	// resources that are no longer needed may boost the performance and reliability of our application.

	// Hence a io.Reader or io.Writer object may implement additional Close method that releases such
	// resources and closes the object for any I/O operations. The io.Closer interface declares the Close method.

	/*
				type Closer interface {
		    	Close() error
				}
	*/

	// The implementation of the Close method is completely up to us. But the first Close call should not
	// return any error and should only perform cleanup operations, while subsequent Close calls can return
	// some error. Any Read or Write call after the object is closed for I/O should return an error.

	// The io packages also provide some interfaces that groups the Read, Write and Close methods as shown below.

	/*
			type ReadCloser interface {
		    Reader
		    Closer
			}

			type WriteCloser interface {
				Writer
				Closer
			}

			type ReadWriteCloser interface {
				Reader
				Writer
				Closer
			}
	*/

	/***********************************************/
	/*      Transferring Data between streams      */
	/***********************************************/

	fmt.Println()
	fmt.Println("Transferring Data between streams")

	// When we have a source src of type io.Reader and a store dst of type io.Writer and we wish to transfer
	// data between these two, there are quite a lot of ways to make that happen.

	// Manually, we can read one packet of data at a time from src using Read() call and write to dst using
	// the a Write() call. But io package provides some function to facilitate this kind of data transfer.

	/*********************/
	/*      io.Copy      */
	/*********************/

	// If we want to write data to an io.Writer object coming from an io.Reader object,
	// then we can use the io.Copy function.

	fmt.Println()
	fmt.Println("io.Copy")

	/*
		func Copy(dst Writer, src Reader) (int64, error)
	*/

	// The Copy function reads data from the src and writes to the dst until io.EOF error is returned by the
	// src (or other uneventful error). It returns the number of bytes copied from the src to the dst.

	// The io.EOF error returned by src is silently dropped. However, if Read or Write operation returns an error,
	// Copy will return that error along with the number of bytes copied to dst until that error has occurred.

	// If src implements the io.WriterTo interface, then Copy function calls src.WriteTo(dst) internally.
	// Otherwise, if dst implements the ReaderFrom interface, the Copy calls dst.ReadFrom(src).

	// create a string `io.Reader` object
	mainSrc4 := strings.NewReader("Hello World! How are you?\n")

	// copy data from `stringsReader` to `os.Stdout` (`io.Writer`)
	io.Copy(os.Stdout, mainSrc4)

	/**********************/
	/*      io.CopyN      */
	/**********************/

	// If we want to write only n bytes from an io.Reader to an io.Writer, we can use io.CopyN function.
	// This function creates a io.LimitReader internally.

	// If dst implements the io.ReaderFrom interface, CopyN function uses dst.ReadFrom method to copy n bytes
	// from src. This function returns the number of bytes copied to dst which ideally should be equal to n.

	/*
		func CopyN(dst Writer,stringReader src Reader, n int64) (int64, error)
	*/

	fmt.Println()
	fmt.Println("io.CopyN")

	// create a string `io.Reader` object
	mainSrc5 := strings.NewReader("Hello World! How are you?\n")

	// copy `12 bytes` data from `mainSrc5`
	// to `os.Stdout` (`io.Writer`)
	io.CopyN(os.Stdout, mainSrc5, 12)

	/**********************/
	/*       io.Pipe      */
	/**********************/

	fmt.Println()
	fmt.Println("io.Pipe")

	// In the previous topic, we discussed that io.Copy function is useful to transfer all the data
	// available with io.Reader to the io.Writer in one go. The Copy function will read data from io.Reader
	// until io.EOF error is returned.

	// The io.Pipe function, on the other hand, creates a synchronous pipe between io.Reader and io.Writer.
	// When some data is written to the io.Writer, it is available to be read by io.Reader instantly.

	/*
		func Pipe() (*PipeReader, *PipeWriter)
	*/

	// The Pipe function returns *io.PipeReader and *io.PipeWriter objects. These are object of struct
	// type that implements ReadCloser and WriteCloser interfaces respectively. Hence we can make Read()
	// call on the PipeReader to read some data and Write() call on PipeWriter to write some data.

	/*
		src, dst := io.Pipe()
	*/

	// The io.Pipe() call returns a io.PipeReader object src and io.PipeWriter object dst. When we read
	// from src using src.Read(p), it will return data available in the pipe. We can make as many Read()
	// calls as possible until no data is left in the pipe.

	// If the Read() doesn’t find any data in the pipe, the current goroutine blocks and Go schedules
	// another goroutine that may write data to dst using dst.Write(data). Each Write() call blocks the
	// current goroutine.

	// Generally, each Read() call should extract all the data available in the pipe but we can make multiple
	// Read() calls until all the data from the pipe is extracted. That’s when the current goroutine will block
	// to schedule other goroutines with a possible Write() call to get additional data from.

	// The dst.Close() method closes the write operations on dst. Any Read operations on closed dst will
	//  return 0 bytes read and io.EOF error. If src is closed for any read using src.Close() call, then
	//  any Write operations on dst will return io.ErrClosedPipe error.

	// create a pipe
	src5, dst5 := io.Pipe()

	// start goroutine that writes data to `dst5`
	go func() {
		dst5.Write([]byte("DATA_1")) // write and block
		dst5.Write([]byte("DATA_2")) // write and block
		dst5.Close()                 // indicate EOF
	}()

	// data transfer packet
	packet4 := make([]byte, 6)

	// read from `src`
	bytesRead7, err7 := src5.Read(packet4)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead7, packet4, err7)

	// read from `src`
	bytesRead8, err8 := src5.Read(packet4)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead8, packet4, err8)

	// read from `src`
	bytesRead9, err9 := src5.Read(packet4)
	fmt.Printf("bytes read: %d, value: %s, err: %v\n", bytesRead9, packet4, err9)

	/*****************************/
	/*      Buffered streams     */
	/*****************************/

	fmt.Println()
	fmt.Println("Buffered streams")

	// A buffer is a region of space in the memory. It can be a fixed or a
	// variable size buffer to read data from or write data to. The bytes
	// built-in package provides Buffer structure type to construct a variable
	// size buffer.

	// We can also use bytes.NewBuffer function which initializes a variable
	// size buffer with initial buffer value. The bytes.NewBufferString
	// function does the same thing but creates a new buffer with bytes of a
	// string as its initial value.

	/*
		func NewBuffer(buf []byte) *bytes.Buffer
		func NewBufferString(s string) *bytes.Buffer
	*/

	// The Buffer structure provides some useful methods to read and write
	// data in multiple fashions. Apart from the usual Read, Write,
	// ReadFrom and WriteTo methods, Buffer implements WriteString method to
	// write data from a string, Reset to empty the buffer, Grow to increase
	// the capacity of the buffer and other useful methods that you can read
	// from here
	// https://golang.org/pkg/bytes/#Buffer

	// create new buffer
	buf := bytes.NewBufferString("Hello World!")

	// write some data to the buffer
	fmt.Print("bytes written => ")
	fmt.Println(buf.WriteString("How are you?"))

	// append data from a `io.Reader` to the buffer
	strReader := strings.NewReader(" Doing Well? ")
	fmt.Print("bytes written => ")
	fmt.Println(buf.ReadFrom(strReader))

	// read first `12 bytes` from the buffer
	fmt.Print("bytes read => ")
	fmt.Println(buf.Read(make([]byte, 12)))

	// read all `unread bytes` to STDOUT
	fmt.Print("bytes read => ")
	fmt.Println(buf.WriteTo(os.Stdout))

	// read another `10 bytes` from the buffer
	fmt.Print("bytes read => ")
	fmt.Println(buf.Read(make([]byte, 10))) // EOF

	// write some more data to the buffer
	fmt.Print("bytes written => ")
	fmt.Println(buf.WriteString("Hello! "))

	// read all `unread bytes` to STDOUT
	fmt.Print("bytes read => ")
	fmt.Println(buf.WriteTo(os.Stdout))

	// read another `10 bytes` from the buffer
	fmt.Print("bytes read => ")
	fmt.Println(buf.Read(make([]byte, 10)))
}
