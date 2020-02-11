package main

import (
	"os"
)

func main() {
	subject := "Continue Pipeline"
	destination := os.Getenv("TO")
	r := NewRequest([]string{destination}, subject)
	r.Send("templates/template.html", map[string]string{"pipelineName": os.Getenv("PIPELINE_NAME"), "link": os.Getenv("ENDPOINT")})
	os.Exit(0)
}
