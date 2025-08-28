package models

type Users struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Sid       string `json:"sid"`
	Ipaddress string `json:"ip_address"`
}
