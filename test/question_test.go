package misc

import (
  "bytes"
  "os"
  "../question"
  "testing"
)

func ExampleTestQuestion() {
  message := "Would you like to create a config.toml? [y/n] "
  question.Question(message)
  //Output: Would you like to create a config.toml? [y/n] 
  bufW := os.Stdout
  other := bytes.NewBufferString("other\n")
  question.Scanner(other, bufW)
  //Output: Would you like to create a config.toml? [y/n] Please enter y(yes) or n(no)
}

func TestScanner(t *testing.T) {
  bufW := &bytes.Buffer{}
  yes := bytes.NewBufferString("yes\n")
  result, _ := question.Scanner(yes, bufW)
  correctResult := true
  if correctResult != result {
    t.Errorf("output result should be %s but %s", correctResult, result)
    t.FailNow()
  }

  no :=bytes.NewBufferString("no\n")
  result2, _ := question.Scanner(no, bufW)
  correctResult = false
  if correctResult != result2 {
    t.Errorf("output result should be %s but %s", correctResult, result2)
    t.FailNow()
  }
}