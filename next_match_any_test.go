package cronexpr

import (
	"log"
	"testing"
	"time"
)

func TestMatchAny(t *testing.T) {
	expressions := []*Expression{
		MustParse("30-59 8 * * *"),
		MustParse("* 9-19 * * *"),
		MustParse("1-30 19 * * *"),
	}

	st, _ := time.Parse(time.RFC3339, "2014-11-12T8:30:00.000Z")
	t1, _ := time.Parse(time.RFC3339, "2014-11-12T10:00:00.000Z")
	t2, _ := time.Parse(time.RFC3339, "2014-11-12T8:00:00.000Z")
	t3, _ := time.Parse(time.RFC3339, "2014-11-12T20:00:00.000Z")
	tom, _ := time.Parse(time.RFC3339, "2014-11-13T8:30:00.000Z")

	if ret, err := NextMatchAny(t1, expressions...); (err != nil) || (!ret.Equal(t1)) {
		log.Printf("couldn't match time that lies within range")
		t.Fail()
	}

	if ret, err := NextMatchAny(t2, expressions...); (err != nil) || (!ret.Equal(st)) {
		log.Printf("couldn't suggest start of range today")
		t.Fail()
	}

	if ret, err := NextMatchAny(t3, expressions...); (err != nil) || (!ret.Equal(tom)) {
		log.Printf("couldn't suggest start of range tomorrow")
		t.Fail()
	}
}
