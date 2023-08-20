```
    OpenAPI -> [generate tables + generate requests]
                        +
    Steampipe
```

# OAPIGen Plugin for Steampipe

Use OpenAPI definitions to generate Steampipe SQL tables.

- **[Get started →](https://hub.steampipe.io/plugins/lyda/oapigen)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/lyda/oapigen/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/lyda/steampipe-plugin-oapigen/issues)

## Quick start

### Install

Download and install the latest OAPIGen plugin:

```bash
steampipe plugin install oapigen
```

Configure your [config file](https://hub.steampipe.io/plugins/lyda/oapigen#configuration) to include directories with OAPIGen definition files.

```hcl
connection "oapigen" {
  plugin = "oapigen"

  # Documents refers to the OpenAPI definition(s) being used.
  documents = ["some/path/to/openapi-definition.yaml"]

  # Version of OpenAPI spec being used.
  Version = 3 

  # Prefix is the prefix for the defined paths.
  prefix = "http://127.0.0.1:1234"
}
```

Run steampipe:

```shell
steampipe query
```

TODO: include a petstore yaml file.

Query all the pets available from the petstore api:

```sql
select
  pet_id,
  pet_name,
  pet_type
from
  oapigen_petstore_pets;
```

```sh
+--------+-----------------+-----------+
| pet_id | pet_name        | pet_type  |
+--------+-----------------+-----------+
| 1      | Fido            | dog       |
| 2      | Fluffy          | cat       |
+--------+-----------------+-----------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)
- [Kin OpenAPI Library](https://github.com/getkin/kin-openapi)

Clone:

```sh
git clone https://github.com/lyda/steampipe-plugin-oapigen.git
cd steampipe-plugin-oapigen
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make
```

Configure the plugin:

```shell
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/oapigen.spc
```

Try it!

```shell
steampipe query
> .inspect oapigen
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/lyda/steampipe-plugin-oapigen/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [OAPIGen Plugin](https://github.com/lyda/steampipe-plugin-oapigen/labels/help%20wanted)
