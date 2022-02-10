package main

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spainishBot struct{}

func main() {
	en := englishBot{}
	sp := spainishBot{}
	printGreeting(en)
	printGreeting(sp)
}
func (englishBot) getGreeting() string {
	return "Hi!"
}
func (spainishBot) getGreeting() string {
	return "Hola!"
}
func printGreeting(b bot) {
	println(b.getGreeting())
}
