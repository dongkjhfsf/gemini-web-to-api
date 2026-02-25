package openai

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewOpenAIService),
	fx.Provide(NewOpenAIController),
	fx.Invoke(RegisterRoutes),
)

func RegisterRoutes(app *fiber.App, c *OpenAIController) {
	// OpenAI routes (prefixed with /openai)
	openaiGroup := app.Group("/openai")
	openaiV1 := openaiGroup.Group("/v1")
	c.Register(openaiV1)

	// Register at root for standard compatibility
	rootV1 := app.Group("/v1")
	c.Register(rootV1)
}
