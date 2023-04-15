package piscine

func ConcatParams(args []string) string {
	var t string
	for a := 0; a < len(args); a++ {
		if a != len(args)-1 {
			t += args[a]
			t += "\n"
		} else {
			t += args[a]
		}
	}
	return t
}
