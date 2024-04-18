package arrays

import "fmt"

//map[key_type]value_type
func Maps() {
	//m := map[string]int{"jan": 1, "feb": 2, "mar": 3, "apr": 4}
	m := make(map[string]int)

	m["jan"] = 1
	m["feb"] = 2
	m["mar"] = 3
	m["apr"] = 4

	//return value using a defined key if it exists in the map
	value, exists := m["may"]
	if exists {
		fmt.Println("Value: ", value)
	} else {
		fmt.Println("Key not found.")
	}

	fmt.Println()
	for key, value := range m {
		fmt.Println(value, "st Month: ", key)
	}

	fmt.Println()
	fmt.Println(m)

	m["may"] = 5

	fmt.Println("Before deletion:", m)
	delete(m, "may")
	fmt.Println("After deletion:", m)
}
