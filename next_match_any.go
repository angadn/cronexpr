package cronexpr

import (
	"fmt"
	"time"
)

var (
	// ErrNoIntersectionPossible if two expressions will never intersect for any time.
	ErrNoIntersectionPossible = fmt.Errorf("no intersection possible among the given expressions")

	// ErrImpossibleExpression is for when a cron-expression is illegal or impossible to
	// satisfy.
	ErrImpossibleExpression = fmt.Errorf("cron expression is illegal or impossible to satisfy")
)

// NextMatchAny finds the closest time instance that is any of the Expressions. All
// expressions are evaluated in the same time-zone as `fromTime`.
func NextMatchAny(fromTime time.Time, expressions ...*Expression) (time.Time, error) {
	var minRet time.Time

	for _, exp := range expressions {
		if nextMatch := exp.Next(fromTime, NextIfNotMatched); nextMatch.Equal(fromTime) {
			return nextMatch, nil
		} else if nextMatch.After(fromTime) {
			if minRet.Equal(time.Time{}) {
				// First time init
				minRet = nextMatch
			} else {
				minRet = MinTime(minRet, nextMatch)
			}
		}
	}

	if minRet.Equal(time.Time{}) {
		// Same as when defined - either no expressions or all invalid
		return minRet, ErrImpossibleExpression
	}

	return minRet, nil
}

// MaxTime returns the maximal value between `a` and `b`.
func MaxTime(a time.Time, b time.Time) time.Time {
	if a.After(b) {
		return a
	}

	return b
}

// MinTime returns the minimal value between `a` and `b`.
func MinTime(a time.Time, b time.Time) time.Time {
	if a.Before(b) {
		return a
	}

	return b
}
