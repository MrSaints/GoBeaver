package main

import (
    "encoding/json"
    "log"
    "sort"
    "github.com/MrSaints/beaverguide"
)

func main() {
    // TODO: Implement concurrency

    var courses beaverguide.Courses
    for program, _ := range beaverguide.PROGRAMMES_URL {
        courses = append(courses, beaverguide.GetCourses(program)...)
    }
    sort.Sort(courses)
    log.Printf("Total courses: %d", len(courses))

    // Test
    for _, test_c := range courses {
        log.Print(test_c)
        test_c.GetProperties()
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