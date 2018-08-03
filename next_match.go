package cronexpr

import (
	"fmt"
	"math"
	"time"
)

// NextMatch finds the closest time instance that is agreeable to all Expressions.
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
			)))
		}

		if maxNextTime == nextTime {
			return nextTime, nil
		}

		nextTime = maxNextTime
	}

	return nextTime, fmt.Errorf("no intersection is impossible among the given expressions")
}
