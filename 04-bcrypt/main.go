package main

import(
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){
	senha := "abobrinha"

	//gerando hash
	senhaerrada := "paçoca"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	//mincost, maxcost, defaultcost
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sb))

	// verificar hash
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senha)))
	fmt.Println(bcrypt.CompareHashAndPassword(sb, []byte(senhaerrada)))
}