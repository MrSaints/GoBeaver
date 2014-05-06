package main

import (
    "log"
    "strings"
    "strconv"
    "github.com/PuerkitoBio/goquery"
)

// Build course properties
func (this *Course) GetProperties() *Course {
    log.Print("--------------------- FIRED")
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

    teachers := strings.Replace(course.Find("#teacherResponsible-Content p").First().Text(), " and", ",", -1)
    this.teachers = strings.Split(teachers, ", ")

    this.availability = FormatProperty(course.Find("#availability-Content p"))
    this.content = FormatProperty(course.Find("#courseContent-Content p"))
    this.teaching = FormatProperty(course.Find("#teaching-Content p"))
    this.formative = FormatProperty(course.Find("#formativeCoursework-Content p"))
    this.readings = FormatProperty(course.Find("#indicativeReading-Content p"))
    this.assessments = FormatProperty(course.Find("#assessment-Content p"))

    return this
}

func FormatProperty(sel *goquery.Selection) string {
    noHTML := strings.Join(sel.Map(func(i int, s *goquery.Selection) string {
        return s.Text()
    }), "\n")
    return strings.TrimSpace(noHTML)
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

func (this *Course) Teachers() []string {
    if len(this.teachers) == 0 {
        this.GetProperties()
    }
    return this.teachers
}

func (this *Course) Availability() string {
    if this.availability == "" {
        this.GetProperties()
    }
    return this.availability
}

func (this *Course) Content() string {
    if this.content == "" {
        this.GetProperties()
    }
    return this.content
}

func (this *Course) Teaching() string {
    if this.teaching == "" {
        this.GetProperties()
    }
    return this.teaching
}

func (this *Course) Formative() string {
    if this.formative == "" {
        this.GetProperties()
    }
    return this.formative
}

func (this *Course) Readings() string {
    if this.readings == "" {
        this.GetProperties()
    }
    return this.readings
}

func (this *Course) Assessments() string {
    if this.assessments == "" {
        this.GetProperties()
    }
    return this.assessments
}

func (this *Course) Program() int {
    return this.program
}