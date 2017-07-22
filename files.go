package globals

import (
	"time"
	"mime/multipart"
)

type File struct {
	ID int64
	Path string
	Name string
	MimeType string
	Thumb string
	SHA1 string
	Size int64
	UpdatedAt time.Time
	CreatedAt time.Time
}


type MultipartFile struct {
	File multipart.File
	Header *multipart.FileHeader
}
