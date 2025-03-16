package code_smells

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 电子商务系统
type ECommerceSystem struct {
	// 用户管理
	Users               map[string]User
	LoggedInUsers       map[string]bool
	UserSessions        map[string]time.Time
	FailedLoginAttempts map[string]int
	UserPreferences     map[string]UserPreference
	UserAddresses       map[string][]Address
	UserPaymentMethods  map[string][]PaymentMethod

	// 产品管理
	Products           map[string]Product
	ProductCategories  map[string][]string
	ProductInventory   map[string]int
	ProductRatings     map[string][]Rating
	FeaturedProducts   []string
	DiscountedProducts map[string]float64

	// 订单管理
	Orders            map[string]Order
	OrderStatuses     map[string]string
	ShippingProviders map[string]ShippingProvider
	ShippingRates     map[string]float64

	// 购物车管理
	Carts          map[string]Cart
	SavedForLater  map[string][]string
	AbandonedCarts map[string]time.Time

	// 促销管理
	Promotions    map[string]Promotion
	CouponCodes   map[string]Coupon
	LoyaltyPoints map[string]int

	// 系统配置
	Configuration          SystemConfig
	APIKeys                map[string]string
	ThirdPartyIntegrations map[string]Integration

	// 统计和分析
	PageViews       map[string]int
	ProductViews    map[string]int
	SearchQueries   map[string]int
	ConversionRates map[string]float64
}

// 用户相关结构
type User struct {
	ID        string
	Username  string
	Email     string
	Password  string
	FirstName string
	LastName  string
	Phone     string
	CreatedAt time.Time
	LastLogin time.Time
	IsActive  bool
}

type UserPreference struct {
	EmailNotifications bool
	SMSNotifications   bool
	LanguagePreference string
	CurrencyPreference string
	ThemePreference    string
}

type Address struct {
	Type      string
	Street    string
	City      string
	State     string
	ZipCode   string
	Country   string
	IsDefault bool
}

type PaymentMethod struct {
	Type       string
	CardNumber string
	ExpiryDate string
	CVV        string
	HolderName string
	IsDefault  bool
}

// 产品相关结构
type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
	ImageURL    string
	Weight      float64
	Dimensions  string
	SKU         string
	Barcode     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Rating struct {
	UserID    string
	Score     int
	Comment   string
	CreatedAt time.Time
}

// 订单相关结构
type Order struct {
	ID              string
	UserID          string
	Products        []OrderItem
	TotalAmount     float64
	ShippingAddress Address
	BillingAddress  Address
	PaymentMethod   PaymentMethod
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type OrderItem struct {
	ProductID  string
	Quantity   int
	UnitPrice  float64
	TotalPrice float64
}

type ShippingProvider struct {
	Name                  string
	TrackingURL           string
	EstimatedDeliveryDays int
}

// 购物车相关结构
type Cart struct {
	UserID    string
	Items     []CartItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartItem struct {
	ProductID string
	Quantity  int
}

// 促销相关结构
type Promotion struct {
	ID          string
	Name        string
	Description string
	Discount    float64
	StartDate   time.Time
	EndDate     time.Time
	IsActive    bool
}

type Coupon struct {
	Code        string
	Discount    float64
	MinPurchase float64
	ExpiryDate  time.Time
	IsActive    bool
}

// 系统配置
type SystemConfig struct {
	SiteName         string
	SiteURL          string
	AdminEmail       string
	MaxLoginAttempts int
	SessionTimeout   time.Duration
	MaintenanceMode  bool
	DebugMode        bool
}

type Integration struct {
	Name        string
	APIEndpoint string
	APIKey      string
	IsActive    bool
}

// 初始化电子商务系统
func NewECommerceSystem() *ECommerceSystem {
	return &ECommerceSystem{
		Users:                  make(map[string]User),
		LoggedInUsers:          make(map[string]bool),
		UserSessions:           make(map[string]time.Time),
		FailedLoginAttempts:    make(map[string]int),
		UserPreferences:        make(map[string]UserPreference),
		UserAddresses:          make(map[string][]Address),
		UserPaymentMethods:     make(map[string][]PaymentMethod),
		Products:               make(map[string]Product),
		ProductCategories:      make(map[string][]string),
		ProductInventory:       make(map[string]int),
		ProductRatings:         make(map[string][]Rating),
		FeaturedProducts:       []string{},
		DiscountedProducts:     make(map[string]float64),
		Orders:                 make(map[string]Order),
		OrderStatuses:          make(map[string]string),
		ShippingProviders:      make(map[string]ShippingProvider),
		ShippingRates:          make(map[string]float64),
		Carts:                  make(map[string]Cart),
		SavedForLater:          make(map[string][]string),
		AbandonedCarts:         make(map[string]time.Time),
		Promotions:             make(map[string]Promotion),
		CouponCodes:            make(map[string]Coupon),
		LoyaltyPoints:          make(map[string]int),
		APIKeys:                make(map[string]string),
		ThirdPartyIntegrations: make(map[string]Integration),
		PageViews:              make(map[string]int),
		ProductViews:           make(map[string]int),
		SearchQueries:          make(map[string]int),
		ConversionRates:        make(map[string]float64),
		Configuration: SystemConfig{
			SiteName:         "我的电子商务网站",
			SiteURL:          "https://myecommerce.com",
			AdminEmail:       "admin@myecommerce.com",
			MaxLoginAttempts: 5,
			SessionTimeout:   24 * time.Hour,
			MaintenanceMode:  false,
			DebugMode:        true,
		},
	}
}

// 用户注册
func (e *ECommerceSystem) RegisterUser(username, email, password, firstName, lastName, phone string) (string, error) {
	// 检查用户名是否已存在
	for _, user := range e.Users {
		if user.Username == username {
			return "", fmt.Errorf("用户名已存在")
		}
		if user.Email == email {
			return "", fmt.Errorf("邮箱已被注册")
		}
	}

	// 创建新用户
	userID := fmt.Sprintf("user_%d", len(e.Users)+1)
	e.Users[userID] = User{
		ID:        userID,
		Username:  username,
		Email:     email,
		Password:  password, // 注意：实际应用中应该哈希密码
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		CreatedAt: time.Now(),
		IsActive:  true,
	}

	// 初始化用户相关数据
	e.UserPreferences[userID] = UserPreference{
		EmailNotifications: true,
		SMSNotifications:   false,
		LanguagePreference: "zh-CN",
		CurrencyPreference: "CNY",
		ThemePreference:    "default",
	}

	e.UserAddresses[userID] = []Address{}
	e.UserPaymentMethods[userID] = []PaymentMethod{}
	e.LoyaltyPoints[userID] = 100 // 注册奖励积分

	return userID, nil
}

// 用户登录
func (e *ECommerceSystem) Login(usernameOrEmail, password string) (string, error) {
	var userID string
	var user User
	found := false

	// 查找用户
	for id, u := range e.Users {
		if u.Username == usernameOrEmail || u.Email == usernameOrEmail {
			userID = id
			user = u
			found = true
			break
		}
	}

	if !found {
		return "", fmt.Errorf("用户不存在")
	}

	// 检查账户是否激活
	if !user.IsActive {
		return "", fmt.Errorf("账户已被禁用")
	}

	// 检查登录失败次数
	if e.FailedLoginAttempts[userID] >= e.Configuration.MaxLoginAttempts {
		return "", fmt.Errorf("登录失败次数过多，账户已被锁定")
	}

	// 验证密码
	if user.Password != password { // 注意：实际应用中应该比较哈希值
		e.FailedLoginAttempts[userID]++
		return "", fmt.Errorf("密码错误")
	}

	// 登录成功，重置失败次数
	e.FailedLoginAttempts[userID] = 0

	// 更新登录状态
	e.LoggedInUsers[userID] = true
	e.UserSessions[userID] = time.Now().Add(e.Configuration.SessionTimeout)

	// 更新最后登录时间
	user.LastLogin = time.Now()
	e.Users[userID] = user

	return userID, nil
}

// 添加产品
func (e *ECommerceSystem) AddProduct(name, description, imageURL, dimensions, sku, barcode string, price, weight float64) (string, error) {
	// 检查SKU是否已存在
	for _, product := range e.Products {
		if product.SKU == sku {
			return "", fmt.Errorf("SKU已存在")
		}
	}

	// 创建新产品
	productID := fmt.Sprintf("prod_%d", len(e.Products)+1)
	e.Products[productID] = Product{
		ID:          productID,
		Name:        name,
		Description: description,
		Price:       price,
		ImageURL:    imageURL,
		Weight:      weight,
		Dimensions:  dimensions,
		SKU:         sku,
		Barcode:     barcode,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// 初始化库存
	e.ProductInventory[productID] = 0

	return productID, nil
}

// 更新产品库存
func (e *ECommerceSystem) UpdateInventory(productID string, quantity int) error {
	if _, exists := e.Products[productID]; !exists {
		return fmt.Errorf("产品不存在")
	}

	e.ProductInventory[productID] = quantity
	return nil
}

// 添加产品到购物车
func (e *ECommerceSystem) AddToCart(userID, productID string, quantity int) error {
	// 验证用户
	if _, exists := e.Users[userID]; !exists {
		return fmt.Errorf("用户不存在")
	}

	// 验证产品
	if _, exists := e.Products[productID]; !exists {
		return fmt.Errorf("产品不存在")
	}

	// 检查库存
	if e.ProductInventory[productID] < quantity {
		return fmt.Errorf("库存不足")
	}

	// 获取或创建购物车
	cart, exists := e.Carts[userID]
	if !exists {
		cart = Cart{
			UserID:    userID,
			Items:     []CartItem{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}

	// 检查购物车中是否已有该产品
	found := false
	for i, item := range cart.Items {
		if item.ProductID == productID {
			cart.Items[i].Quantity += quantity
			found = true
			break
		}
	}

	// 如果购物车中没有该产品，添加新项目
	if !found {
		cart.Items = append(cart.Items, CartItem{
			ProductID: productID,
			Quantity:  quantity,
		})
	}

	cart.UpdatedAt = time.Now()
	e.Carts[userID] = cart

	return nil
}

// 创建订单
func (e *ECommerceSystem) CreateOrder(userID string, shippingAddressID, billingAddressID, paymentMethodID string) (string, error) {
	// 验证用户
	if _, exists := e.Users[userID]; !exists {
		return "", fmt.Errorf("用户不存在")
	}

	// 获取购物车
	cart, exists := e.Carts[userID]
	if !exists || len(cart.Items) == 0 {
		return "", fmt.Errorf("购物车为空")
	}

	// 验证地址
	userAddresses := e.UserAddresses[userID]
	var shippingAddress, billingAddress Address
	foundShipping, foundBilling := false, false

	for _, addr := range userAddresses {
		if addr.Type == shippingAddressID {
			shippingAddress = addr
			foundShipping = true
		}
		if addr.Type == billingAddressID {
			billingAddress = addr
			foundBilling = true
		}
	}

	if !foundShipping {
		return "", fmt.Errorf("送货地址不存在")
	}
	if !foundBilling {
		return "", fmt.Errorf("账单地址不存在")
	}

	// 验证支付方式
	userPaymentMethods := e.UserPaymentMethods[userID]
	var paymentMethod PaymentMethod
	foundPayment := false

	for _, pm := range userPaymentMethods {
		if pm.Type == paymentMethodID {
			paymentMethod = pm
			foundPayment = true
			break
		}
	}

	if !foundPayment {
		return "", fmt.Errorf("支付方式不存在")
	}

	// 创建订单项目并计算总金额
	var orderItems []OrderItem
	totalAmount := 0.0

	for _, item := range cart.Items {
		product := e.Products[item.ProductID]

		// 检查库存
		if e.ProductInventory[item.ProductID] < item.Quantity {
			return "", fmt.Errorf("产品 %s 库存不足", product.Name)
		}

		// 计算价格（考虑折扣）
		unitPrice := product.Price
		if discount, hasDiscount := e.DiscountedProducts[item.ProductID]; hasDiscount {
			unitPrice = unitPrice * (1 - discount)
		}

		totalPrice := unitPrice * float64(item.Quantity)

		orderItems = append(orderItems, OrderItem{
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			UnitPrice:  unitPrice,
			TotalPrice: totalPrice,
		})

		totalAmount += totalPrice

		// 更新库存
		e.ProductInventory[item.ProductID] -= item.Quantity
	}

	// 创建订单
	orderID := fmt.Sprintf("order_%d", len(e.Orders)+1)
	e.Orders[orderID] = Order{
		ID:              orderID,
		UserID:          userID,
		Products:        orderItems,
		TotalAmount:     totalAmount,
		ShippingAddress: shippingAddress,
		BillingAddress:  billingAddress,
		PaymentMethod:   paymentMethod,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// 设置订单状态
	e.OrderStatuses[orderID] = "已创建"

	// 清空购物车
	delete(e.Carts, userID)

	// 添加忠诚度积分
	e.LoyaltyPoints[userID] += int(totalAmount / 10)

	return orderID, nil
}

// 应用优惠券
func (e *ECommerceSystem) ApplyCoupon(userID, couponCode string) (float64, error) {
	// 验证用户
	if _, exists := e.Users[userID]; !exists {
		return 0, fmt.Errorf("用户不存在")
	}

	// 验证优惠券
	coupon, exists := e.CouponCodes[couponCode]
	if !exists {
		return 0, fmt.Errorf("优惠券不存在")
	}

	// 检查优惠券是否有效
	if !coupon.IsActive {
		return 0, fmt.Errorf("优惠券已失效")
	}

	// 检查优惠券是否过期
	if time.Now().After(coupon.ExpiryDate) {
		return 0, fmt.Errorf("优惠券已过期")
	}

	// 获取购物车
	cart, exists := e.Carts[userID]
	if !exists || len(cart.Items) == 0 {
		return 0, fmt.Errorf("购物车为空")
	}

	// 计算购物车总金额
	totalAmount := 0.0
	for _, item := range cart.Items {
		product := e.Products[item.ProductID]
		totalAmount += product.Price * float64(item.Quantity)
	}

	// 检查最低购买金额
	if totalAmount < coupon.MinPurchase {
		return 0, fmt.Errorf("未达到优惠券最低购买金额 %.2f", coupon.MinPurchase)
	}

	// 计算折扣金额
	discountAmount := totalAmount * coupon.Discount

	return discountAmount, nil
}

// 保存系统数据到文件
func (e *ECommerceSystem) SaveToFile(filename string) error {
	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

// 从文件加载系统数据
func (e *ECommerceSystem) LoadFromFile(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, e)
}

// 启动HTTP服务器
func (e *ECommerceSystem) StartServer(port string) {
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		products := make([]Product, 0, len(e.Products))
		for _, product := range e.Products {
			products = append(products, product)
		}

		json.NewEncoder(w).Encode(products)
	})

	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			http.Error(w, "无效的产品ID", http.StatusBadRequest)
			return
		}

		productID := parts[3]
		product, exists := e.Products[productID]
		if !exists {
			http.Error(w, "产品不存在", http.StatusNotFound)
			return
		}

		// 增加产品浏览次数
		e.ProductViews[productID]++

		json.NewEncoder(w).Encode(product)
	})

	http.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "方法不允许", http.StatusMethodNotAllowed)
			return
		}

		var data struct {
			Username  string `json:"username"`
			Email     string `json:"email"`
			Password  string `json:"password"`
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
			Phone     string `json:"phone"`
		}

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "无效的请求数据", http.StatusBadRequest)
			return
		}

		userID, err := e.RegisterUser(data.Username, data.Email, data.Password, data.FirstName, data.LastName, data.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"userID": userID})
	})

	fmt.Printf("服务器启动在端口 %s\n", port)
	http.ListenAndServe(":"+port, nil)
}
