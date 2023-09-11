package contreller

import (
	"github.com/gin-gonic/gin"
	"lss/dao"
	"lss/model"
	"strconv"
)

var User = &user{}

type user struct{}

// Info	godoc
// @Summary		用户详情
// @Description	用户详情
// @Tags	用户
// @Accept	application/json
// @Produce json
// @Success 200 {string} string	"ok"
// @Param id query int true "id"
// @Router	/user/info [get]
func (user) Info(c *gin.Context) model.Res {
	idStr := c.Query("id")
	if idStr == "" {
		return model.RError(-1002, "参数不能为空")
	}

	id, _ := strconv.Atoi(idStr)

	res, err := dao.User.Info(id)
	if err != nil {
		return model.RError(-1002, "获取详情失败")
	}

	return model.RSuccess(res)

}
