package common

// func IfEmptyThrowError(){
// 	var result

// 	return result
// }

func IsEmpty(params string) bool {
	result := false

	if len(params) == 0 {
		result = true
	}

	return result
}
