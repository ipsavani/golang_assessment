package slices

import "fmt"

func Slices() {

	s := make([]int, 4)
	s[0] = 10
	s[1] = 20
	s[2] = 30
	s[3] = 40
	fmt.Println("Slice created with 'make': ", s)

	//add elements to this slice using "append" builtin function:
	s = append(s, 50)
	fmt.Println("Added one element to slice: ", s)

	//append more than one element to the slice:
	s = append(s, 60, 70)
	fmt.Println("Added two elements to slice: ", s)

	//remove from that slice:
	s = append(s[:2], s[2+1:]...)
	fmt.Println("Deleted one element from slice: ", s)

	//dReplace an element with another
	s[2] = s[len(s)-2]
	fmt.Println("Slice with element replaced: ", s)

	//replace the third element now "60" with the last element "70":
	s[2] = s[len(s)-1]
	fmt.Println("Updated Slice with element replaced: ", s)

	//Get particular elements from the slice: [10 20 70 50 60 70]
	//To get the 2nd(index 1) to the 4th(index 3) element, we do:
	s = s[1:4]
	fmt.Println("Slice with second to fourth element: ", s)

	//Get the length of the current slice:
	fmt.Println("Length: ", len(s))

	//Get the capacity of the current slice:
	fmt.Println("Capacity: ", cap(s)) //this is give "7"

	//Copy one slice to another:
	d := make([]int, len(s))
	copy(d, s)
	fmt.Println("This is the new slice: ", d)
}

func Loop_thru_slices() {
	s := []int{10, 20, 30, 40}

	//using "range"
	for key, value := range s {
		fmt.Println(key, value)
	}

	//Using traditional forloop:
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) //get the value at index "i"
	}
}

func Nested_slices() {
	nested := make([][]int, 0, 3)
	for i := 0; i < 3; i++ {
		out := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			out = append(out, j)
		}
		nested = append(nested, out)
	}
	fmt.Println(nested)

	appleLaptops := []string{"MacbookPro", "MacbookAir"}
	hpLaptops := []string{"hp650", "hpEliteBook"}
	laptops := [][]string{appleLaptops, hpLaptops}
	for i, v := range laptops {
		fmt.Println("Record: ", i)
		for _, name := range v {
			fmt.Printf("\t Laptop name: %v \n", name)
		}
	}
}
