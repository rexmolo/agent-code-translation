package main

func AnyInt(x, y, z interface{}) bool {
	// Type assert each parameter to int
	xi, ok1 := x.(int)
	yi, ok2 := y.(int)
	zi, ok3 := z.(int)

	// If any type assertion fails, not all are integers
	if !ok1 || !ok2 || !ok3 {
		return false
	}

	// Check if one equals sum of the other two
	return xi+yi == zi || xi+zi == yi || yi+zi == xi
}
