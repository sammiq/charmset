Charmset: a Hamcrest alike for Go
=================================

[![Go Report Card](https://goreportcard.com/badge/github.com/sammiq/charmset)](https://goreportcard.com/report/github.com/sammiq/charmset) [![GoDoc](https://godoc.org/github.com/sammiq/charmset?status.svg)](https://godoc.org/github.com/sammiq/charmset)

Introduction
------------

[Hamcrest](http://hamcrest.org/) is a matcher framework that is expressive and is used heavily by many test frameworks in various programming languages.

_Hamcrest_ is named for being an anagram of _matchers_ apparently; so is _charmset_ and is a good a reason to name it that as anything else.

Having liked this style of assertion for readability, and realized there is no canonical Go version, I looked around the community for something similar and for one reason or another decided that I would roll my own.

You probably don't want to look at this at the moment:

- It is not well-documented at all
- It is mostly for my personal use and suits my personal tastes
- The APIs are prone to change / breakage until I get happy with it
- I am not a Go expert, there are possibly many dumb things in here

Tiny Examples
-------------

The philosophy was to create a readable fluent API, hence the naming of the matcher packages.

    assert.That(probableString, has.Substring("penguins"))
    assert.That(probableSliceOfNumbers, has.Item(42))
    assert.That(probableMap, has.Key("My Key"))

you can chain them logically

    assert.That(probableSliceOfNumbers, has.Item(42).and(has.Item(24)))
    assert.That(probableString, has.Substring("penguins").or(has.Prefix("seal")))

you get the idea...

Documentation
-------------

Main package documentation is on GoDoc here: [charmset](https://godoc.org/github.com/sammiq/charmset).

Built-in matcher documentation is on GoDoc here: [has](https://godoc.org/github.com/sammiq/charmset/matchers/has), [is](https://godoc.org/github.com/sammiq/charmset/matchers/is) and, [matches](https://godoc.org/github.com/sammiq/charmset/matchers/matches).

