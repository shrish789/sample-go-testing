package testing

import (
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// MockMyDependency is a mock implementation of MyDependency.
type MockMyDependency struct {
    mock.Mock
}

// GetMessage is the mocked implementation of GetMessage for MockMyDependency.
func (m *MockMyDependency) GetMessage() (string, error) {
    args := m.Called()
    return args.String(0), args.Error(1)
}

func TestMyFunction(t *testing.T) {
    // Create a new instance of the mock dependency.
    mockDependency := new(MockMyDependency)

    // Define the expected behavior of the mock dependency.
    mockDependency.On("GetMessage").Return("Mocked message", nil)

    // Call the function we want to test.
    result, err := MyFunction(mockDependency)

    // Assert that the function returned the expected result.
    assert.Equal(t, "Mocked message", result)

    // Assert that the function did not return an error.
    assert.NoError(t, err)

    // Assert that the expected method was called on the mock.
    mockDependency.AssertCalled(t, "GetMessage")
}
