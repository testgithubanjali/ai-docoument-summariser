package ai

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/genai"

	"github.com/testgithubanjali/ai-document-summarizer/internal/config"
)

const model = "gemini-2.5-flash"

func GenerateSummary(text string) (string, error) {

	text = strings.TrimSpace(text)

	if text == "" {
		return "", fmt.Errorf("document contains no readable text")
	}

	client, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey: config.GetEnv("GEMINI_API_KEY"),
	})
	if err != nil {
		return "", err
	}

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

	resp, err := client.Models.GenerateContent(
		context.Background(),
		model,
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(resp.Text()), nil
}
