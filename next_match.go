package cronexpr

import (
	"fmt"
	"log"
	"math"
	"time"
)

var (
	// ErrNoIntersectionPossible if two expressions will never intersect for any time.
	ErrNoIntersectionPossible = fmt.Errorf("no intersection possible among the given expressions")

	// ErrImpossibleExpression is for when a cron-expression is illegal or impossible to
	// satisfy.
	ErrImpossibleExpression = fmt.Errorf("cron expression is illegal or impossible to satisfy")
)

// NextMatch finds the closest time instance that is agreeable to all Expressions. All
// expressions are evaluated in the same time-zone as `fromTime`.
func NextMatch(fromTime time.Time, expressions ...*Expression) (time.Time, error) {
	var (
		n        int
		nextTime time.Time
	)

	n = len(expressions)
	nextTime = fromTime

	for passes := 0; passes <= n; passes++ {
		maxNextTime := nextTime
		for i, e := range expressions {
			// Get the maximal value for the next iteration of nextTime
			var expTime time.Time
			if expTime = e.Next(nextTime, NextIfNotMatched); expTime.Equal(time.Time{}) {
				log.Printf(
					"expression at index %d is illegal or impossible to satisfy\n", i,
				)
				return nextTime, ErrImpossibleExpression
			}

			maxNextTime = MaxTime(maxNextTime, expTime).In(fromTime.Location())
		}

		if maxNextTime == nextTime {
			return nextTime, nil
		}

		nextTime = maxNextTime
	}

	return nextTime, ErrNoIntersectionPossible
}

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

func MaxTime(a time.Time, b time.Time) time.Time {
	return time.Unix(0, int64(math.Max(
		float64(a.UnixNano()),
		float64(b.UnixNano()),
	)))
}

func MinTime(a time.Time, b time.Time) time.Time {
	return time.Unix(0, int64(math.Min(
		float64(a.UnixNano()),
		float64(b.UnixNano()),
	)))
}
