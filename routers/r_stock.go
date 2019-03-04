package routers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	m "ymo.me/sbum-end-v2/models"
	u "ymo.me/sbum-end-v2/utils"
)

func createStock(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Stock

	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.InsertStock(requestData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "ထည့်သွင်းထားရှိပြီးပါပြီ။",
	})
}

func fetchAllStock(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	datas, err := u.FindAllStock()
	fmt.Println(" Stock Datas ", datas)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   datas,
	})
}

func deleteStock(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Stock
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	err := u.DeleteStock(requestData)
	if err != nil {
		fmt.Println(" error ", err.Error())
		context.JSON(http.StatusNotAcceptable, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   "ထိုအတိုင်းအတာယူနစ်အား ဖျက်သိမ်းပြီးပါပြီး။",
	})
}

func feacthTotalPrice(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestData m.Stock
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   requestData.Serial,
	})
}

func featchHistoryPrice(context *gin.Context) {
	//{month: 'Jan', uv: 4000, pv: 2400},
	context.Header("Content-Type", "application/json")
	var requestData m.Stock
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	responseData, err := u.FindAllStockBySerial(requestData.Serial)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   responseData.Price,
	})
}
