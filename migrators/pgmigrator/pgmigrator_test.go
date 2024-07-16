package pgmigrator_test

import (
	"context"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v5/stdlib" // "pgx" driver
	"github.com/peterldowns/pgmigrate"
	"github.com/peterldowns/testy/assert"
	"github.com/peterldowns/testy/check"

	"github.com/special187/pgtestdb"
	"github.com/special187/pgtestdb/migrators/pgmigrator"
	"github.com/special187/pgtestdb/migrators/pgmigrator/migrations"
)

func TestPGMigratorFromDisk(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	dir := os.DirFS("migrations")
	pgm, err := pgmigrator.New(dir)
	assert.Nil(t, err)
	db := pgtestdb.New(t, pgtestdb.Config{
		DriverName: "pgx",
		Host:       "localhost",
		User:       "postgres",
		Password:   "password",
		Port:       "5433",
		Options:    "sslmode=disable",
	}, pgm)
	assert.NotEqual(t, nil, db)

	assert.NoFailures(t, func() {
		var lastAppliedMigration string
		err := db.QueryRowContext(ctx, "select id from public.pgmigrate_migrations order by applied_at desc limit 1").Scan(&lastAppliedMigration)
		assert.Nil(t, err)
		check.Equal(t, "0002_cats", lastAppliedMigration)
	})

	var numUsers int
	err = db.QueryRowContext(ctx, "select count(*) from users").Scan(&numUsers)
	assert.Nil(t, err)
	check.Equal(t, 0, numUsers)

	var numCats int
	err = db.QueryRowContext(ctx, "select count(*) from cats").Scan(&numCats)
	assert.Nil(t, err)
	check.Equal(t, 0, numCats)

	var numBlogPosts int
	err = db.QueryRowContext(ctx, "select count(*) from blog_posts").Scan(&numBlogPosts)
	assert.Nil(t, err)
	check.Equal(t, 0, numBlogPosts)
}

func TestPGMigratorFromFS(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	pgm, err := pgmigrator.New(migrations.FS)
	assert.Nil(t, err)
	db := pgtestdb.New(t, pgtestdb.Config{
		DriverName: "pgx",
		Host:       "localhost",
		User:       "postgres",
		Password:   "password",
		Port:       "5433",
		Options:    "sslmode=disable",
	}, pgm)
	assert.NotEqual(t, nil, db)

	assert.NoFailures(t, func() {
		var lastAppliedMigration string
		err := db.QueryRowContext(ctx, "select id from public.pgmigrate_migrations order by applied_at desc limit 1").Scan(&lastAppliedMigration)
		assert.Nil(t, err)
		check.Equal(t, "0002_cats", lastAppliedMigration)
	})

	var numUsers int
	err = db.QueryRowContext(ctx, "select count(*) from users").Scan(&numUsers)
	assert.Nil(t, err)
	check.Equal(t, 0, numUsers)

	var numCats int
	err = db.QueryRowContext(ctx, "select count(*) from cats").Scan(&numCats)
	assert.Nil(t, err)
	check.Equal(t, 0, numCats)

	var numBlogPosts int
	err = db.QueryRowContext(ctx, "select count(*) from blog_posts").Scan(&numBlogPosts)
	assert.Nil(t, err)
	check.Equal(t, 0, numBlogPosts)
}

func TestPGMigratorWithTableName(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	pgm, err := pgmigrator.New(
		migrations.FS,
		pgmigrator.WithTableName("example_table_name"),
	)
	assert.Nil(t, err)
	db := pgtestdb.New(t, pgtestdb.Config{
		DriverName: "pgx",
		Host:       "localhost",
		User:       "postgres",
		Password:   "password",
		Port:       "5433",
		Options:    "sslmode=disable",
	}, pgm)
	assert.NotEqual(t, nil, db)

	assert.NoFailures(t, func() {
		var lastAppliedMigration string
		err := db.QueryRowContext(ctx, "select id from public.example_table_name order by applied_at desc limit 1").Scan(&lastAppliedMigration)
		assert.Nil(t, err)
		check.Equal(t, "0002_cats", lastAppliedMigration)
	})

	var numUsers int
	err = db.QueryRowContext(ctx, "select count(*) from users").Scan(&numUsers)
	assert.Nil(t, err)
	check.Equal(t, 0, numUsers)

	var numCats int
	err = db.QueryRowContext(ctx, "select count(*) from cats").Scan(&numCats)
	assert.Nil(t, err)
	check.Equal(t, 0, numCats)

	var numBlogPosts int
	err = db.QueryRowContext(ctx, "select count(*) from blog_posts").Scan(&numBlogPosts)
	assert.Nil(t, err)
	check.Equal(t, 0, numBlogPosts)
}

func TestPGMigratorFromFSAndWithOptions(t *testing.T) {
	t.Parallel()
	logger := pgmigrate.NewTestLogger(t)
	pgm, err := pgmigrator.New(
		migrations.FS,
		pgmigrator.WithTableName("example_table_name"),
		pgmigrator.WithLogger(logger),
	)
	assert.Nil(t, err)
	db := pgtestdb.New(t, pgtestdb.Config{
		DriverName: "pgx",
		Host:       "localhost",
		User:       "postgres",
		Password:   "password",
		Port:       "5433",
		Options:    "sslmode=disable",
	}, pgm)
	assert.NotEqual(t, nil, db)
}
