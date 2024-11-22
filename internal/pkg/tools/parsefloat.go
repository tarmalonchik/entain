package tools

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func CentsPrettyPrinted(in int64) string {
	fPart := strconv.FormatInt(in/100, 10)
	lPart := strconv.FormatInt(in%100, 10)

	if len(lPart) == 1 {
		lPart = "0" + lPart
	}
	return fPart + "." + lPart
}

func ConvertNonNegativeFloatToCents(in string) (int64, error) {
	if in == "" {
		return 0, errors.New("empty input")
	}
	if in[0] == '-' {
		return 0, errors.New("is negative")
	}

	items := strings.Split(in, ".")

	if len(items) > 2 || len(items) == 0 {
		return 0, errors.New("invalid float")
	}

	if len(items) == 1 {
		return getFirstPart(items[0])
	}

	fPart, err := getFirstPart(items[0])
	if err != nil {
		return 0, err
	}

	lPart, err := getLastPart(items[1])
	if err != nil {
		return 0, err
	}

	return overflowCheck(fPart, lPart)
}

func getFirstPart(input string) (int64, error) {
	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, err
	}

	if val > 92233720368547758 || val < 0 {
		return 0, errors.New("overflow")
	}

	return val * 100, nil
}

func getLastPart(input string) (int64, error) {
	if len(input) > 2 {
		return 0, nil
	}

	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, err
	}

	if val < 0 {
		return 0, errors.New("invalid")
	}

	if len(input) == 1 {
		return val * 10, nil
	}

	return val, nil
}

func overflowCheck(balance, addition int64) (int64, error) {
	if balance == 0 || addition == 0 {
		return balance + addition, nil
	}
	if addition > 0 {
		if balance > math.MaxInt64-addition {
			return 0, errors.New("overflow")
		}
	} else {
		if balance < math.MaxInt64-addition {
			return 0, errors.New("overflow")
		}
	}
	return balance + addition, nil
}
