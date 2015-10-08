package acorn

import (
	"io"
	"log"
	"os"
	"os/exec"
)

// Output sets up a file output or stdout for generating code and formatting it
func Output(filename *string, formatter *string, callback func(io.Writer)) {
	if filename == nil {
		callback(os.Stdout)
		return
	}

	outputFile, err := os.OpenFile(*filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if formatter != nil {
		defer exec.Command(*formatter, "-w", *filename).Run()
	}
	defer outputFile.Close()
	callback(outputFile)
}
