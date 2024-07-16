package dbmatemigrator

import (
	"context"
	"database/sql"
	"io/fs"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres" // driver

	"github.com/special187/pgtestdb"
	"github.com/special187/pgtestdb/migrators/common"
)

// Option provides a way to configure the DbmateMigrator struct and its behavior.
//
// dbmate documentation: https://github.com/amacneil/dbmate#command-line-options
//
// See:
//   - [WithDir]
//   - [WithTableName]
//   - [WithFS]
type Option func(*DbmateMigrator)

// WithDir specifies the location(s) of the migration files. If you have migrations
// in multiple directories, you should pass each path here instead of passing
// WithDir multiple times.
//
// Default: `"./db/migrations"`
//
// Equivalent to `--migrations-dir`
// https://github.com/amacneil/dbmate#command-line-options
func WithDir(dir ...string) Option {
	return func(m *DbmateMigrator) {
		m.MigrationsDir = dir
	}
}

// WithTableName specifies the name of the table in which dbmate will store its
// migration records.
//
// Default: `"schema_migrations"`
//
// Equivalent to `--migrations-table`
// https://github.com/amacneil/dbmate#command-line-options
func WithTableName(name string) Option {
	return func(m *DbmateMigrator) {
		m.MigrationsTableName = name
	}
}

// WithFS specifies a `fs.FS` from which to read the migration files.
//
// Default: `<nil>` (reads from the real filesystem)
func WithFS(x fs.FS) Option {
	return func(m *DbmateMigrator) {
		m.FS = x
	}
}

// New returns a [DbmateMigrator], which is a pgtestdb.Migrator that
// uses dbmate to perform migrations.
//
// You can configure the behavior of dbmate by passing Options:
//   - [WithDir] is the same as --migrations-dir
//   - [WithTableName] is the same as --migrations-table
//   - [WithFS] allows you to use an embedded filesystem.
func New(opts ...Option) *DbmateMigrator {
	defaults := dbmate.New(nil)
	m := &DbmateMigrator{
		MigrationsDir:       defaults.MigrationsDir,
		MigrationsTableName: defaults.MigrationsTableName,
		FS:                  defaults.FS,
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// DbmateMigrator is a pgtestdb.Migrator that uses dbmate to perform migrations.
//
// DbmateMigrator does not perform any Verify() or Prepare() logic.
type DbmateMigrator struct {
	MigrationsDir       []string
	MigrationsTableName string
	FS                  fs.FS
}

func (m *DbmateMigrator) Hash() (string, error) {
	hash := common.NewRecursiveHash(
		common.Field("MigrationsTableName", m.MigrationsTableName),
	)
	if err := hash.AddDirs(m.FS, "*.sql", m.MigrationsDir...); err != nil {
		return "", err
	}
	return hash.String(), nil
}

// Migrate runs dbmate.CreateAndMigrate() to migrate the template database.
func (m *DbmateMigrator) Migrate(
	_ context.Context,
	_ *sql.DB,
	templateConfig pgtestdb.Config,
) error {
	u, err := url.Parse(templateConfig.URL())
	if err != nil {
		return err
	}
	dbm := dbmate.New(u)
	dbm.MigrationsDir = m.MigrationsDir
	dbm.MigrationsTableName = m.MigrationsTableName
	dbm.FS = m.FS
	return dbm.CreateAndMigrate()
}

// Prepare is a no-op method.
func (*DbmateMigrator) Prepare(
	_ context.Context,
	_ *sql.DB,
	_ pgtestdb.Config,
) error {
	return nil
}

// Verify is a no-op method.
func (*DbmateMigrator) Verify(
	_ context.Context,
	_ *sql.DB,
	_ pgtestdb.Config,
) error {
	return nil
}
