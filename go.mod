module tpro/stage1g

go 1.13

replace (
	tpro/stage1g/handlers/staff => ./lib/handlers/staff
	tpro/stage1g/utils => ./lib/utils
	tpro/stage1g/gql => ./gql
)

require (
	github.com/jackc/pgx/v4 v4.5.0 // indirect
	github.com/jessevdk/go-flags v1.4.0
	tpro/stage1g/handlers/staff v0.0.0-00010101000000-000000000000
	tpro/stage1g/utils v0.0.0-00010101000000-000000000000
	tpro/stage1g/gql v0.0.0-00010101000000-000000000000
)
