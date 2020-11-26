package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

type nbaPlayer struct {
	person
	draftYear     int
	currentTeam   string
	pointsPerGame float64
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "Elijah",
			last:  "Hallan",
			age:   5,
		},
		ltk: false,
	}

	p1 := person{
		first: "Dom",
		last:  "Hallan",
		age:   1,
	}

	nbap1 := nbaPlayer{
		person: person{
			first: "Kobe",
			last:  "Bryant",
			age:   41,
		},
		draftYear:     1996,
		currentTeam:   "Retired from LA Lakers",
		pointsPerGame: 28.6,
	}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(sa1.person.first, sa1.person.last, sa1.person.age)
	fmt.Println(nbap1)

	p2 := struct {
		likesCandy    bool
		hatesOutdoors bool
		ownsMac       bool
	}{
		likesCandy:    true,
		hatesOutdoors: true,
		ownsMac:       true,
	}

	fmt.Println(p2)
}
