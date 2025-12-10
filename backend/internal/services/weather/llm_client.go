package services

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
	"log"

	"github.com/chup1x/weather-stack/internal/config"
)

type llmClient interface {
	Enabled() bool
	RecommendClothes(ctx context.Context, in llmClothesInput) (string, error)
}

type llmClothesInput struct {
	City        string
	Temperature float64
	Description string
	Humidity    float64
	WindSpeed   float64
	UserTemps   map[string]int
}

func NewLLMClient(cfg config.LLMConfig) llmClient {
	if !cfg.Enabled || cfg.URL == "" || cfg.APIKey == "" || cfg.Model == "" {
		return disabledLLM{}
	}

	timeout := time.Duration(cfg.TimeoutSec) * time.Second
	if timeout <= 0 {
		timeout = 50 * time.Second
	}

	return &openAILLM{
		url:         cfg.URL,
		apiKey:      cfg.APIKey,
		model:       cfg.Model,
		temperature: cfg.Temperature,
		maxTokens:   cfg.MaxTokens,
		client:      &http.Client{Timeout: timeout},
	}
}

type disabledLLM struct{}

func (disabledLLM) Enabled() bool { return false }

func (disabledLLM) RecommendClothes(ctx context.Context, _ llmClothesInput) (string, error) {
	return "", errors.New("llm disabled")
}

type openAILLM struct {
	url         string
	apiKey      string
	model       string
	temperature float64
	maxTokens   int
	client      *http.Client
}

func (c *openAILLM) Enabled() bool { return true }

func (c *openAILLM) RecommendClothes(ctx context.Context, in llmClothesInput) (string, error) {
	systemPrompt := "Ты — помощник по подбору одежды по погоде. Отвечай кратко и по делу на русском языке."
	
	userLines := []string{
		fmt.Sprint("Ты — полезный ассистент, который дает рекомендации по выбору одежды исходя из погодных условий."),
		fmt.Sprint("Погода:"),
		fmt.Sprintf("- Город: %s", in.City),
		fmt.Sprintf("- Температура: %.1f°C, описание: %s", in.Temperature, in.Description),
		fmt.Sprintf("- Влажность: %.0f%%, ветер: %.1f м/с", in.Humidity, in.WindSpeed),
		fmt.Sprint("Персональные температурные предпочтения пользователя:"),
		fmt.Sprintf("- Комфортная температура в футболке: %s°C", in.UserTemps["comf"]),
		fmt.Sprintf("- Комфортная температура в толстовке: %s°C", in.UserTemps["tol"]),
		fmt.Sprintf("- Комфортная температура в пуховике: %s°C", in.UserTemps["puh"]),
		fmt.Sprint("Учти эти персональные предпочтения при составлении рекомендаций."),
		fmt.Sprint("Проанализируй погодные условия и дай практичные рекомендации по одежде. Учитывай температуру, осадки, влажность и ветер."),
		fmt.Sprint("Ответь в следующем формате:"),
		fmt.Sprint("- Основная одежда: [рекомендация по верхней одежде]"),
		fmt.Sprint("- Обувь: [рекомендация по обуви]"),
		fmt.Sprint("- Аксессуары: [рекомендация по аксессуарам]"),
		fmt.Sprint("- Общий совет: [краткое итоговое замечание]"),
		fmt.Sprint("Будь кратким и практичным, отвечай на русском языке:"),
	}

	payload := map[string]any{
		"model":       c.model,
		"temperature": c.temperature,
		"max_tokens":  c.maxTokens,
		"messages": []map[string]string{
			{"role": "system", "content": systemPrompt},
			{"role": "user", "content": strings.Join(userLines, "\n")},
		},
	}

	raw, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url, bytes.NewReader(raw))
	if err != nil {
		return "", fmt.Errorf("build llm request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	res, err := c.client.Do(req)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("do llm request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		log.Printf("llm request failed: status %d", res.StatusCode)
		return "", fmt.Errorf("llm request failed: status %d", res.StatusCode)
	}

	body := struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		log.Println(err)
		return "", fmt.Errorf("decode llm response: %w", err)
	}

	if len(body.Choices) == 0 {
		log.Println("llm response empty")
		return "", errors.New("llm response empty")
	}

	return strings.TrimSpace(body.Choices[0].Message.Content), nil
}
