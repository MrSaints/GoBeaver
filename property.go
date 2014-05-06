package main

import (
	"strings"
	"strconv"
)

// Build course properties
func (this *Course) GetProperties() *Course {
	course := GetDocument(this.url)
	key_facts := course.Find("#keyFacts-Content p")

	this.department = strings.TrimPrefix(key_facts.First().Text(), "Department: ")
	this.students, _ = strconv.Atoi(strings.Split(key_facts.Eq(1).Text(), ": ")[1])
	this.class, _ = strconv.Atoi(strings.Split(key_facts.Eq(2).Text(), ": ")[1])

	if strings.Split(key_facts.Eq(3).Text(), " ")[1] == "Half" {
		this.value = 0.5
	} else {
		this.value = 1
	}
	return this
}

// Getters
func (this *Course) Code() string {
	return this.code
}

func (this *Course) Title() string {
	return this.title
}

func (this *Course) URL() string {
	return this.url
}

func (this *Course) Department() string {
	if this.department == "" {
		this.GetProperties()
	}
	return this.department
}

func (this *Course) Students() int {
	if this.students == 0 {
		this.GetProperties()
	}
	return this.students
}

func (this *Course) Class() int {
	if this.class == 0 {
		this.GetProperties()
	}
	return this.class
}

func (this *Course) Value() float32 {
	if this.value == 0 {
		this.GetProperties()
	}
	return this.value
}

func (this *Course) Program() int {
	return this.program
}