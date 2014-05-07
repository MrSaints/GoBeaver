package gobeaver

import (
    "net/url"
    "strings"
    "sort"
    "sync"
    "github.com/PuerkitoBio/goquery"
)

// Course collection
type Courses []*Course

// Course structure
type Course struct {
    Code string
    Title string
    URL string
    Department string
    Students int
    Class int
    Value float32
    Teachers []string // TODO: "Additional Teachers" and paragraphs
    Availability string // TODO: Store in array (compulsory, optional, others)
    Prerequisites string
    Content string // TODO: Handle lists?
    Teaching string // TODO: Store in array (MT, LT, ST)
    Formative string
    Readings string
    Assessments string
    Program int
}

// Build course collection for a specified program
func GetCourses(program_type string) (program_courses Courses) {
    var wg sync.WaitGroup
    program := GetDocument(PROGRAMMES_URL[program_type])
    courses := program.Find("table tr td p a")

    wg.Add(courses.Length())

    courses.Each(func(i int, s *goquery.Selection) {
        go func(s *goquery.Selection) {
            defer wg.Done()
            course_item := strings.SplitN(s.Text(), " ", 2)
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

            program_courses = append(program_courses, course_object)
        }(s)
    })
    wg.Wait()
    return
}

// Build course collection for all programmes
func GetAllCourses() (all_courses Courses) {
    var wg sync.WaitGroup
    wg.Add(len(PROGRAMMES_URL))

    for program, _ := range PROGRAMMES_URL {
        go func(program string) {
            defer wg.Done()
            all_courses = append(all_courses, GetCourses(program)...)
        }(program)
    }

    wg.Wait()
    sort.Sort(all_courses)
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