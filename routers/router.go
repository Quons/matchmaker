package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//启用swagger文档
	_ "github.com/Quons/matchmaker/docs"
	"github.com/Quons/matchmaker/pkg/logging"
	"github.com/Quons/matchmaker/pkg/setting"
	"github.com/Quons/matchmaker/routers/matchmaker"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

/*路由注册*/
func registerRouter(r *gin.Engine) {
	apiv2 := r.Group("/matchmaker/v1")

	//验证token
	{
		//新增候选人
		apiv2.POST("/candidate/add", matchmaker.AddCandidate)
		//修改推送状态
		apiv2.POST("/candidate/updateStatus", matchmaker.UpdatePushStatus)
		apiv2.GET("/candidate/getDemo", matchmaker.GetDemo)
	}
}

/*InitRouter 初始化路由*/
func InitRouter() *gin.Engine {
	gin.SetMode(setting.ServerSetting.RunMode)
	r := gin.New()
	//设置gin日志输出writer
	r.Use(gin.LoggerWithWriter(logging.GetGinLogWriter()))
	//设置gin恢复日志数据writer
	r.Use(gin.RecoveryWithWriter(logging.GetGinLogWriter()))
	//静态目录
	r.StaticFS("/static", http.Dir(setting.AppSetting.LogPath+"/static"))
	//r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	//r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	//swagger自动文档路径
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//注册业务路由
	registerRouter(r)
	return r
}
