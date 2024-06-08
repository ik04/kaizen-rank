package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatih/color"
)

type Subject struct {
	Subject         string `json:"subject"`
	SubjectUUID     string `json:"subject_uuid"`
	AssignmentsCount int    `json:"assignments_count"`
}

type Subjects struct {
	Subjects []Subject `json:"subjects"`
}

func main() {
	res, err := http.Get("https://api.kaizenklass.me/api/v1/get-subjects")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Println(res.StatusCode)
		panic("request Unsuccessful")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var resp Subjects
	err = json.Unmarshal(body, &resp)
	if err != nil {
		panic(err)
	}

	color.Red("Top 3 Subjects:")
	for i := 0; i < 3 && i < len(resp.Subjects); i++ {
		result := fmt.Sprintf("%d) %-25s\n", i+1, resp.Subjects[i].Subject) 
		color.Cyan(result)
	}
}
