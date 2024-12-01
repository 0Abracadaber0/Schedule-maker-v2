package main

import (
	"context"
	"generate/internal/config"
	"generate/internal/router"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("Starting application",
		slog.String("env", cfg.Env),
	)

	app := fiber.New(fiber.Config{
		AppName: "Generate service",
	})
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("logger", log)
		return c.Next()
	})

	router.SetupRoutes(app, log)

	/*data := models.ScheduleGenerator{
		Subjects: map[string][]string{
			"Языки программирования":                               {"Сидорик"},
			"Разработка приложений в визуальных средах":            {"Гурский", "Ермоленков"},
			"Объектно-ориентированные технологии программирования": {"Ошуковская"},
			"Философия": {"Волнистый"},
			"Алгоритмы и структуры данных":   {"Воронич", "Куприянов"},
			"Разработка и анализ требований": {"Барышев", "Козлов"},
			"Физическая культура":            {"Ширяев"},
		},
		Groups: map[string]string{
			"10701122": "1",
			"10701222": "1",
			"10701322": "1",
			"10702122": "2",
			"10702222": "2",
		},
		Plans: map[string]map[string]*models.Course{
			"1": {
				"Разработка и анализ требований": &models.Course{Lectures: 3, Practices: 2, Labs: 1, Stream: "1"},
				"Алгоритмы и структуры данных":   &models.Course{Lectures: 3, Practices: 2, Labs: 1, Stream: "1"},
				"Языки программирования":         &models.Course{Lectures: 2, Practices: 2, Labs: 1, Stream: "1"},
				"Философия":           &models.Course{Lectures: 1, Practices: 1, Labs: 0, Stream: "1"},
				"Физическая культура": &models.Course{Lectures: 1, Practices: 0, Labs: 0, Stream: "1"},
				"Разработка приложений в визуальных средах": &models.Course{Lectures: 2, Practices: 1, Labs: 1, Stream: "1"},
			},
			"2": {
				"Объектно-ориентированные технологии программирования": &models.Course{Lectures: 3, Practices: 2, Labs: 1, Stream: "2"},
				"Алгоритмы и структуры данных":                         &models.Course{Lectures: 3, Practices: 2, Labs: 1, Stream: "2"},
				"Философия": &models.Course{Lectures: 1, Practices: 1, Labs: 0, Stream: "2"},
				"Разработка приложений в визуальных средах": &models.Course{Lectures: 2, Practices: 1, Labs: 1, Stream: "2"},
				"Разработка и анализ требований":            &models.Course{Lectures: 3, Practices: 2, Labs: 1, Stream: "2"},
				"Физическая культура":                       &models.Course{Lectures: 1, Practices: 0, Labs: 0, Stream: "2"},
			},
		},
		Classrooms: []models.Classroom{
			{"507", "Лекция", []string{"Разработка и анализ требований", "Алгоритмы и структуры данных", "Объектно-ориентированные технологии программирования", "Философия", "Разработка приложений в визуальных средах"}},
			{"306", "Лекция", []string{"Разработка и анализ требований", "Алгоритмы и структуры данных", "Объектно-ориентированные технологии программирования", "Разработка приложений в визуальных средах", "Философия"}},
			{"116", "Практика", []string{"Языки программирования", "Философия"}},
			{"105", "Лабораторная", []string{"Объектно-ориентированные технологии программирования", "Разработка и анализ требований", "Алгоритмы и структуры данных"}},
			{"313", "Лабораторная", []string{"Разработка приложений в визуальных средах", "Разработка и анализ требований"}},
			{"329", "Лабораторная", []string{"Разработка приложений в визуальных средах"}},
			{"Спортзал", "Лекция", []string{"Физическая культура"}},
		},
		LessonsPerWeek: 6,
	}

	debug := service.GenerateAllLessons(data)
	fmt.Println(debug)
	*/

	go func() {
		err := app.Listen(":8088")
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	close(quit)
	log.Info("Shutting down...")
	err := app.Shutdown()
	if err != nil {
		panic(err)
	}

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
