package matcher

import (
	"fmt"
	"github.com/Quons/matchmaker/models"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"math"
	"math/rand"
	"time"
)

func DailyMatch() {
	c := cron.New() //精确到秒

	//定时任务
	spec := "*/5 * * * * ?" //cron表达式，每秒一次
	_ = c.AddFunc(spec, func() {
		DoMatch()
	})

	c.Start()
}

// 为每个男生和女生匹配交流对象
// 每个男生匹配一个没有交流过的女生
// 每个女生匹配一个没有交流过的男生
// 每天每个人只匹配一个对象
func DoMatch() {
	// 男生
	maleList, err := models.GetCandidateList(models.GenderMale)
	if err != nil {
		log.Errorf("get_male_list_err:%+v\n", err)
		return
	}

	// 女生
	femaleList, err := models.GetCandidateList(models.GenderFemale)
	if err != nil {
		log.Errorf("get_female_list_err:%+v\n", err)
		return
	}

	// 已匹配过的集合
	// 查询已经匹配过的集合
	matchedRecordList, err := models.GetMatchRecordList()
	if err != nil {
		log.Errorf("get_matched_record_err:%v\n", err)
		return
	}
	// 得到匹配的map，这里应该直接用邮箱来做匹配校验
	matchedMap := make(map[string]struct{})
	for _, record := range matchedRecordList {
		matchedMap[getMatchKey(record.MaleEmail, record.FemaleEmail)] = struct{}{}
	}
	//for {
	Shuffle(femaleList)
	//fmt.Println(femaleList)
	matchingList := match(maleList, femaleList, matchedMap)
	if len(matchingList) == 0 {
		log.Infof("empty_match_result")
		return
	}
	for i := 0; i < len(matchingList); i++ {
		err := models.AddMatchRecord(&matchingList[i])
		if err != nil {
			log.Errorf("add_record_err:%v\n", err)
			return
		}

	}
	fmt.Println(matchingList)
	//// 添加到已匹配的结果中 TODO 修改成批量处理的
	//for _, record := range matchingList {
	//	//record := record
	//	//fmt.Println(record)
	//
	//}

	// TODO 发送邮件
	for _, record := range matchingList {
		fmt.Println(record.ID)
		fmt.Printf("send_email,from:%v,to:%v\n", record.MaleEmail, record.FemaleEmail)
		// 向男生发邮件
		_ = sendEmailAndRecordStatus(record, models.GenderMale)
		// 向女生发送邮件
		_ = sendEmailAndRecordStatus(record, models.GenderFemale)
	}
}

// 发送邮件
func sendEmailAndRecordStatus(record models.MatchRecord, gender models.Gender) error {
	target, matched := getSendPair(record, gender)
	// 获取发送方和接收方
	err := SendEmail(target, matched)
	if err != nil {
		log.Errorf("send_to_male_err:%v", err)
		return err
	}
	err = models.UpdateMatchRecord(record.ID, gender)
	if err != nil {
		log.Errorf("update_send_status_err:%v", err)
		return err
	}
	return nil
}

// 获取待发送的双方邮件
func getSendPair(record models.MatchRecord, gender models.Gender) (target, matched string) {
	if gender == models.GenderMale {
		return record.MaleEmail, record.FemaleEmail
	}
	return record.FemaleEmail, record.MaleEmail
}

// 获取匹配key
func getMatchKey(male, female string) string {
	return fmt.Sprintf("m%v_f%v", male, female)
}

// 男生女生匹配
// 匹配规则：
// 1,没有匹配过
// 2,异性
// 3,年龄差小于5岁
// ...
func match(maleList []models.Candidate, femaleList []models.Candidate, matchedMap map[string]struct{}) []models.MatchRecord {
	var matchingList []models.MatchRecord
	// 遍历每个男生
	for _, male := range maleList {
	inner:
		for _, female := range femaleList {
			// 得到key
			matchedKey := getMatchKey(male.Email, female.Email)
			// 如果已经匹配过，跳过
			_, ok := matchedMap[matchedKey]
			if ok {
				continue
			}
			// 匹配规则引擎进行匹配判断
			isMatched := matchEngine(male, female)
			if !isMatched {
				continue
			}
			// 否则记录下来需要匹配的用户
			matchingList = append(matchingList, models.MatchRecord{
				MaleID:           male.ID,
				FemaleID:         female.ID,
				MaleEmail:        male.Email,
				FemaleEmail:      female.Email,
				CreateTime:       time.Now().Unix(),
				MaleSendStatus:   models.SendStatusNo,
				FemaleSendStatus: models.SendStatusNo,
				ConfirmStatus:    models.ConfirmStatusUnknown,
				MatchResult:      models.AttitudeUnknown,
				MaleAttitude:     models.AttitudeUnknown,
				FemaleAttitude:   models.AttitudeUnknown,
			})
			// 每个男生匹配一次女生就停止，并且移除当前女生
			femaleList = femaleList[1:]
			break inner
		}
	}
	return matchingList
}

// 配对引擎
// 匹配规则：
// 1,没有匹配过
// 2,异性
// 3,年龄差小于5岁
// ...
func matchEngine(male, female models.Candidate) (matched bool) {
	// TODO 增加一些校验逻辑
	// 1,必须为异性，同性匹配后期可以增加~
	if male.Gender == female.Gender {
		return false
	}
	// 2，年龄差需要小于5岁
	if math.Abs(float64(male.Age-female.Age)) >= 5 {
		return false
	}
	return true
}

func Shuffle(slice []models.Candidate) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		// 随机取一个与结尾元素交换
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

// SendEmail 发送邮件
func SendEmail(targetEmail, matchedEmail string) error {
	m := gomail.NewMessage()

	//发送人
	m.SetHeader("From", "cugses@qq.com")
	//接收人
	m.SetHeader("To", targetEmail)
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "吉大群提醒牵线邮件")
	//内容
	m.SetBody("text/html", fmt.Sprintf("<h1>您本次的有缘人邮箱为:%v</h1>", matchedEmail))
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "1506648063@qq.com", "cfuvljpevkjghgfg")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		// TODO 这里需要区分下错误类型，如邮箱地址错误
		log.Errorf("send_email_err:%v,target_candidate:%+v,matched_candiate:%+v", err, targetEmail, matchedEmail)
		return err
	}
	log.Infof("send_email_success:target_candidate:%+v,matched_candiate:%+v", targetEmail, matchedEmail)
	return nil

}
