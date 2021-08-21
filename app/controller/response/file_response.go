package response

type FileResponse struct {
	ID   int    `json:"id"`
	Type string `json:"file_type"`
	Path string `json:"file_path"`
}
