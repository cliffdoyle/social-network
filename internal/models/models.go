package models

type User struct{
	uuid string
	firstName string
	lastName string
	nickName string
	email string
	password string
	profileStatus bool

}

type Post struct{
	puid string
	title string
	description string
	media string
	PostStatus string
}

type Groups struct{
	guid string
	name string
	admins []string
	members []string
	Description string
	Event []string
}

type Event struct{
	euid  string
	title string
	description string
	daytime string
	options []string
}

type chat struct{
cuid string
receiver string
sender string
message []string
}


type Notification struct{
	nuid string
	noticationType string
	sender string 
	receiver string

}

type Comments struct{
	commuid string
	content string
	sender string 

}

type Reaction struct{
	Likes int
	Dislikes int 

}

