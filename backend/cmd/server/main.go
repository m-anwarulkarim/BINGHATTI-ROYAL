package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/yourusername/binghatti-backend/internal/config"
	"github.com/yourusername/binghatti-backend/internal/database"
	"github.com/yourusername/binghatti-backend/internal/handler"
	customMiddleware "github.com/yourusername/binghatti-backend/internal/middleware"
	"github.com/yourusername/binghatti-backend/internal/repository"
	"github.com/yourusername/binghatti-backend/internal/service"
)

func main() {
	log.Println("Starting Binghatti Luxury Real Estate Go Backend Server...")

	cfg := config.LoadConfig()

	// 1. Initialize Database Connection Pool (pgx)
	dbPool, err := database.NewPostgresPool(cfg.DatabaseURL)
	if err != nil {
		log.Printf("Warning: PostgreSQL connection failed (%v). Continuing in standalone mode...", err)
	} else {
		defer dbPool.Close()
	}

	// 2. Repositories
	userRepo := repository.NewUserRepository(dbPool)
	projectRepo := repository.NewProjectRepository(dbPool)
	unitRepo := repository.NewUnitRepository(dbPool)
	leadRepo := repository.NewLeadRepository(dbPool)

	// 3. Services
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	projectService := service.NewProjectService(projectRepo, unitRepo)
	notificationService := service.NewNotificationService()
	leadService := service.NewLeadService(leadRepo, notificationService)

	// 4. Handlers
	authHandler := handler.NewAuthHandler(authService)
	projectHandler := handler.NewProjectHandler(projectService)
	leadHandler := handler.NewLeadHandler(leadService)

	// 5. Router Setup
	r := chi.NewRouter()

	// Built-in Middlewares
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(customMiddleware.CORS(cfg.AllowedOrigins))

	// Health Check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","server":"Binghatti Go API","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`))
	})

	// API v1 Sub-router
	r.Route("/api/v1", func(r chi.Router) {
		// Public Auth
		r.Post("/auth/login", authHandler.Login)

		// Public Leads (VIP Form, Dashboard Listing & Status Update)
		r.Post("/leads", leadHandler.RegisterLead)
		r.Get("/leads", leadHandler.ListLeads)
		r.Put("/leads/{id}/status", leadHandler.UpdateLeadStatus)

		// Public Projects & Units Catalog
		r.Get("/projects", projectHandler.ListProjects)
		r.Get("/projects/{slug}", projectHandler.GetProjectBySlug)

		// Protected Routes Example (Dashboard CRM APIs)
		r.Group(func(r chi.Router) {
			r.Use(customMiddleware.JWTAuth(cfg.JWTSecret))
			r.Get("/dashboard/stats", func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"success":true,"message":"Welcome to protected dashboard API"}`))
			})
		})
	})

	// 6. Graceful Shutdown HTTP Server Setup
	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("Server listening on port :%s", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Wait for interrupt signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced shutdown: %v", err)
	}

	log.Println("Server stopped gracefully.")
}
