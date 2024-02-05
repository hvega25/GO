package main

import (
	"fmt"
	"math/big"
	"time"
)

// Función para calcular el factorial de un número grande
func factorial(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) == 0 || n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}

	return new(big.Int).Mul(n, factorial(new(big.Int).Sub(n, big.NewInt(1))))
}

func main() {
	var numero int64

	// Solicitar al usuario que ingrese un número
	fmt.Print("Ingrese un número para calcular el factorial: ")
	fmt.Scan(&numero)

	// Crear un nuevo big.Int desde el número ingresado
	numeroGrande := big.NewInt(numero)

	// Medir el tiempo de ejecución
	startTime := time.Now()

	// Calcular el factorial
	resultado := factorial(numeroGrande)

	// Obtener la duración del tiempo de ejecución
	duration := time.Since(startTime)

	// Mostrar el resultado y el tiempo de ejecución
	fmt.Printf("El factorial de %d es: %s\n", numero, resultado.String())
	fmt.Printf("Tiempo de ejecución: %s\n", duration)
}
