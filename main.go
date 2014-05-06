package main

import (
    "log"
    "sort"
    "github.com/PuerkitoBio/goquery"
)

const LSE_URL = "http://www.lse.ac.uk"
var PROGRAMMES_URL = map[string]string {
    "Undergraduate":    LSE_URL + "/resources/calendar/courseGuides/undergraduate.htm",
    //"Graduate":         LSE_URL + "/resources/calendar/courseGuides/graduate.htm",
    //"Research":         LSE_URL + "/resources/calendar/courseGuides/research.htm",
}

func getDocument(url string) (program *goquery.Document) {
    var e error
    if program, e = goquery.NewDocument(url); e != nil {
        log.Fatal(e)
    }
    return
}

func main() {
    var courses Courses
    for program, _ := range PROGRAMMES_URL {
        courses = append(courses, getCourses(program)...)
    }
    sort.Sort(courses)
    log.Printf("Total courses: %d", len(courses))

    // Test
    for _, test_c := range courses {
        log.Print(test_c.Code())
    }
}