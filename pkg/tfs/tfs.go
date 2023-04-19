package tfs

type TfsReference struct {
	ID           string
	WorkItemType string
	Title        string
	AssigendTo   string
	Status       string
	LinkComment  string
}

type HistoryItem struct {
	Title string
	Date  string
}

type Data struct {
	EpicTitle              string
	OriginalEstimate       string
	QaOriginalEstimate     string
	RequestedRelease       string
	Priority               string
	ActualEffort           string
	BAOriginalEstimate     string
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

func ReadTfsItem(id int) *Data {

	data := &Data{
		EpicTitle: "Epic title",
	}

	return data
}
