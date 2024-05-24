terraform {
  required_providers {
    elasticsearch = {
      source  = "registry.terraform.io/softwaremind/elasticsearch"
      version = "1.0.0"
    }
  }
}

provider "elasticsearch" {
  url = "http://localhost:9200"
}

resource "elasticsearch_index" "example_index" {
  name    = "example"
  settings = jsonencode({
    settings = {
      number_of_shards = 1
      number_of_replicas = 1
    }
    mappings = {
      properties = {
        name = {
          type = "text"
        }
      }
    }
  })
}