package main

import "time"

type user struct {
	name string
	email string
	ext int
	privileged bool
}

type admin struct {
	person user
	level string
}

type Duration int64

func main()  {
	var bill user

	ted := user {
		name : "ted",
		email : "ted.park@dada.net",
		ext : 111,
		privileged: true,
	}

	cindy := user {"cindy", "cindy@dada.net", 123, false}

	sunny := admin {
		person : user {
			name : "sunny",
			email : "sunny@dada.net",
			ext : 144,
			privileged: true,

		},
		level : "super",
	}

	time.Now()
}
