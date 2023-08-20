connection "oapigen" {
  plugin = "oapigen"

  # Documents refers to the OpenAPI definition(s) being used.
  documents = ["some/path/to/openapi-definition.yaml"]

  # Version of OpenAPI spec being used.
  Version = 3 

  # Prefix is the prefix for the defined paths.
  prefix = "http://127.0.0.1:1234"
}
