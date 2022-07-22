package dao

import (
	"exercise/bookstore/model"
	"exercise/bookstore/utils"
)

//向数据库中添加订单
func AddOrder(order *model.Order) error {
	//sql语句
	sqlStr := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, order.OrderId, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserId)
	if err != nil {
		return err
	}
	return nil
}

//获取所有订单
func GetOrders() ([]*model.Order, error) {
	//sql语句
	sqlStr := "select * from orders"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建Orders切片
	var orders []*model.Order
	//扫描rows得到订单
	for rows.Next() {
		//创建order实例接收信息
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		//将order插入orders
		orders = append(orders, order)
	}
	return orders, nil
}

//获取我的订单
func GetMyOrder(userId int) ([]*model.Order, error) {
	//sql语句
	sqlStr := "select * from orders where user_id=?"
	//执行
	rows, err := utils.Db.Query(sqlStr, userId)
	if err != nil {
		return nil, err
	}
	//创建Orders切片
	var orders []*model.Order
	//扫描rows得到订单
	for rows.Next() {
		//创建order实例接收信息
		order := &model.Order{}
		rows.Scan(&order.OrderId, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.State, &order.UserId)
		//将order插入orders
		orders = append(orders, order)
	}
	return orders, nil
}

//更新订单状态
func UpdateOrderState(orderId string, state int64) error {
	//sql语句
	sqlStr := "update orders set state=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, state, orderId)
	if err != nil {
		return err
	}
	return nil
}
