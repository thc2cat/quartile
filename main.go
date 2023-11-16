package main

// purpose:
// init flags
// read ints from stdin
// sort ints[]
// calc quartile
// print deviant values with adaptable Tukey value
// or 3xMedian if -M is choosed

// v0.1 print green/red values
// v0.2 - dont mess original data, return 1 if found quartile deviance
// v0.3 - trim space and tabs
// v0.4 - float32 and direct input of "sort |uniq -c|sort -rn|head -nn"
// V0.5 - Adding Z-score mod méthod.
// V0.6 - Bug in getNumbers

import (
	"flag"
	"fmt"
	"os"
)

var (
	quietmode, printall, printquartile, printlimits     bool
	useMedianlimit, useBoxplot, useZScore, useZScoreMAD bool
	// printlow, printupper bool
	devianceFactor, ZdevianceFactor, minimalValue float64

	// Version is git tag
	Version string
)

func main() {
	initflags()
	data := readAll()

	switch {
	case useZScore:
		os.Exit(ZScorePrintF32bis(data))
	case useZScoreMAD:
		os.Exit(ZScoreMADPrintF32bis(data))

	default:
		N := getNumbers(data, minimalValue)

		Q := quartileCalcf32(N)
		if Q[0] == 0 && Q[1] == 0 {
			fmt.Printf("Bad quartile (Exiting) : \n")
			printQuartile(Q)
			os.Exit(-1)
		}

		os.Exit(quartileDeviantPrintf32(Q, data, minimalValue))
	}
}

func initflags() {
	flag.BoolVar(&quietmode, "q", false, "quiet mode")
	flag.BoolVar(&printall, "a", false, "print all values")
	flag.BoolVar(&printquartile, "Q", false, "print quartiles values")
	flag.BoolVar(&printlimits, "l", false, "print limits")

	flag.Float64Var(&devianceFactor, "f", 1.5, "Tukey deviance factor")
	flag.Float64Var(&ZdevianceFactor, "z", 3.0, "Z Score deviance factor")
	flag.Float64Var(&minimalValue, "m", 0, "minimal value")

	flag.BoolVar(&useMedianlimit, "M", false, "use Median Limit instead (3x)")
	flag.BoolVar(&useBoxplot, "B", false, "use Boxplot [M-(Q3-Q1),M+(Q3-Q1)]")
	flag.BoolVar(&useZScore, "Z", false, "use Z-Score")
	flag.BoolVar(&useZScoreMAD, "D", false, "use Z-Score mod")

	flag.StringVar(&Version, "v", Version, "show current Version")

	flag.Parse()
}
