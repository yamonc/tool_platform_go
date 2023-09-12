package password

import (
	"biligo/modules/common"
	"biligo/mysql"
	"biligo/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func PasswordList(c *gin.Context) {
	pagination := common.NewPageFromGin(c)
	util.SuccessResult(queryNoteList(pagination)).ToJSON(c)
}

func SavePassword(c *gin.Context) {
	password := Password{
		Uuid:       util.UUID(),
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := c.BindJSON(&password)
	if err != nil {
		util.FailResultWithCodeAndMessage(500, "密码新增失败", nil).ToJSONWithHttpStatus(c)
		return
	}
	mysql.Conn.Create(&password)
	util.SuccessResult(password).ToJSON(c)
}
func UpdatePassword(c *gin.Context) {
	// 查出来是否有该id
	password := Password{}
	oldPassword := Password{}
	var id = c.Param("id")
	fmt.Println(id)
	mysql.Conn.Where("uuid = ?", id).Take(&oldPassword)
	if oldPassword.Uuid == "" {
		util.FailResultWithMessage("当前id不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		// 更新
		err := c.BindJSON(&password)
		if err != nil {
			util.FailResultWithCodeAndMessage(500, "修改密码失败", nil).ToJSONWithHttpStatus(c)
			return
		}
		password.Uuid = id
		password.CreateTime = oldPassword.CreateTime
		//mysql.Conn.Where("uuid = ", id).Updates(&password)
		mysql.Conn.Model(&Password{}).Where("uuid", id).Updates(&password)
		util.SuccessResult(password).ToJSON(c)
	}

}

func DeletePassword(c *gin.Context) {
	id := c.Param("id")
	password := checkPasswordById(id)
	if password.Uuid == "" {
		util.FailResultWithMessage("当前实体不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		mysql.Conn.Model(&password).Where("uuid", id).Delete(&password)
		util.SuccessResultWithMessage("删除成功", password).ToJSON(c)
	}
}
func GetPasswordById(c *gin.Context) {
	id := c.Param("id")
	password := checkPasswordById(id)
	if password.Uuid == "" {
		util.FailResultWithMessage("当前实体不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		util.SuccessResultWithMessage("查询成功", password).ToJSON(c)
	}
}
func checkPasswordById(id string) *Password {
	password := &Password{}
	mysql.Conn.Where("uuid = ?", id).Take(password)
	return password
}

// NoteList - 分页查询 note
func queryNoteList(pagination *common.Pagination) *[]Password {
	passwords := []Password{}
	mysql.Conn.Order("create_time desc").Find(&passwords)
	for i := range passwords {
		parsedTime, err := time.Parse(time.RFC3339, passwords[i].CreateTime)
		if err != nil {
			fmt.Println("解析时间字符串出错:", err)
		}
		passwords[i].CreateTime = parsedTime.Format("2006-01-02 15:04:05")
	}
	return &passwords
}

// MODEL

type Password struct {
	Uuid         string `json:"uuid"`
	PlatformName string `json:"platformName"`
	Url          string `json:"url"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	CreateTime   string `json:"createTime"`
	Remark       string `json:"remark"`
}
