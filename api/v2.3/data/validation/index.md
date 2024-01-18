---
redirect_to:
  - https://docs.fyne.io/api/v2.3/data/validation/

layout: page
tags: [api]
title: Fyne API "validation"
---


# validation
---
```go
import "fyne.io/fyne/v2/data/validation"
```

Package validation provides validation for data inside widgets.

## Usage

#### func  NewAllStrings

```go
func NewAllStrings(validators ...fyne.StringValidator) fyne.StringValidator
```
NewAllStrings creates a validator that requires all of the passed string validators to pass. In short, it combines multiple string validators into one.


<div class="since">Since: <code>
2.2</code></div>

#### func  NewRegexp

```go
func NewRegexp(regexpstr, reason string) fyne.StringValidator
```
NewRegexp creates a new validator that uses regular expression parsing. The validator will return nil if valid, otherwise returns an error with a reason text.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewTime

```go
func NewTime(format string) fyne.StringValidator
```
NewTime creates a new validator that verifies times using time.Parse. The validator will return nil if valid, otherwise returns an error with a reason text. The reference time for the format: Mon Jan 2 15:04:05 -0700 MST 2006. See time.Parse() for more information about the reference time: https://pkg.go.dev/time#Parse


<div class="since">Since: <code>
2.1</code></div>

#### types
