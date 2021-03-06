package models

import (
	"time"
	"zcm_tools/orm"
)

//用户结算购物车中的商品
type AddOrderCar struct {
	Id         int
	Uid        int       //用户id
	Ids        string    //购物车商品的id
	Name       string    //菜名
	Count      string    //点一道菜的数量
	Price      string    //菜的总价
	CreateTime time.Time //创建时间
}

//用户订单表
type UserOrder struct {
	Id          int
	Uid         int       //用户id
	OrderId     string    //订单id
	Name        string    //收件人姓名
	IsProcessed int       //订单是否被处理 1:已处理，2:未处理
	OrderTime   time.Time //订单时间
	OrderPrice  string    //订单总价
	IsBuy       int       //是否付款:1:已支付:2未支付
	AddressId   int       //收货地址id
}

//添加结算购物车商品
func Addshops(uid int, count, name, price, ids string) (err error) {
	sql := `INSERT INTO add_order_car(uid,name,count,price,create_time,ids) VALUE(?, ?, ?, ?, NOW(),?)`
	_, err = orm.NewOrm().Raw(sql, uid, name, count, price, ids).Exec()
	return
}

//查询该用户的结算清单
func GetUserCloseOrder(uid int) (address []AddOrderCar, err error) {
	sql := `SELECT * from add_order_car where uid = ?`
	_, err = orm.NewOrm().Raw(sql, uid).QueryRows(&address)
	return
}

//求结算商品的总价
func SumCount(uid int) (a int, err error) {
	sql := `SELECT SUM(price) allprice
			FROM add_order_car
			WHERE uid = ?`
	err = orm.NewOrm().Raw(sql, uid).QueryRow(&a)
	return
}

//点击提交订单添加用户订单信息到order_table表里
func AddOrderTable(uid, addressId int, orderid, price, name string) (err error) {
	sql := `INSERT INTO order_table(uid, order_id, is_processed, order_time, order_price, is_buy, name, address_id, is_evaluate)
			VALUES(?, ?, 2 , NOW(), ?, 2, ?, ?, 2)`
	_, err = orm.NewOrm().Raw(sql, uid, orderid, price, name, addressId).Exec()
	return
}

//点击提交订单详细信息添加用户订单信息到order_detail表里
func AddOrderDetail(uid int, orderid, dishName, dishCount, dishPrice string) (err error) {
	sql := `INSERT INTO order_detail(uid, order_id, name, count, price, create_time)
			VALUES(?, ?, ?, ?, ?, now())`
	_, err = orm.NewOrm().Raw(sql, uid, orderid, dishName, dishCount, dishPrice).Exec()
	return
}

//修改数据库中菜品表中被点击的次数
func UpdateCountClick(dishName string, count int) (err error) {
	sql := `Update dish_table set order_count = order_count + ? where dish_name = ?`
	_, err = orm.NewOrm().Raw(sql, count, dishName).Exec()
	return
}

//点击支付宝支付修改order_table中 is_buy的状态
func UpdateIsBuy(uid int, orderid string) (err error) {
	sql := `UPDATE order_table SET is_buy = 1, is_evaluate = 2 WHERE uid =? AND order_id = ?`
	_, err = orm.NewOrm().Raw(sql, uid, orderid).Exec()
	return
}

//点击付款时删除已经结算购物车的商品
func DelShops(uid, id int) (err error) {
	sql := `DELETE FROM my_car WHERE uid = ? AND id = ?`
	_, err = orm.NewOrm().Raw(sql, uid, id).Exec()
	return
}

//点击付款时删除结算订单add_order_car里的数据
func DELAddOrderCar(uid, id int) (err error) {
	sql := `DELETE FROM add_order_car WHERE uid = ? AND ids = ?`
	_, err = orm.NewOrm().Raw(sql, uid, id).Exec()
	return
}
