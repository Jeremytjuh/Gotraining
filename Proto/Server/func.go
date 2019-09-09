package main

// import "login/LoginTest"

func Login (username, password string) bool {
	if username == "Admin" && password == "Root" {
		return true
	} else {
		return false
	}
}