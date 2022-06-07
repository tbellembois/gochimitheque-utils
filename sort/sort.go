package sort

import (
	"errors"
	"regexp"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/global"
)

// SortEmpiricalFormula returns the sorted f empirical formula.
func SortEmpiricalFormula(f string) (string, error) {

	logrus.WithFields(logrus.Fields{"f": f}).Debug("SortEmpiricalFormula")

	var (
		err      error
		newf, sp string
	)

	// zero empirical formula
	if f == "XXXX" {
		return f, nil
	}

	// removing spaces
	f = strings.Replace(f, " ", "", -1)

	// if the formula is like abc.def.ghi, spliting it
	splitf := strings.Split(f, ".")
	if len(splitf) == 1 {
		return SortFormula(f)
	}

	for _, p := range splitf {
		if sp, err = SortFormula(p); err != nil {
			return "", err
		}
		newf += "." + sp
	}

	return newf, nil

}

// SortFormula returns the sorted f formula.
func SortFormula(f string) (string, error) {

	logrus.WithFields(logrus.Fields{"f": f}).Debug("SortFormula")

	var (
		hasCatom, hasHatom, hasOatom, hasULatom bool
		upperLowerAtoms, otherAtoms             []string
	)

	// removing spaces
	f = strings.Replace(f, " ", "", -1)

	// checking formula characters
	if !global.FormulaRe.MatchString(f) {
		return "", errors.New("invalid characters in formula")
	}

	// search atoms with and uppercase followed by lowercase letters like Na or Cl
	// return a list of tuples like:
	// [[Cl Cl Cl] [Na Na Na] [Cl3 Cl3 Cl]]
	// for ClNaHCl3
	// the third member of the tupple is used to detect duplicated atoms
	ULAtomsRe := regexp.MustCompile("((?:^[0-9]+)?([A-Z][a-wy-z]{1,3})[0-9,]*)")
	ula := ULAtomsRe.FindAllStringSubmatch(f, -1)

	// detecting wrong UL atoms
	// counting atoms at the same time and leaving on duplicates
	atomcount := make(map[string]int)
	for _, a := range ula {

		// wrong?
		if _, ok := global.Atoms[a[2]]; !ok {
			return "", errors.New("wrong UL atom in formula: " + a[2])
		}
		upperLowerAtoms = append(upperLowerAtoms, a[0])
		// duplicate?
		if _, ok := atomcount[a[2]]; !ok {
			atomcount[a[2]] = 0
		} else {
			// atom already present !
			return "", errors.New("duplicate UL atom in formula")
		}
		// removing from formula for the next steps
		f = strings.Replace(f, a[0], "", -1)

	}
	if len(upperLowerAtoms) > 0 {
		hasULatom = true
	}

	// here we should have only one uppercase letter (and digits) per atom for the rest of
	// the formula

	// searching the C atom
	CAtomRe := regexp.MustCompile("((?:^[0-9]+)?(C)[0-9,]*)")
	ca := CAtomRe.FindAllStringSubmatch(f, -1)
	// will return [[C2 C2 C]] for ClNaC2
	// leaving on duplicated C atom
	if len(ca) > 1 {
		return "", errors.New("duplicate C atom in formula")
	}
	if len(ca) == 1 {
		hasCatom = true
		// removing from formula for the next steps
		f = strings.Replace(f, ca[0][0], "", -1)
	}

	// searching the H atom
	HAtomRe := regexp.MustCompile("((?:^[0-9]+)?(H)[0-9,]*)")
	ha := HAtomRe.FindAllStringSubmatch(f, -1)
	// will return [[H2 H2 H]] for ClNaH2
	// leaving on duplicated C atom
	if len(ha) > 1 {
		return "", errors.New("duplicate H atom in formula")
	}
	if len(ha) == 1 {
		hasHatom = true
		// removing from formula for the next steps
		f = strings.Replace(f, ha[0][0], "", -1)
	}

	// searching the other atoms
	OAtomRe := regexp.MustCompile("((?:^[0-9]+)?([A-Z])[0-9,]*)")
	oa := OAtomRe.FindAllStringSubmatch(f, -1)

	// detecting wrong atoms
	// counting atoms at the same time and leaving on duplicates
	atomcount = make(map[string]int)
	for _, a := range oa {
		// wrong?
		if _, ok := global.Atoms[a[2]]; !ok {
			return "", errors.New("wrong UL atom in formula: " + a[2])
		}
		otherAtoms = append(otherAtoms, a[0])
		// duplicate?
		if _, ok := atomcount[a[2]]; !ok {
			atomcount[a[2]] = 0
		} else {
			// atom already present !
			return "", errors.New("duplicate other atom in formula")
		}
		// removing from formula for the next steps
		f = strings.Replace(f, a[0], "", -1)
	}
	if len(oa) > 0 {
		hasOatom = true
	}

	logrus.WithFields(logrus.Fields{"f": f}).Debug("SortFormula")
	logrus.WithFields(logrus.Fields{"len(f)": len(f)}).Debug("SortFormula")

	// if formula is not emty, this is an error
	if len(f) != 0 {
		return "", errors.New("wrong lowercase atoms in formula: " + f)
	}

	// rebuilding the formula
	newf := ""
	if hasCatom {
		newf += ca[0][0]
	}
	if hasHatom {
		newf += ha[0][0]
	}
	if hasOatom || hasULatom {
		at := append(otherAtoms, upperLowerAtoms...)
		sort.Strings(at)
		for _, a := range at {
			newf += a
		}
	}

	return newf, nil

}
