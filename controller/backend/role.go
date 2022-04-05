package backend

import (
	"electricity/driver"
	"electricity/model"
	"electricity/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type RoleController struct {
	BaseController
}

func (that RoleController) Role(c *gin.Context)  {
	var roles []model.Role
	err := driver.DB.Find(&roles).Error
	if err != nil {
		that.Error(c,"查询失败","/admin/role")
		return
	}
	c.HTML(http.StatusOK,"backend/role/index.html",roles)
}

func (that RoleController) Add(c *gin.Context)  {
	c.HTML(http.StatusOK,"backend/role/add.html",nil)
}

func (that RoleController) DoAdd(c *gin.Context)  {
	title := strings.Trim(c.PostForm("title"),"")
	description := strings.Trim(c.PostForm("description"),"")
	if title == "" {
		that.Error(c,"角色名称不能为空","/admin/role/add")
		return
	}

	var role model.Role
	role.Title = title
	role.Description = description
	role.Status = 1
	role.AddTime = utils.GetUnix()
	err := driver.DB.Create(&role).Error
	if err != nil {
		that.Error(c,"创建用户失败","/admin/role/add")
		return
	}

	that.Success(c,"创建用户成功","/admin/role")
}

