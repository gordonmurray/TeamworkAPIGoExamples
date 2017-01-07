package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type tasks struct {
	STATUS    string `json:"STATUS"`
	Tasks []struct {
		ID                        string `json:"id"`
		CanComplete               bool `json:"canComplete"`
		CommentsCount             int `json:"comments-count"`
		Description               string `json:"description"`
		HasReminders              bool `json:"has-reminders"`
		HasUnreadComments         bool `json:"has-unread-comments"`
		Private                   int `json:"private"`
		Content                   string `json:"content"`
		Order                     int `json:"order"`
		ProjectID                 int `json:"project-id"`
		ProjectName               string `json:"project-name"`
		TodoListID                int `json:"todo-list-id"`
		TodoListName              string `json:"todo-list-name"`
		TasklistPrivate           bool `json:"tasklist-private"`
		TasklistIsTemplate        bool `json:"tasklist-isTemplate"`
		Status                    string `json:"status"`
		CompanyName               string `json:"company-name"`
		CompanyID                 int `json:"company-id"`
		CreatorID                 int `json:"creator-id"`
		CreatorFirstname          string `json:"creator-firstname"`
		CreatorLastname           string `json:"creator-lastname"`
		Completed                 bool `json:"completed"`
		StartDate                 string `json:"start-date"`
		DueDateBase               string `json:"due-date-base"`
		DueDate                   string `json:"due-date"`
		CreatedOn                 time.Time `json:"created-on"`
		LastChangedOn             time.Time `json:"last-changed-on"`
		Position                  int `json:"position"`
		EstimatedMinutes          int `json:"estimated-minutes"`
		Priority                  string `json:"priority"`
		Progress                  int `json:"progress"`
		HarvestEnabled            bool `json:"harvest-enabled"`
		ParentTaskID              string `json:"parentTaskId"`
		LockdownID                string `json:"lockdownId"`
		TasklistLockdownID        string `json:"tasklist-lockdownId"`
		HasDependencies           int `json:"has-dependencies"`
		HasPredecessors           int `json:"has-predecessors"`
		HasTickets                bool `json:"hasTickets"`
		TimeIsLogged              string `json:"timeIsLogged"`
		AttachmentsCount          int `json:"attachments-count"`
		ResponsiblePartyIds       string `json:"responsible-party-ids,omitempty"`
		ResponsiblePartyID        string `json:"responsible-party-id,omitempty"`
		ResponsiblePartyNames     string `json:"responsible-party-names,omitempty"`
		ResponsiblePartyType      string `json:"responsible-party-type,omitempty"`
		ResponsiblePartyFirstname string `json:"responsible-party-firstname,omitempty"`
		ResponsiblePartyLastname  string `json:"responsible-party-lastname,omitempty"`
		ResponsiblePartySummary   string `json:"responsible-party-summary,omitempty"`
		Predecessors              []interface{} `json:"predecessors"`
		CanEdit                   bool `json:"canEdit"`
		ViewEstimatedTime         bool `json:"viewEstimatedTime"`
		CreatorAvatarURL          string `json:"creator-avatar-url"`
		CanLogTime                bool `json:"canLogTime"`
		CommentFollowerSummary    string `json:"commentFollowerSummary,omitempty"`
		CommentFollowerIds        string `json:"commentFollowerIds,omitempty"`
		UserFollowingComments     bool `json:"userFollowingComments"`
		UserFollowingChanges      bool `json:"userFollowingChanges"`
		DLM                       int `json:"DLM"`
		Tags                      []struct {
			ID    string `json:"id"`
			Name  string `json:"name"`
			Color string `json:"color"`
		} `json:"tags,omitempty"`
		CompleterFirstname        string `json:"completer_firstname,omitempty"`
		CompleterLastname         string `json:"completer_lastname,omitempty"`
		CompleterID               string `json:"completer_id,omitempty"`
	} `json:"tasks"`
}

func main() {

	url := fmt.Sprintf("https://xxxxxxxxxx.teamwork.com/completedtasks.json?page=1&pageSize=250&startDate=xxxxxxxxxx&endDate=xxxxxxxxxx&userId=xxxxxxxxxx&type=COMPLETEDBY&dateType=COMPLETEDBETWEEN&includeArchivedProjects=true&sortBy=dateCompleted&sortOrder=asc&filterBy=assignedByAnyoneCompleted")

	req, err := http.NewRequest("GET", url, nil)

	req.SetBasicAuth("xxxxxxxxxx", "")

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// HTTP client
	client := &http.Client{}

	// Send the request via a client
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// close resp.Body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record tasks

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	if (record.STATUS == "OK") {

		for _, element := range record.Tasks {
			fmt.Println(element.ID + " " + element.Content)
		}

	}
}
