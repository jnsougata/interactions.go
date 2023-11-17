package main

type File struct {
	Name    string `json:"name"`
	Content []byte `json:"-"`
	
}
