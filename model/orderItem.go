package model

//订单项结构体
type OrderItem struct {
	OrderItemId int64   //订单项id
	Count       int64   //图书数量
	Amount      float64 //金额小计
	Title       string  //图书名
	Author      string  //作者
	Price       float64 //单价
	ImgPath     string  //封面
	OrderId     string  //所属订单id
}
