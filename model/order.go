package model

//订单结构体
type Order struct {
	OrderId     string  //订单id
	CreateTime  string  //生成订单的时间
	TotalCount  int64   //订单中图书的总数
	TotalAmount float64 //订单总价
	State       int64   //订单状态 0:未发货、1:已发货、2:交易完成
	UserId      int64   //订单所属用户
}

//未发货
func (o *Order) NoSend() bool {
	return o.State == 0
}

//已发货
func (o *Order) IsSend() bool {
	return o.State == 1
}

//已收货
func (o *Order) Complete() bool {
	return o.State == 2
}
