package main

import (
  "fmt"
  "os"
  "github.com/PuerkitoBio/goquery"
  "time"
  "log"
  "encoding/csv"
)

const (
  TARGET = "http://stocks.finance.yahoo.co.jp/stocks/history/?code=7203.T"
)

func makeDirectory() {
  if err := os.Mkdir("data", 0777); err != nil {
    fmt.Println(err)
  }
}

func main() {
  start := time.Now()

  records := [][]string {}

  doc, _ := goquery.NewDocument(TARGET)

  doc.Find(".boardFin").Each(func(_ int, s *goquery.Selection) {
    i := []string{}
    s.Find("th").Each(func(_ int, s *goquery.Selection) {
      i = append(i,s.Text())
      })
    records = append(records, i)

    s.Find("tr").Each(func(_ int, s *goquery.Selection) {
      j := []string{}
      s.Find("td").Each(func(_ int, s *goquery.Selection) {
        j = append(j,s.Text())
        })
      records = append(records, j)
    })
  })

  makeDirectory()

  csvfile, err := os.Create("data/output.csv")
    if err != nil {
            fmt.Println("Error:", err)
            return
    }
  defer csvfile.Close()

  writer := csv.NewWriter(csvfile)
  for _, record := range records {
    err := writer.Write(record)
    if err != nil {
      fmt.Println("error:", err)
      return
    }
  }
  writer.Flush()

  executed := time.Since(start)
  log.Printf("Execution time : %s", executed)
}
