# Beaver Guide

A Go package for programmatically working with the London School of Economics and Political Science (LSE) course guide (unofficial API).

The Beaver Guide library uses [GoQuery](https://github.com/PuerkitoBio/goquery) to extract and synthesise _(web scraping)_ data from the LSE course guide documents. The data are then stored in a Go struct (for each individual course) and a slice (for a collection of courses) which can be programmatically accessed and manipulated.

For example, `generator/generator.go` in this repository idiomatically dumps all the course information that are stored in a slice and struct via Go's JSON serialization library.


## Installation

The Beaver Guide package is comprised of a library and a runnable binary, `generator/generator.go` in this repo.

To install the library:

	$ go get github.com/MrSaints/beaverguide

(Optional) To install and run the generator:

	$ go get github.com/MrSaints/beaverguide/generator
	$ go run generator.go


## Usage / Examples
See `generator/generator.go` in this repo.


## Boring Stuff

The `MrSaints/BeaverGuide` project is NEITHER AFFILIATED WITH NOR ENDORSED BY the LSE university and/or its student union (LSE SU). 
It is a personal project which exposes a rudimentary API for accessing and manipulating [publicly available course information](http://www.lse.ac.uk/resources/calendar/Default.htm). 
Please use this library at your own discretion.


### Licensing

[The MIT License (MIT)](http://ian.mit-license.org/) / LICENSE

Copyright (C) 2014, Ian Lai.