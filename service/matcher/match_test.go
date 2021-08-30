package matcher

import (
	"testing"
)

func TestDoMatch(t *testing.T) {
	DailyMatch()
	select {} //阻塞主线程停止
}
