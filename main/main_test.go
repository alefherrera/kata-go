package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_sum(t *testing.T) {

	assert.Equal(t, 3, sum(1, 2))

}
