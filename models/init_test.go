package models

import (
	"github.com/Quons/matchmaker/pkg/setting"
)

func init() {
	setting.Setup("/Users/didi/go/src/matchmaker/conf/dev.ini")
	Setup()
}
