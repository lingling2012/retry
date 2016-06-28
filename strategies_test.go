package retry

import (
	"testing"
	"time"
)

func TestLimit(t *testing.T) {
	const attemptLimit = 3

	strategy := Limit(attemptLimit)

	if !strategy(1) {
		t.Error("strategy expected to return true")
	}

	if !strategy(2) {
		t.Error("strategy expected to return true")
	}

	if strategy(3) {
		t.Error("strategy expected to return false")
	}
}

func TestDelay(t *testing.T) {
	const delayDuration = time.Duration(10 * time.Millisecond)

	strategy := Delay(delayDuration)

	if now := time.Now(); !strategy(0) || delayDuration > time.Since(now) {
		t.Errorf(
			"strategy expected to return true in %s",
			time.Duration(delayDuration),
		)
	}

	if now := time.Now(); !strategy(5) || 0 != time.Since(now) {
		t.Error("strategy expected to return true in 0 time")
	}
}

func TestWait(t *testing.T) {
	const waitDuration = time.Duration(10 * time.Millisecond)

	strategy := Wait(waitDuration)

	if now := time.Now(); !strategy(0) || 0 != time.Since(now) {
		t.Error("strategy expected to return true in 0 time")
	}

	if now := time.Now(); !strategy(1) || waitDuration > time.Since(now) {
		t.Errorf(
			"strategy expected to return true in %s",
			time.Duration(waitDuration),
		)
	}
}
