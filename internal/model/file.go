package models

type File struct {
	*Model
	ID   string
	Name string
	Type FileType
	Size int64
	Path string
}

type FileType int

const (
	FileTypeImage    FileType = iota // 图片类型
	FileTypeVideo                    // 视频类型
	FileTypeDocument                 // 文档类型
)

// FileTypeNameMap maps file type constants to their corresponding names
var FileTypeNameMap = map[FileType]string{
	FileTypeImage:    "image",
	FileTypeVideo:    "video",
	FileTypeDocument: "document",
	// 其他文件类型的映射
}

func GetFileTypeName(fileType FileType) string {
	return FileTypeNameMap[fileType]
}
