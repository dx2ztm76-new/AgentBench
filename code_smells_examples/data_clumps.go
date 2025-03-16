package code_smells

import (
	"fmt"
	"math"
)

// 计算两点之间的距离
func CalculateDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// 计算三角形面积
func CalculateTriangleArea(x1, y1, x2, y2, x3, y3 float64) float64 {
	// 使用海伦公式计算
	a := CalculateDistance(x1, y1, x2, y2)
	b := CalculateDistance(x2, y2, x3, y3)
	c := CalculateDistance(x3, y3, x1, y1)

	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

// 判断点是否在矩形内
func IsPointInRectangle(px, py, x1, y1, x2, y2 float64) bool {
	return px >= x1 && px <= x2 && py >= y1 && py <= y2
}

// 计算矩形面积
func CalculateRectangleArea(x1, y1, x2, y2 float64) float64 {
	width := math.Abs(x2 - x1)
	height := math.Abs(y2 - y1)
	return width * height
}

// 计算矩形周长
func CalculateRectanglePerimeter(x1, y1, x2, y2 float64) float64 {
	width := math.Abs(x2 - x1)
	height := math.Abs(y2 - y1)
	return 2 * (width + height)
}

// 判断两个矩形是否相交
func DoRectanglesIntersect(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 float64) bool {
	// 检查一个矩形是否在另一个矩形的完全左侧、右侧、上方或下方
	if ax2 < bx1 || ax1 > bx2 || ay2 < by1 || ay1 > by2 {
		return false
	}
	return true
}

// 计算两个矩形相交的面积
func CalculateIntersectionArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 float64) float64 {
	// 如果不相交，返回0
	if !DoRectanglesIntersect(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2) {
		return 0
	}

	// 计算相交矩形的坐标
	ix1 := math.Max(ax1, bx1)
	iy1 := math.Max(ay1, by1)
	ix2 := math.Min(ax2, bx2)
	iy2 := math.Min(ay2, by2)

	// 计算相交矩形的面积
	return CalculateRectangleArea(ix1, iy1, ix2, iy2)
}

// 计算点到线段的最短距离
func PointToLineDistance(px, py, x1, y1, x2, y2 float64) float64 {
	// 线段长度的平方
	lineLength := CalculateDistance(x1, y1, x2, y2)

	if lineLength == 0 {
		// 如果线段长度为0（即点），则直接计算点到点的距离
		return CalculateDistance(px, py, x1, y1)
	}

	// 计算点到线段的投影比例
	t := ((px-x1)*(x2-x1) + (py-y1)*(y2-y1)) / (lineLength * lineLength)

	if t < 0 {
		// 投影点在线段外，最近点是线段的起点
		return CalculateDistance(px, py, x1, y1)
	}
	if t > 1 {
		// 投影点在线段外，最近点是线段的终点
		return CalculateDistance(px, py, x2, y2)
	}

	// 投影点在线段内，计算点到投影点的距离
	projX := x1 + t*(x2-x1)
	projY := y1 + t*(y2-y1)
	return CalculateDistance(px, py, projX, projY)
}

// 计算多边形面积
func CalculatePolygonArea(points [][2]float64) float64 {
	n := len(points)
	if n < 3 {
		return 0 // 不是多边形
	}

	area := 0.0
	for i := 0; i < n; i++ {
		j := (i + 1) % n
		area += points[i][0] * points[j][1]
		area -= points[j][0] * points[i][1]
	}

	area = math.Abs(area) / 2.0
	return area
}

// 判断点是否在多边形内
func IsPointInPolygon(px, py float64, points [][2]float64) bool {
	n := len(points)
	if n < 3 {
		return false // 不是多边形
	}

	inside := false
	j := n - 1

	for i := 0; i < n; i++ {
		xi, yi := points[i][0], points[i][1]
		xj, yj := points[j][0], points[j][1]

		intersect := ((yi > py) != (yj > py)) &&
			(px < (xj-xi)*(py-yi)/(yj-yi)+xi)

		if intersect {
			inside = !inside
		}

		j = i
	}

	return inside
}

// 绘制矩形
func DrawRectangle(x1, y1, x2, y2 float64) string {
	width := math.Abs(x2 - x1)
	height := math.Abs(y2 - y1)

	result := fmt.Sprintf("绘制矩形：\n")
	result += fmt.Sprintf("  左上角坐标: (%.2f, %.2f)\n", x1, y1)
	result += fmt.Sprintf("  右下角坐标: (%.2f, %.2f)\n", x2, y2)
	result += fmt.Sprintf("  宽度: %.2f\n", width)
	result += fmt.Sprintf("  高度: %.2f\n", height)
	result += fmt.Sprintf("  面积: %.2f\n", CalculateRectangleArea(x1, y1, x2, y2))
	result += fmt.Sprintf("  周长: %.2f\n", CalculateRectanglePerimeter(x1, y1, x2, y2))

	return result
}

// 绘制圆形
func DrawCircle(centerX, centerY, radius float64) string {
	result := fmt.Sprintf("绘制圆形：\n")
	result += fmt.Sprintf("  中心坐标: (%.2f, %.2f)\n", centerX, centerY)
	result += fmt.Sprintf("  半径: %.2f\n", radius)
	result += fmt.Sprintf("  面积: %.2f\n", math.Pi*radius*radius)
	result += fmt.Sprintf("  周长: %.2f\n", 2*math.Pi*radius)

	return result
}

// 移动矩形
func MoveRectangle(x1, y1, x2, y2, deltaX, deltaY float64) (float64, float64, float64, float64) {
	return x1 + deltaX, y1 + deltaY, x2 + deltaX, y2 + deltaY
}

// 缩放矩形
func ScaleRectangle(x1, y1, x2, y2, scaleX, scaleY float64) (float64, float64, float64, float64) {
	centerX := (x1 + x2) / 2
	centerY := (y1 + y2) / 2

	halfWidth := math.Abs(x2-x1) / 2
	halfHeight := math.Abs(y2-y1) / 2

	newHalfWidth := halfWidth * scaleX
	newHalfHeight := halfHeight * scaleY

	return centerX - newHalfWidth, centerY - newHalfHeight, centerX + newHalfWidth, centerY + newHalfHeight
}

// 旋转点（围绕原点）
func RotatePoint(x, y, angleDegrees float64) (float64, float64) {
	angleRadians := angleDegrees * math.Pi / 180.0

	cosAngle := math.Cos(angleRadians)
	sinAngle := math.Sin(angleRadians)

	return x*cosAngle - y*sinAngle, x*sinAngle + y*cosAngle
}

// 旋转矩形（围绕中心点）
func RotateRectangle(x1, y1, x2, y2, angleDegrees float64) [][2]float64 {
	centerX := (x1 + x2) / 2
	centerY := (y1 + y2) / 2

	// 矩形的四个角点
	corners := [][2]float64{
		{x1, y1}, // 左上
		{x2, y1}, // 右上
		{x2, y2}, // 右下
		{x1, y2}, // 左下
	}

	// 旋转每个角点
	rotatedCorners := make([][2]float64, 4)
	for i, corner := range corners {
		// 将点平移到原点
		translatedX := corner[0] - centerX
		translatedY := corner[1] - centerY

		// 旋转点
		rotatedX, rotatedY := RotatePoint(translatedX, translatedY, angleDegrees)

		// 将点平移回原位置
		rotatedCorners[i] = [2]float64{rotatedX + centerX, rotatedY + centerY}
	}

	return rotatedCorners
}
