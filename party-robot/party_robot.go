package partyrobot

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, "+name+"!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return "Happy birthday "+name+"! You are now "+string(age)+" years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	return ""
}
