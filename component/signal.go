package component

import "encoding/json"

// 方向
type Direction int

// 电荷量, 电量 : 物体所带电荷的量值 (库仑, C)
// ---------- ----------
//
// ---------- ----------
const (
	EC   = 1.0 // 基本电荷
	Time = 1   // 时间单位(秒)
)

type ElectricCharge struct {
	N int64   `json:"n"` // 电荷量整数倍
	T float64 `json:"t"` // 传输时间间隔
}

func NewElectricCharge(n int64) *ElectricCharge {
	return &ElectricCharge{
		N: n,
	}
}

func (e *ElectricCharge) GetEC() float64 {
	return float64(e.N) * EC
}

// 数据转化为JSON字符串
func (e ElectricCharge) ToJson() string {
	buf, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(buf)
}

// JSON 字符串转化为对象
func (e *ElectricCharge) ByJson(data string) bool {
	err := json.Unmarshal([]byte(data), e)
	if err != nil {
		return false
	}
	return true
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
