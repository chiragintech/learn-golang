package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func SendRequest() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	// body of the response is io.ReadCloser - interface that has Reader and Closer
	// Reader is an interface with Read([]byte) (int, error)
	// Closer is an interface with Close() (error)
	// Reader interface solves the issue of multiple sources of input like body, text file, image file etc
	// Basically, a byte slice output
	// You pass body to a byte slice for Read which can be read by passing a byte slice, outputs the number of bytes that were read

	bs := make([]byte, 99999) // arbitrary large byte slice
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	lw := logWriter{}
	io.Copy(lw, resp.Body)

	io.Copy(os.Stdout, resp.Body) // []byte -> writer -> some form of output
	// copy -> (writer, reader)
	// Stdout -> *File -> writer
	// Writer is similar to a Reader
	// Copy uses a copy buffer that pulls data through Read and Write one by one using the buffer

}

func (logWriter) Write(bs []byte) (int, error) {
	// logWriter is implementing the Writer interface
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}