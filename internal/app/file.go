package app

import "mime/multipart"

func GetFileName(file_header *multipart.FileHeader) string {
	return file_header.Filename
}
