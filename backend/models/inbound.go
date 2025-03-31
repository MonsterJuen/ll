package models

// InboundRequest 入库请求结构体
type InboundRequest struct {
	ProductName string  `json:"productName" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required,min=1"`
	Unit        string  `json:"unit" binding:"required"`
	Temperature float64 `json:"temperature" binding:"required"`
	Location    string  `json:"location" binding:"required"`
}

// InboundResponse 入库响应结构体
type InboundResponse struct {
	Message string         `json:"message"`
	Data    InboundRequest `json:"data"`
}
