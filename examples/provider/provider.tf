
terraform {
  required_providers {
    google-calendar = {
      source = "chewnoill/google-calendar"
      version = "1.0.0"
    }
  }
}

provider "google-calendar" {
    credentials = "./credentials.json"
    token = "./token.json"
}