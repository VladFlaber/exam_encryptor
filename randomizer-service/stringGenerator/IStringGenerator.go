package stringGenerator

type IStringGenerator interface {
	CreateListOfStrings(count int, charset string) *[]string
}