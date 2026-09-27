/*
1. create map of alphabetic counter slices to array of strings
2. for each string
	1. create a alphabetic counter of the string
	2. add to map
3. output the map
*/

func groupAnagrams(strs []string) [][]string {
	counterMap := make(map[string][]string)
	for _, str := range strs {
		countArr := make([]int, 26)
		for _, char := range str {
			charIndex := int(int(char) - int('a'))
			countArr[charIndex] += 1
		}

		key := ""
		for _, count := range countArr {
			key += fmt.Sprintf("%03d", count)
		}

		counterMap[key] = append(counterMap[key], str)
		// for _, s := range countArr {
		// 	fmt.Printf("%d ", s)
		// }
		// fmt.Println("")
	}

	res := [][]string{}
	for _, v := range counterMap {
		res = append(res, v)
	}
	return res
}
