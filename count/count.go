package count

import (
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/global"
)

// AtomCount returns a count of the atoms of the f formula as a map.
// f must be a formula like (XYZ) (XYZ)n or [XYZ] [XYZ]n.
// example:
// (CH3)2 will return "C":2, "H":6
// CH3CH(NO2)CH3 will return "N":1 "O":2
// CH3CH(NO2)(CH3)2 will return "N":1 "O":2 - process only the first match
func AtomCount(f string) map[string]int {

	logrus.WithFields(logrus.Fields{"f": f}).Debug("AtomCount")

	var (
		// the result map
		c = make(map[string]int)
		r = global.OneGroupMolRe
	)
	// Looking for non matching molecules.
	if !r.MatchString(f) {
		return nil
	}

	// sl is a list of 3 elements like
	// [[(CH3Na6CCl5H)2 CH3Na6CCl5H 2]]
	sl := r.FindAllStringSubmatch(f, -1)
	basicMol := sl[0][1]
	multiplier, _ := strconv.Atoi(sl[0][2])

	// if there is no multiplier
	if multiplier == 0 {
		multiplier = 1
	}

	// counting the atoms
	aCount := BaseAtomCount(basicMol)
	for at, nb := range aCount {
		c[at] = nb * multiplier
	}

	return c

}

// BaseAtomCount returns a count of the atoms of the f formula as a map.
// f must be a basic formula with only atoms and numbers.
// example:
// C6H5COC6H4CO2H will return "C1":4, "H":10, "O":3
// CH3CH(NO2)CH3 will return Nil, parenthesis are not allowed
func BaseAtomCount(f string) map[string]int {

	logrus.WithFields(logrus.Fields{"f": f}).Debug("BaseAtomCount")

	var (
		// the result map
		c   = make(map[string]int)
		r   = global.BasicMolRe
		err error
	)
	// Looking for non matching molecules.
	if !r.MatchString(f) {
		return nil
	}

	// sl is a slice like [[Na Na ] [Cl Cl ] [C2 C 2] [Cl3 Cl 3]]
	// for f = NaClC2Cl3
	// [ matchingString capture1 capture2 ]
	// capture1 is the atom
	// capture2 is the its number
	sl := r.FindAllStringSubmatch(f, -1)
	logrus.WithFields(logrus.Fields{"sl": sl}).Debug("BaseAtomCount")

	for _, i := range sl {
		atom := i[1]
		var nbAtom int
		if i[2] != "" {
			nbAtom, err = strconv.Atoi(i[2])
			if err != nil {
				return nil
			}
		} else {
			nbAtom = 1
		}
		if _, ok := c[atom]; ok {
			c[atom] = c[atom] + nbAtom
		} else {
			c[atom] = nbAtom
		}
	}
	return c

}
