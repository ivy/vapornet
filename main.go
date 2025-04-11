package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

//go:embed template/prompt/concept.md
var conceptTemplate string

type RequestData struct {
	Request     RequestInfo
	RequestBody string
}

type RequestInfo struct {
	Method   string
	Path     string
	Host     string
	Language string
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Missing OPENAI_API_KEY")
	}

	client := openai.NewClient(option.WithAPIKey(apiKey))

	tmpl, err := template.New("concept").Parse(conceptTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, _ := httputil.DumpRequest(r, true)
		fmt.Printf("=== Incoming request ===\n%s\n", dump)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Could not read body", http.StatusInternalServerError)
			return
		}

		data := RequestData{
			Request: RequestInfo{
				Method:   r.Method,
				Path:     r.URL.Path,
				Host:     r.Host,
				Language: r.Header.Get("Accept-Language"),
			},
			RequestBody: string(body),
		}

		var promptBuf bytes.Buffer
		if err := tmpl.Execute(&promptBuf, data); err != nil {
			http.Error(w, "Failed to build prompt", http.StatusInternalServerError)
			return
		}

		log.Println("Prompt for OpenAI API:")
		fmt.Println(promptBuf.String())
		resp, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(promptBuf.String()),
			},
			Model: openai.ChatModelGPT4o,
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("OpenAI error: %v", err), http.StatusInternalServerError)
			return
		}

		concept := resp.Choices[0].Message.Content
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, concept)
	})

	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
