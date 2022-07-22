package main

import (
	"exercise/bookstore/controller"
	"net/http"
)

func main() {
	//处理静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))
	//主页
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//登录
	http.HandleFunc("/login", controller.Login)
	//注销
	http.HandleFunc("/logout", controller.Logout)
	//注册
	http.HandleFunc("/regist", controller.Regist)
	//验证用户名
	http.HandleFunc("/checkUsername", controller.CheckUserName)
	//获取所有图书
	//http.HandleFunc("/getBooks", controller.GetBooks)
	//添加图书
	//http.HandleFunc("/addBook", controller.UpdateOrAddBook)
	//获取图书分页
	http.HandleFunc("/getPage", controller.GetPageBooks)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//前往修改图书页面
	http.HandleFunc("/toUpdateBook", controller.ToUpdateBookPage)
	//修改或添加图书
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBook)
	//查询价格范围内的图书
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	//添加图书到购物车
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	//删除购物项
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	//更新购物项
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem)
	//结算订单
	http.HandleFunc("/checkout", controller.Checkout)
	//获取所有订单
	http.HandleFunc("/getOrders", controller.GetOrders)
	//获取订单详情
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo)
	//获取我的订单
	http.HandleFunc("/getMyOrder", controller.GetMyOrder)
	//发货
	http.HandleFunc("/sendOrder", controller.SendOrder)
	//收货
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	http.ListenAndServe(":8080", nil)
}
