package main

import ("testing"
"github.com/golang_project_01_server/practice/"
)
func testNotNil(t *testing.T) {
    total := Sum(5, 5)
    if total != 10 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}