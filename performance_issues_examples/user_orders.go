package performance_examples

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func RunDatabaseExample() {
	// 创建内存数据库
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表并插入测试数据
	setupDatabase(db)

	// 测试不同的查询方法
	fmt.Println("测试不同的查询方法...")

	// 方法1: 单个查询
	start := time.Now()
	method1Results := queryMethod1(db, 100)
	duration1 := time.Since(start)
	fmt.Printf("方法1 (单个查询): %v, 获取了 %d 个订单\n", duration1, len(method1Results))

	// 方法2: 使用JOIN
	start = time.Now()
	method2Results := queryMethod2(db, 100)
	duration2 := time.Since(start)
	fmt.Printf("方法2 (JOIN查询): %v, 获取了 %d 个订单\n", duration2, len(method2Results))

	// 方法3: 使用预处理语句
	start = time.Now()
	method3Results := queryMethod3(db, 100)
	duration3 := time.Since(start)
	fmt.Printf("方法3 (预处理语句): %v, 获取了 %d 个订单\n", duration3, len(method3Results))
}

// 设置数据库表和测试数据
func setupDatabase(db *sql.DB) {
	// 创建用户表
	_, err := db.Exec(`
		CREATE TABLE users (
			id INTEGER PRIMARY KEY,
			name TEXT,
			email TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// 创建订单表
	_, err = db.Exec(`
		CREATE TABLE orders (
			id INTEGER PRIMARY KEY,
			user_id INTEGER,
			product TEXT,
			amount REAL,
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// 插入用户数据
	for i := 1; i <= 100; i++ {
		_, err = db.Exec(
			"INSERT INTO users (id, name, email) VALUES (?, ?, ?)",
			i, fmt.Sprintf("User%d", i), fmt.Sprintf("user%d@example.com", i),
		)
		if err != nil {
			log.Fatal(err)
		}

		// 为每个用户插入10个订单
		for j := 1; j <= 10; j++ {
			_, err = db.Exec(
				"INSERT INTO orders (user_id, product, amount) VALUES (?, ?, ?)",
				i, fmt.Sprintf("Product%d", j), float64(j)*10.0,
			)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// 方法1: 为每个用户单独查询订单
func queryMethod1(db *sql.DB, userCount int) []map[string]interface{} {
	var results []map[string]interface{}

	// 获取所有用户
	rows, err := db.Query("SELECT id, name, email FROM users LIMIT ?", userCount)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 处理每个用户
	for rows.Next() {
		var userId int
		var name, email string
		err = rows.Scan(&userId, &name, &email)
		if err != nil {
			log.Fatal(err)
		}

		// 为每个用户查询订单
		orderRows, err := db.Query("SELECT id, product, amount FROM orders WHERE user_id = ?", userId)
		if err != nil {
			log.Fatal(err)
		}

		// 处理用户的订单
		for orderRows.Next() {
			var orderId int
			var product string
			var amount float64
			err = orderRows.Scan(&orderId, &product, &amount)
			if err != nil {
				log.Fatal(err)
			}

			// 添加到结果集
			results = append(results, map[string]interface{}{
				"user_id":    userId,
				"user_name":  name,
				"user_email": email,
				"order_id":   orderId,
				"product":    product,
				"amount":     amount,
			})
		}
		orderRows.Close()
	}

	return results
}

// 方法2: 使用JOIN查询
func queryMethod2(db *sql.DB, userCount int) []map[string]interface{} {
	var results []map[string]interface{}

	// 使用JOIN一次性获取所有数据
	query := `
		SELECT u.id, u.name, u.email, o.id, o.product, o.amount
		FROM users u
		JOIN orders o ON u.id = o.user_id
		WHERE u.id <= ?
	`

	rows, err := db.Query(query, userCount)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 处理结果
	for rows.Next() {
		var userId, orderId int
		var name, email, product string
		var amount float64

		err = rows.Scan(&userId, &name, &email, &orderId, &product, &amount)
		if err != nil {
			log.Fatal(err)
		}

		// 添加到结果集
		results = append(results, map[string]interface{}{
			"user_id":    userId,
			"user_name":  name,
			"user_email": email,
			"order_id":   orderId,
			"product":    product,
			"amount":     amount,
		})
	}

	return results
}

// 方法3: 使用预处理语句
func queryMethod3(db *sql.DB, userCount int) []map[string]interface{} {
	var results []map[string]interface{}

	// 准备JOIN查询语句
	query := `
		SELECT u.id, u.name, u.email, o.id, o.product, o.amount
		FROM users u
		JOIN orders o ON u.id = o.user_id
		WHERE u.id <= ?
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// 执行预处理语句
	rows, err := stmt.Query(userCount)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// 处理结果
	for rows.Next() {
		var userId, orderId int
		var name, email, product string
		var amount float64

		err = rows.Scan(&userId, &name, &email, &orderId, &product, &amount)
		if err != nil {
			log.Fatal(err)
		}

		// 添加到结果集
		results = append(results, map[string]interface{}{
			"user_id":    userId,
			"user_name":  name,
			"user_email": email,
			"order_id":   orderId,
			"product":    product,
			"amount":     amount,
		})
	}

	return results
}
