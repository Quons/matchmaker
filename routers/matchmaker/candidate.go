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
		appG.Response(nil, e.ERROR_INVALID_PARAMS)
		return
	}
	log.WithField("courseId", candidate).Info()
	err = matcher.AddCandidate(candidate)
	if err != nil {
		log.Errorf("add_candidate_err:%v", err)
		appG.Response(nil, e.ERROR_DATA_ERROR)
		return
	}
	appG.Response(nil, e.SUCCESS)
}
