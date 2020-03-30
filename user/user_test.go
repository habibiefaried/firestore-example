package user

import (
	"testing"
	"fmt"
	"strings"
)

func TestInsertandValidateUser(t *testing.T){
	u, err := InitUser("../sdk.json", "Database")
	if (err != nil){
		t.Fatalf("%v\n",err)
	}
	fmt.Println("Successfully loaded!")

	var um UserModel
	um.Fullname = "Habibie Faried"
	um.Email = "habibiefaried@gmail.com"
	um.Password = "Password"

	err = u.InsertOrUpdateData("habibiefaried",um)
	if (err != nil){
		t.Fatalf("%v\n",err)
	}

	user, err := u.GetData("habibiefaried")
	if (err != nil){
		t.Fatalf("%v\n",err)
	}

	if (strings.Compare(user.Fullname, um.Fullname) != 0){
		t.Fatalf("The fullname is not the same value!")
	}
}

func TestNonExistentData(t *testing.T){
	u, err := InitUser("../sdk.json", "Database")
	if (err != nil){
		t.Fatalf("%v\n",err)
	}
	fmt.Println("Successfully loaded!")

	username := "nonexistent!_+$@)(2020+_'"
	_, err = u.GetData(username)

	if (strings.Compare(err.Error(),"Not found") != 0) {
		t.Fatalf("The error 'Not found' should be here for username %v\n",username)
	} 
}