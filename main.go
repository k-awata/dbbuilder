package main

import (
	_ "embed"
	"errors"
	"flag"
	"fmt"
	"os"
)

const (
	CliName = "dbbuilder"
	Version = "0.1.0"
)

//go:embed sample.json
var sampleJson string

func main() {
	var err error

	// Define flags and usage
	flag.Usage = func() {
		fmt.Fprint(os.Stderr,
			CliName, " outputs a macro for Aveva Administration to set up project database using options in JSON format.\n",
			"\n",
			"Usage: ", CliName, " [-o file] {-s | json_file}\n",
			"Example: ", CliName, " sample.json > export.mac\n",
			"\n",
			"Options:\n",
		)
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr,
			"\n",
			"When json_file is -, read standard input instead of a file.\n",
		)
	}
	vflg := flag.Bool("v", false, "Display version")
	sflg := flag.Bool("s", false, "Output a sample JSON")
	oflg := flag.String("o", "", "Output file")
	flag.Parse()

	// Show version
	if *vflg {
		fmt.Fprintf(os.Stderr, "%s version %s\n", CliName, Version)
		return
	}

	// Specify to output file or stdout
	out := os.Stdout
	if *oflg != "" {
		out, err = os.Create(*oflg)
		chkErr(err)
	}

	// Output sample JSON
	if *sflg {
		fmt.Fprintln(out, sampleJson)
		return
	}

	// Check argument count
	if flag.NArg() == 0 {
		flag.Usage()
		return
	}
	if flag.NArg() >= 2 {
		chkErr(errors.New("too many arguments"))
	}

	// Open input file or use stdin
	var in *os.File
	if flag.Arg(0) == "-" {
		in = os.Stdin
	} else {
		var err error
		in, err = os.Open(flag.Arg(0))
		chkErr(err)
		defer in.Close()
	}

	pj, err := LoadProject(in)
	chkErr(err)
	mac, err := MakeMac(pj)
	chkErr(err)

	fmt.Fprint(out, mac)
}

func chkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}
