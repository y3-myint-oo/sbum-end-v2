package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	m "ymo.me/sbum-end-v2/models"
	u "ymo.me/sbum-end-v2/utils"
)

func fetchAllSupply(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	supplies, err := u.FindAllSupplyers()
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   supplies,
	})
}

func createSupply(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestSupply m.Supply
	if err := context.ShouldBindJSON(&requestSupply); err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	//fmt.Println(" request Supply ---- ", requestSupply)
	err := u.InsertSupplyer(requestSupply)
	if err != nil {
		context.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "တောင်သူအသစ်အား စာရင်း ထည့်သွင်းထားရှိပြီးပါပြီ။",
	})
}

func deleteSupply(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	var requestSupply m.Supply
	if err := context.ShouldBindJSON(&requestSupply); err != nil {
		context.JSON(http.StatusNotAcceptable, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "လုပ်ဆောင်ချက် မအောင်မြင်ပါ",
		})
		return
	}
	//fmt.Println(" request Supply ID ", requestUnit.ID)
	err := u.DeleteSupplyer(requestSupply)
	if err != nil {
		//fmt.Println(" error ", err.Error())
		context.JSON(http.StatusNotAcceptable, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "တောင်သူအားစားရင်းမှ ဖျက်သိမ်းပြီးပါပြီး။",
	})
}
