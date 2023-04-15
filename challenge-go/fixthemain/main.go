package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

type Door struct {
	Name  string
	state bool
}

func OpenDoor(Door Door) {
	PrintStr("Door Opening...")
	z01.PrintRune('\n')
	Door.state = true
}

func CloseDoor(Door Door) {
	PrintStr("Door Closing...")
	z01.PrintRune('\n')
	Door.state = false
}

func IsDoorOpen(Door Door) bool {
	PrintStr("is the Door opened ?")
	z01.PrintRune('\n')
	return Door.state
}

func IsDoorClose(Door Door) bool {
	PrintStr("is the Door closed ?")
	z01.PrintRune('\n')
	return Door.state
}

func main() {
	door := Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == false {
		CloseDoor(door)
	}
}
