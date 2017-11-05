package main

import (
	"Prime/bundles/chatbot"
	"Prime/bundles/wolfram_alpha"
	"errors"
	"fmt"
	"log"
	"os"
	// Autoload environment variables in .env
	_ "github.com/joho/godotenv/autoload"
)

func chatbotProcess(session chatbot.Session, message string) (string, error) {

	// Make sure a history key is defined in the session which points to a slice of strings
	_, historyFound := session["history"]
	if !historyFound {
		session["history"] = []string{}
	}

	// Fetch the history from session and cast it to an array of strings
	history, _ := session["history"].([]string)

	// Add the message in the parsed body to the messages in the session
	history = append(history, message)

	// Form a sentence out of the history in the form Message 1, Message 2, and Message 3
	l := len(history)
	wordsForSentence := make([]string, l)
	copy(wordsForSentence, history)
	if l > 1 {
		wordsForSentence[l-1] = "and " + wordsForSentence[l-1]
	}

	// Save the updated history to the session
	session["history"] = history

	// Initialize controller
	wa := &wolfram_alpha.WolframAlphaController{}

	// Use Wolfram Alpha API
	response, err := wa.HandleQuery(message)

	// Check for errors
	if err != nil || response.QueryResult.Success != true {
		return "", errors.New("Your Query " + message + " is invalid!")
	}

	// Return Query results
	return response.QueryResult.Pods[0].SubPods[0].PlainText, nil
}

func main() {
	// Uncomment the following lines to customize the chatbot
	chatbot.WelcomeMessage = "What's your question?"
	chatbot.ProcessFunc(chatbotProcess)

	// Use the PORT environment variable
	port := os.Getenv("PORT")
	// Default to 3000 if no PORT environment variable was defined
	if port == "" {
		port = "3000"
	}

	// Start the server
	fmt.Printf("Listening on port %s...\n", port)
	log.Fatalln(chatbot.Engage(":" + port))
}
