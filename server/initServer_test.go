package server

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/keshu12345/notes/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"go.uber.org/fx"
	"go.uber.org/fx/fxtest"
)

func TestInitializeServer(t *testing.T) {
	router := gin.Default()
	cfg := &config.Configuration{
		Server: config.Server{
			RestServicePort: 8080,
			ReadTimeout:     10,
			WriteTimeout:    10,
			IdleTimeout:     10,
		},
	}

	lifecycle := fxtest.NewLifecycle(t)
	InitializeServer(router, cfg, lifecycle)

	lifecycle.RequireStart()

	lifecycle.RequireStop()
}

type MockLifecycle struct {
	mock.Mock
}

func (m *MockLifecycle) Append(hook fx.Hook) {
	m.Called(hook)
}

type MockServer struct {
	mock.Mock
}

func (m *MockServer) ListenAndServe() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockServer) Shutdown(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestInitServer(t *testing.T) {

	router := gin.Default()
	cfg := &config.Configuration{
		Server: config.Server{
			RestServicePort: 8080,
			ReadTimeout:     10,
			WriteTimeout:    10,
			IdleTimeout:     10,
		},
	}
	mockLifecycle := new(MockLifecycle)
	mockServer := new(MockServer)

	mockLifecycle.On("Append", mock.Anything).Run(func(args mock.Arguments) {
		hook := args.Get(0).(fx.Hook)
		assert.NotNil(t, hook.OnStart)
		assert.NotNil(t, hook.OnStop)

		err := hook.OnStart(context.Background())
		assert.NoError(t, err)
		err = hook.OnStop(context.Background())
		assert.NoError(t, err)
	})

	InitializeServer(router, cfg, mockLifecycle)

	mockLifecycle.AssertExpectations(t)
	mockServer.AssertExpectations(t)
}

type testCase struct {
	name      string
	testCases func(*testing.T)
}

func Test_InitServer(t *testing.T) {

	tests := []testCase{
		{
			name: "CASE SUCCESS",

			testCases: func(t *testing.T) {
				router := gin.Default()
				cfg := &config.Configuration{
					Server: config.Server{
						RestServicePort: 8080,
						ReadTimeout:     10,
						WriteTimeout:    10,
						IdleTimeout:     10,
					},
				}
				lifecycle := fxtest.NewLifecycle(t)

				InitializeServer(router, cfg, lifecycle)
				lifecycle.RequireStart()
				lifecycle.RequireStop()
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, test.testCases)
	}
}
