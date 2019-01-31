package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		input  = flag.String("input", "input.txt", "File(s) to read.")
		output = flag.String("output", "output.txt", "Output file (required).")
	)
	flag.Parse()
	file, errOpen := os.Open(*input)
	if errOpen != nil {
		log.Fatal(errOpen)
	}
	f, errCreate := os.Create(*output)
	if errCreate != nil {
		log.Fatal(errCreate)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	headerReading := true
	var headerList []string
	var buffer bytes.Buffer
	for {
		line, prefix, err := reader.ReadLine()
		if err != nil {
			break
		}

		if prefix {
			buffer.Write(line)
			continue
		} else {
			buffer.Write(line)
		}
		readText := string(buffer.String())
		buffer.Reset()
		if headerReading {
			headerList = strings.Split(readText, ",")[1:]
			headerReading = false
		} else {
			splitColumn := strings.Split(readText, ",")
			var columnElement string
			for index, element := range splitColumn {
				if index == 0 {
					columnElement = element
				} else {
					f.WriteString(columnElement + "," + headerList[index-1] + "," + element + "\r\n")
				}

			}
		}
	}
	fmt.Println("Finish time is", time.Now())
}
