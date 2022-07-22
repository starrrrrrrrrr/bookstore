package model

//购物车结构体
type Cart struct {
	CartId      string      //购物车id
	CartItems   []*CartItem //购物车中的购物项
	TotalCount  int64       //图书总数
	TotalAmount float64     //金额总数
	UserId      int         //所属用户的id
	UserName    string      //用户名
}

//计算图书总数
func (cart *Cart) GetTotalCount() int64 {
	var totalCount int64

	for _, v := range cart.CartItems {
		totalCount += v.Count
	}

	return totalCount
}

//计算总金额
func (cart *Cart) GetTotalAmount() float64 {
	var totalAmount float64

	for _, v := range cart.CartItems {
		totalAmount += v.GetAmount()
	}

	return totalAmount
}
