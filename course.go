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

// Construct URL to course
func BuildURL(absolute, relative string) string {
    parsed_absolute, _ := url.Parse(absolute)
    parsed_relative, _ := url.Parse(relative)
    return parsed_absolute.ResolveReference(parsed_relative).String()
}

func ProgramAtoi(program_type string) (program_code int) {
    if program_type == "Graduate" {
        program_code = 1
    } else if program_type == "Research" {
        program_code = 2
    } else {
        program_code = 0
    }
    return
}

func BuildCourse(s *goquery.Selection, program_type string) (course *Course) {
    course_item := strings.SplitN(s.Text(), " ", 2)
    course_item_url, _ := s.Attr("href")

    course = new(Course)
    course.Code = course_item[0]
    course.Title = course_item[1]
    course.URL = BuildURL(PROGRAMMES_URL[program_type], course_item_url)
    course.Program = ProgramAtoi(program_type)
    return
}

func GetCourse(code string) (course *Course) {
    done := make(chan bool)

    for program_type, program_url := range PROGRAMMES_URL {
        go func(program_type string, program_url string) {
            courses := GetDocument(program_url).Find("table tr td p a[href*=\""+ code +"\"]")

            if courses.Length() > 0 {
                course = BuildCourse(courses.First(), program_type)
                course.GetProperties()
                done <- true
            }
        }(program_type, program_url)
    }

    <-done
    return
}

// Build course collection for a specified program
func GetCourses(program_type string) (program_courses Courses) {
    var wg sync.WaitGroup
    courses := GetDocument(PROGRAMMES_URL[program_type]).Find("table tr td p a")
    wg.Add(courses.Length())

    courses.Each(func(i int, s *goquery.Selection) {
        go func(s *goquery.Selection) {
            defer wg.Done()
            program_courses = append(program_courses, BuildCourse(s, program_type))
        }(s)
    })

    wg.Wait()
    sort.Sort(program_courses)
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