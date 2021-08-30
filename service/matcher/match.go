package matcher

import (
	"fmt"
	"github.com/Quons/matchmaker/models"
	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
	"math/rand"
	"testing"
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
	// 得到匹配的map
	matchedMap := make(map[string]struct{})
	for _, record := range matchedRecordList {
		matchedMap[getMatchKey(record.MaleID, record.FemaleID)] = struct{}{}
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
		err := models.UpdateMatchRecord(record.ID)
		if err != nil {
			log.Errorf("update_send_status_err:%v", err)
		}
	}

	//}
}

// 获取匹配key
func getMatchKey(maleID, femaleID int64) string {
	return fmt.Sprintf("m%v_f%v", maleID, femaleID)
}

func match(maleList []models.Candidate, femaleList []models.Candidate, matchedMap map[string]struct{}) []models.MatchRecord {
	var matchingList []models.MatchRecord
	// 遍历每个男生
	for _, male := range maleList {
	inner:
		for _, female := range femaleList {
			// 得到key
			matchedKey := getMatchKey(male.ID, female.ID)
			// 如果已经匹配过，跳过
			_, ok := matchedMap[matchedKey]
			if ok {
				continue
			}
			// 否则记录下来需要匹配的用户
			matchingList = append(matchingList, models.MatchRecord{
				MaleID:         male.ID,
				FemaleID:       female.ID,
				MaleEmail:      male.Email,
				FemaleEmail:    female.Email,
				CreateTime:     time.Now().Unix(),
				SendStatus:     models.SendStatusNo,
				ConfirmStatus:  models.ConfirmStatusUnknown,
				MatchResult:    models.AttitudeUnknown,
				MaleAttitude:   models.AttitudeUnknown,
				FemaleAttitude: models.AttitudeUnknown,
			})
			// 每个男生匹配一次女生就停止，并且移除当前女生
			femaleList = femaleList[1:]
			break inner
		}
	}
	return matchingList
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

func TestSendMail(t *testing.T) {
	m := gomail.NewMessage()

	//发送人
	m.SetHeader("From", "cugses@qq.com")
	//接收人
	m.SetHeader("To", "cugjyb@163.com")
	//抄送人
	//m.SetAddressHeader("Cc", "xxx@qq.com", "xiaozhujiao")
	//主题
	m.SetHeader("Subject", "小佩奇")
	//内容
	m.SetBody("text/html", "<h1>新年快乐</h1>")
	//附件
	//m.Attach("./myIpPic.png")

	//拿到token，并进行连接,第4个参数是填授权码
	d := gomail.NewDialer("smtp.qq.com", 587, "1506648063@qq.com", "cfuvljpevkjghgfg")

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("DialAndSend err %v:", err)
		panic(err)
	}
	fmt.Printf("send mail success\n")
}
