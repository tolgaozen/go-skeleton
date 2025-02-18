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

var _ = Describe("DataReader", func() {
	var db database.Database
	var dataWriter *DataWriter
	var dataReader *DataReader

	BeforeEach(func() {
		version := os.Getenv("POSTGRES_VERSION")

		if version == "" {
			version = "14"
		}

		db = instance.PostgresDB(version)
		dataWriter = NewDataWriter(db.(*PQDatabase.Postgres))
		dataReader = NewDataReader(db.(*PQDatabase.Postgres))
	})

	AfterEach(func() {
		err := db.Close()
		Expect(err).ShouldNot(HaveOccurred())
	})

	Context("Read Users", func() {
		It("success", func() {
			ctx := context.Background()

			err := dataWriter.Write(ctx, "user-1")
			Expect(err).ShouldNot(HaveOccurred())

			users, err := dataReader.ReadUsers(ctx, database.NewPagination(database.Size(1), database.Page(1)))
			Expect(err).ShouldNot(HaveOccurred())

			Expect(users[0].Id).Should(Equal(uint64(1)))
			Expect(users[0].Name).Should(Equal("user-1"))
		})
	})
})
