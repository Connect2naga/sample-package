// Package chat_interface contains ...
package chat_interface

import "fmt"

/*
Author : Nagarjuna S
Date : 1/4/22 11:00 AM
Project : sample-package
File : chatting.go
*/

func SayHello(name string) string {
	return fmt.Sprintf("Say hello to %s", name)
}
