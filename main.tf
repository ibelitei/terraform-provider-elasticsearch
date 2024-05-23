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

resource "elasticsearch_index_template" "example" {
  name     = "example"
  template = "{\"index_patterns\": [\"test-*\"], \"settings\": {\"number_of_shards\": 1}}"
}
