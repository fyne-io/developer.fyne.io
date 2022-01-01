---
layout: page
title: Testing

order: 40
---

## Testing Graphical Apps

---

Part of a good test suite is being able to quickly write tests and run them on a regular basis.
Fyne's API is designed to make testing applications easy. By separating component logic from it's rendering definition we can load applications without actually displaying them and test the functionality completely.

### Example

The fyne repository and it's example applications all have test code at various levels.
To show how to test a complete application look at the following code from the
[calculator project](https://github.com/fyne-io/calculator).

```go
package calculator

import (
	"testing"

	"fyne.io/fyne/v2/test"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	calc := newCalculator()
	calc.loadUI(test.NewApp())

	test.Tap(calc.buttons["1"])
	test.Tap(calc.buttons["+"])
	test.Tap(calc.buttons["1"])
	test.Tap(calc.buttons["="])

	assert.Equal(t, "2", calc.output.Text)
}
```

Using the "test" package we load a test application, rather than a regular desktop app (`test.NewApp()` instead of `app.New()`). Then the `test.Tap()` function simulates user action and finally we check the text of the output for the correct answer.

A simple `go test .` will run this as a unit test, loading and testing the application, without ever having to load a window on the screen.
