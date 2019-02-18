package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	m "ymo.me/sbum-end-v2/models"
	u "ymo.me/sbum-end-v2/utils"
)

func createData(context *gin.Context) {
	context.Header("Content-Type", "application/json")

	var requestData m.Data
	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}

	err := u.InsertData(requestData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "အတိုင်းအတာယူနစ်သစ်အား ထည့်သွင်းထားရှိပြီးပါပြီ။",
	})
}

func fetchAllData(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	datas, err := u.FindAllData()
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
