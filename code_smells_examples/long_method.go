package code_smells

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// 用户数据处理函数
func ProcessUserData(userData string) string {
	// 解析用户数据
	parts := strings.Split(userData, ",")
	if len(parts) < 5 {
		return "错误: 数据格式不正确"
	}

	// 提取用户信息
	name := parts[0]
	age := parts[1]
	email := parts[2]
	address := parts[3]
	purchaseHistory := parts[4]

	// 验证用户名
	if len(name) < 2 {
		return "错误: 用户名太短"
	}
	if len(name) > 50 {
		return "错误: 用户名太长"
	}
	if strings.Contains(name, "@") || strings.Contains(name, "#") || strings.Contains(name, "$") {
		return "错误: 用户名包含非法字符"
	}

	// 验证年龄
	ageNum := 0
	fmt.Sscanf(age, "%d", &ageNum)
	if ageNum < 18 {
		return "错误: 用户年龄必须大于18岁"
	}
	if ageNum > 120 {
		return "错误: 用户年龄不合理"
	}

	// 验证邮箱
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return "错误: 邮箱格式不正确"
	}
	emailParts := strings.Split(email, "@")
	if len(emailParts) != 2 || len(emailParts[0]) < 1 || len(emailParts[1]) < 3 {
		return "错误: 邮箱格式不正确"
	}
	if !strings.Contains(emailParts[1], ".") {
		return "错误: 邮箱域名格式不正确"
	}

	// 验证地址
	if len(address) < 10 {
		return "错误: 地址太短"
	}
	if len(address) > 200 {
		return "错误: 地址太长"
	}
	if !strings.Contains(address, " ") {
		return "错误: 地址格式不正确"
	}

	// 处理购买历史
	purchases := strings.Split(purchaseHistory, ";")
	totalSpent := 0.0
	for _, purchase := range purchases {
		if purchase == "" {
			continue
		}
		purchaseParts := strings.Split(purchase, ":")
		if len(purchaseParts) != 2 {
			return "错误: 购买记录格式不正确"
		}

		// 提取商品和价格
		item := purchaseParts[0]
		price := 0.0
		fmt.Sscanf(purchaseParts[1], "%f", &price)

		// 验证商品名称
		if len(item) < 2 {
			return "错误: 商品名称太短"
		}
		if len(item) > 100 {
			return "错误: 商品名称太长"
		}

		// 验证价格
		if price <= 0 {
			return "错误: 商品价格必须为正数"
		}
		if price > 10000 {
			return "错误: 商品价格超出合理范围"
		}

		// 累加总消费
		totalSpent += price
	}

	// 计算用户等级
	userLevel := "普通用户"
	if totalSpent > 1000 {
		userLevel = "银牌用户"
	}
	if totalSpent > 5000 {
		userLevel = "金牌用户"
	}
	if totalSpent > 10000 {
		userLevel = "钻石用户"
	}

	// 计算用户注册时间
	registrationDate := time.Now().AddDate(-ageNum/10, -2, -15)
	daysSinceRegistration := int(math.Floor(time.Since(registrationDate).Hours() / 24))

	// 计算用户活跃度
	activityLevel := "低"
	if daysSinceRegistration < 30 && len(purchases) > 0 {
		activityLevel = "中"
	}
	if daysSinceRegistration < 7 && len(purchases) > 3 {
		activityLevel = "高"
	}

	// 生成用户报告
	report := fmt.Sprintf("用户报告:\n姓名: %s\n年龄: %s\n邮箱: %s\n地址: %s\n总消费: %.2f\n用户等级: %s\n注册日期: %s\n活跃度: %s",
		name, age, email, address, totalSpent, userLevel, registrationDate.Format("2006-01-02"), activityLevel)

	// 添加促销信息
	if userLevel == "普通用户" {
		report += "\n促销信息: 首次购买满100元立减10元"
	} else if userLevel == "银牌用户" {
		report += "\n促销信息: 任意商品9折优惠"
	} else if userLevel == "金牌用户" {
		report += "\n促销信息: 任意商品8折优惠"
	} else {
		report += "\n促销信息: 任意商品7折优惠，另赠送会员积分"
	}

	// 添加活动推荐
	if activityLevel == "低" {
		report += "\n活动推荐: 久未光顾，现在购买有特别优惠"
	} else if activityLevel == "中" {
		report += "\n活动推荐: 感谢您的支持，参与新品试用可获得积分"
	} else {
		report += "\n活动推荐: 尊贵的活跃用户，您有专属VIP活动邀请"
	}

	return report
}
