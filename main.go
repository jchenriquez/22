package main

import "fmt"

func generateParenthesisH(leftCount, rightCount int, str string, result *[]string) {
  if leftCount == 0 && rightCount == 0 {
    *result = append(*result, str)
    return
  }

  if leftCount > 0 {
    generateParenthesisH(leftCount-1, rightCount, fmt.Sprintf("%s%c", str, '('), result)
  }

  if rightCount > 0 && leftCount < rightCount{
    generateParenthesisH(leftCount, rightCount-1, fmt.Sprintf("%s%c", str, ')'), result)
  }
}

func generateParenthesis(n int) []string {
    res := &[]string{} 
    generateParenthesisH(n, n, "", res)

    return *res
}

func main() {
  fmt.Printf("all parenthesis %v\n", generateParenthesis(0))
}