package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/htwais/morse/convert"
)

func main() {
	// reduce noise in log messages:
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	// without options, flags default usage looks sad:
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "usage: morse path/to/input.{txt,morse} [path/to/output/file]\n")
	}
	flag.Parse()

	// first handle invalid args:
	if flag.NArg() < 1 || flag.NArg() > 2 {
		flag.Usage()
		os.Exit(1)
	}

	// arg1 is mandatory - path/to/input/file:
	inputPath := flag.Arg(0)
	in, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("FAILED: %v", err)
	}

	// suffix must be .txt or .morse - choose appropriate conversion function:
	conv := convert.Encode
	if strings.HasSuffix(inputPath, "morse") {
		conv = convert.Decode
	} else if !strings.HasSuffix(inputPath, "txt") {
		log.Fatalf("invalid path to input (%s) - must end with either .txt or .morse", inputPath)
	}

	// arg2 is optional - path/to/output/file, use stdout if none:
	var out io.Writer = os.Stdout
	if flag.NArg() == 2 {
		var f *os.File
		if f, err = os.Create(flag.Arg(1)); err != nil {
			log.Fatalf("FAILED: %v", err)
		}
		out = f
	}

	// and now do the conversion:
	if err = conv(out, in); err != nil {
		log.Fatalf("FAILED: %v", err)
	}
}
