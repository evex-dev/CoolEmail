package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("source : https://www.reddit.com/r/Design/comments/14c4nw/reddit_this_is_my_business_card/?st=JF7OQ2FR&sh=58e05f66")

	fmt.Printf("Email >")
	var Email string
	fmt.Scan(&Email)

	Name := strings.Split(Email, "@")[0]
	Domain := strings.Split(Email, "@")[1]
	Discord := "@" + strings.Split(strings.Split(Email, "@")[1], ".")[0]

	fmt.Println(Name, Domain, Discord)

	Etitle := "Email"
	Ntitle := "Name"
	Dtitle := "Domain"
	Distitle := "Discord"
	fmt.Println(Email)

	EmailDesc := GenDesc(Email, Etitle, true)
	NameDesc := GenDesc(Name, Ntitle, true)
	DomainDesc := GenDesc(Domain, Dtitle, true)
	DiscordDesc := GenDesc(Discord, Distitle, false)

	DiscordL := strings.Repeat(" ", strings.Index(Email, "@")) + DiscordDesc
	NameDomainL := NameDesc + " " + DomainDesc

	fmt.Println(DiscordL)
	fmt.Println(Email)
	fmt.Println(NameDomainL)
	fmt.Println(EmailDesc)
}

func GenDesc(text, title string, under bool) string {
	var Desc string
	Hed := "└"
	Line := "─"
	Bottom := "┘"
	if !under {
		Hed = "┌"
		Bottom = "┐"
	}
	if len(text) < len(title)+2 {
		title = string([]rune(title)[:len(text)-2])
		Desc = Hed + title + Bottom
	} else {
		padding := strings.Repeat(Line, (len(text)-len(title)-2)/2)
		Desc = Hed + padding + title + padding
		if (len(text)-len(title))%2 != 0 {
			Desc += Line
		}
		Desc += Bottom
	}

	return Desc
}
