package interfaces

type FileStorageProviderInterface interface {
	Upload(fileName string, filePath string) error
}
