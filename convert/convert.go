package convert

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/global"
)

type AtomCount struct {
	Atom         string
	BlockIDStack []int
	Count        int
}

func ToEmpiricalFormula(formula string) (empiricalFormula string, err error) {

	var suffix string

	dotSplittedFormula := strings.Split(formula, ".")

	if len(dotSplittedFormula) > 1 {
		formula = dotSplittedFormula[0]

		for i := range dotSplittedFormula {
			if i == 0 {
				continue
			}

			suffix += dotSplittedFormula[i]
		}

		suffix = "." + suffix
	}

	f := []rune(formula)

	var (
		// Current parenthesis block while parsing the formula.
		currentBlockID int
		// Block ID stack while parsing the formula.
		blockIDStack []int
		// Current index while parsing the formula.
		currentIndex int
		// Current char while parsing the formula.
		currentChar rune
		// Char after currentChar.
		currentCharPlusOne rune

		uppercaseCharRe *regexp.Regexp
		lowercaseCharRe *regexp.Regexp

		atomCountList []*AtomCount
	)

	uppercaseCharRe = regexp.MustCompile("[A-Z]")
	lowercaseCharRe = regexp.MustCompile("[a-z]")

	for currentIndex < len(f) {
		currentChar = f[currentIndex]

		// fmt.Printf("%d:%s\n", currentIndex, string(currentChar))

		if currentIndex < len(f)-1 {
			currentCharPlusOne = f[currentIndex+1]
		} else {
			currentCharPlusOne = ' '
		}

		switch string(currentChar) {
		case "(", "[":
			blockIDStack = append(blockIDStack, currentBlockID)
			currentBlockID++
		case ")", "]":
			if unicode.IsDigit(currentCharPlusOne) {
				var (
					multiplier int
				)

				if multiplier, err = strconv.Atoi(string(currentCharPlusOne)); err != nil {
					logrus.Errorln(err)
					return
				}

				for _, a := range atomCountList {
					for _, bid := range a.BlockIDStack {
						if bid == blockIDStack[len(blockIDStack)-1] {
							a.Count *= multiplier
						}
					}
				}
			}

			if len(blockIDStack) == 0 {
				err = errors.New("invalid parenthesis")
				return
			}

			blockIDStack = blockIDStack[:len(blockIDStack)-1]

		default:
			var (
				atom string
			)

			multiplier := 1

			// Is the current char the beginning of an atom?
			if uppercaseCharRe.MatchString(string(currentChar)) {
				atom = string(currentChar)

				// Is the following char still part of the current atom.
				if lowercaseCharRe.MatchString(string(currentCharPlusOne)) {
					// Two chars atom.
					atom += string(currentCharPlusOne)

					// Finding possible multiplier.
					var (
						multiplierString  string
						hasMultiplier     bool
						hasIncreasedIndex bool
					)

					if currentIndex < len(f)-1 {
						currentIndex++
						hasIncreasedIndex = true
					}

					for unicode.IsDigit(f[currentIndex]) {
						hasMultiplier = true
						multiplierString += string(f[currentIndex])

						// fmt.Println("multiplierString:" + multiplierString)
						if currentIndex < len(f)-1 {
							currentIndex++
						} else {
							break
						}
					}

					if hasIncreasedIndex {
						currentIndex--
					}

					if hasMultiplier {
						if multiplier, err = strconv.Atoi(multiplierString); err != nil {
							logrus.Errorln(err)
							return
						}
					}

				} else {
					// One char atom.
					// Finding possible multiplier.

					var (
						multiplierString  string
						hasMultiplier     bool
						hasIncreasedIndex bool
					)

					if currentIndex < len(f)-1 {
						currentIndex++
						hasIncreasedIndex = true
					}
					for unicode.IsDigit(f[currentIndex]) {
						hasMultiplier = true
						multiplierString += string(f[currentIndex])

						// fmt.Println("multiplierString:" + multiplierString)
						if currentIndex < len(f)-1 {
							currentIndex++
						} else {
							break
						}
					}
					if hasIncreasedIndex {
						currentIndex--
					}

					if hasMultiplier {
						if multiplier, err = strconv.Atoi(multiplierString); err != nil {
							logrus.Errorln(err)
							return
						}
					}
				}
			} else {
				// Any other char.
				currentIndex++

				continue
			}

			// Validating the atom.
			var found bool

			for _, a := range global.SortedByLengthAtoms {
				if a == atom {
					found = true
					break
				}
			}

			if !found {
				err = errors.New("invalid atom")
				return
			}

			blockIDStackCopy := make([]int, len(blockIDStack))

			copy(blockIDStackCopy, blockIDStack)

			atomCountList = append(atomCountList, &AtomCount{
				Atom:         atom,
				BlockIDStack: blockIDStackCopy,
				Count:        multiplier,
			})
		}

		currentIndex++
	}

	atomCount := map[string]int{}

	// Counting atoms.
	for _, a := range atomCountList {
		atom := a.Atom
		count := a.Count

		if currentCount, ok := atomCount[atom]; ok {
			atomCount[atom] = currentCount + count
		} else {
			atomCount[atom] = count
		}
	}

	// Building empirical formula.
	// C, H and then in alphabetical order.
	if CCount, ok := atomCount["C"]; ok {
		count := ""
		if CCount != 1 {
			count = strconv.Itoa(CCount)
		}

		empiricalFormula = fmt.Sprintf("%sC%s", empiricalFormula, count)

		delete(atomCount, "C")
	}

	if HCount, ok := atomCount["H"]; ok {
		count := ""
		if HCount != 1 {
			count = strconv.Itoa(HCount)
		}

		empiricalFormula = fmt.Sprintf("%sH%s", empiricalFormula, count)

		delete(atomCount, "H")
	}

	atomCountKeys := make([]string, 0, len(atomCount))
	for k := range atomCount {
		atomCountKeys = append(atomCountKeys, k)
	}

	sort.Strings(atomCountKeys)

	for _, k := range atomCountKeys {
		count := ""
		if atomCount[k] != 1 {
			count = strconv.Itoa(atomCount[k])
		}

		empiricalFormula = fmt.Sprintf("%s%s%s", empiricalFormula, k, count)
	}

	empiricalFormula += suffix

	logrus.WithFields(logrus.Fields{"empiricalFormula": empiricalFormula}).Debug("ToEmpiricalFormula")

	return
}
