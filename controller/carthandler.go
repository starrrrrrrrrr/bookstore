package controller

import (
	"encoding/json"
	"exercise/bookstore/dao"
	"exercise/bookstore/model"
	"exercise/bookstore/utils"
	"html/template"
	"net/http"
	"strconv"
)

//添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session := dao.IsLogin(r)
	if flag {
		//获取bookId
		bookId := r.FormValue("bookId")
		//获取book
		book, _ := dao.GetBookById(bookId)

		//获取userId
		userId := session.UserId
		//判断数据库中是否已有该用户的购物车
		cart, _ := dao.GetCartByUserId(strconv.Itoa(userId))
		if cart != nil {
			//该用户已有购物车
			//判断购物车中是否已添加该图书
			cartItem, _ := dao.GetCartItemByBookIdAndCartId(bookId, cart.CartId)
			if cartItem != nil {
				//购物车中已有该图书
				//得到购物车的所有购物项
				cartItems := cart.CartItems
				//遍历切片找到对应的购物项
				for _, v := range cartItems {
					if v.Book.Id == cartItem.Book.Id {
						//将数量增加
						v.Count++
						//更新数据库中的购物项
						dao.UpdateBookCount(v)
					}
				}

			} else {
				//购物车中还没有此图书
				//创建购物项
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartId: cart.CartId,
				}
				//将购物项添加到购物车中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将购物项添加到数据库中
				dao.AddCartItem(cartItem)
			}
			//无论是否已有该图书,都要更新购物车
			dao.UpdateCart(cart)
		} else {
			//该用户没有购物车
			//创建购物车
			cartId := utils.CreateUUID()
			cart := &model.Cart{
				CartId: cartId,
				UserId: userId,
			}
			//创建购物项
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartId: cartId,
			}
			cartItems = append(cartItems, cartItem)

			cart.CartItems = cartItems
			//将购物车添加到数据库中
			dao.AddCart(cart)
		}

		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车"))
	} else {
		w.Write([]byte("请先登录..."))
	}

}

//获取购物车信息
func GetCartInfo(w http.ResponseWriter, r *http.Request) {
	//获取session和userid
	_, session := dao.IsLogin(r)
	userId := session.UserId
	//从数据库中获取购物车
	cart, _ := dao.GetCartByUserId(strconv.Itoa(userId))

	if cart != nil {
		cart.UserName = session.UserName
		session.Cart = cart
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	} else {
		//购物车为空
		//解析模板文件
		t := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		//执行
		t.Execute(w, session)
	}
}

//清空购物车
func DeleteCart(w http.ResponseWriter, r *http.Request) {
	//获取cartId
	cartId := r.FormValue("cartId")
	//从数据库中删除购物车
	dao.DeleteCartByCartId(cartId)

	//刷新购物车页面
	GetCartInfo(w, r)
}

//删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	//获取cartItemId
	cartItemId := r.FormValue("cartItemId")
	IcartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	//获取session
	_, session := dao.IsLogin(r)
	//获取userId
	userId := session.UserId
	//获取cart
	cart, _ := dao.GetCartByUserId(strconv.Itoa(userId))
	//获取cartItems切片
	cartItems := cart.CartItems
	//找到要删除的购物项
	for k, v := range cartItems {
		if v.CartItemId == IcartItemId {
			cartItems = append(cartItems[:k], cartItems[k+1:]...)
			//将删除后的切片重新赋给cart
			cart.CartItems = cartItems
			//在数据库中删除购物项
			dao.DeleteCartItemById(cartItemId)
		}
	}
	//更新数据库中的购物车
	dao.UpdateCart(cart)
	//刷新购物车页面
	GetCartInfo(w, r)
}

//更新购物项
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	//获取cartItemId
	cartItemId := r.FormValue("cartItemId")
	IcartItemId, _ := strconv.ParseInt(cartItemId, 10, 64)
	//获取用户输入的图书数量
	count := r.FormValue("count")
	icount, _ := strconv.ParseInt(count, 10, 64)
	//获取session
	_, session := dao.IsLogin(r)
	//获取userId
	userId := session.UserId
	//获取cart
	cart, _ := dao.GetCartByUserId(strconv.Itoa(userId))
	//获取cartItems切片
	cartItems := cart.CartItems
	//找到要更新的购物项
	for _, v := range cartItems {
		if v.CartItemId == IcartItemId {
			v.Count = icount
			//在数据库中更新
			dao.UpdateBookCount(v)
		}
	}
	//更新数据库中的购物车
	dao.UpdateCart(cart)
	//获取更新后的购物车
	cart, _ = dao.GetCartByUserId(strconv.Itoa(userId))
	//获取购物车中更新的数据
	var amount float64
	totalCount := cart.TotalCount
	totalAmount := cart.TotalAmount
	cartItems = cart.CartItems
	for _, v := range cartItems {
		if IcartItemId == v.CartItemId {
			amount = v.Amount
		}
	}
	//创建data实例
	data := &model.Data{
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
		Amount:      amount,
	}

	json, _ := json.Marshal(data)
	//响应到服务器
	w.Write(json)
}
