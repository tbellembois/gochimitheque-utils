package convert_test

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/convert"
)

func TestToEmpiricalFormula(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	var (
		csvFile *os.File
		err     error
	)

	if csvFile, err = os.Open("../test_empiricalformula.csv"); err != nil {
		t.Errorf("can not open csv file: %v", err)
	}

	defer csvFile.Close()

	var (
		records [][]string
	)

	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = '\t'

	if records, err = csvReader.ReadAll(); err != nil {
		t.Errorf("unable to parse csv file: %v", err)
	}

	for _, record := range records {

		fmt.Println(record)

		empiricalFormula := record[0]
		linearFormula := record[1]

		logrus.Printf("linear formula: %s", linearFormula)

		converted, err := convert.ToEmpiricalFormula(linearFormula)

		if err != nil {
			t.Errorf(err.Error())
		}

		if converted != empiricalFormula {
			t.Errorf("%s not converted, expected %s, got %s", linearFormula, converted, empiricalFormula)
		}

	}

}
