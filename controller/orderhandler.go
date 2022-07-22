package controller

import (
	"exercise/bookstore/dao"
	"exercise/bookstore/model"
	"exercise/bookstore/utils"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

//结算订单
func Checkout(w http.ResponseWriter, r *http.Request) {
	//获取session
	_, session := dao.IsLogin(r)
	//获取userId
	userId := session.UserId
	//获取购物车信息
	cart, _ := dao.GetCartByUserId(strconv.Itoa(userId))
	//生成订单号
	orderId := utils.CreateUUID()
	//将购物车信息保存到订单
	order := &model.Order{
		OrderId:     orderId,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserId:      int64(userId),
	}
	//将订单保存到数据库
	dao.AddOrder(order)
	//获取购物车中的购物项
	cartItems := cart.CartItems
	//遍历购物项
	for _, v := range cartItems {
		//将购物项中的信息保存到订单项
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderId: orderId,
		}
		//将订单项保存到数据库
		dao.AddOrderItem(orderItem)
		//更新图书信息，库存，销量
		book := v.Book
		book.Sales += int(v.Count)
		book.Stock -= int(v.Count)
		dao.UpdateBook(book)
	}
	//清空购物车
	dao.DeleteCartByCartId(cart.CartId)
	//将订单号赋给session
	session.Order = order

	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	//执行
	t.Execute(w, session)
}

//获取所有订单
func GetOrders(w http.ResponseWriter, r *http.Request) {
	//获取订单
	orders, _ := dao.GetOrders()
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	//执行
	t.Execute(w, orders)
}

//获取订单详情
func GetOrderInfo(w http.ResponseWriter, r *http.Request) {
	//获取订单号
	orderId := r.FormValue("orderId")
	//获取对应的所有订单项
	orderItems, _ := dao.GetOrderItemsByOrderId(orderId)
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	//执行
	t.Execute(w, orderItems)
}

//获取我的订单
func GetMyOrder(w http.ResponseWriter, r *http.Request) {
	//获取session和userId
	_, session := dao.IsLogin(r)
	userId := session.UserId
	//获取对应的所有订单
	orders, _ := dao.GetMyOrder(userId)
	session.Orders = orders
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/order/order.html"))
	//执行
	t.Execute(w, session)
}

//发货
func SendOrder(w http.ResponseWriter, r *http.Request) {
	//获取订单号
	orderId := r.FormValue("orderId")
	//更新订单状态
	dao.UpdateOrderState(orderId, 1)
	//刷新订单页面
	GetOrders(w, r)
}

//收货
func TakeOrder(w http.ResponseWriter, r *http.Request) {
	//获取订单号
	orderId := r.FormValue("orderId")
	//更新订单状态
	dao.UpdateOrderState(orderId, 2)
	//刷新我的订单页面
	GetMyOrder(w, r)
}
