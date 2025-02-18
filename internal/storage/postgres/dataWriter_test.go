package postgres

import (
	"context"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/tolgaOzen/go-skeleton/internal/storage/postgres/instance"
	"github.com/tolgaOzen/go-skeleton/pkg/database"
	PQDatabase "github.com/tolgaOzen/go-skeleton/pkg/database/postgres"
)

var _ = Describe("DataWriter", func() {
	var db database.Database
	var dataWriter *DataWriter

	BeforeEach(func() {
		version := os.Getenv("POSTGRES_VERSION")

		if version == "" {
			version = "14"
		}

		db = instance.PostgresDB(version)
		dataWriter = NewDataWriter(db.(*PQDatabase.Postgres))
	})

	AfterEach(func() {
		err := db.Close()
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("Write", func() {
		It("success", func() {
			ctx := context.Background()
			err := dataWriter.Write(ctx, "user-1")
			Expect(err).ShouldNot(HaveOccurred())
		})
	})
})
