package code_smells

import (
	"fmt"
	"time"
)

// 创建用户报告
func GenerateUserReport(
	userID string,
	username string,
	firstName string,
	lastName string,
	email string,
	phoneNumber string,
	address string,
	city string,
	state string,
	zipCode string,
	country string,
	birthDate time.Time,
	registrationDate time.Time,
	lastLoginDate time.Time,
	accountType string,
	accountStatus string,
	totalOrders int,
	totalSpent float64,
	averageOrderValue float64,
	loyaltyPoints int,
	preferredPaymentMethod string,
	preferredShippingMethod string,
	newsletterSubscribed bool,
	marketingPreferences []string,
	deviceType string,
	browser string,
	operatingSystem string,
	lastIPAddress string,
) string {
	// 构建用户基本信息
	report := fmt.Sprintf("用户报告 - 生成于: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	report += "== 用户信息 ==\n"
	report += fmt.Sprintf("ID: %s\n", userID)
	report += fmt.Sprintf("用户名: %s\n", username)
	report += fmt.Sprintf("姓名: %s %s\n", firstName, lastName)
	report += fmt.Sprintf("邮箱: %s\n", email)
	report += fmt.Sprintf("电话: %s\n", phoneNumber)

	// 地址信息
	report += "\n== 地址信息 ==\n"
	report += fmt.Sprintf("地址: %s\n", address)
	report += fmt.Sprintf("城市: %s\n", city)
	report += fmt.Sprintf("州/省: %s\n", state)
	report += fmt.Sprintf("邮编: %s\n", zipCode)
	report += fmt.Sprintf("国家: %s\n", country)

	// 账户信息
	report += "\n== 账户信息 ==\n"
	report += fmt.Sprintf("出生日期: %s\n", birthDate.Format("2006-01-02"))
	report += fmt.Sprintf("注册日期: %s\n", registrationDate.Format("2006-01-02"))
	report += fmt.Sprintf("最后登录: %s\n", lastLoginDate.Format("2006-01-02 15:04:05"))
	report += fmt.Sprintf("账户类型: %s\n", accountType)
	report += fmt.Sprintf("账户状态: %s\n", accountStatus)

	// 购买信息
	report += "\n== 购买信息 ==\n"
	report += fmt.Sprintf("订单总数: %d\n", totalOrders)
	report += fmt.Sprintf("总消费: %.2f\n", totalSpent)
	report += fmt.Sprintf("平均订单金额: %.2f\n", averageOrderValue)
	report += fmt.Sprintf("忠诚度积分: %d\n", loyaltyPoints)
	report += fmt.Sprintf("首选支付方式: %s\n", preferredPaymentMethod)
	report += fmt.Sprintf("首选配送方式: %s\n", preferredShippingMethod)

	// 偏好设置
	report += "\n== 偏好设置 ==\n"
	report += fmt.Sprintf("订阅新闻通讯: %t\n", newsletterSubscribed)
	report += "营销偏好: "
	for i, pref := range marketingPreferences {
		if i > 0 {
			report += ", "
		}
		report += pref
	}
	report += "\n"

	// 技术信息
	report += "\n== 技术信息 ==\n"
	report += fmt.Sprintf("设备类型: %s\n", deviceType)
	report += fmt.Sprintf("浏览器: %s\n", browser)
	report += fmt.Sprintf("操作系统: %s\n", operatingSystem)
	report += fmt.Sprintf("最后IP地址: %s\n", lastIPAddress)

	// 计算用户活跃度
	daysSinceLastLogin := int(time.Since(lastLoginDate).Hours() / 24)
	activityLevel := "高"
	if daysSinceLastLogin > 30 {
		activityLevel = "低"
	} else if daysSinceLastLogin > 7 {
		activityLevel = "中"
	}

	// 计算用户价值
	userValue := "普通"
	if totalSpent > 10000 {
		userValue = "VIP"
	} else if totalSpent > 5000 {
		userValue = "高价值"
	} else if totalSpent > 1000 {
		userValue = "中等价值"
	}

	// 添加分析结果
	report += "\n== 分析结果 ==\n"
	report += fmt.Sprintf("用户活跃度: %s\n", activityLevel)
	report += fmt.Sprintf("用户价值: %s\n", userValue)

	// 生成推荐
	report += "\n== 推荐操作 ==\n"
	if activityLevel == "低" {
		report += "- 发送重新激活邮件\n"
		report += "- 提供特别折扣\n"
	}
	if userValue == "VIP" || userValue == "高价值" {
		report += "- 提供专属客户服务\n"
		report += "- 邀请参加VIP活动\n"
	}
	if !newsletterSubscribed {
		report += "- 鼓励订阅新闻通讯\n"
	}

	return report
}

// 创建产品报告
func GenerateProductReport(
	productID string,
	productName string,
	description string,
	category string,
	subCategory string,
	brand string,
	supplier string,
	sku string,
	barcode string,
	price float64,
	cost float64,
	weight float64,
	dimensions string,
	color string,
	material string,
	creationDate time.Time,
	lastUpdateDate time.Time,
	inStock int,
	reorderLevel int,
	reorderQuantity int,
	totalSold int,
	averageRating float64,
	reviewCount int,
	returnRate float64,
	isActive bool,
	isFeatured bool,
	isDiscounted bool,
	discountPercentage float64,
	taxRate float64,
	shippingCategory string,
) string {
	// 构建产品基本信息
	report := fmt.Sprintf("产品报告 - 生成于: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	report += "== 产品信息 ==\n"
	report += fmt.Sprintf("ID: %s\n", productID)
	report += fmt.Sprintf("名称: %s\n", productName)
	report += fmt.Sprintf("描述: %s\n", description)
	report += fmt.Sprintf("类别: %s > %s\n", category, subCategory)
	report += fmt.Sprintf("品牌: %s\n", brand)
	report += fmt.Sprintf("供应商: %s\n", supplier)

	// 产品标识
	report += "\n== 产品标识 ==\n"
	report += fmt.Sprintf("SKU: %s\n", sku)
	report += fmt.Sprintf("条形码: %s\n", barcode)

	// 价格信息
	report += "\n== 价格信息 ==\n"
	report += fmt.Sprintf("售价: %.2f\n", price)
	report += fmt.Sprintf("成本: %.2f\n", cost)
	profit := price - cost
	margin := (profit / price) * 100
	report += fmt.Sprintf("利润: %.2f (%.2f%%)\n", profit, margin)

	// 物理特性
	report += "\n== 物理特性 ==\n"
	report += fmt.Sprintf("重量: %.2f\n", weight)
	report += fmt.Sprintf("尺寸: %s\n", dimensions)
	report += fmt.Sprintf("颜色: %s\n", color)
	report += fmt.Sprintf("材质: %s\n", material)

	// 库存信息
	report += "\n== 库存信息 ==\n"
	report += fmt.Sprintf("创建日期: %s\n", creationDate.Format("2006-01-02"))
	report += fmt.Sprintf("最后更新: %s\n", lastUpdateDate.Format("2006-01-02"))
	report += fmt.Sprintf("库存数量: %d\n", inStock)
	report += fmt.Sprintf("补货水平: %d\n", reorderLevel)
	report += fmt.Sprintf("补货数量: %d\n", reorderQuantity)

	// 销售信息
	report += "\n== 销售信息 ==\n"
	report += fmt.Sprintf("总销量: %d\n", totalSold)
	report += fmt.Sprintf("平均评分: %.1f\n", averageRating)
	report += fmt.Sprintf("评论数: %d\n", reviewCount)
	report += fmt.Sprintf("退货率: %.2f%%\n", returnRate*100)

	// 状态信息
	report += "\n== 状态信息 ==\n"
	report += fmt.Sprintf("是否激活: %t\n", isActive)
	report += fmt.Sprintf("是否推荐: %t\n", isFeatured)
	report += fmt.Sprintf("是否折扣: %t\n", isDiscounted)
	if isDiscounted {
		report += fmt.Sprintf("折扣比例: %.2f%%\n", discountPercentage*100)
		discountedPrice := price * (1 - discountPercentage)
		report += fmt.Sprintf("折扣价: %.2f\n", discountedPrice)
	}

	// 其他信息
	report += "\n== 其他信息 ==\n"
	report += fmt.Sprintf("税率: %.2f%%\n", taxRate*100)
	report += fmt.Sprintf("配送类别: %s\n", shippingCategory)

	// 分析结果
	report += "\n== 分析结果 ==\n"

	// 库存状态
	stockStatus := "正常"
	if inStock <= 0 {
		stockStatus = "缺货"
	} else if inStock < reorderLevel {
		stockStatus = "需要补货"
	}
	report += fmt.Sprintf("库存状态: %s\n", stockStatus)

	// 产品表现
	performance := "一般"
	if totalSold > 1000 && averageRating >= 4.5 {
		performance = "优秀"
	} else if totalSold > 500 || averageRating >= 4.0 {
		performance = "良好"
	} else if totalSold < 100 && averageRating < 3.0 {
		performance = "不佳"
	}
	report += fmt.Sprintf("产品表现: %s\n", performance)

	// 推荐操作
	report += "\n== 推荐操作 ==\n"
	if stockStatus == "缺货" || stockStatus == "需要补货" {
		report += "- 补充库存\n"
	}
	if performance == "不佳" {
		report += "- 考虑降价或促销\n"
		report += "- 改进产品质量\n"
	}
	if returnRate > 0.1 {
		report += "- 调查高退货率原因\n"
	}
	if performance == "优秀" && !isFeatured {
		report += "- 考虑将产品设为推荐\n"
	}

	return report
}
