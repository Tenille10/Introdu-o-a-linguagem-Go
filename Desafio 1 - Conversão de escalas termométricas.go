// declarar meu pacote principal
package main

//importar função fmt
import "fmt"

//declaração da variável CONST do ponto de ebulição da água em F
const ebulicaoF = 212

//função principal
func main() {

	tempFF := 32
	tempFC := 0
	tempF := ebulicaoF            // variável do valor da temperatura em graus F
	tempC := (tempF - 32) * 5 / 9 //Conversão de F para C
	//apareça na tela o resultado
	fmt.Printf("A temperatura de ebulição da água em °F = %v (%T), temperatura de ebulição da água em °C =%v (%T) .", tempF, tempF, tempC, tempC)
	fmt.Printf("A temperatura de fusão da água em °F = %v (%T), e a temperatura de fusão da água em °C é = %v (%T). ", tempFF, tempFF, tempFC, tempFC)
	// A gente espera que apareça F = 212 e C = 100

}
