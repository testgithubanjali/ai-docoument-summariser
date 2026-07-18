package dto

type ProcessDocumentResponse struct {
	ID       uint   `json:"id"`
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	Summary  string `json:"summary"`
	Message  string `json:"message"`
}
