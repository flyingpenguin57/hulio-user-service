package utils

import (
	"testing"
)
func TestHashPwd(t *testing.T) {
    password := "123ggg456"
    hashedPwd, err := HashPassword(password)
    if err != nil {
        t.Fatal(err)
    }

    if hashedPwd == "" {
        t.Fatal("hashed password should not be empty")
    }

    t.Log("Hashed password:", hashedPwd) 

	b := CheckPassword("$2a$10$/ScZVG4qyNcjoP3cofdvc.Z0pTv4wwSRgsGrdAx0g.1fr1OMbUtuu", password)

	if !b {
        t.Fatal("password error")
    }

    t.Log("passwor correct") 
}

