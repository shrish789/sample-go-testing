package services

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

type MockMyDependency struct {
    mock.Mock
}

func (m *MockMyDependency) Example() bool {
    args := m.Called()
    return args.Bool(0)
}

func TestTrueExample(t *testing.T) {
	mockMyDependency := new(MockMyDependency)

	mockMyDependency.On("Example").Return(true)

	w := &Workflow{
		Db: mockMyDependency,
	}
	result := w.Example()

	assert.Equal(t, "Hello World! Example function", result)
}

func TestFalseExample(t *testing.T) {
	mockMyDependency := new(MockMyDependency)

	mockMyDependency.On("Example").Return(false)

	w := &Workflow{
		Db: mockMyDependency,
	}
	result := w.Example()

	assert.Equal(t, "Yo!", result)
}
