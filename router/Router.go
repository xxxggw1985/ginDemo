package router

import (
	"ginDemo/common/function"
	"ginDemo/controller/v1"
	"ginDemo/controller/v2"
	"ginDemo/middleware/sign"
	"github.com/gin-gonic/gin"
	"net/url"
	"strconv"
)

func InitRouter(r *gin.Engine) {

	r.GET("/sn", SignDemo)

	// v1 版本
	GroupV1 := r.Group("/v1")
	{

		GroupV1.Any("/member/add", v1.AddMember)
		GroupV1.Any("/member/edit", v1.EditMember)
	}

	// v2 版本
	GroupV2 := r.Group("/v2", sign.Sign())
	{

		GroupV2.Any("/member/add", v2.AddMember)
	}
}

func SignDemo(c *gin.Context) {
	ts := strconv.FormatInt(function.GetTimeUnix(), 10)
	res := map[string]interface{}{}
	params := url.Values{
		"name":  []string{"a"},
		"price": []string{"10"},
		"ts":    []string{ts},
	}
	res["sn"] = function.CreateSign(params)
	res["ts"] = ts
	function.RetJson("200", "", res, c)
}
