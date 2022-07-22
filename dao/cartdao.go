package dao

import (
	"exercise/bookstore/model"
	"exercise/bookstore/utils"
)

//向数据库中插入购物车
func AddCart(cart *model.Cart) error {
	//sql语句
	sqlStr := "insert into carts (id, total_count, total_amount, user_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.CartId, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserId)
	if err != nil {
		return err
	}
	//获取购物项
	cartItems := cart.CartItems
	for _, cartItem := range cartItems {
		//将购物项添加到数据库中
		AddCartItem(cartItem)
	}
	return nil
}

//根据userId获取购物车
func GetCartByUserId(userId string) (*model.Cart, error) {
	//sql语句
	sqlStr := "select * from carts where user_id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, userId)
	//创建cart结构体
	cart := &model.Cart{}
	//扫描row得到cart的字段的值
	err := row.Scan(&cart.CartId, &cart.TotalCount, &cart.TotalAmount, &cart.UserId)
	if err != nil {
		return nil, err
	}
	//获取cart的cartItems
	cartItems, _ := GetCartItemsByCartId(cart.CartId)
	cart.CartItems = cartItems
	return cart, nil
}

//更新购物车的总数、总金额
func UpdateCart(cart *model.Cart) error {
	//sql语句
	sqlStr := "update carts set total_count=?, total_amount=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartId)
	if err != nil {
		return err
	}
	return nil
}

//删除购物车
func DeleteCartByCartId(cartId string) error {
	//先删除购物车中的所有购物项
	DeleteCartItemsByCartId(cartId)
	//sql语句
	sqlStr := "delete from carts where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartId)
	if err != nil {
		return err
	}
	return nil
}
