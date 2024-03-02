package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"lab2/coder"
)

func main() {
	var (
		optDecode  bool
		optInPath  string
		optOutPath string
		optKeyPath string
	)

	flag.BoolVar(&optDecode, "d", false, "decode")
	flag.StringVar(&optInPath, "i", "", "input filepath")
	flag.StringVar(&optOutPath, "o", "", "output filepath")
	flag.StringVar(&optKeyPath, "k", "", "key filepath")
	flag.Parse()

	var (
		key []byte
		err error
	)

	if optKeyPath != "" {
		if key, err = os.ReadFile(optKeyPath); err != nil {
			exit(err)
		}
	} else {
		key = []byte(strings.Join(flag.Args(), " "))
	}

	var (
		encoder   = coder.New(key)
		operation = encoder.Encode

		input  = os.Stdin
		output = os.Stdout
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
