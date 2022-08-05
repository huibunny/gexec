# gexec

## sql

```sql

select
	column_name,
	data_type
from
	information_schema."columns" c
where
	table_name = 't_user'

```

## build

```bash

#  create docs.go at  docs/docs.go
#  create swagger.json at  docs/swagger.json
#  create swagger.yaml at  docs/swagger.yaml
$ swag init --parseDependency --parseInternal -g cmd/app/main.go
# $ swag init -g internal/controller/http/v1/router.go

# run build.sh and it will output gexec file.
$ ./build.sh

```

## run

```bash

$ ./goclean -consul localhost:8500 -name hello -listen :9090

```

## test

```bash

curl -X POST -d '{"tableinfo": [{"table": "t_app", "columns": ["name", "icon", "cover"], "values": [["note", "https://bn.com/note.png", "https://bn.com/cover.png"]]}, {"table": "t_apptype_rel", "columns": ["app_id", "type_id"], "values": [[1, 2], [2, 3]]}]}' 'http://localhost:8830/v1/exec/save'

```

## function

* remove http port from config file and use the one from cmd arguments.
* support consul(register/deregister/kv)

