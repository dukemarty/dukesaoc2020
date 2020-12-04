package main

type Passport struct {
	Fields map[string]string
}

func MakePassport() Passport {
	res := Passport{
		Fields: make(map[string]string),
	}

	return res
}
