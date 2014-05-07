package gobeaver

import (
    "strconv"
    "strings"
    "sync"
    "github.com/PuerkitoBio/goquery"
)

// Build properties for a course
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
    this.Prerequisites = FormatProperty(course.Find("#preRequisites-Content p"))
    this.Content = FormatProperty(course.Find("#courseContent-Content p"))
    this.Teaching = FormatProperty(course.Find("#teaching-Content p"))
    this.Formative = FormatProperty(course.Find("#formativeCoursework-Content p"))
    this.Readings = FormatProperty(course.Find("#indicativeReading-Content p"))
    this.Assessments = FormatProperty(course.Find("#assessment-Content p"))

    return this
}

// Build properties for all courses
func (this Courses) GetProperties() Courses {
    var wg sync.WaitGroup
    wg.Add(len(this))

    for _, course := range this {
        go func(course *Course) {
            defer wg.Done()
            course.GetProperties()
        }(course)
    }

    wg.Wait()
    return this
}

func FormatProperty(paragraphs *goquery.Selection) string {
    no_html := strings.Join(paragraphs.Map(func(i int, s *goquery.Selection) string {
        return s.Text()
    }), "\n")
    return strings.TrimSpace(no_html)
}