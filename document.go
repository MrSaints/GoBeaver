package gobeaver

import (
    "log"
    "github.com/PuerkitoBio/goquery"
)

const LSE_URL = "http://www.lse.ac.uk"
var PROGRAMMES_URL = map[string]string {
    "Undergraduate":    LSE_URL + "/resources/calendar/courseGuides/undergraduate.htm",
    "Graduate":         LSE_URL + "/resources/calendar/courseGuides/graduate.htm",
    "Research":         LSE_URL + "/resources/calendar/courseGuides/research.htm",
}

func GetDocument(url string) (program *goquery.Document) {
    var e error
    if program, e = goquery.NewDocument(url); e != nil {
        log.Fatal(e)
    }
    return
}