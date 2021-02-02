package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lkingland/gridd"
	"github.com/lkingland/gridd/boson"
)

var usage = `gridd

Plug the current node into the Grid.
`
var (
	Version = flag.Bool("version", false, "Print version [$GRIDD_VERSION]")
	Verbose = flag.Bool("verbose", false, "Print verbose logs [$GRIDD_VERBOSE]")

	date string
	vers string
	hash string
)

func parseEnv() {
	parseBool("GRIDD_VERSION", Version)
	parseBool("GRIDD_VERBOSE", Verbose)
}

func printCfg() {
	fmt.Printf("GRIDD_VERSION=%v\n", *Version)
	fmt.Printf("GRIDD_VERBOSE=%v\n", *Verbose)
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
	}
	parseEnv()
	flag.Parse()

	if *Verbose {
		printCfg()
	}

	if *Version || (len(os.Args) > 1 && os.Args[1] == "version") {
		fmt.Println(version())
		return
	}

	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() (err error) {
	fmt.Println("start", version())

	provider := boson.NewProvider(*Verbose)
	_ = gridd.New(provider, gridd.WithVerbose(*Verbose))

	return nil
}

func parseBool(key string, value *bool) {
	if val, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(val)
		if err != nil {
			panic(err)
		}
		*value = b
	}
}

func parseString(key string, value *string) {
	if val, ok := os.LookupEnv(key); ok {
		*value = val
	}
}

func version() string {
	// If 'vers' is not a semver already, then the binary was built either
	// from an untagged git commit or was built directly from source
	// (set semver to v0.0.0)

	var elements = []string{}
	if strings.HasPrefix(vers, "v") {
		elements = append(elements, vers) // built via make with a tagged commit
	} else {
		elements = append(elements, "v0.0.0") // from source or untagged commit
	}

	if date != "" {
		elements = append(elements, date)
	}

	if hash != "" {
		elements = append(elements, hash)
	}

	return strings.Join(elements, "-")

}
