provider "aws" {
  region  = "ap-northeast-1"
  profile = "private"
}

resource "aws_s3_bucket" "deploy-gin-helloworld" {
  bucket = "deploy-gin-helloworld"
}

resource "aws_api_gateway_rest_api" "gin-helloworld" {
  name = "gin-helloworld"
  endpoint_configuration {
    types = ["REGIONAL"]
  }
}
