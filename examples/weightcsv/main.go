// weightcsv parses a collection of exported Fitbit weight data files to a
// single ordered CSV log.
//
// These files seem to come in roughly monthly increments, but not exactly.
// Shrug.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"

	"github.com/mroth/fitdump"
)

var (
	flagVerbose   = flag.Bool("v", false, "verbose output")
	flagSkipNoFat = flag.Bool("fatonly", false, "skip entries missing body fat %")
)

func usage() {
	fmt.Fprintf(
		flag.CommandLine.Output(),
		"Usage: %s [FLAGS] <files...>\n\nOptions:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	var combined fitdump.WeightLog
	for _, filepath := range flag.Args() {
		wl, err := parseFile(filepath)
		if err != nil {
			log.Fatal(err)
		}
		if *flagVerbose {
			log.Printf("processed %s: %d entries\n", filepath, len(*wl))
		}
		combined = append(combined, *wl...)
	}

	sort.Slice(combined, func(a, b int) bool {
		return combined[a].RecordedAt.Before(combined[b].RecordedAt)
	})

	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()

	f.WriteString("Timestamp,Weight,BMI,BodyFat\n")
	for _, r := range combined {
		if !(*flagSkipNoFat && r.Fat == 0.0) {
			f.WriteString(fmt.Sprintf(
				"%s,%.1f,%.2f,%.3f\n",
				r.RecordedAt.Format(time.RFC3339), r.Weight, r.BMI, r.Fat,
			))
		}
	}
}

func parseFile(path string) (*fitdump.WeightLog, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %w", path, err)
	}

	var wl fitdump.WeightLog
	err = json.Unmarshal(dat, &wl)
	if err != nil {
		return nil, fmt.Errorf("error decoding %s: %w", path, err)
	}

	return &wl, nil
}
