package tfs

import "time"

type TfsWorkItem struct {
	ID     int `json:"id,omitempty"`
	Rev    int `json:"rev,omitempty"`
	Fields struct {
		SystemID                                            int       `json:"System.Id"`
		SystemAreaID                                        int       `json:"System.AreaId"`
		SystemAreaPath                                      string    `json:"System.AreaPath"`
		SystemTeamProject                                   string    `json:"System.TeamProject"`
		SystemNodeName                                      string    `json:"System.NodeName"`
		SystemAreaLevel1                                    string    `json:"System.AreaLevel1"`
		SystemRev                                           int       `json:"System.Rev"`
		SystemAuthorizedDate                                time.Time `json:"System.AuthorizedDate"`
		SystemRevisedDate                                   time.Time `json:"System.RevisedDate"`
		SystemIterationID                                   int       `json:"System.IterationId"`
		SystemIterationPath                                 string    `json:"System.IterationPath"`
		SystemIterationLevel1                               string    `json:"System.IterationLevel1"`
		SystemIterationLevel2                               string    `json:"System.IterationLevel2"`
		SystemIterationLevel3                               string    `json:"System.IterationLevel3"`
		SystemWorkItemType                                  string    `json:"System.WorkItemType"`
		SystemState                                         string    `json:"System.State"`
		SystemReason                                        string    `json:"System.Reason"`
		SystemAssignedTo                                    string    `json:"System.AssignedTo"`
		SystemCreatedDate                                   time.Time `json:"System.CreatedDate"`
		SystemCreatedBy                                     string    `json:"System.CreatedBy"`
		SystemChangedDate                                   time.Time `json:"System.ChangedDate"`
		SystemChangedBy                                     string    `json:"System.ChangedBy"`
		SystemAuthorizedAs                                  string    `json:"System.AuthorizedAs"`
		SystemPersonID                                      int       `json:"System.PersonId"`
		SystemWatermark                                     int       `json:"System.Watermark"`
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
		RetalixCustomTemplateComments                       string    `json:"Retalix.CustomTemplate.Comments"`
		RetalixCustomTemplateImpactAnnalysisDetail          string    `json:"Retalix.CustomTemplate.ImpactAnnalysisDetail"`
		SystemTags                                          string    `json:"System.Tags"`
	} `json:"fields"`
	Relations []struct {
		Rel string `json:"rel"`
		URL string `json:"url"`
		// Attributes []struct {
		// 	IsLocked             bool      `json:"isLocked,omitempty"`
		// 	AuthorizedDate       time.Time `json:"authorizedDate,omitempty"`
		// 	ID                   int       `json:"id,omitempty"`
		// 	ResourceCreatedDate  time.Time `json:"resourceCreatedDate,omitempty"`
		// 	ResourceModifiedDate time.Time `json:"resourceModifiedDate,omitempty"`
		// 	RevisedDate          time.Time `json:"revisedDate,omitempty"`
		// 	Comment              string    `json:"comment,omitempty"`
		// 	Name                 string    `json:"name,omitempty"`
		// } `json:"attributes,omitempty"`
	} `json:"relations"`
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

type WorkItemHistory struct {
	Count int `json:"count"`
	Value []struct {
		Rev       int    `json:"rev"`
		Value     string `json:"value"`
		RevisedBy struct {
			ID   string `json:"id"`
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"revisedBy"`
		RevisedDate time.Time `json:"revisedDate"`
		URL         string    `json:"url"`
	} `json:"value"`
}
