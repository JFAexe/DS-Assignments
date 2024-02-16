package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"lab1/coder"
)

func main() {
	var (
		optDecode  bool
		optInPath  string
		optOutPath string
	)

	flag.BoolVar(&optDecode, "d", false, "decode")
	flag.StringVar(&optInPath, "i", "", "input filepath")
	flag.StringVar(&optOutPath, "o", "", "output filepath")
	flag.Parse()

	key := strings.Join(flag.Args(), " ")

	if len(key) < 1 {
		exit("no key set")
	}

	var (
		encoder   = coder.New(key)
		operation = encoder.Encode

		input  = os.Stdin
		output = os.Stdout

		err error
	)

	if optDecode {
		operation = encoder.Decode
	}

	if optInPath != "" {
		input, err = os.Open(optInPath)
		if err != nil {
			exit(err)
		}
		defer input.Close()
	}

	if optOutPath != "" {
		output, err = os.Create(optOutPath)
		if err != nil {
			exit(err)
		}
		defer output.Close()
	}

	if err = operation(input, output); err != nil {
		exit(err)
	}
}

func exit(msg ...any) {
	fmt.Println(msg...)

	os.Exit(1)
}
