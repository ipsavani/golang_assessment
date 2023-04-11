package structs

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type animal struct {
	name            string
	characteristics []string
}

//A herbivore is an animal, so it can have the animal struct as a field
type herbivore struct {
	animal
	eatHuman bool
}

func (a animal) run() {
	fmt.Println(a.name, "is a lazy animal hence cannot run")
}

func Structs() {
	//Assigning values to the fields in the person struct:
	p1 := person{
		firstName: "Mark",
		lastName:  "Kedu",
		age:       30,
	}

	fmt.Println("The is the person: ", p1)
}

func Nested_structs() {

	herb := herbivore{
		animal: animal{
			name: "Goat",
			characteristics: []string{"Lacks sense",
				"Lazy",
				"Eat grass",
			},
		},
		eatHuman: false, //maybe
	}

	//We use dot(.) to acces each field in the struct
	fmt.Println("Animal name:", herb.animal.name)
	fmt.Println("Eats human? ", herb.eatHuman)
	fmt.Println("Characteristics:")
	for _, v := range herb.animal.characteristics {
		fmt.Printf("\t %v\n", v)
	}
}

func Recieverfunc() {

	animal1 := animal{
		name: "Elephant",
	}

	animal1.run()
}
