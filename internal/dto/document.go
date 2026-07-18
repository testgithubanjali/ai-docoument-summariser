package dto

type UploadResponse struct {
	ID       uint   `json:"id"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	Message  string `json:"message"`
}
