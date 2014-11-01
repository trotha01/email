package emain

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
)

func main() {

	to := parseCommandLine()
	subject := subjectPrompt()
	text := textPrompt()
	// text := "Hello, test message here. And link here, https://www.sendgrid.com/"

	printEmail(to, text, subject)
	fmt.Print("Y/n: ")

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		fmt.Errorf("Error pasring line. Error: %s", err.Error())
	}
	if response == "y" {
		sg := sendgrid.NewSendGridClient("trevor.rothaus@sendgrid.com", "DemoPassword!")
		message := sendgrid.NewMail()
		message.AddTo(to)
		message.AddToName("trevor")
		message.SetSubject(subject)
		message.SetText(text)
		message.SetFrom("trevor.rothaus@sendgrid.com")
		if r := sg.Send(message); r == nil {
			fmt.Println("Email sent!")
		} else {
			fmt.Println(r)
		}
	}
}

func printEmail(to string, subject string, text string) {
	fmt.Println("")
	fmt.Println("To: " + to)
	fmt.Println("Subject: " + text)
	fmt.Println("Body: " + subject)
	fmt.Println("")
}

func subjectPrompt() string {
	subject := ""
	fmt.Print("Subject: ")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		subject = scanner.Text()
		break
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("Error scanning subject. Error: %s", err.Error())
	}
	return subject
}

func textPrompt() string {
	text := ""
	currentline := ""
	fmt.Print("Body:\n")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		currentline = scanner.Text()
		if currentline == "" {
			break
		}
		text += currentline

	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("Error scanning subject. Error: %s", err.Error())
	}
	return text

}

// Returns to address
func parseCommandLine() string {
	to := os.Args[1]
	return to
}
