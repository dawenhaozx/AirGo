package admin_api

import (
	"github.com/gin-gonic/gin"
	"github.com/ppoonk/AirGo/constant"
	"github.com/ppoonk/AirGo/global"
	"github.com/ppoonk/AirGo/model"
	"github.com/ppoonk/AirGo/service/common_logic"
	"github.com/ppoonk/AirGo/utils/response"
)

// 获取全部订单，分页获取
func GetOrderList(ctx *gin.Context) {
	var params model.QueryParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	res, total, err := common_logic.CommonSqlFindWithFieldParams(&params)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail("GetOrderList error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("GetOrderList success", model.CommonDataResp{
		Total: total,
		Data:  res,
	}, ctx)

}

// 获取订单统计
func GetOrderStatistics(ctx *gin.Context) {
	var params model.QueryParams
	err := ctx.ShouldBind(&params)
	res, err := orderService.GetMonthOrderStatistics(&params)
	if err != nil {
		global.Logrus.Error(err.Error())
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	response.OK("GetOrderStatistics success", res, ctx)
}

// 更新用户订单
func UpdateOrder(ctx *gin.Context) {
	var order model.Order
	err := ctx.ShouldBind(&order)
	if err != nil {
		global.Logrus.Error(err)
		response.Fail(constant.ERROR_REQUEST_PARAMETER_PARSING_ERROR+err.Error(), nil, ctx)
		return
	}
	err = orderService.UpdateOrder(&order) //更新数据库状态
	if err != nil {
		global.Logrus.Error(err)
		response.Fail("UpdateOrder error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("UpdateOrder success", nil, ctx)
}
