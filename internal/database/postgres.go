package database

import (
	"task-manager-app/backend/internal/domain"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase initializes the PostgreSQL database connection and returns a *gorm.DB instance.
func NewDatabase(connectionString string) (*gorm.DB, error) {
	// Open database connection
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(25)                 // Limit maximum simultaneous connections
	sqlDB.SetMaxIdleConns(5)                  // Keep some connections ready
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Refresh connections periodically

	// Verify connection is working
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	// Perform migrations
	if err := db.AutoMigrate(&domain.User{}, &domain.Task{}); err != nil {
		return nil, err
	}

	// Insert initial data
	if err := insertInitialData(db); err != nil {
		return nil, err
	}

	return db, nil
}

func insertInitialData(db *gorm.DB) error {
	// Check if initial data already exists
	var count int64
	db.Model(&domain.User{}).Count(&count)
	if count == 0 {
		// Insert initial user
		user := domain.User{
			Email:        "admin@example.com",
			PasswordHash: "admin", // Note: In a real application, make sure to hash the password
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		if err := user.HashPassword("admin"); err != nil {
			return err
		}
		if err := db.Create(&user).Error; err != nil {
			return err
		}

		// Insert initial tasks
		tasks := []domain.Task{
			{Title: "Investigar a origem do Véu de Névoa", Description: "Estudar os fenômenos misteriosos que envolvem a névoa na região de Fog Hill e seus impactos sobre os cinco elementos.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Dominar o Elemento do Fogo", Description: "Treinar e aprimorar as habilidades de manipulação do Fogo, buscando controlar esse poder com mais precisão e força.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Recolher fragmentos de Cristais Elementais", Description: "Viajar pelas colinas de Fog Hill e coletar fragmentos de cristais que podem conter a essência dos elementos.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Investigação das Ruínas dos Antigos", Description: "Explorar ruínas antigas, em busca de artefatos que possam revelar segredos sobre o antigo império que dominava a região de Fog Hill.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Estudar as criaturas das neblinas", Description: "Observar e estudar criaturas nativas que surgem das névoas de Fog Hill, descobrindo seu papel no equilíbrio dos elementos.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Treinamento de combate com os Elementais", Description: "Desafiar e treinar contra os poderosos elementais das colinas, aprimorando suas táticas de combate contra essas entidades.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Explorar o Vale da Nebulosa", Description: "Viajar até o misterioso Vale da Nebulosa e procurar por qualquer sinal de distúrbios nos elementos que possam afetar a região.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Estudar o Livro das Cinzas", Description: "Pesquisar um livro antigo que detalha os rituais e segredos dos mestres elementais que viveram há séculos.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Conquistar o domínio do Ar", Description: "Treinar para controlar o Elemento do Ar, aprendendo a manipular ventos e tempestades em batalhas e estratégias.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Ajudar os habitantes de Fog Hill", Description: "Prestar auxílio a aldeões e caçadores de névoa que estão sendo afetados pelos desequilíbrios nos elementos.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Derrotar um monstro elemental corrompido", Description: "Enfrentar e derrotar um monstro que foi corrompido pelos elementos, protegendo as aldeias e mantendo o equilíbrio.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Pesquisar os portais para outras dimensões", Description: "Investigar os portais elementais que conectam Fog Hill a outras dimensões, e buscar pistas para o resgate de Elys.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Conquistar o domínio da Terra", Description: "Aprender a manipular o Elemento da Terra, criando barreiras e manipulando o solo e as rochas para a defesa e ataque.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Resgatar um aliado perdido na névoa", Description: "Rastrear um aliado que desapareceu nas névoas e resgatá-lo de uma prisão dimensional, possivelmente envolvendo forças arcanas.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Atravessar o Labirinto das Correntes", Description: "Vencer o Labirinto das Correntes, um local místico onde o tempo e o espaço se distorcem, tentando recuperar um artefato importante.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Entender o Código das Chamas Eternas", Description: "Decifrar um antigo código que pode revelar um poder ancestral relacionado à manipulação do Fogo eterno, guardado pelos mestres do Fogo.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Descobrir a origem do Monólito do Vento", Description: "Explorar as montanhas de Fog Hill em busca do Monólito do Vento, uma formação que se diz ser a chave para controlar os ventos mais fortes.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Treinar com um mestre elemental", Description: "Envolver-se em um intenso treinamento com um mestre elemental para aprimorar as habilidades de manipulação dos cinco elementos.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Ajudar a restaurar o equilíbrio dos elementos", Description: "Trabalhar com outros guerreiros e sábios para restaurar o equilíbrio dos elementos que foi perdido devido aos distúrbios recentes.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
			{Title: "Investigação da Névoa Mortal", Description: "Investigar a causa do surgimento da Névoa Mortal, uma névoa corrompida que está afetando a vida e os elementos em Fog Hill.", IsCompleted: false, UserID: user.ID, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		}

		for _, task := range tasks {
			if err := db.Create(&task).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
