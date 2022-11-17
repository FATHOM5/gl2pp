package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/xanzy/go-gitlab"
)

func translate(GitlabIssue []*gitlab.Issue) [][]string {
	var planningpokerissues [][]string
	for _, GitIssue := range GitlabIssue {
		var PlanningPoker []string
		GitIssueIIDToString := strconv.Itoa(GitIssue.IID)
		PlanningPoker = append(PlanningPoker, GitIssueIIDToString)
		PlanningPoker = append(PlanningPoker, GitIssue.Title)
		PlanningPoker = append(PlanningPoker, GitIssue.WebURL)
		PlanningPoker = append(PlanningPoker, "")
		GitIssueWeightToString := strconv.Itoa(GitIssue.Weight)
		PlanningPoker = append(PlanningPoker, GitIssueWeightToString)
		planningpokerissues = append(planningpokerissues, PlanningPoker)
	}
	return planningpokerissues
}

func Exporttocsv(planningpokerissues [][]string, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	planningPokerissuesheader := []string{"Issue Key", "Summary", "Description", "Acceptance Criteria", "Story Points"}
	if err != nil {
		return err
	}
	csvwriter := csv.NewWriter(file)
	err = csvwriter.Write(planningPokerissuesheader)
	if err != nil {
		return err
	}
	for _, issues := range planningpokerissues {
		err = csvwriter.Write(issues)
	}
	csvwriter.Flush()
	return nil
}
