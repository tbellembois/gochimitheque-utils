package validator_test

import (
	"encoding/csv"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/tbellembois/gochimitheque-utils/validator"
)

func TestIsCasNumber(t *testing.T) {

	logrus.SetLevel(logrus.DebugLevel)

	var (
		csvFile *os.File
		err     error
	)

	if csvFile, err = os.Open("../test_casnumber.csv"); err != nil {
		t.Errorf("can not open csv file: %v", err)
	}

	defer csvFile.Close()

	var (
		records [][]string
	)

	csvReader := csv.NewReader(csvFile)
	csvReader.Comma = ' '

	if records, err = csvReader.ReadAll(); err != nil {
		t.Errorf("unable to parse csv file: %v", err)
	}

	for _, record := range records {

		cas := record[1]

		logrus.Printf("cas number: %s", cas)

		if !validator.IsCasNumber(cas) {
			t.Errorf("%s is not a valid cas number", cas)
		}

	}
}
