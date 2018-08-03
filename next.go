package cronexpr

import (
	"fmt"
	"math"
	"time"
)

// Next finds the closest time instance that is agreeable to all Expressions.
func Next(fromTime time.Time, expressions ...Expression) (time.Time, error) {
	var (
		n        int
		nextTime time.Time
	)

	n = len(expressions)
	for passes := 0; passes <= n; passes++ {
		maxNextTime := nextTime
		for _, e := range expressions {
			// Get the maximal value for the next iteration of nextTime
			maxNextTime = time.Unix(0, int64(math.Max(
				float64(maxNextTime.UnixNano()),
				float64(e.Next(nextTime).UnixNano()),
			)))
		}

		if maxNextTime == nextTime {
			return nextTime, nil
		}

		nextTime = maxNextTime
	}

	return nextTime, fmt.Errorf("no intersection is impossible among the given expressions")
}
