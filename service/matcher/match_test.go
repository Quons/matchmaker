package matcher

import (
	"math"
	"testing"
)

func TestDoMatch(t *testing.T) {
	DailyMatch()
	select {} //阻塞主线程停止
}

func TestAbs(t *testing.T) {
	t.Log(math.Abs(-1.5))

}
