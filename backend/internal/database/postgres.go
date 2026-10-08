package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(dbURL string) (*pgxpool.Pool, error) {
	if err := ensureDatabaseExists(dbURL); err != nil {
		log.Printf("Auto database check warning: %v", err)
	}

	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	config.MaxConns = 25
	config.MinConns = 5
	config.MaxConnLifetime = 30 * time.Minute
	config.MaxConnIdleTime = 5 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx connection pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL via pgxpool")

	// Ensure lead_status enum has 'viewing_scheduled'
	_, _ = pool.Exec(ctx, "ALTER TYPE lead_status ADD VALUE IF NOT EXISTS 'viewing_scheduled';")

	if err := autoRunMigrationsAndSeeds(ctx, pool); err != nil {
		log.Printf("Auto migration check warning: %v", err)
	}

	return pool, nil
}

func ensureDatabaseExists(dbURL string) error {
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return err
	}

	targetDB := config.ConnConfig.Database
	if targetDB == "" || targetDB == "postgres" {
		return nil
	}

	config.ConnConfig.Database = "postgres"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.ConnectConfig(ctx, config.ConnConfig)
	if err != nil {
		return fmt.Errorf("failed to connect to root postgres DB: %w", err)
	}
	defer conn.Close(ctx)

	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1);"
	if err := conn.QueryRow(ctx, query, targetDB).Scan(&exists); err != nil {
		return fmt.Errorf("failed to query database existence: %w", err)
	}

	if !exists {
		log.Printf("Target database '%s' does not exist. Auto-creating database now...", targetDB)
		createSQL := fmt.Sprintf("CREATE DATABASE %s;", pgx.Identifier{targetDB}.Sanitize())
		if _, err := conn.Exec(ctx, createSQL); err != nil {
			return fmt.Errorf("failed to auto-create database '%s': %w", targetDB, err)
		}
		log.Printf("Database '%s' created successfully!", targetDB)
	}

	return nil
}

func autoRunMigrationsAndSeeds(ctx context.Context, pool *pgxpool.Pool) error {
	var tableExists bool
	checkQuery := "SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'projects');"
	if err := pool.QueryRow(ctx, checkQuery).Scan(&tableExists); err != nil {
		return err
	}

	if tableExists {
		return nil
	}

	log.Println("Initializing database schema & seeds automatically...")

	migrationPaths := []string{
		"../db/migrations/000001_init_schema.up.sql",
		"db/migrations/000001_init_schema.up.sql",
		"../../db/migrations/000001_init_schema.up.sql",
	}

	seedPaths := []string{
		"../db/seeds/000001_seed_data.sql",
		"db/seeds/000001_seed_data.sql",
		"../../db/seeds/000001_seed_data.sql",
	}

	executeFile := func(paths []string, label string) error {
		for _, p := range paths {
			absPath, _ := filepath.Abs(p)
			content, err := os.ReadFile(absPath)
			if err == nil {
				log.Printf("Auto-executing %s from %s...", label, p)
				if _, err := pool.Exec(ctx, string(content)); err != nil {
					if !strings.Contains(err.Error(), "already exists") {
						return fmt.Errorf("failed executing %s: %w", label, err)
					}
				}
				return nil
			}
		}
		return fmt.Errorf("could not locate %s file", label)
	}

	if err := executeFile(migrationPaths, "Schema Migration"); err != nil {
		log.Printf("Warning: %v", err)
	}

	if err := executeFile(seedPaths, "Seed Data"); err != nil {
		log.Printf("Warning: %v", err)
	}

	log.Println("Database schema & seed data initialized successfully!")
	return nil
}
