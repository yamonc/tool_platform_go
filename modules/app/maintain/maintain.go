package maintain

import (
	"biligo/modules/app/car"
	"biligo/modules/common"
	"biligo/mysql"
	"fmt"

	"biligo/util"
	"github.com/gin-gonic/gin"
	"time"
)

func SaveCarMaintainInfo(c *gin.Context) {
	car := CarMaintain{}
	err := c.BindJSON(&car)
	car.Uuid = util.UUID()
	car.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		util.FailResultWithCodeAndMessage(500, "保养记录添加失败", nil).ToJSONWithHttpStatus(c)
		return
	}
	// 注意：这里表中并没有car_name这个字段，这里只是为了忽略CarMaintain中的carName字段
	mysql.Conn.Omit("car_name").Create(&car)
	util.SuccessResult(car).ToJSON(c)
}

func CarMaintainInfoList(c *gin.Context) {
	pagination := common.NewPageFromGin(c)
	util.SuccessResult(queryNoteList(pagination)).ToJSON(c)
}
func GetMaintainInfoById(c *gin.Context) {
	id := c.Param("id")
	carMaintain := checkCarMaintainInfoById(id)
	car := car.Car{}
	// 查carInfo表
	mysql.Conn.Where("uuid = ?", carMaintain.CarId).Take(&car)
	if carMaintain.Uuid == "" {
		util.FailResultWithMessage("当前实体不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		carMaintain.CarName = car.Name
		util.SuccessResultWithMessage("查询成功", carMaintain).ToJSON(c)
	}
}

func DeleteMaintainInfoById(c *gin.Context) {
	id := c.Param("id")
	maintainInfoById := checkCarMaintainInfoById(id)
	if maintainInfoById.Uuid == "" {
		util.FailResultWithMessage("当前实体不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		mysql.Conn.Model(&maintainInfoById).Where("uuid", id).Delete(&maintainInfoById)
		util.SuccessResultWithMessage("删除成功", maintainInfoById).ToJSON(c)
	}
}

func checkCarMaintainInfoById(id string) *CarMaintain {
	carMaintainInfo := &CarMaintain{}
	mysql.Conn.Where("uuid = ?", id).Take(carMaintainInfo)
	return carMaintainInfo
}

// NoteList - 分页查询 note
func queryNoteList(pagination *common.Pagination) *[]CarMaintain {
	carMaintains := []CarMaintain{}
	mysql.Conn.Order("create_time desc").Find(&carMaintains)
	for i := range carMaintains {
		// 根据uuid获取car信息
		car := car.Car{}
		mysql.Conn.Where("uuid = ?", carMaintains[i].CarId).Take(&car)
		carMaintains[i].CarName = car.Name
		parsedTime, err := time.Parse(time.RFC3339, carMaintains[i].CreateTime)
		lastTimeParsed, err := time.Parse(time.RFC3339, carMaintains[i].LastTime)
		if err != nil {
			fmt.Println("解析时间字符串出错:", err)
		}
		carMaintains[i].CreateTime = parsedTime.Format("2006-01-02 15:04:05")
		carMaintains[i].LastTime = lastTimeParsed.Format("2006-01-02 15:04:05")
	}
	return &carMaintains
}

type CarMaintain struct {
	Uuid       string `json:"uuid"`
	CarId      string `json:"carId"`
	CarName    string `json:"carName"`
	LastKm     string `json:"lastKm"`
	LastTime   string `json:"lastTime"`
	NowKm      string `json:"nowKm"`
	Remark     string `json:"remark"`
	CreateTime string `json:"createTime"`
}
