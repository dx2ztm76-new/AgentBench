package code_smells

import (
	"fmt"
	"math"
)

// 计算圆形区域的面积
func CalculateCircleArea(radius float64) float64 {
	// 验证半径是否为正数
	if radius <= 0 {
		fmt.Println("错误: 半径必须为正数")
		return 0
	}

	// 计算面积
	area := math.Pi * radius * radius

	// 打印结果
	fmt.Printf("圆形面积: %.2f\n", area)

	return area
}

// 计算圆形区域的周长
func CalculateCirclePerimeter(radius float64) float64 {
	// 验证半径是否为正数
	if radius <= 0 {
		fmt.Println("错误: 半径必须为正数")
		return 0
	}

	// 计算周长
	perimeter := 2 * math.Pi * radius

	// 打印结果
	fmt.Printf("圆形周长: %.2f\n", perimeter)

	return perimeter
}

// 计算矩形区域的面积
func CalculateRectangleArea(width, height float64) float64 {
	// 验证宽度是否为正数
	if width <= 0 {
		fmt.Println("错误: 宽度必须为正数")
		return 0
	}

	// 验证高度是否为正数
	if height <= 0 {
		fmt.Println("错误: 高度必须为正数")
		return 0
	}

	// 计算面积
	area := width * height

	// 打印结果
	fmt.Printf("矩形面积: %.2f\n", area)

	return area
}

// 计算矩形区域的周长
func CalculateRectanglePerimeter(width, height float64) float64 {
	// 验证宽度是否为正数
	if width <= 0 {
		fmt.Println("错误: 宽度必须为正数")
		return 0
	}

	// 验证高度是否为正数
	if height <= 0 {
		fmt.Println("错误: 高度必须为正数")
		return 0
	}

	// 计算周长
	perimeter := 2 * (width + height)

	// 打印结果
	fmt.Printf("矩形周长: %.2f\n", perimeter)

	return perimeter
}

// 计算三角形区域的面积
func CalculateTriangleArea(a, b, c float64) float64 {
	// 验证边长是否为正数
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("错误: 边长必须为正数")
		return 0
	}

	// 验证三角形是否有效
	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("错误: 不满足三角形条件")
		return 0
	}

	// 使用海伦公式计算面积
	s := (a + b + c) / 2
	area := math.Sqrt(s * (s - a) * (s - b) * (s - c))

	// 打印结果
	fmt.Printf("三角形面积: %.2f\n", area)

	return area
}

// 计算三角形区域的周长
func CalculateTrianglePerimeter(a, b, c float64) float64 {
	// 验证边长是否为正数
	if a <= 0 || b <= 0 || c <= 0 {
		fmt.Println("错误: 边长必须为正数")
		return 0
	}

	// 验证三角形是否有效
	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("错误: 不满足三角形条件")
		return 0
	}

	// 计算周长
	perimeter := a + b + c

	// 打印结果
	fmt.Printf("三角形周长: %.2f\n", perimeter)

	return perimeter
}
