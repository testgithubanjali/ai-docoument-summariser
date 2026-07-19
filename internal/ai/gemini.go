package ai

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/testgithubanjali/ai-document-summarizer/internal/config"
	"google.golang.org/genai"
)

const model = "gemini-2.5-flash"

func GenerateSummary(text string) (string, error) {

	log.Println("========== Gemini Summary Generation ==========")

	text = strings.TrimSpace(text)

	if text == "" {
		log.Println("Document contains no readable text")
		return "", fmt.Errorf("document contains no readable text")
	}

	apiKey := config.GetEnv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Println("GEMINI_API_KEY is missing")
		return "", fmt.Errorf("GEMINI_API_KEY is not configured")
	}

	log.Println("Gemini API Key Loaded")

	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		log.Println("Failed to create Gemini client:", err)
		return "", err
	}

	log.Println("Gemini client created successfully")

	prompt := fmt.Sprintf(`
You are an expert AI document summarizer.

Create a professional summary of the following document.

Instructions:
- Keep the summary under 200 words.
- Use bullet points.
- Focus on the most important information.
- Ignore repeated content.
- Return only the summary.

Document:

%s
`, text)

	log.Println("Sending request to Gemini...")

	resp, err := client.Models.GenerateContent(
		context.Background(),
		model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Println("Gemini API Error:", err)
		return "", err
	}

	summary := strings.TrimSpace(resp.Text())

	log.Println("Gemini response received")
	log.Printf("Summary Length: %d characters\n", len(summary))

	if summary == "" {
		log.Println("Gemini returned an empty summary")
		return "", fmt.Errorf("Gemini returned an empty summary")
	}

	log.Println("========== Gemini Summary Generation Completed ==========")

	return summary, nil
}
