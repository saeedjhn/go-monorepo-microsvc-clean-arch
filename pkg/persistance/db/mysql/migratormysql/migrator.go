package migratormysql

import (
	"fmt"
	"log"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/saeedjhn/go-monorepo-microsvc-clean-arch/pkg/persistance/db/pq"
)

const dialect = "mysql"

type Migrator struct {
	conn       pq.DB
	dialect    string
	migrations *migrate.FileMigrationSource
}

// TODO - set migration table name
// TODO - add limit to Up and Down method

func New(conn pq.DB, absolutePath string) Migrator {
	// Read migrations from a folder:
	migrations := &migrate.FileMigrationSource{
		Dir: absolutePath,
	}

	return Migrator{conn: conn, dialect: dialect, migrations: migrations}
}

func (m Migrator) Up() error {
	n, err := migrate.Exec(m.conn.Conn(), m.dialect, m.migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("can't apply migrations: %w", err)
	}

	log.Printf("Applied %d migrations!\n", n)

	return nil
}

func (m Migrator) Down() error {
	n, err := migrate.Exec(m.conn.Conn(), m.dialect, m.migrations, migrate.Down)
	if err != nil {
		return fmt.Errorf("can't rollback migrations: %w", err)
	}

	log.Printf("Rollbacked %d migrations!\n", n)

	return nil
}

func (m Migrator) Status() {
	// TODO - add status
	panic("IMPLEMENT ME")
}