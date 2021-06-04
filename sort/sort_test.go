package sort_test

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/sort"
)

func TestSortEmpiricalFormula(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	var (
		sortedf string
		err     error
	)
	f := "Al2(SO4)3"
	if sortedf, err = sort.SortEmpiricalFormula(f); err != nil {
		t.Errorf("%s is not a valid formula: %v", f, err)
	}
	if sortedf != "Al2(SO4)3" {
		t.Errorf("%s was not sorted - output: %s", f, sortedf)
	}

}
