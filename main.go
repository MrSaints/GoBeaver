package main // Rename in the future

import (
    "log"
    "sort"
    "encoding/json"
    "github.com/PuerkitoBio/goquery"
)

const LSE_URL = "http://www.lse.ac.uk"
var PROGRAMMES_URL = map[string]string {
    "Undergraduate":    LSE_URL + "/resources/calendar/courseGuides/undergraduate.htm",
    //"Graduate":         LSE_URL + "/resources/calendar/courseGuides/graduate.htm",
    //"Research":         LSE_URL + "/resources/calendar/courseGuides/research.htm",
}

func GetDocument(url string) (program *goquery.Document) {
    var e error
    if program, e = goquery.NewDocument(url); e != nil {
        log.Fatal(e)
    }
    return
}

func main() {
    var courses Courses
    for program, _ := range PROGRAMMES_URL {
        courses = append(courses, GetCourses(program, false)...)
    }
    sort.Sort(courses)
    log.Printf("Total courses: %d", len(courses))

    // Test
    for _, test_c := range courses {
        log.Print(test_c)
        //test_c.GetProperties()
        /*log.Print(test_c.Teachers())
        log.Print(test_c.Availability())
        log.Print(test_c.Content())
        log.Print(test_c.Teaching())
        log.Print(test_c.Formative())
        log.Print(test_c.Readings())
        log.Print(test_c.Assessments())*/
        test_json, e := json.Marshal(test_c)
        if e != nil {
            log.Fatal(e)
        }
        log.Print(string(test_json))
        break
    }
}