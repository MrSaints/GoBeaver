package main

import (
    "strings"
    "net/url"
    "github.com/PuerkitoBio/goquery"
)

// Course collection
type Courses []Course

// Course structure
type Course struct {
    code string
    title string
    url string
    department string
    students int
    class int
    value int
    program string
}

// Sort interface
func (slice Courses) Len() int {
    return len(slice)
}

func (slice Courses) Less(i, j int) bool {
    return slice[i].code < slice[j].code;
}

func (slice Courses) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}

// Build course collection
func getCourses(Type string) (program_courses Courses) {
    program := getDocument(PROGRAMMES_URL[Type])

    program.Find("table tr td p a").Each(func(i int, s *goquery.Selection) {
        course_item := strings.Split(s.Text(), " ")
        course_item_url, _ := s.Attr("href")
        parsed_url, _ := url.Parse(PROGRAMMES_URL[Type])
        parsed_relative, _ := url.Parse(course_item_url)

        course_object := new(Course)
        course_object.code = course_item[0]
        course_object.title = course_item[1]
        course_object.url = parsed_url.ResolveReference(parsed_relative).String()
        course_object.program = Type
        program_courses = append(program_courses, *course_object)
    })
    return
}