package validator

import (
	"regexp"
	"strconv"
)

// IsCeNumber returns true if c is a valid ce number
func IsCeNumber(c string) bool {

	var (
		err                error
		checkdigit, checkd int
	)

	if c == "000-000-0" {
		return true
	}

	// compiling regex
	r := regexp.MustCompile("^(?P<groupone>[0-9]{3})-(?P<grouptwo>[0-9]{3})-(?P<groupthree>[0-9]{1})$")
	// finding group names
	n := r.SubexpNames()
	// finding matches
	ms := r.FindAllStringSubmatch(c, -1)
	if len(ms) == 0 {
		return false
	}
	m := ms[0]
	// then building a map of matches
	md := map[string]string{}
	for i, j := range m {
		md[n[i]] = j
	}

	if len(m) > 0 {
		numberpart := md["groupone"] + md["grouptwo"]

		// converting the check digit into int
		if checkdigit, err = strconv.Atoi(string(md["groupthree"])); err != nil {
			return false
		}

		// calculating the check digit
		counter := 1  // loop counter
		currentd := 0 // current processed digit in c

		for i := 0; i < len(numberpart); i++ {
			// converting digit into int
			if currentd, err = strconv.Atoi(string(numberpart[i])); err != nil {
				return false
			}
			checkd += counter * currentd
			counter++
			//fmt.Printf("counter: %d currentd: %d checkd: %d\n", counter, currentd, checkd)
		}
	}

	return checkd%11 == checkdigit
}

// IsCasNumber returns true if c is a valid cas number
func IsCasNumber(c string) bool {

	var (
		err                error
		checkdigit, checkd int
	)

	if c == "0000-00-0" {
		return true
	}

	// compiling regex
	r := regexp.MustCompile("^(?P<groupone>[0-9]{1,7})-(?P<grouptwo>[0-9]{2})-(?P<groupthree>[0-9]{1})$")
	// finding group names
	n := r.SubexpNames()
	// finding matches
	ms := r.FindAllStringSubmatch(c, -1)
	if len(ms) == 0 {
		return false
	}
	m := ms[0]
	// then building a map of matches
	md := map[string]string{}
	for i, j := range m {
		md[n[i]] = j
	}

	if len(m) > 0 {
		numberpart := md["groupone"] + md["grouptwo"]

		// converting the check digit into int
		if checkdigit, err = strconv.Atoi(string(md["groupthree"])); err != nil {
			return false
		}
		//fmt.Printf("checkdigit: %d\n", checkdigit)

		// calculating the check digit
		counter := 1  // loop counter
		currentd := 0 // current processed digit in c

		for i := len(numberpart) - 1; i >= 0; i-- {
			// converting digit into int
			if currentd, err = strconv.Atoi(string(numberpart[i])); err != nil {
				return false
			}
			checkd += counter * currentd
			counter++
			//fmt.Printf("counter: %d currentd: %d checkd: %d\n", counter, currentd, checkd)
		}
	}
	return checkd%10 == checkdigit
}
