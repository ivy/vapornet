package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

//go:embed template/prompt/concept.md
var conceptTemplate string

//go:embed template/prompt/html_renderer.md
var htmlRendererTemplate string

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

type RenderHTMLData struct {
	Concept     string
	Request     RequestInfo
	RequestBody string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("💖 Error loading .env file")
	}

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("💖 Missing OPENAI_API_KEY")
	}

	client := openai.NewClient(option.WithAPIKey(apiKey))

	conceptTmpl := template.Must(template.New("concept").Parse(conceptTemplate))
	htmlTmpl := template.Must(template.New("html_renderer").Parse(htmlRendererTemplate))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, _ := httputil.DumpRequest(r, true)
		log.Printf("💖 === Incoming request ===\n%s", dump)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("💖 Failed to read request body: %v", err)
			http.Error(w, "Could not read body", http.StatusInternalServerError)
			return
		}

		reqData := RequestData{
			Request: RequestInfo{
				Method:   r.Method,
				Path:     r.URL.Path,
				Host:     r.Host,
				Language: r.Header.Get("Accept-Language"),
			},
			RequestBody: string(body),
		}

		// Step 1: Generate concept prompt
		var conceptBuf bytes.Buffer
		if err := conceptTmpl.Execute(&conceptBuf, reqData); err != nil {
			log.Printf("💖 Failed to build concept prompt: %v", err)
			http.Error(w, "Failed to build concept prompt", http.StatusInternalServerError)
			return
		}
		log.Printf("💖 Concept Prompt:\n%s", conceptBuf.String())

		conceptResp, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(conceptBuf.String()),
			},
			Model: openai.ChatModelGPT4o,
		})
		if err != nil {
			log.Printf("💖 OpenAI error (concept): %v", err)
			http.Error(w, fmt.Sprintf("OpenAI error (concept): %v", err), http.StatusInternalServerError)
			return
		}

		concept := conceptResp.Choices[0].Message.Content
		log.Printf("💖 Concept Output:\n%s", concept)

		// Step 2: Generate HTML from concept
		htmlData := RenderHTMLData{
			Concept:     concept,
			Request:     reqData.Request,
			RequestBody: reqData.RequestBody,
		}

		var htmlPromptBuf bytes.Buffer
		if err := htmlTmpl.Execute(&htmlPromptBuf, htmlData); err != nil {
			log.Printf("💖 Failed to build HTML prompt: %v", err)
			http.Error(w, "Failed to build HTML prompt", http.StatusInternalServerError)
			return
		}
		log.Printf("💖 HTML Prompt:\n%s", htmlPromptBuf.String())

		htmlResp, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(htmlPromptBuf.String()),
			},
			Model: openai.ChatModelGPT4o,
		})
		if err != nil {
			log.Printf("💖 OpenAI error (HTML): %v", err)
			http.Error(w, fmt.Sprintf("OpenAI error (HTML): %v", err), http.StatusInternalServerError)
			return
		}

		finalHTML := htmlResp.Choices[0].Message.Content
		log.Printf("💖 Final HTML Output Length: %d", len(finalHTML))

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, finalHTML)
	})

	fmt.Println("💖 Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
