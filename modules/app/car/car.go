package car

import (
	"biligo/modules/app/model"
	_ "biligo/modules/app/model"
	"biligo/modules/common"
	"biligo/mysql"
	"biligo/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func SaveCarInfo(c *gin.Context) {
	car := model.Car{}
	err := c.ShouldBindJSON(&car)
	car.Uuid = util.UUID()
	// TODO:开启追踪功能
	if err != nil {
		fmt.Println(err)
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
func queryNoteList(pagination *common.Pagination) *[]model.APICar {
	carInfoList := []model.Car{}
	mysql.Conn.Order("created_at desc").Find(&carInfoList)
	result := make([]model.APICar, 0, 10)
	for i := range carInfoList {
		//lastMaintainInfo := getLastTimeMaintainInfo(carInfoList[i].ID)
		//statusMap := getCarStatus(lastMaintainInfo)
		//var builder strings.Builder
		//for k, v := range statusMap {
		//	builder.WriteString(k)
		//	builder.WriteString(":")
		//	builder.WriteString(strconv.FormatInt(int64(v), 10))
		//	builder.WriteString(";")
		//}
		//carInfoList[i].Status = builder.String()
		// 复制其他字段并将返回值分配给 result
		result = append(result, model.APICar{
			Uuid:       carInfoList[i].Uuid,
			Name:       carInfoList[i].Name,
			DailyKm:    carInfoList[i].DailyKm,
			IsAlarm:    carInfoList[i].IsAlarm,
			NowKm:      carInfoList[i].NowKm,
			CreateTime: carInfoList[i].CreatedAt.Format("2006-01-02T15:04:05"),
			Remark:     carInfoList[i].Remark,
		})

		//result[i] = model.APICar{}

		//parsedTime, err := time.Parse(time.RFC3339, carInfoList[i].CreatedAt)
		//if err != nil {
		//	fmt.Println("解析时间字符串出错:", err)
		//}
		//carInfoList[i].CreatedAt = carInfoList[i].CreatedAt.Format("2006-01-02 15:04:05")
		//carInfoList[i].CreatedAt = parsedTime.Format("2006-01-02 15:04:05")
	}
	return &result
}

func TestCar(c *gin.Context) {
	carMaintain := model.CarMaintain{
		"1", "1", "1", "25600", "2023-09-07 00:00:00", "30000", "5000km一保养", "2023-10-11 15:56:46",
	}
	getCarStatus(carMaintain)
}

func getCarStatus(carMaintainInfo model.CarMaintain) map[string]int {
	record := []model.CarRecord{}
	mysql.Conn.Find(&record)
	result := make(map[string]int)
	for i := range record {
		var temp = record[i].Km
		if temp != "" {
			// 开始计算
			km, _ := strconv.ParseInt(record[i].Km, 10, 64)
			fmt.Println(km)
			lastKm, _ := strconv.ParseInt(carMaintainInfo.LastKm, 10, 64)
			nowKm, _ := strconv.ParseInt(carMaintainInfo.NowKm, 10, 64)
			// 固定计算
			var expectKm = lastKm + km
			var retainKm = expectKm - nowKm
			if retainKm < 1000 {
				result[record[i].Item] = int(expectKm)
			}
		}
	}
	fmt.Println(result)
	return result
}

func getLastTimeMaintainInfo(id string) model.CarMaintain {
	carMaintain := model.CarMaintain{}
	//mysql.Conn.Where("uuid = ?", id).Order("last_time desc").First(&carMaintain)
	mysql.Conn.Order("last_time desc").First(&carMaintain)
	return carMaintain
}

func GetCarInfoById(c *gin.Context) {
	id := c.Param("id")
	carInfo := checkCarInfoById(id)
	if carInfo.Uuid == "" {
		util.FailResultWithMessage("当前实体不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		util.SuccessResultWithMessage("查询成功", model.APICar{
			Uuid:       carInfo.Uuid,
			Name:       carInfo.Name,
			DailyKm:    carInfo.DailyKm,
			IsAlarm:    carInfo.IsAlarm,
			NowKm:      carInfo.NowKm,
			CreateTime: carInfo.CreatedAt.Format("2006-01-02T15:04:05"),
			Remark:     carInfo.Remark,
		}).ToJSON(c)
	}
}

func UpdateCarInfoById(c *gin.Context) {
	// 查出来是否有该id
	carInfo := model.Car{}
	oldCarInfo := model.Car{}
	var id = c.Param("id")
	fmt.Println(id)
	mysql.Conn.Where("uuid = ?", id).Take(&oldCarInfo)
	if oldCarInfo.Uuid == "" {
		util.FailResultWithMessage("当前id不存在", nil).ToJSONWithHttpStatus(c)
	} else {
		// 更新
		err := c.BindJSON(&carInfo)
		if err != nil {
			util.FailResultWithCodeAndMessage(500, "更新失败", err).ToJSONWithHttpStatus(c)
			return
		}
		if err := mysql.Conn.Model(&model.Car{}).Where("uuid", id).Updates(map[string]interface{}{
			"name":     carInfo.Name,
			"daily_km": carInfo.DailyKm,
			"now_km":   carInfo.NowKm,
			"is_alarm": carInfo.IsAlarm,
			"remark":   carInfo.Remark,
		}).Error; err != nil {
			util.FailResultWithCodeAndMessage(500, "更新失败", err).ToJSONWithHttpStatus(c)
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
	cars := []model.Car{}
	mysql.Conn.Order("create_time desc").Find(&cars)
	util.SuccessResultWithMessage("查询成功", cars).ToJSON(c)
}

func checkCarInfoById(id string) *model.Car {
	car := &model.Car{}
	mysql.Conn.Where("uuid = ?", id).Take(car)
	return car
}
