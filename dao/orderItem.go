package dao

import (
	"exercise/bookstore/model"
	"exercise/bookstore/utils"
)

//向数据库中添加订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	//sql语句
	sqlStr := "insert into order_Items (count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderId)
	if err != nil {
		return err
	}
	return nil
}

//获取订单的所有订单项
func GetOrderItemsByOrderId(orderId string) ([]*model.OrderItem, error) {
	//sql语句
	sqlStr := "select * from order_items where order_id=?"
	//执行
	rows, err := utils.Db.Query(sqlStr, orderId)
	if err != nil {
		return nil, err
	}
	//创建订单项切片
	var orderItems []*model.OrderItem
	//遍历rows
	for rows.Next() {
		//创建订单项实例接收信息
		orderItem := &model.OrderItem{}
		rows.Scan(&orderItem.OrderItemId, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderId)
		//将订单项插入切片
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}
