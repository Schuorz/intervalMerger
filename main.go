package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"

	"intervalMerger/intervalparser"
)

// arrayIsInterval checks wether the slice can be treated as and interval
// -> only two values, the first must be smaller or equal than the second
func arrayIsInterval(a []int) bool {
	if len(a) != 2 {
		return false
	}
	if a[0] > a[1] {
		return false
	}

	return true
}

// mergeIntervals tries to merge two slices which represent an interval.
// both arrays need to be of the form [lower_bound, upper_bound].
// There are three possible results.
// 1. The intervals are not overlapping -> return both intervals (slices) in order
// 2. The intervals are partially overlapping -> return the lowest lower bound and the highest upper bound as one new interval (slice)
// 3. One interval encases the other -> return the encasing interval (slice)
func mergeIntervals(a, b []int) ([][]int, error) {
	if !arrayIsInterval(a) || !arrayIsInterval(b) {
		return nil, fmt.Errorf(
			"at least one slice is not a correct interval, a = %+v, b = %+v",
			a, b)
	}

	// find interval with the smaller lower bound f(first) the other one will be s(second)
	var f, s *[]int
	if a[0] <= b[0] {
		f, s = &a, &b
	} else {
		f, s = &b, &a
	}

	// the first intervall ends before the second begins,
	// return both in order [f1,f2],[s1,s2]
	if (*f)[1] <= (*s)[0] {
		return [][]int{*f, *s}, nil
	}

	// the first interval ends in or at the end of the second,
	// new intervall is [f1,s2]
	if (*f)[1] <= (*s)[1] {
		return [][]int{{(*f)[0], (*s)[1]}}, nil
	}

	// since the former cases did not apply,
	// the upper bound of f is bigger than the upper bound of s,
	// f encases s, return [f1,f2]
	return [][]int{*f}, nil
}

// worst case gaussian -> (n²+2)/2 -> O(n²)
// best case O(n)
// robust -> sort für laufzeit, speichern für platz + zeit
func MERGE(intervals [][]int) ([][]int, error) {
	if len(intervals) < 2 {
		return nil, fmt.Errorf("at least two intervals required")
	}
	// mergedIntervals hold all the intervals in ascending order
	// that were merged/ for which a merge was already attempted.
	// populate with the first two intervals (/their merge)
	mergedIntervals, err := mergeIntervals(intervals[0], intervals[1])
	if err != nil {
		return nil, err
	}

	for i := 2; i < len(intervals); i++ {
		for j, mI := range mergedIntervals {
			subMerge, err := mergeIntervals(intervals[i], mI)
			if err != nil {
				return nil, err
			}

			// two cases:
			// 1. if two intervals where in fact merged, replace the old one with the new one
			// nothing further needs to be done, since the intervals in mergedIntervals are sorted
			// 2. if two intervals were not merged, continue ONLY if the one that is currently merged
			// into mergedIntervals is returned as the latter interval. Since the intervals are sorted
			// merges only need to be attempted until we are out of bound
			if len(subMerge) == 1 || reflect.DeepEqual(subMerge[0], intervals[i]) {
				mergedIntervals = slices.Replace(mergedIntervals, j, j+1, subMerge...)
				break
			}

			// if we are in the last run and the current interval was returned second, it needs to be appended at the end
			if j == len(mergedIntervals)-1 {
				// since the last if did not run, this needs to be true
				// if not then the logic is broken, check anyways
				if !reflect.DeepEqual(subMerge[1], intervals[i]) {
					return nil, fmt.Errorf("interval to be merged could not be sorted into the intervals, something is serioulsy wrong")
				}
				mergedIntervals = slices.Replace(mergedIntervals, j, j+1, subMerge...)
			}
		}
	}

	return mergedIntervals, nil
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("please provide either a string in quotes with intervals")
	}

	firstArg := os.Args[1]

	intervals, err := intervalparser.ParseIntervalsFromString(firstArg)
	if err != nil {
		log.Fatalf("unable to parse intervals: %s", err.Error())
	}

	merged, err := MERGE(intervals)
	if err != nil {
		log.Fatalf("could not merge intervals: %w", err)
	}

	fmt.Printf("Merged intervals:\n %v", merged)
}
