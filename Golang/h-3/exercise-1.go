package exercise1

import "fmt"

func exercise1() {
	fmt.Println(ArrayMerge([]string{"Alterra", "Academy", "malang"}, []string{"academy", "jakarta"}))
}

func ArrayMerge(arrayA []string, arrayB []string) []string {
	mergedArr := append(arrayA, arrayB...)
	uniqueArr := []string{}

	for i, v := range mergedArr {
		if len(uniqueArr) == 0 {
			uniqueArr = append(uniqueArr, v)
		} else {
			isFound = false
			for j := i + 1; j < len(mergedArr); j++ {
				if v == mergedArr[j] {
					isFound = true
					// jika isFound tidak ada yang ketemu
					// break;
				}
			}
			if !isFound {
				uniqueArr = append(uniqueArr, v)
			}
		}
	}
	return uniqueArr
}
