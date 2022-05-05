package models

type MyNode struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}
