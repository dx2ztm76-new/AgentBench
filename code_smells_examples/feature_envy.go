package code_smells

import (
	"fmt"
	"math"
	"time"
)

// 用户结构体
type User struct {
	ID               string
	Name             string
	Email            string
	Address          string
	City             string
	State            string
	ZipCode          string
	Country          string
	PhoneNumber      string
	RegistrationDate time.Time
	LastLoginDate    time.Time
	TotalOrders      int
	TotalSpent       float64
}

// 订单结构体
type Order struct {
	ID           string
	UserID       string
	Items        []OrderItem
	OrderDate    time.Time
	Status       string
	ShippingCost float64
	TaxAmount    float64
}

// 订单项结构体
type OrderItem struct {
	ProductID   string
	ProductName string
	Quantity    int
	UnitPrice   float64
}

// 特性依恋示例：这个函数过度使用User对象的数据，应该移到User结构体中
func CalculateUserLoyaltyScore(user User, orders []Order) float64 {
	// 计算用户的忠诚度分数

	// 基于注册时间的分数（最多30分）
	daysSinceRegistration := time.Since(user.RegistrationDate).Hours() / 24
	registrationScore := math.Min(daysSinceRegistration/30, 30)

	// 基于订单数量的分数（最多40分）
	orderCountScore := math.Min(float64(user.TotalOrders)*2, 40)

	// 基于消费金额的分数（最多30分）
	spendingScore := math.Min(user.TotalSpent/100, 30)

	// 计算最近一次订单的时间
	var mostRecentOrderDate time.Time
	for _, order := range orders {
		if order.UserID == user.ID && (mostRecentOrderDate.IsZero() || order.OrderDate.After(mostRecentOrderDate)) {
			mostRecentOrderDate = order.OrderDate
		}
	}

	// 基于最近订单时间的加分或减分（最多±20分）
	var recencyScore float64
	if !mostRecentOrderDate.IsZero() {
		daysSinceLastOrder := time.Since(mostRecentOrderDate).Hours() / 24
		if daysSinceLastOrder < 30 {
			// 如果30天内有订单，加分
			recencyScore = 20
		} else if daysSinceLastOrder < 90 {
			// 如果90天内有订单，少量加分
			recencyScore = 10
		} else if daysSinceLastOrder < 180 {
			// 如果180天内有订单，不加不减
			recencyScore = 0
		} else if daysSinceLastOrder < 365 {
			// 如果一年内有订单，少量减分
			recencyScore = -10
		} else {
			// 如果超过一年没有订单，减分
			recencyScore = -20
		}
	}

	// 计算总分
	totalScore := registrationScore + orderCountScore + spendingScore + recencyScore

	// 确保分数在0-100之间
	return math.Max(0, math.Min(totalScore, 100))
}

// 特性依恋示例：这个函数过度使用Order对象的数据，应该移到Order结构体中
func CalculateOrderTotal(order Order) float64 {
	// 计算订单的总金额

	var subtotal float64
	for _, item := range order.Items {
		subtotal += float64(item.Quantity) * item.UnitPrice
	}

	// 添加运费和税费
	total := subtotal + order.ShippingCost + order.TaxAmount

	return total
}

// 特性依恋示例：这个函数过度使用OrderItem对象的数据，应该移到OrderItem结构体中
func CalculateItemTotal(item OrderItem) float64 {
	// 计算订单项的总金额
	return float64(item.Quantity) * item.UnitPrice
}

// 打印用户信息和忠诚度分数
func PrintUserLoyaltyInfo(user User, orders []Order) {
	loyaltyScore := CalculateUserLoyaltyScore(user, orders)
	fmt.Printf("用户 %s 的忠诚度分数: %.2f\n", user.Name, loyaltyScore)

	if loyaltyScore >= 90 {
		fmt.Println("忠诚度等级: 钻石")
	} else if loyaltyScore >= 70 {
		fmt.Println("忠诚度等级: 黄金")
	} else if loyaltyScore >= 50 {
		fmt.Println("忠诚度等级: 白银")
	} else if loyaltyScore >= 30 {
		fmt.Println("忠诚度等级: 青铜")
	} else {
		fmt.Println("忠诚度等级: 普通")
	}
}

// 打印订单信息
func PrintOrderInfo(order Order) {
	total := CalculateOrderTotal(order)
	fmt.Printf("订单 %s 的总金额: ¥%.2f\n", order.ID, total)
	fmt.Println("订单项:")

	for _, item := range order.Items {
		itemTotal := CalculateItemTotal(item)
		fmt.Printf("  - %s (x%d): ¥%.2f\n", item.ProductName, item.Quantity, itemTotal)
	}
}

func main() {
	// 创建用户
	user := User{
		ID:               "U12345",
		Name:             "张三",
		Email:            "zhangsan@example.com",
		Address:          "北京路123号",
		City:             "上海",
		State:            "上海",
		ZipCode:          "200000",
		Country:          "中国",
		PhoneNumber:      "13800138000",
		RegistrationDate: time.Now().AddDate(-2, 0, 0), // 2年前注册
		LastLoginDate:    time.Now().AddDate(0, 0, -5), // 5天前登录
		TotalOrders:      25,
		TotalSpent:       4500.0,
	}

	// 创建订单
	orders := []Order{
		{
			ID:           "O12345",
			UserID:       "U12345",
			OrderDate:    time.Now().AddDate(0, 0, -20), // 20天前的订单
			Status:       "已完成",
			ShippingCost: 15.0,
			TaxAmount:    25.0,
			Items: []OrderItem{
				{
					ProductID:   "P1",
					ProductName: "笔记本电脑",
					Quantity:    1,
					UnitPrice:   5999.0,
				},
				{
					ProductID:   "P2",
					ProductName: "无线鼠标",
					Quantity:    1,
					UnitPrice:   99.0,
				},
			},
		},
		{
			ID:           "O12346",
			UserID:       "U12345",
			OrderDate:    time.Now().AddDate(0, -3, 0), // 3个月前的订单
			Status:       "已完成",
			ShippingCost: 0.0,
			TaxAmount:    10.0,
			Items: []OrderItem{
				{
					ProductID:   "P3",
					ProductName: "蓝牙耳机",
					Quantity:    1,
					UnitPrice:   299.0,
				},
			},
		},
	}

	// 打印用户忠诚度信息
	PrintUserLoyaltyInfo(user, orders)

	// 打印订单信息
	for _, order := range orders {
		fmt.Println("\n订单详情:")
		PrintOrderInfo(order)
	}
}
