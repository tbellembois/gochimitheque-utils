package count_test

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/count"
)

func TestBaseAtomCount(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	var (
		countedf map[string]int
	)
	f := "H10O3C12"
	countedf = count.BaseAtomCount(f)
	if countedf["C"] != 12 || countedf["H"] != 10 || countedf["O"] != 3 || len(countedf) != 3 {
		t.Errorf("%s count wrong - output: %+v len:%d", f, countedf, len(countedf))
	}

}
