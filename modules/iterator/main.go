package main

import "fmt"

// Итератор позволяет клиенту обходить разные коллекции одним и тем же способом,
// используя единый интерфейс итераторов (без раскрытия деталей реализации коллекции).

type Collection interface {
	createIterator() Iterator
}

type UserCollection struct {
	users []*User
}

func (c *UserCollection) createIterator() Iterator {
	return &UserIterator{
		collection: c,
	}
}

type Iterator interface {
	hasNext() bool
	getNext() *User
}

type UserIterator struct {
	index      int
	collection *UserCollection
}

func (u *UserIterator) hasNext() bool {
	if u.index < len(u.collection.users) {
		return true
	}
	return false

}
func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.collection.users[u.index]
		u.index++
		return user
	}
	return nil
}

type User struct {
	name string
	age  int
}

func main() {

	user1 := &User{
		name: "a",
		age:  30,
	}
	user2 := &User{
		name: "b",
		age:  20,
	}

	userCollection := &UserCollection{
		users: []*User{user1, user2},
	}

	iterator := userCollection.createIterator()

	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
