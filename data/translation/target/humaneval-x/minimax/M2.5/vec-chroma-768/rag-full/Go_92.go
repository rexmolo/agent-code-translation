package main

func AnyInt(x, y, z interface{}) bool {
	// Use type assertions to check if each parameter is an int
	xi, ok1 := x.(int)
	yi, ok2 := y.(int)
	zi, ok3 := z.(int)

	// If any of them is not an int, return false
	if !ok1 || !ok2 || !ok3 {
		return false
	}

	// Check if one number equals the sum of the other two
	return (xi+yi == zi) || (xi+zi == yi) || (yi+zi == xi)
}