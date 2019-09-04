package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"

	"github.com/fatih/color"
)

// purpose:
// read ints from stdin
// sort ints[]
// print quartile
// print deviant values

func main() {
	data := readints()
	datas := sortints(data)
	quartileDeviantPrint(datas)
}

// Byvaltype : used for sorting ds values
type Byvaltype []int32

func (a Byvaltype) Len() int           { return len(a) }
func (a Byvaltype) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Byvaltype) Less(i, j int) bool { return a[i] < a[j] }

func readints() []int32 {
	var ds []int32
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		N, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Error converting %s to int", line)
		} else {
			ds = append(ds, (int32)(N))
		}
	}
	return ds
}

func sortints(d []int32) []int32 {
	var s = make([]int32, len(d))
	copy(s, d)
	sort.Sort(Byvaltype(s))
	return s
}

func isDecimal(n float64) (x bool) {
	if math.Floor(n) == n {
		x = true
	}
	return
}
func quartileDeviantPrint(d []int32) {

	green := color.New(color.FgGreen).FprintfFunc()
	red := color.New(color.FgRed).FprintfFunc()
	R := quartileCalc(sortints(d))
	if R[0] == -1 && R[1] == 0 && R[2] == 0 {
		log.Print("bad quartile calc")
	}

	limitsup := R[2] + (R[2]-R[1])*1.5 // borne sup + ecart interquartile
	limitinf := R[1] - (R[2]-R[1])*1.5
	for _, v := range d {
		if (float32)(v) < limitinf {
			green(os.Stdout, " %d", v)
		} else if (float32)(v) > limitsup {
			red(os.Stdout, " %d", v)
		} else {
			fmt.Printf(" %d", v)
		}

	}
	fmt.Printf("\n")
}

// Exemple de calcul
// 10, 25, 30, 40, 41, 42, 50, 55, 70, 101, 110, 111 => 12 Valeurs
//             40      42+50       70
// => M = 42+50/2 => 46
// => Q1 = 12/4=> 4 eme valeur => 40
// => Q2 = 12*3/4 => 9 eme valeur => 70

func quartileCalc(d []int32) [3]float32 {
	var Q [3]float32
	var n float64
	var i int

	N := len(d)
	if N < 4 {
		log.Printf("bad quartileCalc %v", d)
		return [3]float32{-1, 0, 0}
	}

	// Calcul de la mediane Q[0]
	if N%2 == 0 { // Pair
		Q[0] = (float32)((d[N/2-1] + d[N/2])) / 2
	} else {
		Q[0] = (float32)(d[(int)(math.Ceil((float64)(N)/2))-1])
	}

	// Calc Q1
	n = (float64)(N) / 4
	if isDecimal(n) { // entier
		i = (N / 4) + 1
	} else {
		i = (int)(math.Ceil(n))
	}
	Q[1] = (float32)(d[i-1])

	// calc Q3
	n = (float64)(N) * 3 / 4
	if isDecimal(n) { // entier
		i = ((N * 3) / 4)
	} else {
		i = (int)(math.Ceil(((float64)(N) * 3 / 4)))
	}
	Q[2] = (float32)(d[i-1])

	return Q
}
