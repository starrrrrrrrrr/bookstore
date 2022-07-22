package model

//购物项结构体
type CartItem struct {
	CartItemId int64   //购物项id
	Book       *Book   //图书信息
	Count      int64   //图书数量
	Amount     float64 //金额小计
	CartId     string  //所属的购物车的id
}

//计算金额小计
func (c *CartItem) GetAmount() float64 {
	price := c.Book.Price
	return float64(c.Count) * price
}
