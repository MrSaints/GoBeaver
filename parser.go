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
    "Graduate":         LSE_URL + "/resources/calendar/courseGuides/graduate.htm",
    "Research":         LSE_URL + "/resources/calendar/courseGuides/research.htm",
}

type Course struct {
    Code string
    Title string
    URL string
    Department string
    Students int
    Class int
    Value int
    Program int
}

func main() {
    var program *goquery.Document
    var e error

    if program, e = goquery.NewDocument(PROGRAMMES_URL["Undergraduate"]); e != nil {
        log.Fatal(e)
    }

    var courses []interface{}
    program.Find("table tr td p a").Each(func(i int, s *goquery.Selection) {
        course_item := strings.Split(s.Text(), " ")
        course_item_url, _ := s.Attr("href")
        parsed_url, _ := url.Parse(PROGRAMMES_URL["Undergraduate"])
        parsed_relative, _ := url.Parse(course_item_url)

        course_object := new(Course)
        course_object.Code = course_item[0]
        course_object.Title = course_item[1]
        course_object.URL = parsed_url.ResolveReference(parsed_relative).String()
        course_object.Program = 0
        courses = append(courses, course_object)
    })
}