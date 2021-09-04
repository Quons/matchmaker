package matcher

import (
	"github.com/Quons/matchmaker/models"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

func RetrySendMailCycle() {
	c := cron.New() //精确到秒

	//定时任务 每周三中午12点整
	spec := "0 0/10 * * * ? " //cron表达式，
	_ = c.AddFunc(spec, func() {
		RetrySendMail()
	})

	c.Start()
}

// RetrySendMail 重新发送邮件
func RetrySendMail() {
	// 先查询所有发送失败的记录
	failedList, err := models.GetFailedMatchRecordList()
	if err != nil {
		logrus.WithField("err", err).Error("get_failed_match_record_err")
		return
	}
	// 判断如果是男生男生发送失败，就重发男生的
	for _, record := range failedList {
		if record.FemaleSendStatus == models.SendStatusOK && record.MaleSendStatus == models.SendStatusOK {
			continue
		}
		retrySend(record)
	}
}

func retrySend(record models.MatchRecord) {
	if record.MaleSendStatus != models.SendStatusOK {
		// 给男生发送邮件
		err := sendEmailAndRecordStatus(record, models.GenderMale)
		if err != nil {
			logrus.WithField("err", err).Error("retry_send_mail_err")
		}
	}

	if record.FemaleSendStatus != models.SendStatusOK {
		// 给女生发送邮件
		err := sendEmailAndRecordStatus(record, models.GenderFemale)
		if err != nil {
			logrus.WithField("err", err).Error("retry_send_mail_err")
		}
	}
}
