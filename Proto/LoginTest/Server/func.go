package main

import pb "github.com/Jeremytjuh/GoTraining/Proto/LoginTest/LoginTest"

// Login used to login
func Login(m *pb.LoginRequest) bool {
	var tof bool
	if m.Username == "Admin" && m.Password == "Root" {
		tof = true
	} else {
		tof = false
	}
	return tof
}
