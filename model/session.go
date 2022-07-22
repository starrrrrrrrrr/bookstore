package model

//session 结构体
type Session struct {
	SessionId string
	UserName  string
	UserId    int
	Cart      *Cart
	Order     *Order
	Orders    []*Order
}
