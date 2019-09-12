{{Login}}
Login

# RPC Login

Here I have some useful information about the [*Remote Procedure Call*](https://en.wikipedia.org/wiki/Remote_procedure_call "Definition of RPC") **Login**.\
This *RPC* uses the argument LoginRequest and returns a LoginReply.\
Here is a code snippit of the login function with some dummy login credentials:
```go
package main

import pb "login/LoginTest"

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
```
# Gopher
While we're at it, let's add a dancing gopher!\
![alt text](https://camo.githubusercontent.com/c70f18274a81ee98dca1c116b68d5a35847b2e65/687474703a2f2f7374617469632e76656c76657463616368652e6f72672f70616765732f323031382f30362f31332f70617274792d676f706865722f64616e63696e672d676f706865722e676966 "Dancing Gopher")

# Documentation
All of this **documentation** is stored inside of the **documentation.md** file and gets generated together with its corresponding **proto** file.\
![alt text](https://github.com/ "Proto file")
{{end}}
{{end}}