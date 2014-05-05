package main

import (
    "log"
    "strings"
    "net/url"
    "github.com/PuerkitoBio/goquery"
)

const LSE_URL = "http://www.lse.ac.uk"
var PROGRAMMES_URL = map[string]string {
    "Undergraduate":    LSE_URL + "/resources/calendar/courseGuides/undergraduate.htm",
    //"Graduate":         LSE_URL + "/resources/calendar/courseGuides/graduate.htm",
    //"Research":         LSE_URL + "/resources/calendar/courseGuides/research.htm",
}

type Course struct {
    code string
    title string
    url string
    department string
    students int
    class int
    value int
    program int
}

func getDocument(url string) (program *goquery.Document) {
    var e error
    if program, e = goquery.NewDocument(url); e != nil {
        log.Fatal(e)
    }
    return
}

func getCourses(Type string) (program_courses []Course) {
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
        course_object.program = 0 // TODO
        program_courses = append(program_courses, *course_object)
    })
    return
}

func main() {
    var courses []Course
    for program, _ := range PROGRAMMES_URL {
        courses = append(courses, getCourses(program)...)
    }
    log.Printf("Total courses: %d", len(courses))

    // Test
    for _, test_c := range courses {
        log.Print(test_c.Code())
    }
}