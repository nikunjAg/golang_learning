package main

import (
	"fmt"
	"time"

	"example.com/app/user"
)

func getInput(label string) string {
	fmt.Printf("%v: ", label)
	var value string
	fmt.Scanln(&value)

	return value
}

func acceptUserDetails() {

	firstName := getInput("Please enter First name")
	lastName := getInput("Please enter Last name")
	birthDate := getInput("Please enter Birth Date(DD/MM/YYYY)")

	var userA, err = user.New(firstName, lastName, birthDate, time.Now())
	var studentA *Student
	studentA, err = NewStudent(1)

	var userB = user.User{
		FirstName: "F1",
		LastName:  "L1",
	}

	if err != nil {
		fmt.Println(err)
		return
	} else {

		userA.DisplayDetails()
		userB.DisplayDetails()
		// userA.ClearUsername()
		// userA.DisplayDetails()
		studentA.User.DisplayDetails()
	}

}

// func main() {
// 	// acceptUserDetails()

// 	// var book = book{
// 	// 	name: "Harry Potter and the Deathly Hollows",
// 	// }

// 	// fmt.Println(book)

// 	var books [4]book
// 	books[0] = book{
// 		name: "Book 1",
// 	}
// 	books[1] = book{
// 		name: "Book 2",
// 	}
// 	books[2] = book{
// 		name: "Book 3",
// 	}
// 	books[3] = book{
// 		name: "Book 4",
// 	}
// 	fmt.Println(books)

// 	selectedBooks := books[:2]
// 	specialBooks := selectedBooks[:1]
// 	fmt.Println(selectedBooks)

// 	selectedBooks[0].name = "Updated Book"
// 	fmt.Println("After Update----------")

// 	fmt.Println(books)
// 	fmt.Println(selectedBooks)
// 	fmt.Println(specialBooks)
// 	fmt.Println(len(selectedBooks), cap(selectedBooks))
// 	fmt.Println(len(specialBooks), cap(specialBooks))

// 	specialBooks = specialBooks[:3]
// 	fmt.Println(specialBooks)

// }

// func main() {
// 	a := [4]int{1, 2, 3, 4}

// 	b := a[:2]
// 	// Gives IndexOutOfRange for accessing an index outside of len
// 	// b[2] = 5

// 	c := b[:2]

// 	c = append(c, 5)

// 	fmt.Println(a, b, c)
// 	fmt.Println(len(c))

// }

/* func main() {

	// Task 1
	hobbies := [3]string{"problem-solving", "cricket", "watch"}
	fmt.Println(hobbies)

	// Task 2
	fmt.Println(hobbies[0], hobbies[1:3])

	// Task 3
	slicedArr := hobbies[:2]
	fmt.Println(slicedArr)

	// Task 4
	reslicedArr := slicedArr[1:3]
	fmt.Println(reslicedArr)

	// Task 5
	courseGoals := []string{"Learn Golang", "Rest Apis"}
	fmt.Println(courseGoals)

	// Task 6
	courseGoals[1] = "Api Development"
	fmt.Println(courseGoals)
	courseGoals = append(courseGoals, "Upskilling")
	fmt.Println(courseGoals)

	// Task 7
}
*/

// func main() {
// 	websites := map[string]string{
// 		"A": "apple.com",
// 		"B": "bira.com",
// 		"C": "cocacola.com",
// 	}

// 	fmt.Println(websites)
// 	fmt.Println(websites["C"])
// 	delete(websites, "B")
// 	fmt.Println(websites)
// }

func updateSlice(slice []int) {
	for index := range slice {
		if index == 1 {
			slice[index] = 5
		}
	}
}

func mapFn[T comparable](arr [4]T, cb func(T, [4]T, int) T) [4]T {

	updatedArr := [4]T{}

	for index, value := range arr {
		updatedArr[index] = cb(value, arr, index)
	}

	return updatedArr
}

func double[T int | float64](val T, arr [4]T, index int) T {
	return 2 * val
}

func sumAll(a int, b ...int) int {

	total := 0

	for _, val := range b {
		total += val
	}

	return a + total
}

func main() {

	arr := []int{1, 2, 3, 4}

	fmt.Println(sumAll(1, arr...))

}
