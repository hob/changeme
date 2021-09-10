package main

import "fmt"

// Work in progress

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func main() {
	tdoc()
	// println(fmt.Sprintf("%sHello, %s%s", Blue, os.Args[1], Reset))
}

func tdoc() {
	type gift struct {
		day  string
		gift string
	}
	giftsMap := []gift{
		{"First", "A Partridge in a Pear Tree"},
		{"Second", "Two Turtle Doves"},
		{"Third", "Three French Hens"},
		{"Fourth", "Four Calling Birds"},
		{"Fifth", "Five Gold Rings"},
		{"Sixth", "Six Geese-a-Laying"},
		{"Seventh", "Seven Swans-a-Swimming"},
		{"Eighth", "Eight Maids-a-Milking"},
		{"Ninth", "Nine Ladies Dancing"},
		{"Tenth", "Ten Lords-a-Leaping"},
		{"Eleventh", "Eleven Pipers Piping"},
		{"Twelfth", "Twelve Drummers Drumming"},
	}
	for i, g := range giftsMap {
		println(fmt.Sprintf("On the %s day of Christmas\nMy true love sent to me", g.day))
		for j := i; j >= 0; j-- {
			lineEnd := "."
			if j == 1 {
				lineEnd = ", and"
			} else if j > 1 {
				lineEnd = ","
			}
			println(fmt.Sprintf("%s%s", giftsMap[j].gift, lineEnd))
		}
		println()
	}
}
