package main

type Error struct {
	Code    int    `json:"code"`
	Err     string `json:"error"`
	Message string `json:"description"`
}
