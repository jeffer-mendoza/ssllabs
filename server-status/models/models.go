package models

type Server struct {
	Address  string
	SslGrade string
	Country  string
	Owner    string
}

type HostInfo struct {
	ServersChanged   bool
	SslGrade         string
	PreviousSslGrade string
	Logo             string
	Title            string
	IsDown           bool
	Servers          []Server
}

type Response struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Protocol        string `json:"protocol"`
	IsPublic        bool   `json:"isPublic"`
	Status          string `json:"status"`
	StartTime       int64  `json:"startTime"`
	TestTime        int64  `json:"testTime"`
	EngineVersion   string `json:"engineVersion"`
	CriteriaVersion string `json:"criteriaVersion"`
	Endpoints       []struct {
		IPAddress         string `json:"ipAddress"`
		ServerName        string `json:"serverName"`
		StatusMessage     string `json:"statusMessage"`
		Grade             string `json:"grade"`
		GradeTrustIgnored string `json:"gradeTrustIgnored"`
		HasWarnings       bool   `json:"hasWarnings"`
		IsExceptional     bool   `json:"isExceptional"`
		Progress          int    `json:"progress"`
		Duration          int    `json:"duration"`
		Delegation        int    `json:"delegation"`
	} `json:"endpoints"`
}
