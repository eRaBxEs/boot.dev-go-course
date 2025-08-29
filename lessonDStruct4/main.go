package main

import (
	"fmt"
)

type authenticationInfo struct {
	username string
	password string
}

// create the method below
// method: defined struct with its own function
// where (authI authenticationInfo) is the receiver struct
// authI is the placeholder
func (authI authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		authI.username, authI.password,
	)
}

func test(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("==========================================")
}

func main() {
	auth1 := authenticationInfo{
		username: "Bing",
		password: "9872356",
	}

	auth2 := authenticationInfo{
		username: "BDG",
		password: "563774",
	}

	test(auth1)
	test(auth2)
}
