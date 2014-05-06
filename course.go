package main

import (
    "log"
    "net/url"
    "strings"
    "sync"
    "github.com/PuerkitoBio/goquery"
)

// Course collection
type Courses []Course

// Course structure
type Course struct {
    Code string
    Title string
    URL string
    Department string
    Students int
    Class int
    Value float32
    Teachers []string // TODO: "Additional Teachers"
    Availability string // TODO: Store in array (compulsory, optional, others)
    Content string // TODO: Handle lists?
    Teaching string // TODO: Store in array (MT, LT, ST)
    Formative string
    Readings string
    Assessments string
    Program int
}

// Build course collection
func GetCourses(program_type string, lazyload bool) (program_courses Courses) {
    var wg sync.WaitGroup
    program := GetDocument(PROGRAMMES_URL[program_type])
    courses := program.Find("table tr td p a")

    if !lazyload {
        wg.Add(courses.Length())
    }

    courses.Each(func(i int, s *goquery.Selection) {
        course_item := strings.Split(s.Text(), " ")
        course_item_url, _ := s.Attr("href")
        parsed_url, _ := url.Parse(PROGRAMMES_URL[program_type])
        parsed_relative, _ := url.Parse(course_item_url)

        course_object := new(Course)
        course_object.Code = course_item[0]
        course_object.Title = course_item[1]
        course_object.URL = parsed_url.ResolveReference(parsed_relative).String()

        if program_type == "Graduate" {
            course_object.Program = 1
        } else if program_type == "Research" {
            course_object.Program = 2
        } else {
            course_object.Program = 0
        }

        if !lazyload {
            go func(course_object *Course) {
                defer wg.Done()
                course_object.GetProperties()
            }(course_object)
        }

        program_courses = append(program_courses, *course_object)
    })

    if !lazyload {
        wg.Wait()
    }

    return
}

// Sort interface
func (slice Courses) Len() int {
    return len(slice)
}

func (slice Courses) Less(i, j int) bool {
    return slice[i].Code < slice[j].Code;
}

func (slice Courses) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}