package iomanager

type IOManager interface {
	ReadLinesFromFile() ([]string, error)
	WriteJsonToFile(data interface{}) error
}
