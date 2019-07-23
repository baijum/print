package print

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var logfile *os.File

// Println write the messaage to a file in the TEMP directory.
func Println(msg string) {
	_, err := fmt.Fprintf(logfile, msg+"\n")
	if err != nil {
		log.Fatal(err)
	}
}

// Close the log file descriptor
func Close() error {
	return logfile.Close()
}

func init() {
	pid := strconv.Itoa(os.Getpid())
	var err error
	logfile, err = ioutil.TempFile("", pid+".txt")
	if err != nil {
		log.Fatal(err)
	}
}
