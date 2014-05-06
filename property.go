package main

// Build course properties
func (this *Course) getProperties() {

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
	return this.department
}

func (this *Course) Students() int {
	return this.students
}

func (this *Course) Class() int {
	return this.class
}

func (this *Course) Value() int {
	return this.value
}

func (this *Course) Program() string {
	return this.program
}