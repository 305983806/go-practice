package chatbot

type Talk interface {
	Hello() string
	Talk(heard string) (saying string, end bool, err error)
}

type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

var chatbotMap = map[string]Chatbot{}

