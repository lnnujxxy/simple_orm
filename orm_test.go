package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ormDB *sql.DB
	users func(...Dba) *Query
)

func init() {
	ormDB, err := Connect("root@tcp(127.0.0.1:3306)/orm_db?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	users = Table(ormDB, "user")
}
func Test_sKV(t *testing.T) {
	type man struct {
		User   *User
		Gender string `json:"gender"`
	}
	now := time.Now()
	s := reflect.ValueOf(man{
		&User{
			Age:       30,
			FirstName: "Tom",
			LastName:  "Cat",
			Email:     "Tom@test.com",
			CreatedAt: now,
		},
		"男",
	})
	wantKeys := []string{"age", "first_name", "last_name", "email", "created_at", "gender"}
	wantValues := []string{"30", "'Tom'", "'Cat'", "'Tom@test.com'", fmt.Sprintf("FROM_UNIXTIME(%d)", now.Unix()), "'男'"}
	keys, values := sKV(s)
	if !reflect.DeepEqual(wantKeys, keys) {
		t.Fatal("sKV keys error")
	}
	if !reflect.DeepEqual(wantValues, values) {
		t.Fatal("sKV values error")
	}
}

func Test_mKV(t *testing.T) {
	m := map[string]interface{}{
		"string": "hello world",
		"int":    10068,
		"time":   time.Now(),
	}
	keys, values := mKV(reflect.ValueOf(m))
	if len(keys) != len(values) ||
		len(keys) != len(m) {
		t.Fatal("mKv error")
	}
}

func TestQuery_Insert(t *testing.T) {
	user1 := &User{
		Age:       30,
		FirstName: "Tom",
		LastName:  "Cat",
	}
	user2 := User{
		Age:       30,
		FirstName: "Tom",
		LastName:  "Curise",
	}
	user3 := User{
		Age:       30,
		FirstName: "Tom",
		LastName:  "Hanks",
	}
	user4 := map[string]interface{}{
		"age":        30,
		"first_name": "Tom",
		"last_name":  "Zzy",
	}
	_, err := users().Insert([]interface{}{user1, user2})
	if err != nil {
		t.Fatal(err)
	}
	_, err = users().Insert(user3)
	if err != nil {
		t.Fatal(err)
	}
	_, err = users().Insert(user4)
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuery_Select(t *testing.T) {
	var user User
	err := users().Where("first_name = 'Tom'").Only("last_name").Select(&user)
	if err != nil {
		t.Fatal(err)
	}

	var userMore []User
	err = users().Where("first_name = 'Tom'").Order("id desc").Select(&userMore)
	if err != nil {
		t.Fatal(err)
	}

	var userMoreP []*User
	err = users().Where("first_name = 'Tom'").Select(&userMoreP)
	if err != nil {
		t.Fatal(err)
	}

	var lastName string
	err = users().Where(&User{FirstName: "Tom"}).Only("last_name").Select(&lastName)
	if err != nil {
		t.Fatal(err)
	}

	var lastNames []string
	err = users().Where(map[string]interface{}{
		"first_name": "Tom",
	}).Only("last_name").Select(&lastNames)
	if err != nil {
		t.Fatal(err)
	}

	var userM map[string]interface{}
	err = users().Where(&User{FirstName: "Tom"}).Only("last_name").Select(&userM)
	if err != nil {
		t.Fatal(err)
	}

	var userMS []map[string]interface{}
	err = users().Where("age > 10").Only("last_name", "age").Limit(100).Select(&userMS)
	if err != nil {
		t.Fatal(err)
	}
}

func TestQuery_Update(t *testing.T) {
	u1 := "age = 100"
	u2 := map[string]interface{}{
		"age":        100,
		"first_name": "z",
		"last_name":  "zy",
	}
	u3 := &User{
		Age:       100,
		FirstName: "z",
		LastName:  "zy",
	}
	_, _ = users().Where("age > 10").Update(u1)
	_, _ = users().Where("age > 10").Update(u2)
	_, _ = users().Where("age > 10").Update(u3)
}

func TestQuery_Delete(t *testing.T) {
	w := map[string]interface{}{
		"id": []int{1, 2, 3, 4},
	}
	_, _ = users().Where(w, "age > 10").Delete()

	users().Delete()
	var count int
	users().Only("count(1)").Select(&count)
	if count > 0 {
		t.Fatal()
	}
}
