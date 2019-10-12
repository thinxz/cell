package component

// 针脚信号定义 [电信号]
// 为 nil 为任意匹配值 (不参与计算)
// ---------- ----------
type Signal struct {
	R float32 // 电阻
	V float32 // 电压
	I float32 // 电流
}
