1º algo
package main

import "fmt"

func main() {
	var name string
	fmt.Println("qual o seu nome?")
	fmt.Scan(&name)
	var idade int
	fmt.Println("sua idade:")
	fmt.Scan(&idade)
	var peso float64
	fmt.Println("Seu peso:")
	fmt.Scan(&peso)

	fmt.Println("Seu nome é:", name)
	fmt.Println("Voce tem", idade, "anos")
	fmt.Println("Voce pesa", peso)
  ___________________________________________________
  2º algoritimo 
  package main

import "fmt"

func main() {
	var altura float64
	fmt.Println("Altura do retangulo:")
	fmt.Scan(&altura)
	var base float64
	fmt.Println("Base do retangulo:")
	fmt.Scan(&base)
	var area float64
	area = altura * base
	fmt.Println("A area é:", area)
}
_______________________________________________________
3º algoritimo
package main

import "fmt"

func main() {
	var altura float64
	fmt.Println("Altura da Caixa:")
	fmt.Scan(&altura)
	var base float64
	fmt.Println("Base do Caixa:")
	fmt.Scan(&base)
	var profundidade float64
	fmt.Println("Profundidade da Caixa:")
	fmt.Scan(&profundidade)
	var area float64
	area = altura * base * profundidade
	fmt.Println("A area é:", area)
  _____________________________________________________
  4º algoritimo
  package main

import "fmt"

func main() {
	var numero_1 float64
	fmt.Println("Primeiro numero:")
	fmt.Scan(&numero_1)
	var numero_2 float64
	fmt.Println("Segundo numero:")
	fmt.Scan(&numero_2)
	var numero_3 float64
	fmt.Println("Terceiro numero:")
	fmt.Scan(&numero_3)
	var numero_4 float64
	fmt.Println("Quarto numero:")
	fmt.Scan(&numero_4)
	var media float64
	media = (numero_1 + numero_2 + numero_3 + numero_4) / 4
	fmt.Println("A media é:", media)

}
_________________________________________________________
5º algoritimo
  package main

import "fmt"

func main() {
	var Real float64
	fmt.Println("Quantos reais deseja converter para Dolar:")
	fmt.Scan(&Real)
	var Dolar float64
	fmt.Println("Quantos dolares deseja converter para Real")
	fmt.Scan(&Dolar)
	var ConvertorRd float64
	ConvertorRd = (Real * 0.19)
	fmt.Println("O valor de real para dolar é:", ConvertorRd)
	var ConvertorDr float64
	ConvertorDr = (Dolar * 5.26)
	fmt.Println("O valor de Dolar para real é:", ConvertorDr)
}
