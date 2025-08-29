package main

import (
	"fmt"
	"reflect"
)

// fields of a struct should be arranged in such a way that the fields that use more size should be at the top level

type stats struct {
	Reach    uint16
	NumPosts uint8
	NumLikes uint8
}

type stats2 struct { // This ordering uses up more memory
	NumPosts uint8
	Reach    uint16
	NumLikes uint8
}

type contact struct {
	sendingLimit int32
	userID       string
	age          int32
}

// re-arranged struct
type contact2 struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	canSend         bool
	canReceive      bool
	permissionLevel int
	canManage       bool
}

type perms2 struct {
	canSend         bool
	canReceive      bool
	canManage       bool
	permissionLevel int
}

func main() {
	typ := reflect.TypeOf(stats{})
	fmt.Printf("Struct is %d bytes\n", typ.Size())
	fmt.Println("=============================================")

	typ2 := reflect.TypeOf(stats2{})
	fmt.Printf("Struct is %d bytes\n", typ2.Size())
	fmt.Println("=============================================")

	typ3 := reflect.TypeOf(contact{})
	fmt.Printf("Struct is %d bytes\n", typ3.Size())
	fmt.Println("=============================================")

	typ4 := reflect.TypeOf(contact2{})
	fmt.Printf("Struct is %d bytes\n", typ4.Size())
	fmt.Println("=============================================")

	typ5 := reflect.TypeOf(perms{})
	fmt.Printf("Struct is %d bytes\n", typ5.Size())
	fmt.Println("=============================================")

	typ6 := reflect.TypeOf(perms2{})
	fmt.Printf("Struct is %d bytes\n", typ6.Size())
	fmt.Println("=============================================")

}
