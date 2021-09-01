package matcher

import "testing"

func TestRetrySendMail(t *testing.T) {
	RetrySendMail()
}

func TestRetrySendMailCycle(t *testing.T) {
	RetrySendMailCycle()
	select {}
}
