package resolvers

import (
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func (m *MockConfigTransaction) Transaction(f func(transactionID string) (*http.Response, error)) (*http.Response, error) {
	args := m.Called(f)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *MockConfigResolvers) GetAResolversConfiguration(name string, transactionID string) (*http.Response, error) {
	args := m.Called(name, transactionID)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *MockConfigResolvers) AddResolversConfiguration(payload []byte, transactionID string) (*http.Response, error) {
	args := m.Called(payload, transactionID)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *MockConfigResolvers) UpdateResolversConfiguration(name string, payload []byte, transactionID string) (*http.Response, error) {
	args := m.Called(name, payload, transactionID)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (m *MockConfigResolvers) DeleteResolversConfiguration(name string, transactionID string) (*http.Response, error) {
	args := m.Called(name, transactionID)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestHaproxyResolversGet(t *testing.T) {
	mockTransaction := &MockConfigTransaction{}
	mockResolversConfig := &MockConfigResolvers{}

	// Mock response for GetAResolversConfiguration
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
	}
	mockTransaction.On("Transaction", mock.Anything).Return(mockResponse, nil)

	// Create a test ResourceData with necessary fields
	testData := schema.TestResourceDataRaw(t, ResourceHaproxyResolvers().Schema, map[string]interface{}{
		"name": "test-resolver",
	})

	// Call the function to be tested
	err := resourceHaproxyResolversRead(testData, map[string]interface{}{
		"transaction": mockTransaction,
		"resolvers":   mockResolversConfig,
	})

	// Check if the function produced the expected result
	assert.NoError(t, err)
	assert.Equal(t, "test-resolver", testData.Id())
}

func TestHaproxyResolversCreate(t *testing.T) {
	mockTransaction := &MockConfigTransaction{}
	mockResolversConfig := &MockConfigResolvers{}

	// Mock response for AddResolversConfiguration
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
	}
	mockTransaction.On("Transaction", mock.Anything).Return(mockResponse, nil)

	// Create a test ResourceData with necessary fields
	testData := schema.TestResourceDataRaw(t, ResourceHaproxyResolvers().Schema, map[string]interface{}{
		"name":                  "test-resolver",
		"accepted_payload_size": 1500,
		// Add other fields as needed
	})

	// Call the function to be tested
	err := resourceHaproxyResolversCreate(testData, map[string]interface{}{
		"transaction": mockTransaction,
		"resolvers":   mockResolversConfig,
	})

	// Check if the function produced the expected result
	assert.NoError(t, err)
	assert.Equal(t, "test-resolver", testData.Id())
}

func TestHaproxyResolversUpdate(t *testing.T) {
	mockTransaction := &MockConfigTransaction{}
	mockResolversConfig := &MockConfigResolvers{}

	// Mock response for UpdateResolversConfiguration
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
	}
	mockTransaction.On("Transaction", mock.Anything).Return(mockResponse, nil)

	// Create a test ResourceData with necessary fields
	testData := schema.TestResourceDataRaw(t, ResourceHaproxyResolvers().Schema, map[string]interface{}{
		"name":                  "test-resolver",
		"accepted_payload_size": 1600,
		// Add other fields as needed
	})

	// Call the function to be tested
	err := resourceHaproxyResolversUpdate(testData, map[string]interface{}{
		"transaction": mockTransaction,
		"resolvers":   mockResolversConfig,
	})

	// Check if the function produced the expected result
	assert.NoError(t, err)
	assert.Equal(t, "test-resolver", testData.Id())
}

func TestHaproxyResolversDelete(t *testing.T) {
	mockTransaction := &MockConfigTransaction{}
	mockResolversConfig := &MockConfigResolvers{}

	// Mock response for DeleteResolversConfiguration
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
	}
	mockTransaction.On("Transaction", mock.Anything).Return(mockResponse, nil)

	// Create a test ResourceData with necessary fields
	testData := schema.TestResourceDataRaw(t, ResourceHaproxyResolvers().Schema, map[string]interface{}{
		"name": "test-resolver",
	})

	// Call the function to be tested
	err := resourceHaproxyResolversDelete(testData, map[string]interface{}{
		"transaction": mockTransaction,
		"resolvers":   mockResolversConfig,
	})

	// Check if the function produced the expected result
	assert.NoError(t, err)
	assert.Equal(t, "", testData.Id())
}
