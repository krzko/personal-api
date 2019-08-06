package main

type Creation struct {
	Project
	URL string `json:url`
}

type Donations struct {
	Currencies []Project
}

type Person struct {
	FirstName string `json:first_name`
	LastName  string `json:last_name`
	Uptime    string `json:uptime`
}
type Project struct {
	Name    string `json:name`
	Address string `json:address`
}

type Projects struct {
	Projects []Project
}

type Social struct {
	GitHub    string `json:github`
	Instagram string `json:instagram`
	LinkedIn  string `json:linkedin`
	Strava    string `json:strava`
	Twitter   string `json:twitter`
}
