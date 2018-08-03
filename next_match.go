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

			maxNextTime = time.Unix(0, int64(math.Max(
				float64(maxNextTime.UnixNano()),
				float64(expTime.UnixNano()),
			))).In(fromTime.Location())
		}

		if maxNextTime == nextTime {
			return nextTime, nil
		}

		nextTime = maxNextTime
	}

	return nextTime, ErrNoIntersectionPossible
}
