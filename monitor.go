package http

import (
	"github.com/infrago/base"
	"github.com/infrago/infra"
)

func (m *Module) Ready() bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.started && len(m.instances) > 0
}

func (m *Module) Health() infra.ModuleHealth {
	m.mutex.Lock()
	started := m.started
	connections := len(m.instances)
	routers := len(m.routers)
	m.mutex.Unlock()
	return infra.NewModuleHealth("http", started && connections > 0, nil, base.Map{
		"connections": connections,
		"routers":     routers,
	})
}

func (m *Module) Stats() infra.ModuleStats {
	m.mutex.Lock()
	started := m.started
	connections := len(m.instances)
	routers := len(m.routers)
	m.mutex.Unlock()
	return infra.NewModuleStats("http", started && connections > 0, base.Map{
		"connections": connections,
		"routers":     routers,
	})
}
