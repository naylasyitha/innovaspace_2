package storage

import (
	// "mime/multipart"

	"innovaspace/internal/infra/env"
	"mime/multipart"

	supabasestorageuploader "github.com/adityarizkyramadhan/supabase-storage-uploader"
)

type StorageSupabase struct {
	Client *supabasestorageuploader.Client
}

func NewStorageSupabase() *StorageSupabase {
	return &StorageSupabase{
		Client: supabasestorageuploader.New(
			env.GetEnv("SUPABASE_ENDPOINT"),
			env.GetEnv("SUPABASE_TOKEN"),
			env.GetEnv("SUPABASE_BASKET_NAME"),
		),
	}
}

func (s StorageSupabase) UploadProfilePicture(UserId string, file *multipart.FileHeader) (string, error) {
	url, err := s.Client.Upload(file)
	if err != nil {
		return "", err
	}

	return url, nil
}
