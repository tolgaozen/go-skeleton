package factories

import (
	"github.com/tolgaOzen/go-skeleton/internal/storage"
	PQRepository "github.com/tolgaOzen/go-skeleton/internal/storage/postgres"
	"github.com/tolgaOzen/go-skeleton/pkg/database"
	PQDatabase "github.com/tolgaOzen/go-skeleton/pkg/database/postgres"
)

// DataReaderFactory creates and returns a DataReader based on the database engine type.
func DataReaderFactory(db database.Database) (repo storage.DataReader) {
	switch db.GetEngineType() {
	case "postgres":
		// If the database engine is Postgres, create a new DataReader using the Postgres implementation
		return PQRepository.NewDataReader(db.(*PQDatabase.Postgres))
	default:
		// For any other type, use the in-memory implementation as a default
		return PQRepository.NewDataReader(db.(*PQDatabase.Postgres))
	}
}

// DataWriterFactory creates and returns a DataWriter based on the database engine type.
func DataWriterFactory(db database.Database) (repo storage.DataWriter) {
	switch db.GetEngineType() {
	case "postgres":
		// If the database engine is Postgres, create a new DataWriter using the Postgres implementation
		return PQRepository.NewDataWriter(db.(*PQDatabase.Postgres))
	default:
		// For any other type, use the in-memory implementation as a default
		return PQRepository.NewDataWriter(db.(*PQDatabase.Postgres))
	}
}
