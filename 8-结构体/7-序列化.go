package main

type Student struct {
	ID     int
	Name   string
	Gender string
}

type Class struct {
	Title   string
	Student []*Student
}
