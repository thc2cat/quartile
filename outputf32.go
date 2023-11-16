package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/constraints"
)

func quartileDeviantPrintf32(Q [3]float32, data []entries, minf float64) int {
	var hasDeviant int
	var limitsup, limitinf float32

	if Q[0] == -1 && Q[1] == 0 && Q[2] == 0 {
		log.Print("bad quartile calc ")
		os.Exit(-1)
	}

	if printquartile {
		printQuartile(Q)
	}

	switch {
	case useMedianlimit:
		limitsup = Q[0] + 3*Q[0]
		limitinf = Q[0] - 3*Q[0]
	case useBoxplot: // Méthode des moustaches
		limitsup = Q[0] + Q[2] - Q[1]
		limitinf = Q[0] - (Q[2] - Q[1])
	default:
		// Quartile : borne sup + ecart interquartile
		limitsup = Q[2] + (Q[2]-Q[1])*(float32)(devianceFactor)
		limitinf = Q[1] - (Q[2]-Q[1])*(float32)(devianceFactor)
	}

	if printlimits {
		fmt.Printf("== Limits are [ %s , %s ] ==\n", betterFormat(limitinf), betterFormat(limitsup))
	}

	for k := range data {
		if data[k].value < limitinf {
			if !quietmode {
				fmt.Printf("< %s %s\n", betterFormat(data[k].value), data[k].text)
			}
			hasDeviant++
		} else if data[k].value > limitsup {
			if !quietmode {
				fmt.Printf("> %s %s\n", betterFormat(data[k].value), data[k].text)
			}
			hasDeviant++
		} else {
			if printall && !quietmode {
				min := (float32)(minf)
				if min == 0 || data[k].value >= min {
					fmt.Printf("  %s %s\n", betterFormat(data[k].value), data[k].text)
				}
			}
		}
	}
	return hasDeviant
}

func betterFormat(num float32) string {
	s := fmt.Sprintf("%.4f", num)
	return strings.TrimRight(strings.TrimRight(s, "0"), ".")
}

// Using generic functions
func sortSlice[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func printQuartile(R [3]float32) {
	if !printquartile {
		return
	}
	fmt.Printf("== Q1=%v", R[1])
	fmt.Printf(" Median=%s", betterFormat(R[0]))
	fmt.Printf(" Q3=%v ==\n", R[2])
}

func ZScorePrintF32bis(data []entries) (flag int) {
	//
	// ZScore => Calculer la moyenne, l'ecart type
	// si l'ecart type > 3x l'ecart moyen alors suspect
	//
	var sign string

	moyenne, ecartmoyen := ZScoreCalF32(data, minimalValue)
	if printlimits {
		fmt.Printf("Zscore : Moyenne μ=%.2f Ecart Type moyen σ=%.2f (~%.0f%%)\n",
			moyenne, ecartmoyen, ecartmoyen*100/moyenne)
	}
	for k, v := range data {
		if v.value >= float32(minimalValue) {
			sign = ">"
			if myAbs(moyenne-v.value) > ((float32)(ZdevianceFactor) * ecartmoyen) {

				if v.value < moyenne {
					sign = "<"
				}
				if !quietmode {
					fmt.Printf("%s %s %s\n", sign, betterFormat(data[k].value), data[k].text)
				}
				flag++
			} else {
				if !quietmode && printall {
					sign = " "
					fmt.Printf("%s %s %s\n", sign, betterFormat(data[k].value), data[k].text)
				}
			}
		}
	}

	return flag
}

func ZScoreMADPrintF32bis(data []entries) (flag int) {
	//
	// ZScore => Calculer la moyenne, l'ecart type
	// si l'ecart type > 3x l'ecart moyen alors suspect
	//
	var sign string

	mediane, ecartmoyen := ZScoreMADCalF32(data, minimalValue)
	if mediane == 0 && ecartmoyen == 0 {
		fmt.Printf("Error: No values\n")
		return -1
	}
	if printlimits {
		fmt.Printf("Zscore MAD : %.2f Ecart Type moyen σ=%.2f (~%.0f%%)\n",
			mediane, ecartmoyen, ecartmoyen*100/mediane)
	}
	for k, v := range data {
		if v.value >= float32(minimalValue) {
			sign = ">"
			if myAbs(mediane-v.value) > ((float32)(ZdevianceFactor) * ecartmoyen) {
				if v.value < mediane {
					sign = "<"
				}
				if !quietmode {
					fmt.Printf("%s %s %s\n", sign, betterFormat(data[k].value), data[k].text)
				}
				flag++
			} else {
				if !quietmode && printall {
					sign = " "
					fmt.Printf("%s %s %s\n", sign, betterFormat(data[k].value), data[k].text)
				}
			}
		}
	}

	return flag
}

// func ZScorePrint(data []entries, MyFunc func([]entries, float64) (float32 float32)) (flag int) {
// 	//
// 	// ZScore => Calculer la moyenne, l'ecart type
// 	// si l'ecart type > 3x l'ecart moyen alors suspect
// 	//
// 	var sign string

// 	moyenne, ecartmoyen := MyFunc(data, minimalValue)
// 	if printlimits {
// 		fmt.Printf("Zscore : Moyenne μ=%.2f Ecart Type moyen σ=%.2f (~%.0f%%)\n",
// 			moyenne, ecartmoyen, ecartmoyen*100/moyenne)
// 	}
// 	for k, v := range data {
// 		if v.value >= float32(minimalValue) {
// 			sign = ">"
// 			if myAbs(moyenne-v.value) > ((float32)(ZdevianceFactor) * ecartmoyen) {

// 				if v.value < moyenne {
// 					sign = "<"
// 				}
// 				if !quietmode {
// 					fmt.Printf("%s %s %s\n", sign, betterFormat(data[k].value), data[k].text)
// 				}
// 				flag++
// 			} else {
// 				if !quietmode && printall {
// 					sign = " "
// 					fmt.Printf("%s %s %s\n", sign, betterFormat(data[k].value), data[k].text)
// 				}
// 			}
// 		}
// 	}

// 	return flag
// }
