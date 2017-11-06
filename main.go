// Tiny log merging tool. Suitable for merging multiple log files of different size. Reads streams, so should not use a lot of RAM.
// Logs are merged based on parsed timestamp, so it must be the same in all input files.
// You should provide '-ts' timestamp format that corresponds the one in your logs. Default is: 'Jan 2 15:04:05'.
// Please refer to time package documentation for more info. https://golang.org/pkg/time/#pkg-constants
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	start := time.Now()

	var (
		output     string
		inputDir   string
		timeFormat string
	)

	flag.StringVar(&output, "o", "output.log", "Path where result must be written")
	flag.StringVar(&inputDir, "i", "input", "Path to the folder with logs to merge")
	flag.StringVar(&timeFormat, "tf", "Jan  2 15:04:05", "Time format of log entries")

	flag.Parse()

	outFile, err := os.Create(output)
	if err != nil {
		log.Fatal("Failed create merged file:", err)
	}
	defer outFile.Close()

	if err := merge(inputDir, bufio.NewWriter(outFile), timeFormat); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Job done in:", time.Since(start))
}

func merge(inputDir string, out *bufio.Writer, timeFormat string) error {
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("failed to read input dir content: %s", err)
	}

	nFiles := len(files)

	var (
		buf     = make([]string, nFiles)
		readers = make([]*bufio.Scanner, nFiles)
	)

	var file *os.File

	for i, fi := range files {
		filePath := filepath.Join(inputDir, fi.Name())
		file, err = os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed opening file: %q, %s", filePath, err)
		}
		defer file.Close()
		readers[i] = bufio.NewScanner(file)
	}

	var readersDone int

	// Load first events from all files.
	for i, reader := range readers {
		if ok := reader.Scan(); !ok {
			if reader.Err() != nil {
				return fmt.Errorf("failed reading file: %s", err)
			}
			readersDone++
			continue
		}
		buf[i] = reader.Text()
	}

	sorter := newEventSorter(nFiles, timeFormat)

	for {
		if len(readers) == readersDone {
			break
		}

		i, err := sorter.firstEventIndex(buf)
		if err != nil {
			return err
		}

		if _, err := out.WriteString(buf[i] + "\n"); err != nil {
			return fmt.Errorf("failed writing sting to bufio: %s", err)
		}

		buf[i] = ""

		if err := out.Flush(); err != nil {
			return fmt.Errorf("failed to write data to output file: %s", err)
		}

		if ok := readers[i].Scan(); !ok {
			if readers[i].Err() != nil {
				return fmt.Errorf("failed reading file: %s", err)
			}

			readersDone++
			continue
		}

		buf[i] = readers[i].Text()
	}

	return nil
}
