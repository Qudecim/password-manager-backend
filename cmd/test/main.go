package main

import (
	"crypto/rand"
	"crypto/sha512"
	"database/sql"
	"fmt"
	"log"
	"math/big"

	_ "github.com/go-sql-driver/mysql"
	"github.com/qudecim/password-manager-backend/internal/enycrypt"
	"github.com/qudecim/password-manager-backend/internal/models/mysql"
)

func main() {
	test3()
}

func test1() {
	db, err := sql.Open("mysql",
		"main:main@tcp(127.0.0.1:3306)/main")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id   int
		name string
	)
	rows, err := db.Query("select PersonID, LastName from Persons where PersonID = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func test2() {
	db, err := sql.Open("mysql",
		"main:main@tcp(127.0.0.1:3306)/main")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	d := mysql.UserModel{}
	fmt.Print(d.Get(""))
}

func test3() {

	str := "email:password"

	h := sha512.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	s1, s2 := enycrypt.SplitBytes(bs)

	n1 := new(big.Int)
	p := n1.SetBytes(s1)

	n2 := new(big.Int)
	q := n2.SetBytes(s2)

	fmt.Println("p")
	fmt.Println(p)

	fmt.Println("q")
	fmt.Println(q)

	k, _ := rand.Prime(rand.Reader, 512)
	fmt.Println("test")
	fmt.Println(k)

	b := []byte("ABC€")

	pub, pri, err := enycrypt.GenerateByPassword(str)

	//pub, pri, err := enycrypt.GenerateKeys(1024)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pri)

	en, err := enycrypt.EncryptRSA(pub, b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(2)

	fmt.Println(en)

	den, err := enycrypt.DecryptRSA(pri, en)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(3)
	fmt.Println(string(den))
}

func encrypt() {

	b := []byte("ABC€")

	pub, pri, err := enycrypt.GenerateKeys(1024)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pri)

	en, err := enycrypt.EncryptRSA(pub, b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(2)

	fmt.Println(en)

	den, err := enycrypt.DecryptRSA(pri, en)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(den))
}
