package car

import (
	"biligo/modules/common"
	"biligo/mysql"
	"biligo/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func SaveCarInfo(c *gin.Context) {
	car := Car{}
	err := c.BindJSON(&car)
	car.Uuid = util.UUID()
	car.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		util.FailResultWithCodeAndMessage(500, "车辆信息新增失败", nil).ToJSONWithHttpStatus(c)
		return
	}
	mysql.Conn.Create(&car)
	util.SuccessResult(car).ToJSON(c)
}

func CarInfoList(c *gin.Context) {
	pagination := common.NewPageFromGin(c)
	util.SuccessResult(queryNoteList(pagination)).ToJSON(c)
}

// NoteList - 分页查询 note
func queryNoteList(pagination *common.Pagination) *[]Car {
	carInfoList := []Car{}
	mysql.Conn.Order("create_time desc").Find(&carInfoList)
	for i := range carInfoList {
		parsedTime, err := time.Parse(time.RFC3339, carInfoList[i].CreateTime)
		if err != nil {
			fmt.Println("解析时间字符串出错:", err)
		}
		carInfoList[i].CreateTime = parsedTime.Format("2006-01-02 15:04:05")
	}
	return &carInfoList
}

func GetCarInfoById(c *gin.Context) {
	id := c.Param("id")
	carInfo := checkCarInfoById(id)
	if carInfo.Uuid == "" {
		util.FailResultWithMessage("当前实体不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		util.SuccessResultWithMessage("查询成功", carInfo).ToJSON(c)
	}
}

func UpdateCarInfoById(c *gin.Context) {
	// 查出来是否有该id
	carInfo := Car{}
	oldCarInfo := Car{}
	var id = c.Param("id")
	fmt.Println(id)
	mysql.Conn.Where("uuid = ?", id).Take(&oldCarInfo)
	if oldCarInfo.Uuid == "" {
		util.FailResultWithMessage("当前id不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		// 更新
		err := c.BindJSON(&carInfo)
		if err != nil {
			util.FailResultWithCodeAndMessage(500, "修改密码失败", nil).ToJSONWithHttpStatus(c)
			return
		}
		carInfo.Uuid = id
		carInfo.CreateTime = oldCarInfo.CreateTime
		//mysql.Conn.Where("uuid = ", id).Updates(&carMaintainInfo)
		if err := mysql.Conn.Model(&Car{}).Where("uuid", id).Updates(&carInfo).Error; err != nil {
			util.FailResultWithMessage("更新失败", err).ToJSONWithHttpStatus(c)
			return
		}
		util.SuccessResult(carInfo).ToJSON(c)
	}
}

func DeleteCarInfoById(c *gin.Context) {
	id := c.Param("id")
	maintainInfoById := checkCarInfoById(id)
	if maintainInfoById.Uuid == "" {
		util.FailResultWithMessage("当前实体不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		mysql.Conn.Model(&maintainInfoById).Where("uuid", id).Delete(&maintainInfoById)
		util.SuccessResultWithMessage("删除成功", maintainInfoById).ToJSON(c)
	}
}

func GetAllCarIdAndName(c *gin.Context) {
	cars := []Car{}
	mysql.Conn.Order("create_time desc").Find(&cars)
	util.SuccessResultWithMessage("查询成功", cars).ToJSON(c)
}

func checkCarInfoById(id string) *Car {
	car := &Car{}
	mysql.Conn.Where("uuid = ?", id).Take(car)
	return car
}

type Car struct {
	Uuid       string `json:"uuid"`
	Name       string `json:"name"`
	DailyKm    string `json:"dailyKm"`
	IsAlarm    bool   `json:"isAlarm"`
	Remark     string `json:"remark"`
	CreateTime string `json:"createTime"`
}
