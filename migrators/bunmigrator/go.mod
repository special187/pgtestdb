module github.com/special187/pgtestdb/migrators/bunmigrator

go 1.18

replace github.com/special187/pgtestdb => ../../

require (
	github.com/special187/pgtestdb v0.0.14
	github.com/peterldowns/testy v0.0.1
	github.com/uptrace/bun v1.2.1
	github.com/uptrace/bun/dialect/pgdialect v1.2.1
	github.com/uptrace/bun/driver/pgdriver v1.2.1
)

require (
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df // indirect
	golang.org/x/sys v0.18.0 // indirect
	mellium.im/sasl v0.3.1 // indirect
)
