package convert

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tbellembois/gochimitheque-utils/count"
	"github.com/tbellembois/gochimitheque-utils/global"
)

// LinearToEmpiricalFormula returns the empirical formula from the linear formula f.
// example: [(CH3)2SiH]2NH
//          (CH3)2C[C6H2(Br)2OH]2
func LinearToEmpiricalFormula(f string) string {
	var ef string

	s := "-"
	nf := ""

	// Finding the first (XYZ)n match
	reg := global.OneGroupMolRe

	for s != "" {
		s = reg.FindString(f)

		// Counting the atoms and rebuilding the molecule string
		m := count.AtomCount(s)
		ms := "" // molecule string
		for k, v := range m {
			ms += k
			if v != 1 {
				ms += fmt.Sprintf("%d", v)
			}
		}

		// Then replacing the match with the molecule string - nf is for "new f"
		nf = strings.Replace(f, s, ms, 1)
		f = nf
	}

	// Counting the atoms
	bAc := count.BaseAtomCount(nf)

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
