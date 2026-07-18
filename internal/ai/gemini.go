package ai

import (
	"context"
	"fmt"

	"google.golang.org/genai"

	"github.com/testgithubanjali/ai-document-summarizer/internal/config"
)

func GenerateSummary(text string) (string, error) {

	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey: config.GetEnv("GEMINI_API_KEY"),
	})
	if err != nil {
		return "", err
	}

	prompt := fmt.Sprintf(`
You are an AI document summarizer.

Summarize the following document.

Requirements:
- Keep the summary concise.
- Use bullet points.
- Highlight important information.
- Maximum 200 words.

Document:

%s
`, text)

	resp, err := client.Models.GenerateContent(
		context.Background(),
		"gemini-2.5-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
