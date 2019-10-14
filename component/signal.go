package component

// 方向
type Direction int

// 电荷量, 电量 : 物体所带电荷的量值 (库仑, C)
// ---------- ----------
//
// ---------- ----------
const (
	EC = 1.0 // 基本电荷
)

type ElectricCharge struct {
	n int // 电荷量整数倍
}

func (e *ElectricCharge) GetEC() float64 {
	return float64(e.n) * EC
}

// 针脚信号 [电信号]
// ---------- ----------
type Signal struct {
	ElectricCharge // 电荷
}

func NewSignal() *Signal {
	return &Signal{
		ElectricCharge: ElectricCharge{},
	}
}
