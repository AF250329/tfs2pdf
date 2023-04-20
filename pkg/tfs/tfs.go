// Package provide API and data structs to work with TFS
package tfs

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const (
	URL = "http://susday2582.corp.ncr.com:8080/tfs/DefaultCollection/_apis/wit/workitems/%d"
)

// Data structure that describe 'reference' in TFS
type TfsReference struct {
	ID           string
	WorkItemType string
	Title        string
	AssigendTo   string
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
	AcceptanceCriteria     string
	History                []HistoryItem
	AllLinks               []TfsReference
	ImpactAnalysis         string
}

// Function will communicate with TFS server and return technical details about provided TFS item ID
func ReadTfsItem(id int, tfsAddress, tfsToken string) (*Data, error) {

	tfsWorkItem, err := loadDataFromServer(id, tfsToken)
	if err != nil {
		return nil, err
	}

	data := convert(tfsWorkItem)

	return data, nil
}

func loadDataFromServer(id int, tfsToken string) (*TfsWorkItem, error) {

	tfsHttpClient := &TfsHttpClient{
		Token: tfsToken,
	}

	url := fmt.Sprintf(URL, id)

	resp, err := tfsHttpClient.GetJson(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	tfsWorkItem := &TfsWorkItem{}

	err = json.Unmarshal(body, tfsWorkItem)
	if err != nil {
		return nil, err
	}

	return tfsWorkItem, nil
}

func convert(t *TfsWorkItem) *Data {
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
