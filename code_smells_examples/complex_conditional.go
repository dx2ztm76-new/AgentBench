package code_smells

import (
	"fmt"
	"time"
)

// 用户类型
type UserType int

const (
	Regular UserType = iota
	Premium
	Enterprise
	Admin
	Guest
)

// 支付方式
type PaymentMethod int

const (
	CreditCard PaymentMethod = iota
	DebitCard
	BankTransfer
	DigitalWallet
	GiftCard
	StoredCredit
)

// 订单状态
type OrderStatus int

const (
	Pending OrderStatus = iota
	Processing
	Shipped
	Delivered
	Cancelled
	Returned
	Refunded
)

// 产品类型
type ProductType int

const (
	Physical ProductType = iota
	Digital
	Subscription
	Service
)

// 过度复杂的条件逻辑示例：计算订单折扣
func CalculateOrderDiscount(
	orderTotal float64,
	userType UserType,
	userRegistrationDate time.Time,
	previousOrdersCount int,
	totalSpent float64,
	paymentMethod PaymentMethod,
	isHoliday bool,
	hasPromoCode bool,
	promoCodeValue float64,
	orderStatus OrderStatus,
	productType ProductType,
	quantity int,
	isFirstPurchase bool,
	hasCoupon bool,
	couponValue float64,
	isWeekend bool,
	isFlashSale bool,
) float64 {
	// 基础折扣
	var discount float64 = 0

	// 根据用户类型计算折扣
	if userType == Regular {
		// 普通用户
		if previousOrdersCount > 10 {
			discount += 0.05
		} else if previousOrdersCount > 5 {
			discount += 0.02
		}

		if totalSpent > 1000 {
			discount += 0.03
		} else if totalSpent > 500 {
			discount += 0.01
		}

		// 注册时间超过1年的老用户
		if time.Since(userRegistrationDate).Hours() > 24*365 {
			discount += 0.02
		}
	} else if userType == Premium {
		// 高级用户固定折扣
		discount += 0.10

		// 额外的忠诚度折扣
		if previousOrdersCount > 20 {
			discount += 0.05
		} else if previousOrdersCount > 10 {
			discount += 0.03
		}

		// 高消费额外奖励
		if totalSpent > 2000 {
			discount += 0.05
		} else if totalSpent > 1000 {
			discount += 0.02
		}
	} else if userType == Enterprise {
		// 企业用户固定折扣
		discount += 0.15

		// 大订单额外折扣
		if orderTotal > 5000 {
			discount += 0.10
		} else if orderTotal > 1000 {
			discount += 0.05
		}
	} else if userType == Admin {
		// 管理员折扣
		discount += 0.25
	} else if userType == Guest {
		// 游客没有基础折扣
		discount += 0

		// 但如果是首次购买，给予小折扣以鼓励注册
		if isFirstPurchase {
			discount += 0.01
		}
	}

	// 根据支付方式计算额外折扣
	if paymentMethod == CreditCard {
		discount += 0.02
	} else if paymentMethod == DebitCard {
		discount += 0.01
	} else if paymentMethod == BankTransfer {
		// 银行转账没有额外折扣
	} else if paymentMethod == DigitalWallet {
		discount += 0.015
	} else if paymentMethod == GiftCard {
		// 礼品卡没有额外折扣
	} else if paymentMethod == StoredCredit {
		discount += 0.03
	}

	// 节假日特殊折扣
	if isHoliday {
		if userType == Regular || userType == Guest {
			discount += 0.02
		} else if userType == Premium || userType == Enterprise {
			discount += 0.03
		} else if userType == Admin {
			// 管理员已经有很高折扣，不再提供节假日折扣
		}

		// 节假日期间使用信用卡额外折扣
		if paymentMethod == CreditCard {
			discount += 0.01
		}
	}

	// 周末折扣
	if isWeekend {
		if productType == Physical {
			discount += 0.01
		} else if productType == Digital {
			discount += 0.02
		} else if productType == Subscription {
			// 订阅产品周末没有额外折扣
		} else if productType == Service {
			discount += 0.015
		}
	}

	// 闪购折扣
	if isFlashSale {
		// 闪购基础折扣
		discount += 0.05

		// 不同用户类型在闪购时的额外折扣
		if userType == Premium {
			discount += 0.02
		} else if userType == Enterprise {
			discount += 0.03
		}

		// 闪购期间大量购买的额外折扣
		if quantity > 5 {
			discount += 0.03
		} else if quantity > 2 {
			discount += 0.01
		}
	}

	// 促销码折扣
	if hasPromoCode {
		// 如果折扣码提供的折扣更高，则使用折扣码的折扣
		if promoCodeValue > discount {
			discount = promoCodeValue
		} else {
			// 否则，在现有折扣基础上增加一点额外折扣
			discount += 0.01
		}
	}

	// 优惠券折扣
	if hasCoupon {
		// 优惠券折扣与其他折扣叠加
		discount += couponValue
	}

	// 根据订单状态调整折扣
	if orderStatus == Cancelled || orderStatus == Returned {
		// 取消或退货的订单没有折扣
		discount = 0
	} else if orderStatus == Pending {
		// 待处理订单正常享受折扣
	} else if orderStatus == Processing {
		// 处理中订单正常享受折扣
	} else if orderStatus == Shipped {
		// 已发货订单正常享受折扣
	} else if orderStatus == Delivered {
		// 已送达订单正常享受折扣
	} else if orderStatus == Refunded {
		// 已退款订单没有折扣
		discount = 0
	}

	// 根据产品类型调整折扣
	if productType == Physical {
		// 实体产品正常享受折扣
	} else if productType == Digital {
		// 数字产品额外折扣
		discount += 0.01
	} else if productType == Subscription {
		// 订阅产品额外折扣
		discount += 0.02

		// 长期用户订阅额外折扣
		if previousOrdersCount > 5 && userType != Guest {
			discount += 0.01
		}
	} else if productType == Service {
		// 服务类产品正常享受折扣
	}

	// 大订单额外折扣
	if orderTotal > 10000 {
		discount += 0.10
	} else if orderTotal > 5000 {
		discount += 0.07
	} else if orderTotal > 1000 {
		discount += 0.05
	} else if orderTotal > 500 {
		discount += 0.02
	}

	// 大量购买额外折扣
	if quantity > 20 {
		discount += 0.08
	} else if quantity > 10 {
		discount += 0.05
	} else if quantity > 5 {
		discount += 0.02
	}

	// 确保折扣不超过最大值
	if discount > 0.50 {
		discount = 0.50
	}

	// 计算最终折扣金额
	discountAmount := orderTotal * discount

	return discountAmount
}

// 打印订单折扣信息
func PrintOrderDiscountInfo(
	orderTotal float64,
	userType UserType,
	userRegistrationDate time.Time,
	previousOrdersCount int,
	totalSpent float64,
	paymentMethod PaymentMethod,
	isHoliday bool,
	hasPromoCode bool,
	promoCodeValue float64,
	orderStatus OrderStatus,
	productType ProductType,
	quantity int,
	isFirstPurchase bool,
	hasCoupon bool,
	couponValue float64,
	isWeekend bool,
	isFlashSale bool,
) {
	discount := CalculateOrderDiscount(
		orderTotal,
		userType,
		userRegistrationDate,
		previousOrdersCount,
		totalSpent,
		paymentMethod,
		isHoliday,
		hasPromoCode,
		promoCodeValue,
		orderStatus,
		productType,
		quantity,
		isFirstPurchase,
		hasCoupon,
		couponValue,
		isWeekend,
		isFlashSale,
	)

	fmt.Printf("订单总额: ¥%.2f\n", orderTotal)
	fmt.Printf("折扣金额: ¥%.2f\n", discount)
	fmt.Printf("折扣后金额: ¥%.2f\n", orderTotal-discount)
	fmt.Printf("折扣比例: %.2f%%\n", (discount/orderTotal)*100)
}

func main() {
	// 示例：计算一个高级用户的订单折扣
	orderTotal := 1200.0
	userType := Premium
	userRegistrationDate := time.Now().AddDate(-2, 0, 0) // 2年前注册
	previousOrdersCount := 15
	totalSpent := 3000.0
	paymentMethod := CreditCard
	isHoliday := true
	hasPromoCode := true
	promoCodeValue := 0.05
	orderStatus := Processing
	productType := Physical
	quantity := 3
	isFirstPurchase := false
	hasCoupon := true
	couponValue := 0.03
	isWeekend := true
	isFlashSale := false

	PrintOrderDiscountInfo(
		orderTotal,
		userType,
		userRegistrationDate,
		previousOrdersCount,
		totalSpent,
		paymentMethod,
		isHoliday,
		hasPromoCode,
		promoCodeValue,
		orderStatus,
		productType,
		quantity,
		isFirstPurchase,
		hasCoupon,
		couponValue,
		isWeekend,
		isFlashSale,
	)
}
