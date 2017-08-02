package misc

import (
  "bytes"
  "../question"
  "testing"
)

func ExampleTestQuestion() {
  message := "Would you like to create a config.toml? [y/n] "
  question.Question(message)
  //Output: Would you like to create a config.toml? [y/n] 
}

func TestScanner(t *testing.T) {
  bufW := &bytes.Buffer{}
  yes := bytes.NewBufferString("yes\n")
  result,_ :=question.Scanner(yes, bufW)
  correctResult := true
  if correctResult != result {
    t.Errorf("output result should be %s but %s", correctResult, result)
    t.FailNow()
  }
}