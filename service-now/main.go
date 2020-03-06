package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const serviceNowInstance = "https://dev100979.service-now.com"
const tableApiPath = "/api/now/table"
const serviceName = "/incident"
const userName = "admin"
const password = "cxEgPTkRw5Y4"
const GET = "GET"
const POST = "POST"

func main() {
	fmt.Println("Ready to connect with Service Now")

	for {
		fmt.Println("What operation would like to do? Please enter")
		fmt.Println("1. Get Call")
		fmt.Println("2. Post Call")
		fmt.Println("3. To Exit")
		answerCh := make(chan string)
		go func() {
			var answer string
			_, err := fmt.Scanf("%s\n", &answer)
			errorMessage(err, "Error occured while reading input")
			answerCh <- answer
		}()
		select {
		case answer := <-answerCh:
			if answer == "1" {
				var limit string
				fmt.Println("Please enter expected count")
				fmt.Scanf("%s\n", &limit)
				getTableResponse(limit)
			} else if answer == "2" {
				var shortDesc string
				fmt.Println("Please enter short problem description")
				fmt.Scanf("%s\n", &shortDesc)
				postResponse(shortDesc)
			} else if answer == "3" {
				os.Exit(1)
			} else {
				fmt.Println("Please enter correct value")
			}

		}
	}

}

func getTableResponse(limit string) {
	url := fmt.Sprintf(serviceNowInstance+tableApiPath+serviceName+"?sysparm_limit=%s", limit)

	request, err := http.NewRequest(GET, url, nil)
	request.SetBasicAuth(userName, password)
	client := &http.Client{}
	response, err := client.Do(request)
	errorMessage(err, "The HTTP GET request failed with error")

	data, _ := ioutil.ReadAll(response.Body)

	var result GetResponse
	err = json.Unmarshal(data, &result)
	errorMessage(err, "Error occurred while json unmarshalling ")
	if err != nil {
		// Optional
		//PrettyPrint(&result)
		for _, value := range result.Result {
			fmt.Println("Sys ID :: ", value.Sysid)
			fmt.Println("Short Description :: ", value.Shortdescription)
			fmt.Println("Opened at :: ", value.Openedat)
			fmt.Println("Opened By :: ", value.OpenedBy.Value)
			fmt.Println("Impact ::", value.Impact)
			fmt.Println("Priority:: ", value.Priority)
		}
	}
}

func postResponse(shortDesc string) {
	incidentData := IncidentRequest{
		ShortDescription: shortDesc,
		AssignmentGroup:  "287ebd7da9fe198100f92cc8d1d2154e",
		Urgency:          "2",
		Impact:           "2",
	}
	jsonValue, _ := json.Marshal(incidentData)
	buffer := bytes.NewBuffer(jsonValue)

	url := fmt.Sprintf(serviceNowInstance + tableApiPath + serviceName)

	request, err := http.NewRequest(POST, url, buffer)
	request.SetBasicAuth(userName, password)
	client := &http.Client{}
	response, err := client.Do(request)

	errorMessage(err, "The HTTP POST request failed with error")

	data, _ := ioutil.ReadAll(response.Body)

	var result Result
	err = json.Unmarshal(data, &result)
	errorMessage(err, "Error occurred while json unmarshalling ")
	if err != nil {
		// Optional
		//PrettyPrint(&result)
		fmt.Println("Incident ID :: ", result.Result.Correlationid)
		fmt.Println("Short Description :: ", result.Result.Shortdescription)
		fmt.Println("Opened at :: ", result.Result.Openedat)
		fmt.Println("Opened By :: ", result.Result.OpenedBy.Value)
		fmt.Println("Impact ::", result.Result.Impact)
		fmt.Println("Priority:: ", result.Result.Priority)
	}

}

func errorMessage(err error, msg string) {
	if err != nil {
		fmt.Printf(msg+"%s\n", err)
	}
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

type IncidentRequest struct {
	ShortDescription string `json:"short_description"`
	AssignmentGroup  string `json:"assignment_group"`
	Urgency          string `json:"urgency"`
	Impact           string `json:"impact"`
}

type GetResponse struct {
	Result []IncidentResponse `json:"result"`
}

type Result struct {
	Result IncidentResponse `json:"result"`
}

type IncidentResponse struct {
	Uponapproval           string       `json:"upon_approval"`
	Location               string       `json:"location"`
	Expectedstart          string       `json:"expected_start"`
	Reopencount            string       `json:"reopen_count"`
	Closenotes             string       `json:"close_notes"`
	Additionalassigneelist string       `json:"additional_assignee_list"`
	Impact                 string       `json:"impact"`
	Urgency                string       `json:"urgency"`
	Correlationid          string       `json:"correlation_id"`
	SysDomain              LinkAndValue `json:"sys_domain"`
	Description            string       `json:"description"`
	Grouplist              string       `json:"group_list"`
	Priority               string       `json:"priority"`
	Deliveryplan           string       `json:"delivery_plan"`
	Sysmodcount            string       `json:"sys_mod_count"`
	Worknoteslist          string       `json:"work_notes_list"`
	Businessservice        LinkAndValue `json:"business_service"`
	Followup               string       `json:"follow_up"`
	Closedat               string       `json:"closed_at"`
	Sladue                 string       `json:"sla_due"`
	Deliverytask           string       `json:"delivery_task"`
	Sysupdatedon           string       `json:"sys_updated_on"`
	Parent                 string       `json:"parent"`
	Workend                string       `json:"work_end"`
	Number                 string       `json:"number"`
	Closedby               LinkAndValue `json:"closed_by"`
	Workstart              string       `json:"work_start"`
	Calendarstc            string       `json:"calendar_stc"`
	Category               string       `json:"category"`
	Businessduration       string       `json:"business_duration"`
	Incidentstate          string       `json:"incident_state"`
	Activitydue            string       `json:"activity_due"`
	Correlationdisplay     string       `json:"correlation_display"`
	Company                LinkAndValue `json:"company"`
	Active                 string       `json:"active"`
	Duedate                string       `json:"due_date"`
	AssingmentGroup        LinkAndValue `json:"assignment_group"`
	Callerid               LinkAndValue `json:"caller_id"`
	Knowledge              string       `json:"knowledge"`
	Madesla                string       `json:"made_sla"`
	Commentsandworkn       string       `json:"comments_and_work_notes"`
	Parentincident         string       `json:"parent_incident"`
	State                  string       `json:"state"`
	Userinput              string       `json:"user_input"`
	Syscreatedon           string       `json:"sys_created_on"`
	Approvalset            string       `json:"approval_set"`
	Reassignmentcoun       string       `json:"reassignment_count"`
	Rfc                    string       `json:"rfc"`
	Childincidents         string       `json:"child_incidents"`
	Openedat               string       `json:"opened_at"`
	Shortdescription       string       `json:"short_description"`
	Order                  string       `json:"order"`
	Sysupdatedby           string       `json:"sys_updated_by"`
	Resolvedby             LinkAndValue `json:"resolved_by"`
	Notify                 string       `json:"notify"`
	Uponreject             string       `json:"upon_reject"`
	Approvalhistory        string       `json:"approval_history"`
	Problemid              string       `json:"problem_id"`
	Worknotes              string       `json:"work_notes"`
	Calendarduration       string       `json:"calendar_duration"`
	Closecode              string       `json:"close_code"`
	Sysid                  string       `json:"sys_id"`
	Approval               string       `json:"approval"`
	Causedby               string       `json:"caused_by"`
	Severity               string       `json:"severity"`
	Syscreatedby           string       `json:"sys_created_by"`
	Resolvedat             string       `json:"resolved_at"`
	Assignedto             LinkAndValue `json:"assigned_to"`
	Businessstc            string       `json:"business_stc"`
	Wfactivity             string       `json:"wf_activity"`
	Sysdomainpath          string       `json:"sys_domain_path"`
	Cmdbci                 LinkAndValue `json:"cmdb_ci"`
	OpenedBy               LinkAndValue `json:"opened_by"`
	Subcategory            string       `json:"subcategory"`
	Rejectiongoto          string       `json:"rejection_goto"`
	Sysclassname           string       `json:"sys_class_name"`
	Watchlist              string       `json:"watch_list"`
	Timeworked             string       `json:"time_worked"`
	Contacttype            string       `json:"contact_type"`
	Escalation             string       `json:"escalation"`
	Comments               string       `json:"comments"`
}

type LinkAndValue struct {
	Link  string `json:"link"`
	Value string `json:"value"`
}
