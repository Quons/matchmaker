package matcher

import (
	"fmt"
	"github.com/robfig/cron"
	"testing"
)

func TestDoMatch(t *testing.T) {
	DailyMatch()
	//DoMatch()
	select {} //阻塞主线程停止
}

func TestAbs(t *testing.T) {
	//定时任务
	c := cron.New()
	//spec := "* * * * * ?" //cron表达式， 0 3 21 ? * 3 每周2的晚上21点3分执行一次
	spec := "0 31 21 ? * 2" //cron表达式， 0 3 21 ? * 3 每周2的晚上21点3分执行一次
	err := c.AddFunc(spec, func() {
		fmt.Println("................")
		//DoMatch()
	})
	if err != nil {
		t.Error(err)
		return
	}
	c.Start()

	select {}
}

func TestSendEmail(t *testing.T) {

	//err := SendEmail("rongquanwang@ustb.edu.cn", "565905306@qq.com")
	err := SendEmail("1459170090@qq.com", "690106233@qq.com")
	//err := SendEmail("rongquanwang@ustb.edu.cn", "ermh1712@163.com")
	if err != nil {
		t.Error(err)
		return
	}

}
