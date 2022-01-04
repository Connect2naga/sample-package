// Package chat_interface contains ...
package chat_interface

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Author : Nagarjuna S
Date : 1/4/22 11:01 AM
Project : sample-package
File : chatting_test.go
*/

func TestSayHello(t *testing.T) {
	assert.Equal(t, "Say hello to Nagarjuna",SayHello("Nagarjuna"))
}