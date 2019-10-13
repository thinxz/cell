package man

import (
	"../component"
)

type Manager struct {
	// 器件管理器, 管理所有器件对象
	enManager *component.EnManager
}

func (m *Manager) Start() {
	// 初始化实例管理器
	m.enManager = &component.EnManager{}
	m.enManager.Init()
	// 启动
	m.enManager.Run()
}
