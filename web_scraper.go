// REF : https://www.youtube.com/watch?v=3KsE7zMm-AI

package main

import (
	"github.com/golocally/colly"
	"encoding/csv"
	"log"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)

	if err != nil {
		log.Fatalf("Could not create file. err : %q", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowDomains("internshala.com")
	)

	c.OnHTML(".internship_meta", func(e *colly.HTMLElement) {
		writer.Write( []strring{
			e.ChildText("a"),
			e,ChildText("span"),
		})
	})

	for i:=0; i<312; i++ {
		fmt.Printf("Scraping Page : %d\n", i)
		c.Visit("https://internshala.com/internships/page-" + strconv.Itoa(i))
	}

	log.Printf("Scraping Complete\n")
	log.Println(c)

}