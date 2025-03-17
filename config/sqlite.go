package config

import (
	"os"

	"github.com/thiago-kersting/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	// Usar caminho absoluto para o banco de dados
	dbPath := "./db/main.db"

	// Criar diretório db se não existir
	err := os.MkdirAll("./db", os.ModePerm)
	if err != nil {
		logger.Errorf("error creating database directory: %v", err)
		return nil, err
	}

	// Tentar abrir conexão com o banco
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
