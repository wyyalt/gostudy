package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Natthew", "Sarah", "Augustus", "Heidi", "Emilie", 
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func countCoins(username string) int {
	sum := 0
	for _, char := range username{
		switch char {
			case 'a', 'A', 'e', 'E':
				sum += 1
			case 'i', 'I':
				sum += 2
			case 'o', 'O':
				sum += 3
			case 'u', 'U':
				sum += 5
		}
	}
	return sum
}

func dispatchCoins() int {
	totalNum := 100
	for _, user := range users{
		allCoins := countCoins(user)
		totalNum = totalNum - allCoins
		value, exists := distribution[user]
		if !exists{
			distribution[user] = allCoins
		} else {
			distribution[user] = allCoins + value
		}
	}
	return totalNum

}

func main() {
	left := dispatchCoins()
	fmt.Println(distribution)
	fmt.Println(left)
}
