package main

import "fmt"

// TODO 1: 定义一个 Car 结构体，包含以下字段：
// - Brand (品牌，string)
// - Model (型号，string)
// - Price (价格，float64)

type Car struct {
	Brand string
	Model string
	Price float64
}

type Rectangle struct {
	Width  int
	Height int
}

type User struct {
	Name  string
	Email string
	age   int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// TODO 1: 写一个函数 updateEmail，接收 User 和 新邮箱，修改邮箱
// 注意：这个函数应该能修改原值，所以需要传指针
func upadteEmail(u *User, newEmail string) {
	u.Email = newEmail
}

// TODO 2: 写一个函数 showUser，接收 User，打印用户信息（不需要修改原值）
func showUser(u User) {
	fmt.Println(u.Name, u.Email, u.age)
}
func main() {
	// TODO 2: 创建一个 Car 实例，品牌为"特斯拉"，型号为"Model 3"，价格为 250000
	car := Car{Brand: "特斯拉", Model: "3", Price: 250000}
	// TODO 3: 打印这辆车的所有信息
	// 格式："品牌:特斯拉, 型号:Model 3, 价格:250000"
	fmt.Printf("这辆车的品牌是%v，型号是%v，价格是%v \n", car.Brand, car.Model, car.Price)

	//TODO 1: 给 Rectangle 绑定一个 Area 方法，返回面积（宽 × 高）
	rect := Rectangle{Width: 3, Height: 4}
	area := rect.Area()
	fmt.Println(area)
	//创建实例
	u := User{Name: "梁金山", Email: "34666", age: 10}

	fmt.Println("=== 修改前 ===")
	showUser(u)
	upadteEmail(&u, "newemail@example.com")
	fmt.Println("=== 修改后 ===")
	showUser(u)

}
