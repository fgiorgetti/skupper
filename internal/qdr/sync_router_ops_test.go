package qdr

import (
	"errors"
	"testing"

	"gotest.tools/v3/assert"
)

type fakeBridgeConfigAgent struct {
	getResponses []*BridgeConfig
	getErrAt     int
	getCalls     int
	updateCalls  int
	updateErr    error
	lastDiff     *BridgeConfigDifference
}

func (f *fakeBridgeConfigAgent) GetLocalBridgeConfig() (*BridgeConfig, error) {
	f.getCalls++
	if f.getErrAt == f.getCalls {
		return nil, errors.New("get failed")
	}
	idx := f.getCalls - 1
	if idx >= len(f.getResponses) {
		idx = len(f.getResponses) - 1
	}
	return f.getResponses[idx], nil
}

func (f *fakeBridgeConfigAgent) UpdateLocalBridgeConfig(changes *BridgeConfigDifference) error {
	f.updateCalls++
	f.lastDiff = changes
	return f.updateErr
}

func TestSyncBridgeConfigNoChanges(t *testing.T) {
	bridge := NewBridgeConfig()
	bridge.AddTcpListener(TcpEndpoint{Name: "web", Address: "web:80", Host: "0.0.0.0", Port: "80"})

	agent := &fakeBridgeConfigAgent{
		getResponses: []*BridgeConfig{&bridge},
	}

	actual, synced, err := syncBridgeConfig(agent, &bridge, true)
	assert.NilError(t, err)
	assert.Assert(t, synced)
	assert.Equal(t, agent.getCalls, 1)
	assert.Equal(t, agent.updateCalls, 0)
	assert.Equal(t, len(actual.TcpListeners), 1)
	assert.Equal(t, actual.TcpListeners["web"].Port, "80")
}

func TestSyncBridgeConfigUpdatesAndRefetches(t *testing.T) {
	initial := NewBridgeConfig()
	desired := NewBridgeConfig()
	desired.AddTcpListener(TcpEndpoint{Name: "api", Address: "api:8080", Host: "0.0.0.0", Port: "8080"})
	desired.AddTcpConnector(TcpEndpoint{Name: "db", Address: "db:5432", Host: "10.0.0.2", Port: "5432"})

	updated := NewBridgeConfig()
	updated.AddTcpListener(TcpEndpoint{Name: "api", Address: "api:8080", Host: "0.0.0.0", Port: "8080"})
	updated.AddTcpConnector(TcpEndpoint{Name: "db", Address: "db:5432", Host: "10.0.0.2", Port: "5432"})
	agent := &fakeBridgeConfigAgent{
		getResponses: []*BridgeConfig{&initial, &updated},
	}

	actual, synced, err := syncBridgeConfig(agent, &desired, true)
	assert.NilError(t, err)
	assert.Assert(t, !synced)
	assert.Equal(t, agent.getCalls, 2)
	assert.Equal(t, agent.updateCalls, 1)
	assert.Equal(t, len(actual.TcpListeners), 1)
	assert.Equal(t, len(actual.TcpConnectors), 1)
	assert.Equal(t, actual.TcpListeners["api"].Address, "api:8080")
	assert.Equal(t, actual.TcpConnectors["db"].Host, "10.0.0.2")
}

func TestSyncBridgeConfigSkipsRefetchWhenNotRequested(t *testing.T) {
	initial := NewBridgeConfig()
	desired := NewBridgeConfig()
	desired.AddTcpListener(TcpEndpoint{Name: "api", Address: "api:8080", Host: "0.0.0.0", Port: "8080"})

	agent := &fakeBridgeConfigAgent{
		getResponses: []*BridgeConfig{&initial},
	}

	actual, synced, err := syncBridgeConfig(agent, &desired, false)
	assert.NilError(t, err)
	assert.Assert(t, !synced)
	assert.Assert(t, actual == nil)
	assert.Equal(t, agent.getCalls, 1)
	assert.Equal(t, agent.updateCalls, 1)
}

func TestSyncBridgeConfigErrors(t *testing.T) {
	initial := NewBridgeConfig()
	desired := NewBridgeConfig()
	desired.AddTcpListener(TcpEndpoint{Name: "api", Address: "api:8080", Host: "0.0.0.0", Port: "8080"})

	t.Run("initial get fails", func(t *testing.T) {
		agent := &fakeBridgeConfigAgent{getErrAt: 1}
		_, _, err := syncBridgeConfig(agent, &desired, true)
		assert.ErrorContains(t, err, "Error retrieving bridges")
	})

	t.Run("update fails", func(t *testing.T) {
		agent := &fakeBridgeConfigAgent{
			getResponses: []*BridgeConfig{&initial},
			updateErr:    errors.New("update boom"),
		}
		_, _, err := syncBridgeConfig(agent, &desired, true)
		assert.ErrorContains(t, err, "Error syncing bridges")
	})

	t.Run("post-update get fails", func(t *testing.T) {
		agent := &fakeBridgeConfigAgent{
			getResponses: []*BridgeConfig{&initial},
			getErrAt:     2,
		}
		_, _, err := syncBridgeConfig(agent, &desired, true)
		assert.ErrorContains(t, err, "Error retrieving updated bridges")
	})
}

type fakeBridgeRuntimeAgent struct {
	fakeBridgeConfigAgent
	runtimeState *BridgeRuntimeState
	runtimeCalls int
	runtimeErr   error
}

func (f *fakeBridgeRuntimeAgent) GetLocalBridgeRuntimeState() (*BridgeRuntimeState, error) {
	f.runtimeCalls++
	if f.runtimeErr != nil {
		return nil, f.runtimeErr
	}
	return f.runtimeState, nil
}

func TestSyncBridgeRuntimeStateAndRecordMapping(t *testing.T) {
	rec := Record{
		"name":              "web",
		"address":           "web:80",
		"host":              "0.0.0.0",
		"port":              "80",
		"operStatus":        "up",
		"connectionsOpened": uint64(12),
		"connectionsClosed": uint64(10),
		"bytesIn":           uint64(4096),
		"bytesOut":          uint64(8192),
	}
	rt := asTcpEndpointRuntime(rec)
	assert.Equal(t, rt.Name, "web")
	assert.Equal(t, rt.OperStatus, "up")
	assert.Equal(t, rt.ConnectionsOpened, uint64(12))
	assert.Equal(t, rt.ConnectionsClosed, uint64(10))
	assert.Equal(t, rt.BytesIn, uint64(4096))
	assert.Equal(t, rt.BytesOut, uint64(8192))

	initial := NewBridgeConfig()
	desired := NewBridgeConfig()
	desired.AddTcpListener(TcpEndpoint{Name: "web", Address: "web:80", Host: "0.0.0.0", Port: "80"})

	live := NewBridgeRuntimeState()
	live.TcpListeners["web"] = rt

	agent := &fakeBridgeRuntimeAgent{
		fakeBridgeConfigAgent: fakeBridgeConfigAgent{
			getResponses: []*BridgeConfig{&initial},
		},
		runtimeState: &live,
	}

	got, synced, err := syncBridgeRuntimeState(agent, &desired)
	assert.NilError(t, err)
	assert.Assert(t, !synced)
	assert.Equal(t, agent.updateCalls, 1)
	assert.Equal(t, agent.runtimeCalls, 1)
	assert.Equal(t, got.TcpListeners["web"].OperStatus, "up")
	assert.Equal(t, got.TcpListeners["web"].BytesIn, uint64(4096))
}
