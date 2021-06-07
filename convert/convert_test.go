package convert_test

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/convert"
)

func TestLinearToEmpiricalFormula(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	var (
		convertedf string
	)
	f := "(C6H5O)2P(O)N3"
	convertedf = convert.LinearToEmpiricalFormula(f)
	if convertedf != "C12H10N3O3P" {
		t.Errorf("%s was not converted - output: %s", f, convertedf)
	}

}