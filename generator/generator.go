package main

import (
    "encoding/json"
    "io/ioutil"
    "log"
    "github.com/MrSaints/gobeaver"
)

func DumpUndergraduate() {
    //courses := gobeaver.GetAllCourses()
    courses := gobeaver.GetCourses("Undergraduate")
    log.Printf("Total courses: %d", len(courses))
    courses.GetProperties()

    json_courses, json_error := json.Marshal(courses)
    if json_error != nil {
        log.Fatal(json_error)
    }

    file_error := ioutil.WriteFile("courses_dump.json", json_courses, 0644)
    if file_error != nil {
        panic(file_error)
    }

    log.Print("Complete!")
}

func main() {
    EC599 := gobeaver.GetCourse("EC599")
    log.Print(EC599)
}