package main

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/xanzy/go-gitlab"
)

func PrintGroupsTable(groups []*gitlab.Group) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"group-id", "name", "url"})
	for _, group := range groups {
		t.AppendRow(table.Row{group.ID, group.Name, group.WebURL})
		t.AppendSeparator()
	}
	t.Render()
}

func PrintItterationsTable(iterations []*gitlab.GroupIteration) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Title", "ID", "URL"})
	for _, iterations := range iterations {
		t.AppendRow(table.Row{iterations.Title, iterations.ID, iterations.WebURL})
		t.AppendSeparator()
	}
	t.Render()

}

func PrintIssuesTable(issues []*gitlab.Issue) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.AppendHeader(table.Row{"Title", "ID", "IID", "Description", "ProjectID", "Epic", "Weight", "webURL"})
	for _, issue := range issues {
		t.AppendRow(table.Row{issue.Title, issue.ID, issue.IID, issue.Description, issue.ProjectID, issue.Epic, issue.Weight, issue.WebURL})
		t.AppendSeparator()
	}
	t.Render()

}
