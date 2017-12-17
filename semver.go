package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	"github.com/blang/semver"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %v <condition> <version>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "\nOPTIONS:")
		fmt.Fprintln(os.Stderr, "	-q	Do not output anything")
		fmt.Fprintln(os.Stderr, "\nEXAMPLES:")
		fmt.Fprintf(os.Stderr, "	%v \">=1.0.0\" \"1.1.0\"\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "	%v \">1.0.0 <2.0.0 || >3.0.0 !4.2.1\" \"3.1.1\"\n\n", os.Args[0])
	}

	quiet := flag.Bool("q", false, "Do not output anything!")

	flag.Parse()

	if len(flag.Args()) != 2 {
		if !*quiet {
			fmt.Fprintln(os.Stderr, "Not enough parameters.")
			flag.Usage()
		}
		os.Exit(128)
	}

	condition, err := semver.ParseRange(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Condition %v is not a valid condition.", flag.Arg(0))
		fmt.Fprintf(os.Stderr, "Error: %q", err)
		os.Exit(128)
	}

	var re = regexp.MustCompile(`^(\d+\.\d+)(\-.*)?$`)
	version, err := semver.Parse(re.ReplaceAllString(flag.Arg(1), `$1.0$2`))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Version %v is not a valid version.", flag.Arg(1))
		fmt.Fprintf(os.Stderr, "Error: %q", err)
		os.Exit(128)
	}

	if condition(version) {
		if !*quiet {
			fmt.Printf("%v meets the condition.\n", flag.Arg(1))
		}
		os.Exit(0)
	} else {
		if !*quiet {
			fmt.Printf("%v does not meet the condition!\n", flag.Arg(1))
		}
		os.Exit(1)
	}
}
