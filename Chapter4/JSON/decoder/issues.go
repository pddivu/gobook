package main

import (
	"html/template"
	"log"
	"os"
	"time"

	"github.com/pddivu/LearningGo/gobook/Chapter4/JSON/decoder/github"
)

const temp1 = `{{.Totalcount}} issues:
{{range .Items}}------------------------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title| printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

/*report,err := template.New("report").
Funcs(template.FuncMap{"daysAgo" :daysAgo}).
Parse(temp1)

if err!=nil{
	log.Fatal(err)
}*/

var report = template.Must(template.New("issuelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(temp1))

func main() {

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	/*fmt.Printf("%d issues: \n", result.Totalcount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d  %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}*/

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}

}
