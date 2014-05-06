package main

import (
    "strconv"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

// Build course properties
func (this *Course) GetProperties() *Course {
    course := GetDocument(this.URL)
    key_facts := course.Find("#keyFacts-Content p")

    this.Department = strings.TrimPrefix(key_facts.First().Text(), "Department: ")
    this.Students, _ = strconv.Atoi(strings.Split(key_facts.Eq(1).Text(), ": ")[1])
    this.Class, _ = strconv.Atoi(strings.Split(key_facts.Eq(2).Text(), ": ")[1])

    if strings.Split(key_facts.Eq(3).Text(), " ")[1] == "Half" {
        this.Value = 0.5
    } else {
        this.Value = 1
    }

    teachers := strings.Replace(course.Find("#teacherResponsible-Content p").First().Text(), " and", ",", -1)
    this.Teachers = strings.Split(teachers, ", ")

    this.Availability = FormatProperty(course.Find("#availability-Content p"))
    this.Content = FormatProperty(course.Find("#courseContent-Content p"))
    this.Teaching = FormatProperty(course.Find("#teaching-Content p"))
    this.Formative = FormatProperty(course.Find("#formativeCoursework-Content p"))
    this.Readings = FormatProperty(course.Find("#indicativeReading-Content p"))
    this.Assessments = FormatProperty(course.Find("#assessment-Content p"))

    return this
}

func FormatProperty(paragraphs *goquery.Selection) string {
    no_html := strings.Join(paragraphs.Map(func(i int, s *goquery.Selection) string {
        return s.Text()
    }), "\n")
    return strings.TrimSpace(no_html)
}