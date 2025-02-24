package types

type BranchIssue struct {
	BranchName  string `json:"BranchName"`
	IssueNumber string `json:"IssueNumber"`
}

type ProjectInfo struct {
	ProjectPath string        `json:"ProjectPath"`
	Branches    []BranchIssue `json:"Branches"`
}
