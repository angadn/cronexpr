package cronexpr

import (
	"fmt"
	"math"
	"time"
)

var (
	// ErrNoIntersectionPossible if two expressions will never intersect for any time.
	ErrNoIntersectionPossible = fmt.Errorf("no intersection possible among the given expressions")
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
		for _, e := range expressions {
			// Get the maximal value for the next iteration of nextTime
			maxNextTime = time.Unix(0, int64(math.Max(
				float64(maxNextTime.UnixNano()),
				float64(e.Next(nextTime, NextIfNotMatched).UnixNano()),
			))).In(fromTime.Location())
		}

		if maxNextTime == nextTime {
			return nextTime, nil
		}

		nextTime = maxNextTime
	}

	return nextTime, ErrNoIntersectionPossible
}
