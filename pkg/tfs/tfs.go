// Package provide API and data structs to work with TFS
package tfs

import (
	"fmt"
	"strings"
	"time"
)

const (
	URL = "http://susday2582.corp.ncr.com:8080/tfs/DefaultCollection/_apis/wit/workitems/%d?$expand=all"
)

// Data structure that describe 'reference' in TFS
type TfsReference struct {
	ID           int
	WorkItemType string
	Title        string
	AssignedTo   string
	Status       string
	LinkComment  string
}

// Data structure that describe item from 'History'
type HistoryItem struct {
	Title string
	Date  string
}

// Data structure that contain all elements to fill template HTML page
type Data struct {
	WorkItemType           string
	EpicTitle              string
	OriginalEstimate       float64
	QaOriginalEstimate     float64
	RequestedRelease       string
	Priority               int
	ActualEffort           int
	BAOriginalEstimate     float64
	AgreedRelease          string
	EpicOwner              string
	ArchitectureOwner      string
	AssignedTo             string
	POOwner                string
	DevManager             string
	State                  string
	Impediment             string
	ProductSuite           string
	Reason                 string
	PSPriority             string
	Area                   string
	IsFullyElaborated      string
	Classification         string
	So                     string
	Ui                     string
	PiName                 string
	Theme                  string
	Description            string
	HasImplementationItems bool
	ImplementationItems    []TfsReference
	HasFeatureItems        bool
	FeatureItems           []TfsReference
	HasAcceptanceCriteria  bool
	AcceptanceCriteria     string
	HasHistoryItems        bool
	History                []HistoryItem
	HasAllLinks            bool
	AllLinks               []TfsReference
	HasImpactAnalysis      bool
	ImpactAnalysis         string
}

// Function will communicate with TFS server and return technical details about provided TFS item ID
func ReadTfsItem(id int, tfsAddress, tfsToken string) (*Data, error) {

	tfsHttpClient := &TfsHttpClient{
		Token: tfsToken,
	}

	url := fmt.Sprintf(URL, id)

	tfsWorkItem, err := tfsHttpClient.GetWorkItem(url)
	if err != nil {
		return nil, err
	}

	data := convert(tfsWorkItem, tfsHttpClient)

	return data, nil
}

func convert(t *TfsWorkItem, client *TfsHttpClient) *Data {
	data := &Data{
		WorkItemType:       t.Fields.SystemWorkItemType,
		EpicTitle:          fmt.Sprintf("%s %d: %s", t.Fields.SystemWorkItemType, t.ID, t.Fields.SystemTitle),
		OriginalEstimate:   t.Fields.MicrosoftVSTSSchedulingOriginalEstimate,
		ActualEffort:       t.Fields.RetalixCustomTemplateActualEffort,
		QaOriginalEstimate: t.Fields.RetalixCustomTemplateQAOriginalEstimation,
		BAOriginalEstimate: t.Fields.RetalixCustomTemplateBAOriginalEstimation,
		RequestedRelease:   t.Fields.RetalixCustomTemplateRequestedRelease,
		AgreedRelease:      t.Fields.RetalixCustomTemplateAgreedRelease,
		Priority:           t.Fields.MicrosoftVSTSCommonPriority,
		EpicOwner:          t.Fields.RetalixCustomTemplateEpicOwner,
		POOwner:            t.Fields.RetalixCustomTemplatePOOwner,
		ArchitectureOwner:  t.Fields.RetalixCustomTemplateArchitectureOwner,
		DevManager:         t.Fields.RetalixCustomTemplateDevManager,
		AssignedTo:         t.Fields.SystemAssignedTo,
		State:              t.Fields.SystemState,
		Reason:             t.Fields.SystemReason,
		IsFullyElaborated:  t.Fields.RetalixCustomTemplateFullyElaborated,
		Ui:                 t.Fields.RetalixCustomTemplateUI,
		Impediment:         t.Fields.RetalixCustomTemplateImpediment,
		PSPriority:         "", // ??
		Classification:     t.Fields.RetalixCustomTemplateClassification,
		PiName:             t.Fields.RetalixCustomTemplatePIName,
		ProductSuite:       t.Fields.RetalixCustomTemplateProductName,
		Area:               t.Fields.SystemAreaPath,
		So:                 "", // ???
		Theme:              t.Fields.RetalixCustomTemplateTheme,
		Description:        strings.ReplaceAll(t.Fields.SystemDescription, "susday2582", "susday2582.corp.ncr.com"),
		AcceptanceCriteria: strings.ReplaceAll(t.Fields.MicrosoftVSTSCommonAcceptanceCriteria, "susday2582", "susday2582.corp.ncr.com"),
	}

	if t.Fields.RetalixCustomTemplateImpactAnnalysisDetail != "" {
		data.HasImpactAnalysis = true
		data.ImpactAnalysis = t.Fields.RetalixCustomTemplateImpactAnnalysisDetail
	}

	for _, val := range t.Relations {

		switch val.Rel {

		case "System.LinkTypes.Hierarchy-Forward":

			sourceItem, err := client.GetWorkItem(val.URL)
			if err != nil {
				panic(err)
			}

			// This is 'Implementation' link
			// https://learn.microsoft.com/en-us/azure/devops/boards/queries/link-type-reference?view=azure-devops
			data.ImplementationItems = append(data.ImplementationItems, TfsReference{
				ID:           sourceItem.ID,
				WorkItemType: sourceItem.Fields.SystemWorkItemType,
				Title:        sourceItem.Fields.SystemTitle,
				AssignedTo:   sourceItem.Fields.SystemAssignedTo,
				Status:       sourceItem.Fields.SystemState,
				LinkComment:  "", // ??
			})

			data.HasImplementationItems = true

		case "Hyperlink":
			// This is link to external resource (sharepoint for example)
			data.AllLinks = append(data.AllLinks, TfsReference{
				ID:           0,
				WorkItemType: "Hyperlink",
				Title:        "Hyperlink",
				AssignedTo:   "",
				Status:       "",
				LinkComment:  val.URL, // ??
			})

		case "ArtifactLink":
			// This is special link - directly to source code
			// Don't need to get this item

		default:

			sourceItem, err := client.GetWorkItem(val.URL)
			if err != nil {
				panic(err)
			}

			// All links
			data.AllLinks = append(data.AllLinks, TfsReference{
				ID:           sourceItem.ID,
				WorkItemType: sourceItem.Fields.SystemWorkItemType,
				Title:        sourceItem.Fields.SystemTitle,
				AssignedTo:   sourceItem.Fields.SystemAssignedTo,
				Status:       sourceItem.Fields.SystemState,
				LinkComment:  "", // ??
			})
		}
	}

	historyCollection, err := client.GetHistoryLinks(t.Links.WorkItemHistory.Href)
	if err != nil {
		panic(err)
	}

	if len(historyCollection.Value) > 0 {
		data.HasHistoryItems = true
	}

	for _, val := range historyCollection.Value {
		data.History = append(data.History, HistoryItem{
			Title: val.Value,
			Date:  val.RevisedDate.Format(time.RFC3339),
		})
	}

	return data
}

// func openConnection(tfsAddress, tfsToken string) (*tfvc.Client, error) {
// 	tfsConnection := azuredevops.NewPatConnection(tfsAddress, tfsToken)
//
// 	client, err := tfvc.NewClient(context.Background(), tfsConnection)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not open connection to TFS server at %v.\nAdditional information: %w", tfsAddress, err)
// 	}
//
// 	return &client, nil
// }

// func getData(client tfvc.Client) (*Data, error) {
//
// 	item, err := client.GetItem(context.Background(), tfvc.GetItemArgs{})
//
// 	data := &Data{
// 		EpicTitle: "Epic title",
// 	}
//
// 	return data, nil
// }
