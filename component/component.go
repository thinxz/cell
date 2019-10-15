package component

import (
	"../evt"
	"bytes"
	"fmt"
	"strings"
)

// 器件最大针脚数量
const SMaxNum = 2

// 器件接口定义
// ---------- ---------
type IComponent interface {
	Transmission(event evt.Event)                           // 事件传输到该器件, 触发器发生计算
	GetStitchPoint(name string, no int) (string, int, bool) // 获取器件针脚关联的唯一点
	GetRelation(point string) ([]Relation, bool)            // 获取点关联的所有器件及其对应针脚
	Name() string                                           // 获取器件名称, 唯一标识符
	Describe() string                                       // 器件相关信息描述
}

//  事件传输到对应器件
// ---------- ----------
func (c *Component) Transmission(event evt.Event) {
	fmt.Println(fmt.Sprintf("IComponent [%s] -> not calculate ing .......... ......... ", c.name))
}

// 获取器件针脚, 关联的所有器件
func (c *Component) GetStitchPoint(name string, no int) (string, int, bool) {
	return c.man.GetStitchPoint(name, no)
}

func (c *Component) GetRelation(point string) ([]Relation, bool) {
	return c.man.GetRelation(point)
}

// 器件发布事件
// ---------- ----------
// event    事件类型
// source   发布器件名称
// sourceNo 发布器件针脚
// target   接收器件名称
// targetNo 接收器件针脚
// ---------- ----------
func (c *Component) Event(event, source string, sourceNo int, target string, targetNo int, data ElectricCharge) {
	c.man.PutEvent(event, source, sourceNo, target, targetNo, data.ToJson())
}

// 获取器件名称, 唯一标识符
// ---------- ----------
func (c *Component) Name() string {
	return c.name
}

// 器件描述信息, 返回JSON格式描述
// ---------- ----------
func (c *Component) Describe() string {
	var buff bytes.Buffer
	//for _, v := range c.sts {
	//	var buffer bytes.Buffer
	//	// 针脚编号
	//	buffer.WriteString(fmt.Sprintf("\"%d\":[", v.GetNo()))
	//	//for i := 0; i < len(v.Relation); i++ {
	//	//	// 器件名称:器件针脚编号,
	//	//	buffer.WriteString("\"")
	//	//	buffer.WriteString(v.Relation[i].Component.Name())
	//	//	buffer.WriteString("\"")
	//	//	buffer.WriteString(":")
	//	//	buffer.WriteString(strconv.Itoa(v.Relation[i].No))
	//	//	buffer.WriteString(",")
	//	//}
	//	buffer.WriteString("]")
	//	// 关联器件名称, 关联器件编号
	//	buff.WriteString(fmt.Sprintf("%s],", strings.TrimRight(buffer.String(), ",]")))
	//}
	return fmt.Sprintf("{%s}", strings.TrimRight(buff.String(), ","))
}

/********** ********** ********** ********** ********** ********** ********** ********** **********/

// 器件定义
// ---------- ---------
type Component struct {
	name   string     // 器件名称
	stitch int        // 针脚数量
	man    *EnManager // 器件管理器
}

// 初始化器件数据
// ---------- ----------
// name   器件名称
// stitch 针脚数量
// man    器件管理器
func (c *Component) Init(name string, stitch int, man *EnManager) {
	c.name = name
	c.stitch = stitch
	c.man = man
}
