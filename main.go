
package main

import "fmt"

// User representa um usuário.
type User struct {
	Name string
	Age  int
}

// mostrarMaioresDeIdade exibe os usuários com 18 anos ou mais.
func mostrarMaioresDeIdade(users []User) {
	fmt.Println("Usuários maiores de idade:")

	for _, user := range users {
		if user.Age >= 18 {
			fmt.Println(user.Name)
		}
	}
}

// contarMenoresDeIdade conta usuários com menos de 18 anos.
func contarMenoresDeIdade(users []User) int {
	contador := 0

	for _, user := range users {
		if user.Age < 18 {
			contador++
		}
	}

	return contador
}

func main() {
	users := []User{
		{Name: "Ana", Age: 17},
		{Name: "Marcos", Age: 42},
		{Name: "João", Age: 15},
		{Name: "Maria", Age: 22},
	}

	mostrarMaioresDeIdade(users)

	total := contarMenoresDeIdade(users)
	fmt.Println("Quantidade de menores de idade:", total)
}
