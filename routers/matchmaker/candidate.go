package matchmaker

import (
	"github.com/Quons/matchmaker/pkg/app"
	"github.com/Quons/matchmaker/pkg/e"
	"github.com/Quons/matchmaker/service/matcher"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// @Summary 增加候选人
// @Produce  json
// @Param image post file true "图片文件"
// @Success 200 {string} json "{"code":200,"data":{"image_save_url":"upload/images/96a.jpg", "image_url": "http://..."}"
// @Router /api/v1/tags/import [post]
func AddCandidate(c *gin.Context) {
	appG := app.Gin{C: c}
	candidate := matcher.Candidate{}
	err := c.ShouldBind(&candidate)
	if err != nil {
		log.Info(err)
		appG.Response("参数错误", e.ERROR_INVALID_PARAMS)
		return
	}
	log.WithField("courseId", candidate).Info()
	err = matcher.AddCandidate(candidate)
	if err != nil {
		log.Errorf("add_candidate_err:%v", err)
		appG.Response(err.Error(), e.ERROR_DATA_ERROR)
		return
	}
	appG.Response("提交信息成功", e.SUCCESS)
}

// @Summary 修改推送状态
// @Produce  json
// @Param image post file true "图片文件"
// @Success 200 {string} json "{"code":200,"data":{"image_save_url":"upload/images/96a.jpg", "image_url": "http://..."}"
// @Router /api/v1/tags/import [post]
func UpdatePushStatus(c *gin.Context) {
	appG := app.Gin{C: c}
	pushStatus := matcher.PushStatus{}
	err := c.ShouldBind(&pushStatus)
	if err != nil {
		log.Info(err)
		appG.Response("状态修改失败：参数错误", e.ERROR_INVALID_PARAMS)
		return
	}
	log.WithField("courseId", pushStatus).Info()
	err = matcher.UpdatePushStatus(pushStatus)
	if err != nil {
		log.Errorf("update_candidate_err:%v", err)
		appG.Response(err.Error(), e.ERROR_DATA_ERROR)
		return
	}
	appG.Response("更新推送状态成功", e.SUCCESS)
}

// @Summary 测试接口
// @Produce  json
// @Param image post file true "图片文件"
// @Success 200 {string} json "{"code":200,"data":{"image_save_url":"upload/images/96a.jpg", "image_url": "http://..."}"
// @Router /api/v1/tags/import [post]
func GetDemo(c *gin.Context) {
	appG := app.Gin{C: c}
	data := make(map[string]interface{})
	data["name"] = "张三"
	data["age"] = 18
	data["sex"] = "男"
	data["id"] = "001"
	appG.Response(data, e.SUCCESS)
}
