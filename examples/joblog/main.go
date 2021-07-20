package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/acsellers/jams"
)

var (
	needsZero = regexp.MustCompile(`^[0-9]:[0-9][0-9]`)
)

func main() {
	ctx := context.Background()
	c := jams.NewAPIClient(&jams.Configuration{
		BasePath:  "/JAMS/",
		Host:      os.Getenv("HOST"),
		Scheme:    "http",
		UserAgent: "Golang Client",
	})
	_, err := c.Login(ctx, jams.LoginData{
		Username: os.Getenv("JUSER"),
		Password: os.Getenv("JPASS"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to JAMS")
	jobs, err := c.Job.Jobs(ctx)
	fmt.Println("Got", len(jobs), "Jobs")
	if err != nil {
		log.Fatal("Jobs: ", err)
	}

	for _, job := range jobs {
		if strings.Contains(job.JobName, "Test Client Job") {
			fmt.Println(job.JobID, job.JobName)

			history, err := c.History.History(ctx, jams.HistorySearch{
				JobName: job.JobName,
			})
			if err != nil {
				log.Fatal("History: ", err)
			}
			if len(history) == 0 {
				fmt.Println("No history for test job, running")
				info, err := c.Submit.SubmitInfoByJobID(ctx, job.JobID)
				if err != nil {
					log.Fatal("SubmitInfo: ", err)
				}
				_, err = c.Submit.SubmitJob(ctx, info)
				if err != nil {
					if je, ok := err.(jams.Error); ok {
						f, err := os.Create("error.html")
						if err != nil {
							log.Fatal(err)
						}
						_, err = f.Write(je.Body)
						if err != nil {
							log.Fatal("WriteError: ", err)
						}
						f.Close()
					}
					log.Fatal(err)
				}
				fmt.Println("Waiting a little while for the job to go")
				time.Sleep(10 * time.Second)
				history, err = c.History.History(ctx, jams.HistorySearch{
					JobName: job.JobName,
				})
			}
			for _, record := range history {
				joblog, err := c.History.JobLog(ctx, record.JobName, record.Ron, record.RestartCount, true)
				if err != nil {
					log.Fatal("JobLog: ", err)
				}
				fmt.Println(joblog)
			}
		}
	}

}
