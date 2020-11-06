package main

import (
	"fmt"
	"strings"
)

type User struct {
	priority int64
	UserName string
	UserPurchase int64
	nextUser, previousUser *User
}


type UserList struct {
	Users []User
	FirstUser User
	LastUser User
}

///////////Implementing the Bubble Sort
//////////Sorted by priority
func (sort UserList) Sort() {
	for index := 0; index < len(sort.Users); index++ {
		for newIndex := 1 ; newIndex < len(sort.Users); newIndex++ {
			if sort.Users[newIndex - 1].priority > sort.Users[newIndex].priority {
				value := sort.Users[newIndex - 1]
				sort.Users[newIndex - 1] = sort.Users[newIndex]
				sort.Users[newIndex] = value
			}
		}
	}
	for newUser := 0; newUser  < len(sort.Users)-1; newUser ++ {
		sort.Users[newUser].nextUser = &sort.Users[newUser +1]
	}
	for newUser  := 1; newUser  < len(sort.Users); newUser ++ {
		sort.Users[newUser].previousUser = &sort.Users[newUser -1]
	}
	sort.FirstUser = sort.Users[0]
	sort.LastUser = sort.Users[len(sort.Users)-1]
}
///////Sorted by Purchase
func (sortPurchase UserList) SortPurchase() {
	for index := 0; index < len(sortPurchase.Users); index++ {
		for newIndex := 1 ; newIndex < len(sortPurchase.Users); newIndex++ {
			if sortPurchase.Users[newIndex - 1].UserPurchase < sortPurchase.Users[newIndex].UserPurchase {
				value := sortPurchase.Users[newIndex - 1]
				sortPurchase.Users[newIndex - 1] = sortPurchase.Users[newIndex]
				sortPurchase.Users[newIndex] = value
			}
		}
	}
	for newUser := 0; newUser  < len(sortPurchase.Users)-1; newUser ++ {
		sortPurchase.Users[newUser].nextUser = &sortPurchase.Users[newUser +1]
	}
	for newUser  := 1; newUser  < len(sortPurchase.Users); newUser ++ {
		sortPurchase.Users[newUser].previousUser = &sortPurchase.Users[newUser -1]
	}
	sortPurchase.FirstUser = sortPurchase.Users[0]
	sortPurchase.LastUser = sortPurchase.Users[len(sortPurchase.Users)-1]
}





/////AddNewUser
func (AddUser *UserList) AddNewUser(newValue *[]User, newUser User) {
	if AddUser.Users != nil {
		AddUser.Sort()
		*newValue = append(*newValue, newUser)
	}
	AddUser.Sort()
}


/////////////PopFunction
/////////////////////////////
func (del *UserList) Delete() {
	if del.Users != nil {
		del.Users = del.Users[1:]
		del.FirstUser.previousUser = nil
	}

	del.Sort()
}



func main() {
	dashes := strings.Repeat("-", 280)
	Userq := UserList{
		Users: []User{
			{
				UserPurchase: 200,
				UserName: "Ansor",
				priority: 15,
			},
			{
				UserPurchase: 152,
				UserName: "Siyovush",
				priority: 38,
			},
			{
				UserPurchase: 800,
				UserName: "Khurshed",
				priority: 90,
			},
			{
				UserPurchase: 923,
				UserName: "Ramazon",
				priority: 2,
			},
		},
		FirstUser: User{},
		LastUser: User{},
	}

	fmt.Println(dashes)
	Userq.Sort()
	fmt.Println(Userq)
	//////////////////////////////////////////////


	////New User!!!

	AddAgainUser := User{
		priority:     1,
		UserName:     "Mahmud",
		UserPurchase: 15200,
		nextUser:         nil,
		previousUser:     nil,
	}
	fmt.Println(dashes)
	Userq.AddNewUser(&Userq.Users, AddAgainUser)
	Userq.Sort()
	fmt.Println(Userq)

	/////////////////////////Delete
	fmt.Println(dashes)
	Userq.Delete()
	fmt.Println(&Userq.Users) ///We've deleted


	///Finally Sorted by UserPurchase
	fmt.Println(dashes)
	fmt.Println("--------------Sorted by Purchase--------------")
	Userq.SortPurchase()
	fmt.Println(&Userq.Users) ///Sorted by Purchase
}

