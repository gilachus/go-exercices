func main() {
	// Declaración de constantes
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Pi: ", pi)
	fmt.Println("Pi2: ", pi2)

	// Declaración de varaibles
	// Go no compila si las variables no son usadas
	base := 12          // Primera forma
	var altura int = 14 // Segunda forma
	var area int        // Tercera forma

	fmt.Println(base, altura, area)

	// Zero values
	// Go asigna valores a variables vacías
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)
}