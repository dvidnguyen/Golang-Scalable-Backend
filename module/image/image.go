package image

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// CREATE TABLE `images` (
// `id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
// `title` varchar(150) DEFAULT NULL,
// `file_name` varchar(255) DEFAULT NULL,
// `file_size` int DEFAULT NULL,
// `file_type` varchar(20) DEFAULT NULL,
// `storage_provider` enum('aws_s3','local') DEFAULT NULL,
// `status_enum` enum('activated','uploaded','deleted') DEFAULT 'uploaded',
// `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
// `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
// PRIMARY KEY (`id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0
const (
	TbName          = "images"
	ProviderAWSS3   = "aws_s3"
	StatusUploaded  = "uploaded"
	StatusActivated = "activated"
	StatusDeleted   = "deleted"
)

type Image struct {
	Id              uuid.UUID `json:"id"`
	Title           string    `json:"title"`
	FileName        string    `json:"file_name"`
	FileUrl         string    `json:"file_url" gorm:"-"`
	FileSize        int       `json:"file_size"`
	FileType        string    `json:"file_type"`
	StorageProvider string    `json:"storage_provider"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func NewImage(id uuid.UUID, title string, fileName string, fileUrl string, fileSize int, fileType string, storageProvider string, status string, createdAt time.Time, updatedAt time.Time) *Image {
	return &Image{Id: id, Title: title, FileName: fileName, FileUrl: fileUrl, FileSize: fileSize, FileType: fileType, StorageProvider: storageProvider, Status: status, CreatedAt: createdAt, UpdatedAt: updatedAt}
}

func (img Image) SetCDNDomain(domain string) bool {
	img.FileUrl = fmt.Sprintf("%s/%s", domain, img.FileName)
}
