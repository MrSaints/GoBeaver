package main

import (
    "log"
    "fmt"
    "strings"
    "github.com/PuerkitoBio/goquery"
)

var BASE_URL = "http://www.lse.ac.uk"
var COURSES_URL = map[string]string {
    "Undergraduate":    BASE_URL + "/resources/calendar/courseGuides/undergraduate.htm",
    "Graduate":         BASE_URL + "/resources/calendar/courseGuides/graduate.htm",
    "Research":         BASE_URL + "/resources/calendar/courseGuides/research.htm",
}

type Undergraduate []*Course
type Graduate []*Course
type Research []*Course
type Course struct {
    Code string
    Title string
    URL string
    Department string
    Students int
    Class int
    Value int
}

func main() {
    var program *goquery.Document
    var e error

    if program, e = goquery.NewDocument(COURSES_URL["Undergraduate"]); e != nil {
        log.Fatal(e)
    }

    program.Find("table tr td p a").Each(func(i int, s *goquery.Selection) {
        course_item := strings.Split(s.Text(), " ")
        course_object := new(Course)
        course_object.Code = course_item[0]
        course_object.Title = course_item[1]
        course_object.URL, _ = s.Attr("href")
        fmt.Printf("%s: %s (%s)\n", course_object.Code, course_object.Title, course_object.URL)
    })
}