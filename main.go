package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/odanaraujo/golang/solid/dependency-inversion-principle/good"
)

func main() {
	mySQLRepo := &good.MySQLRepository{}
	services := &good.UserServices{Repository: mySQLRepo}

	name := "Dan Araújo"
	ID := strconv.Itoa(rand.Intn(100000))
	user_inserted := services.Repository.InsertUser(ID, name)
	fmt.Printf("User inserted: %#v \n", user_inserted)
	user_returned := services.Repository.GetUserByID(user_inserted.ID)
	fmt.Printf("User returned: %#v \n", user_returned)

	// user_inserted := bad.InsertUserServices(name)
	// fmt.Printf("User inserted: %#v \n", user_inserted)

	// user_returned := bad.GetUserByIDServices(user_inserted.ID)
	// fmt.Printf("User returned: %#v \n", user_returned)
}
