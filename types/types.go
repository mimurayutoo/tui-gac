package types

// json形式で保存されている情報の型
type BranchIssue struct {
	BranchName  string `json:"BranchName"`
	IssueNumber string `json:"IssueNumber"`
}

type ProjectInfo struct {
	ProjectPath string        `json:"ProjectPath"`
	Branches    []BranchIssue `json:"Branches"`
}
