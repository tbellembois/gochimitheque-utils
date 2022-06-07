package sort_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/sort"
	"github.com/tbellembois/gochimitheque-utils/test"
)

func TestSortEmpiricalFormula(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	var (
		jsonFile *os.File
		err      error
	)

	if jsonFile, err = os.Open("../testdata.json"); err != nil {
		t.Errorf("can not open json file: %v", err)
	}

	defer jsonFile.Close()

	var (
		byteValue []byte
	)

	if byteValue, err = ioutil.ReadAll(jsonFile); err != nil {
		t.Errorf("can not read json file: %v", err)
	}

	var (
		PCCompounds test.PCCompounds
	)

	if json.Unmarshal(byteValue, &PCCompounds); err != nil {
		t.Errorf("can not unmarshal json: %v", err)
	}

	logrus.Println(len(PCCompounds.PC_Compounds))

	var (
		sortedf string
	)

	for _, compound := range PCCompounds.PC_Compounds {
		for _, prop := range compound.Props {

			if prop.Urn.Label == "Molecular Formula" {

				f := prop.Value.Sval

				logrus.Printf("formula: %s", f)

				if sortedf, err = sort.SortEmpiricalFormula(f); err != nil {
					t.Errorf("%s is not a valid formula: %v", f, err)
				}

				if sortedf != f {
					t.Errorf("%s was not sorted - output: %s", f, sortedf)
				}

			}

		}
	}

}
