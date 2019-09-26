package services

var documents [2]string

func GetAllDocuments() [2]string {
	return documents
}

func CreateDocument(newDocument string) string {
	documents[0] = newDocument
	return newDocument
}
