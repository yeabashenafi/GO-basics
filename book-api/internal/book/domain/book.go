package domain

import "io"

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// FileUploadReq holds the file content and any metadata sent alongside it
type FileUploadReq struct {
	FileName string
	Size     int64
	Content  io.Reader
}
