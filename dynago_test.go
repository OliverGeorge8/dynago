package dynago

import (
	"fmt"
	"testing"
)

type User struct {
	*Model
}

func (u User) SetPK(id string) {
	//TODO implement me
	panic("implement me")
}

func (u User) SetSK(id string) {
	//TODO implement me
	panic("implement me")
}

func (User) SkPrefix() string {
	return "USER#"
}

func (User) PkPrefix() string {
	return "USER#"
}

func TestDynago(t *testing.T) {

	dynago, err := Initialize()

	if err != nil {
		fmt.Println(err)
	}
	key := Key("Saf", nil)
	var user User
	err = dynago.GetItem(key, &user)
	if err != nil {
		fmt.Println(err)
	}

	err = dynago.PutItem(&User{})
	if err != nil {
		return
	}

	update := Update().Set("email", "oliver@gmail.com")
	updatedUser := &User{}
	_, err = dynago.UpdateItem(Key("USER#", nil), update, updatedUser)
	if err != nil {
		fmt.Println(err)

	}

	keyCond := KeyConditions().EqualTo("PK", "USER#"+updatedUser.PK).EqualTo("SK", updatedUser.SK)

	gotUser := &User{}
	err = dynago.GetItem(Key(updatedUser.PK, updatedUser.SK), gotUser)
	if err != nil {
		fmt.Println(err)

	}

	expr := NewExpr().WithKeyConditions(keyCond)
	var models []User
	err = dynago.Query(expr, &models)
	if err != nil {
		fmt.Println(err)
	}
}
