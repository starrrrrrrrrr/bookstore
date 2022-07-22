package dao

import (
	"exercise/bookstore/model"
	"fmt"
	"testing"
	"time"
)

// func TestUser(t *testing.T) {
// 	fmt.Println("测试userdao中的函数")
// }

// func TestSaveUser(t *testing.T) {
// 	SaveUser("admin", "123456", "admin@163.com")
// }

// func TestCheckUsernameAndPassword(t *testing.T) {
// 	user, _ := CheckUsernameAndPassword("admin", "123456")
// 	fmt.Println("获取的用户内容:", user)
// }

// func TestCheckUsername(t *testing.T) {
// 	user, _ := CheckUsername("admin")
// 	fmt.Println("获取的用户内容2:", user)
// }

// func TestGetBooks(t *testing.T) {
// 	fmt.Println("测试GetBooks函数")

// 	books, _ := GetBooks()

// 	for k, v := range books {
// 		fmt.Printf("第%v本图书的信息是%v\n", k+1, v)
// 	}
// }

// func TestAddBook(t *testing.T) {
// 	fmt.Println("测试AddBook函数")

// 	book := &model.Book{
// 		Title:   "三国演义",
// 		Author:  "罗贯中",
// 		Price:   88.88,
// 		Sales:   100,
// 		Stock:   100,
// 		ImgPath: "static/img/default.jpg",
// 	}

// 	AddBook(book)
// }
// func TestDeleteBook(t *testing.T) {
// 	fmt.Println("测试DeleteBook函数")
// 	DeleteBook(33)
// }

// func TestGetBookById(t *testing.T) {
// 	fmt.Println("测试GetBookById函数")
// 	book, _ := GetBookById(31)
// 	fmt.Println(book)
// }

// func TestUpdateBook(t *testing.T) {
// 	fmt.Println("测试UpdateBook函数")
// 	b := &model.Book{
// 		Id:      31,
// 		Title:   "三国演义",
// 		Author:  "罗贯中",
// 		Price:   99.99,
// 		Sales:   200,
// 		Stock:   200,
// 		ImgPath: "static/img/default.jpg",
// 	}
// 	UpdateBook(b)
// }

// func TestGetPage(t *testing.T) {
// 	fmt.Println("测试GetPage函数:")
// 	page, _ := GetPage("1")
// 	fmt.Println(page)
// 	for i, v := range page.Books {
// 		fmt.Println(i, v)
// 	}
// }

// func TestGetPageBooksByPrice(t *testing.T) {
// 	fmt.Println("测试GetPageBooksByPrice函数:")
// 	page, _ := GetPageBooksByPrice("1", "10", "30")
// 	fmt.Println(page)
// 	for i, v := range page.Books {
// 		fmt.Println(i, v)
// 	}
// }

// func testAddSession(t *testing.T) {
// 	fmt.Println("开始测试AddSession函数")
// 	sess := &model.Session{
// 		SessionId: 111,
// 		UserName:  "admin",
// 		UserId:    1,
// 	}
// 	AddSession(sess)
// }

// func testDelSession(t *testing.T) {
// 	fmt.Println("开始测试DelSession函数")
// 	DelSession("111")
// }

// func testGetSession(t *testing.T) {
// 	fmt.Println("开始测试GetSession函数")
// 	session, _ := GetSession("dc7fca13-0688-4ee5-43f4-164ce326eb60")
// 	fmt.Println("获取的session:", session)
// }

// func TestAddCart(t *testing.T) {
// 	fmt.Println("测试AddCart函数")
// 	//book实例
// 	book1 := &model.Book{
// 		Id:    1,
// 		Price: 27.20,
// 	}
// 	book2 := &model.Book{
// 		Id:    2,
// 		Price: 23.00,
// 	}
// 	var cartItems []*model.CartItem
// 	//cartItem实例
// 	cartItem1 := &model.CartItem{
// 		Book:   book1,
// 		Count:  10,
// 		CartId: "666666",
// 	}
// 	cartItems = append(cartItems, cartItem1)
// 	cartItem2 := &model.CartItem{
// 		Book:   book2,
// 		Count:  10,
// 		CartId: "666666",
// 	}
// 	cartItems = append(cartItems, cartItem2)
// 	//cart实例
// 	cart := &model.Cart{
// 		CartId:    "666666",
// 		CartItems: cartItems,
// 		UserId:    1,
// 	}
// 	AddCart(cart)
// }

// func TestGetCartItemByBookId(t *testing.T) {
// 	cartItem, _ := GetCartItemByBookId("1")

// 	fmt.Println("bookId=1的购物项是", cartItem)
// }

// func TestGetCartItemsByCartId(t *testing.T) {
// 	cartItems, _ := GetCartItemsByCartId("666666")
// 	fmt.Println("cartId=666666的购物项有:")
// 	for _, v := range cartItems {
// 		fmt.Println(v)
// 	}
// }

// func TestGetCartByUserId(t *testing.T) {
// 	cart, _ := GetCartByUserId("1")

// 	fmt.Println("UserId=1的购物车信息有:")
// 	fmt.Printf("购物车Id:%v, 总数量:%v, 总价:%v\n", cart.CartId, cart.TotalCount, cart.TotalAmount)
// 	for k, v := range cart.CartItems {
// 		fmt.Printf("第:%v个购物项:%v\n", k, v)
// 	}
// }

// func TestDeleteCartByCartId(t *testing.T) {
// 	fmt.Println("测试删除购物车")
// 	DeleteCartByCartId("9849abef-c528-4948-41ee-d50ff3ef6dbc")
// }

// func TestDeleteCartItemById(t *testing.T) {
// 	fmt.Println("测试删除购物项")
// 	DeleteCartItemById("27")
// }

func TestAddOrder(t *testing.T) {
	fmt.Println("测试添加订单")
	//创建订单
	order := &model.Order{
		OrderId:     "11111",
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  2,
		TotalAmount: 300,
		State:       0,
		UserId:      1,
	}
	//创建订单项
	orderItem := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   100,
		ImgPath: "/default.jpg",
		OrderId: "11111",
	}

	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  200,
		Title:   "西游记",
		Author:  "吴承恩",
		Price:   200,
		ImgPath: "/default.jpg",
		OrderId: "11111",
	}
	//添加到数据库
	AddOrder(order)
	AddOrderItem(orderItem)
	AddOrderItem(orderItem2)
}

// func TestGetOrders(t *testing.T) {
// 	fmt.Println("测试获取所有订单")
// 	orders, _ := GetOrders()

// 	for _, v := range orders {
// 		fmt.Println("获取的订单:", v)
// 	}
// }

// func TestGetOrderItemsByOrderId(t *testing.T) {
// 	fmt.Println("测试获取订单的所有订单项")
// 	orderItems, _ := GetOrderItemsByOrderId("a216034b-84d1-4e5d-6de6-8b5c4cba3d59")
// 	for _, v := range orderItems {
// 		fmt.Println("订单信息:", v)
// 	}
// }

// func TestGetOrders(t *testing.T) {
// 	fmt.Println("测试获取我的订单")
// 	orders, _ := GetMyOrder(1)

// 	for _, v := range orders {
// 		fmt.Println("获取的订单:", v)
// 	}
// }

// func TestUpdateOrderState(t *testing.T) {
// 	fmt.Println("测试更新订单状态")
// 	UpdateOrderState("4c6099fb-e9b1-46c5-5adc-48212481ce92", 1)
// }
