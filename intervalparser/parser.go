package intervalparser

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	START int = iota
	FIRST
	SECOND
	CLOSE
)

// A custom error so the user can decide wether he wants to ignore it or not
var ErrParsingFinishedAbrubtly = errors.New("the interval string seems to have ended abruptly")

// ParseIntervalsFromString parses all intervals from the intervalString of the form [a,b]
// with a and b being integers. Whitespaces and commas between the intervals are possible but not
// mandatory.
func ParseIntervalsFromString(intervalString string) ([][]int, error) {
	var intervals [][]int
	status := START
	var firstInt, secondInt string
	for _, c := range intervalString {
		// accept whitespace no matter the status
		if string(c) == " " {
			continue
		}
		switch status {
		case START:
			if string(c) == "," {
				continue
			}
			if string(c) != "[" {
				return nil, fmt.Errorf("could not find opening bracket for intervall")
			}
			status = FIRST
		case FIRST:
			// first number complete
			if string(c) == "," && firstInt != "" {
				status = SECOND
				break
			}
			if _, err := strconv.Atoi(string(c)); err != nil {
				return nil, fmt.Errorf("%q is not a number\n", string(c))
			}
			firstInt = firstInt + string(c)
		case SECOND:
			// in contrast to the first number we do not want to consume another token
			// after completing the number, instead we fallthrough into the CLOSE case
			// second number complete
			if string(c) == "]" && secondInt != "" {
				status = CLOSE
			} else {
				if _, err := strconv.Atoi(string(c)); err != nil {
					return nil, fmt.Errorf("%q is not a number\n", string(c))
				}
				secondInt = secondInt + string(c)
				break
			}
			fallthrough
		case CLOSE:
			fI, err := strconv.Atoi(firstInt)
			if err != nil {
				return nil, fmt.Errorf("unable to convert int: %w", err)
			}
			sI, err := strconv.Atoi(secondInt)
			if err != nil {
				return nil, fmt.Errorf("unable to convert int: %w", err)
			}
			intervals = append(intervals, []int{fI, sI})
			firstInt = ""
			secondInt = ""
			status = START
		default:
			return nil, fmt.Errorf("unknown status")
		}
	}
	if status != START {
		return intervals, ErrParsingFinishedAbrubtly
	}
	return intervals, nil
}
