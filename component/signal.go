package component

// 方向
type Direction int

// 源指向目标(对针脚, 从自身指向外部, 同向为正极)
// ---------- ----------
// 直流电路中, 电流电压方向相同
// ---------- ----------
const (
	Forward Direction = 1  // 正方向 [正极 -> 负极]
	Reverse Direction = -1 // 反方向
	Invert            = -1 // 取反[相乘]
)

type I struct {
	Value     float64 // 电流值
	Direction         // 电流方向
}

func (i I) Equal(other I) bool {
	if i.Value == other.Value && i.Direction == other.Direction*Invert {
		return true
	}
	return false
}

func (i I) PowerEqual(other I) bool {
	if i.Value == other.Value && i.Direction == other.Direction*Invert {
		return true
	}
	return false
}

type V struct {
	Value     float64 // 电压值
	Direction         // 电压方向
}

func (v V) Equal(other V) bool {
	if v.Value == other.Value && v.Direction == other.Direction*Invert {
		return true
	}
	return false
}

func (v V) PowerEqual(other V) bool {
	if v.Value == other.Value && v.Direction == other.Direction {
		return true
	}
	return false
}

// 针脚信号定义 [电信号]
// 为 nil 为任意匹配值 (不参与计算)
// ---------- ----------
// 正方向为
type Signal struct {
	V // 电压
	I // 电流
}
