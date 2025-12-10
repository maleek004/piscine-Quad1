package checkpoint 

// Basically in this function, we are repeatedly going to find modulo of the 2 numbers (bigger to smaller ; smaller to first modulo , new smaller to the next second modulo... and so on )
// going by the mathematical principal : The greatest common divisor of two numbers, a and b, is the same as the greatest common divisor of the smaller number, b, and the remainder of their division
func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
