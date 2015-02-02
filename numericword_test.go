package numericword

import (
	"testing"
)

func verify(t *testing.T, i int64, s string) {
	res := ToWords(i)
	if res != s {
		t.Errorf("Converting %d\nExpected '%s'\nReceived '%s'", i, s, res)
	}
}

func TestSmallPositiveValues(t *testing.T) {
	verify(t, 0, "zero")
	verify(t, 10, "ten")
	verify(t, 55, "fifty five")
	verify(t, 100, "one hundred")
	verify(t, 101, "one hundred one")
	verify(t, 111, "one hundred eleven")
	verify(t, 859, "eight hundred fifty nine")
}

func TestLargePositiveValues(t *testing.T) {
	verify(t, 1000, "one thousand")
	verify(t, 1001, "one thousand one")
	verify(t, 1141, "one thousand one hundred forty one")
	verify(t, 9400, "nine thousand four hundred")
	verify(t, 900001, "nine hundred thousand one")
	verify(t, 34058689, "thirty four million fifty eight thousand six hundred eighty nine")
}

func TestNegativeValues(t *testing.T) {
	verify(t, -10, "negative ten")
	verify(t, -55, "negative fifty five")
	verify(t, -100, "negative one hundred")
	verify(t, -101, "negative one hundred one")
	verify(t, -111, "negative one hundred eleven")
	verify(t, -859, "negative eight hundred fifty nine")
}
