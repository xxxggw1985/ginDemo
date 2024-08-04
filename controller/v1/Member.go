package v1

import (
	"ginDemo/entity"
	"ginDemo/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddMember(c *gin.Context) {
	//获取参数
	res := entity.Result{}
	mem := entity.Member{}

	if err := c.ShouldBind(&mem); err != nil {
		res.SetCode(entity.CODE_ERROR)
		res.SetMessage(err.Error())
		c.JSON(http.StatusForbidden, res)
		c.Abort()
		return
	}

	// 处理业务(下次再分享)

	data := map[string]interface{}{
		"name": mem.Name,
		"age":  mem.Age,
	}
	res.SetCode(entity.CODE_ERROR)
	res.SetData(data)
	c.JSON(http.StatusOK, res)
}
func EditMember(c *gin.Context) {

	mem := entity.Member{}

	if err := c.ShouldBind(&mem); err != nil {

		param := c.Query("trace-id")
		c.JSON(http.StatusForbidden, errno.ErrParam.WithID(param))
		c.Abort()
		return
	}

	// 处理业务(下次再分享)

	data := map[string]interface{}{
		"name": mem.Name,
		"age":  mem.Age,
	}

	c.JSON(http.StatusOK, errno.OK.WithData(data))

}
