package main

import (
  "os"
  "fmt"
  "strings"
  "bufio"
  "strconv"
)
func check(e error) {
  if e != nil {
      panic(e)
  }
}

func contains(e int, s []int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

type coordinates struct {
  x, y []int
}

func expander(coord string) []int {
  i1, err := strconv.Atoi(strings.Split(coord, "..")[0])
  check(err)
  i2, err := strconv.Atoi(strings.Split(coord, "..")[1])
  check(err)
  expanded := make([]int, i2-i1+1)
  for i := range expanded {
    expanded[i] = i1+i
  }
  return expanded
}

func drawmap(lines []coordinates) [][]string {


  max_x := 0
  min_x := 100000
  max_y := 0
  min_y := 100000
  for _, line := range lines {
    fmt.Println(line)
    for _, value := range line.x {
      if value > max_x {
        max_x = value
      }
      if value < min_x {
        min_x = value
      }
    }
    for _, value := range line.y {
      if value > max_y {
        max_y = value
      }
      if value < min_y {
        min_y = value
      }
    }
  }
  mapped := make([][]string, max_x-min_x+1)
  for i, _ := range mapped {
    mapped[i] = make([]string, max_y-min_y+1)
    for j, _ := range mapped[i] {
      for _, line := range lines {
        x_val := make([]int, max_x-min_x+1)
        y_val := make([]int, max_y-min_y+1)

        for h, _ := range make([]int, len(lines)-1) {
          x_val[h] = line.x[h]-min_x
          y_val[h] = line.y[h]-min_y
        }
        
        if contains(i, x_val)&&contains(j, y_val) {
          mapped[i][j] = "#"
        } else {
          mapped[i][j] = "."
        }
      }
    }
  }


  fmt.Println(max_x)
  fmt.Println(max_y)
  fmt.Println(min_x)
  fmt.Println(min_y)
  fmt.Println(mapped)
  return mapped
}

// Take in a set of coordinates and build a map out of it

func main() {
  f, err := os.Create("/Users/dloiko/Downloads/advent/day17/generated.map")
  check(err)
  defer f.Close()

  inputFile, err := os.Open("./sample.txt")
  scanner := bufio.NewScanner(inputFile)

  var lines []coordinates

  for scanner.Scan() {
    line := scanner.Text()
    parsed1 := strings.Split(line, " ")[0]
    parsed2 := strings.Split(line, " ")[1]
    var x, y []int

    if strings.Contains(parsed1, "x") {
      x_num, err := strconv.Atoi(strings.Split(strings.Split(parsed1, "=")[1], ",")[0])
      check(err)
      x = append(x, x_num)


      y_raw := strings.Split(parsed2, "=")[1]
      y = expander(y_raw)
    } else if strings.Contains(parsed1, "y") {
      y_num, err := strconv.Atoi(strings.Split(strings.Split(parsed1, "=")[1], ",")[0])
      check(err)
      y = append(y, y_num)

      x_raw := strings.Split(parsed2, "=")[1]
      x = expander(x_raw)
    }

    lines = append(lines, coordinates{x, y})

    // Draw line here


  }


  fmt.Println(lines)
  drawmap(lines)

  // write to file
  d1 := []byte("hello\ngo\n")
  n2, err := f.Write(d1)
  check(err)
  fmt.Printf("wrote %d bytes\n", n2)

}
