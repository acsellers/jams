package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

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
	days := make([][]string, 7)
	for _, job := range jobs {
		if job.AutoSubmit {
			if needsZero.MatchString(job.ScheduledTime) {
				job.ScheduledTime = "0" + job.ScheduledTime
			}
			indexes := GetDays(job.ScheduledDate, job.ExceptForDate)
			for _, ix := range indexes {
				days[ix] = append(days[ix], fmt.Sprintf("%s: %d J::%s\\%s", job.ScheduledTime, job.JobID, job.ParentFolderName, job.JobName))
			}
		}
	}
	setups, err := c.Setup.Setups(ctx)
	if err != nil {
		log.Fatal("Setups: ", err)
	}
	for _, setup := range setups {
		fmt.Println(setup.SetupName, setup.Jobs)
		if setup.AutoSubmit {
			fullSetup, _ := c.Setup.SetupByID(ctx, setup.SetupID)
			if needsZero.MatchString(setup.ScheduledTime) {
				setup.ScheduledTime = "0" + setup.ScheduledTime
			}
			indexes := GetDays(setup.ScheduledDate, setup.ExceptForDate)
			for _, ix := range indexes {
				for _, job := range fullSetup.Jobs {
					days[ix] = append(days[ix], fmt.Sprintf("%s: %03d S::%s\\%s", setup.ScheduledTime, job.JobID, setup.ParentFolderName, job.JobName))
				}

			}
		}
	}
	for _, jobs := range days {
		sort.Strings(jobs)
	}
	fmt.Println("\nSunday\n==================")
	for _, job := range days[0] {
		fmt.Println(job)
	}
	fmt.Println("\nMonday\n==================")
	for _, job := range days[1] {
		fmt.Println(job)
	}
	fmt.Println("\nTuesday\n==================")
	for _, job := range days[2] {
		fmt.Println(job)
	}
	fmt.Println("\nWednesday\n==================")
	for _, job := range days[3] {
		fmt.Println(job)
	}
	fmt.Println("\nThursday\n==================")
	for _, job := range days[4] {
		fmt.Println(job)
	}
	fmt.Println("\nFriday\n==================")
	for _, job := range days[5] {
		fmt.Println(job)
	}
	fmt.Println("\nSaturday\n==================")
	for _, job := range days[6] {
		fmt.Println(job)
	}

}

func GetDays(date, except string) []int {
	date = strings.ToLower(date)
	except = strings.ToLower(except)
	indexes := make(map[int]bool)
	if date == "daily" {
		indexes = map[int]bool{0: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true}
	} else {
		days := strings.Split(date, ",")
		for _, day := range days {
			indexes[ToIndex(day)] = true
		}
	}
	exceptions := strings.Split(except, ",")
	for _, exc := range exceptions {
		indexes[ToIndex(exc)] = false
	}
	days := []int{}
	for d, v := range indexes {
		if v {
			days = append(days, d)
		}
	}

	return days
}
func ToIndex(day string) int {
	day = strings.TrimSpace(day)
	switch day {
	case "sunday":
		return 0
	case "monday":
		return 1
	case "tuesday":
		return 2
	case "wednesday":
		return 3
	case "thursday":
		return 4
	case "friday":
		return 5
	case "saturday":
		return 6
	}
	return -1
}
