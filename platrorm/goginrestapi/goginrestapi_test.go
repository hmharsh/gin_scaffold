package goginrestapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	restApi := New()
	restApi.Add(Item{})
	assert.Equal(t, len(restApi.Items), 1, "Item was not added")

}

func TestGetAll(t *testing.T) {
	restApi := New()
	restApi.Add(Item{})
	results := restApi.GetAll()
	assert.Equal(t, len(results), 1, "Item was not added")

}
