package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/**
 * struc entries is used for spliting "value text" strings
 *
 */
type entries struct {
	value float32
	text  string
}

/**
 * Read stdin searching for "value text", return a slice of
 * entries struct containing value and text
 *
 */
func readAll() []entries {
	R := make([]entries, 0)
	re := regexp.MustCompile(`^[\s\t]*(?P<value>[\d]*)[\s\t]{1}(?P<text>.*)$`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 1 {
			founds := getParams(re, strings.TrimSpace(text))
			if len(founds) > 0 {
				var tmp entries
				x, err := strconv.ParseFloat(founds["value"], 32)
				if err != nil {
					continue
				}
				tmp.value = (float32)(x)
				tmp.text = strings.Clone(strings.TrimSuffix(founds["text"], "\n"))
				R = append(R, tmp)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return R
}

/**
 * getNumbers return slice of float32 from []entries value
 * where []entries > min
 */
func getNumbers(In []entries, min float64) []float32 {
	minf := (float32)(min)
	R := make([]float32, 0)
	for k := range In {
		if minf == 0 || (In[k].value >= minf) {
			R = append(R, In[k].value)
		}
	}
	return R
}

/**
 * Parses text with the given regular expression and returns the
 * group values defined in the expression.
 *
 */
func getParams(compRegEx *regexp.Regexp, text string) (paramsMap map[string]string) {

	match := compRegEx.FindStringSubmatch(text)
	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
