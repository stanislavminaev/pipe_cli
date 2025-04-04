package jira_wick

type JiraWickClient interface {
	CloseRelease(req *CloseReleaseRequest) (*CloseReleaseResponse, error)
}
