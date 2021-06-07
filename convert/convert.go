package convert

import (
	"fmt"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/count"
	"github.com/tbellembois/gochimitheque-utils/global"
)

// LinearToEmpiricalFormula returns the empirical formula from the linear formula f.
// example: [(CH3)2SiH]2NH
//          (CH3)2C[C6H2(Br)2OH]2
func LinearToEmpiricalFormula(f string) string {

	logrus.WithFields(logrus.Fields{"f": f}).Debug("LinearToEmpiricalFormula")

	var ef string

	// s := "-"
	nf := f

	// Finding the first (XYZ)n match
	reg := global.OneGroupMolRe

	atomCount := make(map[string]int)
	for _, s := range reg.FindAllString(f, -1) {

		logrus.WithFields(logrus.Fields{"s": s}).Debug("LinearToEmpiricalFormula")

		// Removing the match from the original formula.
		nf = strings.Replace(nf, s, "", -1)

		// Counting the atoms and rebuilding the molecule string
		m := count.AtomCount(s)
		logrus.WithFields(logrus.Fields{"m": m}).Debug("LinearToEmpiricalFormula")

		for k, v := range m {
			if _, ok := atomCount[k]; ok {
				atomCount[k] = atomCount[k] + v
			} else {
				atomCount[k] = v
			}
		}

	}

	ms := "" // molecule string
	for k, v := range atomCount {
		ms += k
		if v != 1 {
			ms += fmt.Sprintf("%d", v)
		}
	}
	logrus.WithFields(logrus.Fields{"ms": ms}).Debug("LinearToEmpiricalFormula")

	// Then replacing the match with the molecule string - nf is for "new f"
	nf = ms + nf

	// Counting the atoms
	bAc := count.BaseAtomCount(nf)
	logrus.WithFields(logrus.Fields{"bAc": bAc}).Debug("LinearToEmpiricalFormula")

	// Sorting the atoms
	// C, H and then in alphabetical order
	var ats []string // atoms
	hasC := false    // C atom present
	hasH := false    // H atom present

	for k := range bAc {
		switch k {
		case "C":
			hasC = true
		case "H":
			hasH = true
		default:
			ats = append(ats, k)
		}
	}
	sort.Strings(ats)

	if hasH {
		ats = append([]string{"H"}, ats...)
	}
	if hasC {
		ats = append([]string{"C"}, ats...)
	}

	for _, at := range ats {
		ef += at
		nb := bAc[at]
		if nb != 1 {
			ef += fmt.Sprintf("%d", nb)
		}
	}

	return ef
}
