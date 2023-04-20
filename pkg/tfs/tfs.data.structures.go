package tfs

import "time"

type TfsWorkItem struct {
	ID     int `json:"id,omitempty"`
	Rev    int `json:"rev,omitempty"`
	Fields struct {
		SystemAreaPath                                      string    `json:"System.AreaPath"`
		SystemTeamProject                                   string    `json:"System.TeamProject"`
		SystemIterationPath                                 string    `json:"System.IterationPath"`
		SystemWorkItemType                                  string    `json:"System.WorkItemType"`
		SystemState                                         string    `json:"System.State"`
		SystemReason                                        string    `json:"System.Reason"`
		SystemAssignedTo                                    string    `json:"System.AssignedTo"`
		SystemCreatedDate                                   time.Time `json:"System.CreatedDate"`
		SystemCreatedBy                                     string    `json:"System.CreatedBy"`
		SystemChangedDate                                   time.Time `json:"System.ChangedDate"`
		SystemChangedBy                                     string    `json:"System.ChangedBy"`
		SystemTitle                                         string    `json:"System.Title"`
		SystemBoardColumn                                   string    `json:"System.BoardColumn"`
		SystemBoardColumnDone                               bool      `json:"System.BoardColumnDone"`
		MicrosoftVSTSCommonStateChangeDate                  time.Time `json:"Microsoft.VSTS.Common.StateChangeDate"`
		MicrosoftVSTSCommonActivatedDate                    time.Time `json:"Microsoft.VSTS.Common.ActivatedDate"`
		MicrosoftVSTSCommonActivatedBy                      string    `json:"Microsoft.VSTS.Common.ActivatedBy"`
		MicrosoftVSTSCommonPriority                         int       `json:"Microsoft.VSTS.Common.Priority"`
		MicrosoftVSTSCommonStackRank                        float64   `json:"Microsoft.VSTS.Common.StackRank"`
		MicrosoftVSTSBuildIntegrationBuild                  string    `json:"Microsoft.VSTS.Build.IntegrationBuild"`
		MicrosoftVSTSSchedulingRemainingWork                float64   `json:"Microsoft.VSTS.Scheduling.RemainingWork"`
		MicrosoftVSTSSchedulingOriginalEstimate             float64   `json:"Microsoft.VSTS.Scheduling.OriginalEstimate"`
		RetalixCustomTemplateProductName                    string    `json:"Retalix.CustomTemplate.ProductName"`
		RetalixCustomTemplateRequestedRelease               string    `json:"Retalix.CustomTemplate.RequestedRelease"`
		RetalixCustomTemplateAgreedRelease                  string    `json:"Retalix.CustomTemplate.AgreedRelease"`
		RetalixCustomTemplateInterestCustomerNames          string    `json:"Retalix.CustomTemplate.InterestCustomerNames"`
		RetalixCustomTemplateDevManager                     string    `json:"Retalix.CustomTemplate.DevManager"`
		RetalixCustomTemplateCustomerRefNumber              string    `json:"Retalix.CustomTemplate.CustomerRefNumber"`
		RetalixCustomTemplateBudgetConstraint               string    `json:"Retalix.CustomTemplate.BudgetConstraint"`
		RetalixCustomTemplateFullyElaborated                string    `json:"Retalix.CustomTemplate.FullyElaborated"`
		RetalixCustomTemplateArchitectureOwner              string    `json:"Retalix.CustomTemplate.ArchitectureOwner"`
		RetalixCustomTemplateEpicOwner                      string    `json:"Retalix.CustomTemplate.EpicOwner"`
		RetalixCustomTemplateImpactAnnalysisRequired        string    `json:"Retalix.CustomTemplate.ImpactAnnalysisRequired"`
		RetalixCustomTemplateUI                             string    `json:"Retalix.CustomTemplate.UI"`
		RetalixCustomTemplateTheme                          string    `json:"Retalix.CustomTemplate.Theme"`
		RetalixCustomTemplateProbability                    int       `json:"Retalix.CustomTemplate.Probability"`
		RetalixCustomTemplateImpediment                     string    `json:"Retalix.CustomTemplate.Impediment"`
		WEF92978C6FEDFA4821A3FACDA41FB422ECKanbanColumn     string    `json:"WEF_92978C6FEDFA4821A3FACDA41FB422EC_Kanban.Column"`
		RetalixCustomTemplateQAOriginalEstimation           float64   `json:"Retalix.CustomTemplate.QAOriginalEstimation"`
		RetalixCustomTemplateBAOriginalEstimation           float64   `json:"Retalix.CustomTemplate.BAOriginalEstimation"`
		RetalixCustomTemplateExpectedDevEndIteration        time.Time `json:"Retalix.CustomTemplate.ExpectedDevEndIteration"`
		RetalixCustomTemplateCustomerEnhancement            string    `json:"Retalix.CustomTemplate.CustomerEnhancement"`
		RetalixCustomTemplateDetailsChangeDate              time.Time `json:"Retalix.CustomTemplate.DetailsChangeDate"`
		WEF92978C6FEDFA4821A3FACDA41FB422ECKanbanColumnDone bool      `json:"WEF_92978C6FEDFA4821A3FACDA41FB422EC_Kanban.Column.Done"`
		RetalixCustomTemplateActualEffort                   int       `json:"Retalix.CustomTemplate.Actual_Effort"`
		RetalixCustomTemplateClassification                 string    `json:"Retalix.CustomTemplate.Classification"`
		RetalixCustomTemplatePIName                         string    `json:"Retalix.CustomTemplate.PIName"`
		RetalixCustomTemplatePOOwner                        string    `json:"Retalix.CustomTemplate.POOwner"`
		RetalixCustomTemplateBackwardsCompatibility         string    `json:"Retalix.CustomTemplate.BackwardsCompatibility"`
		RetalixCustomTemplateAPIChanges                     string    `json:"Retalix.CustomTemplate.APIChanges"`
		SystemDescription                                   string    `json:"System.Description"`
		MicrosoftVSTSCommonAcceptanceCriteria               string    `json:"Microsoft.VSTS.Common.AcceptanceCriteria"`
		RetalixCustomTemplateImpactAnnalysisDetail          string    `json:"Retalix.CustomTemplate.ImpactAnnalysisDetail"`
		SystemTags                                          string    `json:"System.Tags"`
	} `json:"fields"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		WorkItemUpdates struct {
			Href string `json:"href"`
		} `json:"workItemUpdates"`
		WorkItemRevisions struct {
			Href string `json:"href"`
		} `json:"workItemRevisions"`
		WorkItemHistory struct {
			Href string `json:"href"`
		} `json:"workItemHistory"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		WorkItemType struct {
			Href string `json:"href"`
		} `json:"workItemType"`
		Fields struct {
			Href string `json:"href"`
		} `json:"fields"`
	} `json:"_links"`
	URL string `json:"url"`
}
