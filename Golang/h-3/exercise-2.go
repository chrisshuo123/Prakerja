package exercise2

import "fmt"

func exercise2() {
	fmt.Println(Mapping([]string{"alta", "malang", "alta", "jakarta", "malang"}))
}

// func mapping(memanggil Mapping[]string diatas dalam bentuk slice) ..ini Returnnya.. {
	// Direturn disini yg terbawah sblm "}"
//}

func mappping(slice []string) map[]int {
	var resultMapping = map[string]int[]

	for _, keyMap := range slice {
		_, isValid := resultMapping[keyMap]

		if !isValid {
			resultMapping[keyMap] = 1
		} else {
			resultMapping[keyMap] += 1
		}
	}

	return resultMapping
}