package cronexpr

import (
	"log"
	"testing"
	"time"
)

func TestIntersecting(t *testing.T) {
	exp1 := MustParse("* * * * *")
	exp2 := MustParse("*/15 * * * *")
	exp3 := MustParse("0 * * * *")
	exp4 := MustParse("* 9-17 * * FRI")
	exp5 := MustParse("* 8-19 * * SUN-SAT,")

	if _, err := NextMatch(
		time.Now(), exp1, exp2, exp3, exp4, exp5,
	); err != nil {
		log.Printf(err.Error())
		t.Fail()
	}
}

func TestNonIntersecting(t *testing.T) {
	exp1 := MustParse("* * * * *")
	exp2 := MustParse("* */2 * * *")
	exp3 := MustParse("* */3 * * *")
	exp4 := MustParse("* * * * MON")
	exp5 := MustParse("* * * * TUE")

	if _, err := NextMatch(
		time.Now(), exp1, exp2, exp3, exp4, exp5,
	); err == nil {
		// Must fail if no error!
		t.Fail()
	} else {
		log.Printf("test passed by error as expected: %s\n", err.Error())
	}
}
