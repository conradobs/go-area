package matharea

import "math"

// Pi é uma proporcao numerica
const Pi = 3.1416

// Circ calcula area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula a area de um retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Não é visivel, comentario n é exigido
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
