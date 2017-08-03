package question

import (
  "bufio"
  "fmt"
  "io"
)

func Question(q string) {
  fmt.Print(q)
}

func Scanner(r io.Reader, w io.Writer) (bool, error) {
  result := false
  scanner := bufio.NewScanner(r)
  for scanner.Scan() {
    i := scanner.Text()
    if i == "Y" || i == "y" || i == "yes" || i == "" {
      result = true
      break
    } else if i == "N" || i == "n" || i == "no" {
      result = false
      break
    } else {
      fmt.Fprint(w, "Please enter y(yes) or n(no)")
    }
  }
  err := scanner.Err()
  return result, err
}