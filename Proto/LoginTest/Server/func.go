package main

import pb "login/LoginTest"

// Login used to login
func Login(m *pb.LoginRequest) bool {
	if m.Username == "Admin" && m.Password == "Root" {
		return true
	} else {
		return false
	}
}
