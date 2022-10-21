package handler

import (
	"covid19-api/internal/pkg/response"
	"covid19-api/internal/service"
	"github.com/gin-gonic/gin"
)

//overall handlers

func GetOverall(c *gin.Context) {
	overallService := service.NewOverallService(c)
	result, err := overallService.GetLatestOverall()
	if err != nil {
		response.ErrorByCode(c, 200, err.Error())
	}
	response.Success(c, result)
}

func ListOverall(c *gin.Context) {
	overallService := service.NewOverallService(c)
	result, err := overallService.ListOverall()
	if err != nil {
		response.ErrorByCode(c, 200, err.Error())
	}
	response.Success(c, result)
}

//area handler

func ListProvince(c *gin.Context) {
	provinceService := service.NewAreaService(c)
	result, err := provinceService.ListProvince()
	if err != nil {
		response.ErrorByCode(c, 200, err.Error())
	}
	response.Success(c, result)
}

//abroad handler

func ListAbroad(c *gin.Context) {
	abroadService := service.NewAbroadService(c)
	result, err := abroadService.GetLatest()
	if err != nil {
		response.ErrorByCode(c, 200, err.Error())
	}
	response.Success(c, result)
}

func ListTimeline(c *gin.Context) {
	timelineService := service.NewTimelineService(c)
	result, err := timelineService.GetLatest()
	if err != nil {
		response.ErrorByCode(c, 200, err.Error())
	}
	response.Success(c, result)
}

func ListProvinceName(c *gin.Context) {
	provinceName := []string{
		"上海市", "云南省", "内蒙古自治区", "北京市", "台湾", "吉林省", "四川省", "天津市", "宁夏回族自治区",
		"安徽省", "山东省", "山西省", "广东省", "广西壮族自治区", "新疆维吾尔自治区", "江苏省",
		"江西省", "河北省", "河南省", "浙江省",
	}
	response.Success(c, provinceName)
}

func ListProvinceShortName(c *gin.Context) {
	provinceShortName := []string{
		"上海", "云南", "内蒙古", "北京", "台湾", "吉林", "四川", "天津",
		"宁夏", "安徽", "山东", "山西", "广东", "广西", "新疆", "江苏",
		"江西", "河北", "河南", "浙江",
	}
	response.Success(c, provinceShortName)
}
