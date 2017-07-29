package globals

import (
	"time"
	"mime/multipart"
)

type File struct {
	ID        int64
	Path      string
	Name      string
	MimeType  string
	Thumb     string
	SHA1      string
	Size      int64
	UpdatedAt time.Time
	CreatedAt time.Time
}

type MultipartFile struct {
	File   multipart.File
	Header *multipart.FileHeader
}

type Fileser interface {
	Get(fileID int64) File
	GetThumb(fileID, width, height int64) (string, error)
	ImgSize(path string) (width, height int, err error)
	UploadRemote(url string) (File, error)
	Upload(mFile MultipartFile) (File, error)
	Clear(fileID int64)
}
