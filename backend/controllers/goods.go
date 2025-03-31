package controllers

import (
	"fmt"
	"net/http"
	"time"

	"coldchain/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// 模拟数据存储
var goodsList = make([]models.Goods, 0)

// 创建中文时间格式模板
const chineseTimeLayout = "2006年1月2日15时04分"

// GetGoodsList 获取货物列表
func GetGoodsList(c *gin.Context) {
	// var filter models.GoodsFilter
	// if err := c.ShouldBindQuery(&filter); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // 过滤数据
	// filteredList := make([]models.Goods, 0)
	// for _, goods := range goodsList {
	// 	if filter.GoodsID != "" && goods.ID != filter.GoodsID {
	// 		continue
	// 	}
	// 	if filter.Type != "" && goods.Type != filter.Type {
	// 		continue
	// 	}
	// 	if filter.Status != "" && goods.Status != filter.Status {
	// 		continue
	// 	}
	// 	filteredList = append(filteredList, goods)
	// }

	// c.JSON(http.StatusOK, models.GoodsListResponse{
	// 	Total: int64(len(filteredList)),
	// 	Items: filteredList,
	// })
}

// CreateGoods 创建货物
func CreateGoods(c *gin.Context) {
	var goods models.Goods
	if err := c.ShouldBindJSON(&goods); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 生成批次编号
	goods.ID = "SF" + time.Now().Format("20060102") + uuid.New().String()[:3]
	goods.InTime = time.Now().Format(chineseTimeLayout)
	goods.Status = "in_stock"

	fmt.Println(goods)

	c.JSON(http.StatusOK, goods)
}

// UpdateGoods 更新货物信息
func UpdateGoods(c *gin.Context) {
	id := c.Param("id")
	var updatedGoods models.Goods
	if err := c.ShouldBindJSON(&updatedGoods); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, goods := range goodsList {
		if goods.ID == id {
			updatedGoods.ID = id               // 保持ID不变
			updatedGoods.InTime = goods.InTime // 保持入库时间不变
			goodsList[i] = updatedGoods
			c.JSON(http.StatusOK, updatedGoods)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "goods not found"})
}

// DeleteGoods 删除货物
func DeleteGoods(c *gin.Context) {
	id := c.Param("id")

	for i, goods := range goodsList {
		if goods.ID == id {
			goodsList = append(goodsList[:i], goodsList[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "goods deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "goods not found"})
}

// OutGoods 货物出库
func OutGoods(c *gin.Context) {
	// id := c.Param("id")

	// for i, goods := range goodsList {
	// 	if goods.ID == id {
	// 		if goods.Status != models.StatusInStock {
	// 			c.JSON(http.StatusBadRequest, gin.H{"error": "goods is not in stock"})
	// 			return
	// 		}
	// 		goodsList[i].Status = models.StatusOutStock
	// 		c.JSON(http.StatusOK, goodsList[i])
	// 		return
	// 	}
	// }

	// c.JSON(http.StatusNotFound, gin.H{"error": "goods not found"})
}
