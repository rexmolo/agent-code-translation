package main

func AnyInt(x, y, z interface{}) bool {
	// Type assert all three parameters to int
	ix, ok1 := x.(int)
niy, ok2 := y.(int)
	iz, ok3 := z.(int)
	
	// If any of them is not an int, return false
	if !ok1 || !ok2 || !ok3 {
		return false
	}
	
	// Check if one equals the sum of the other two
	if ix+iy == iz || ix+iz == iy || iy+iz == ix {
		return true
	}
	
	return false
}