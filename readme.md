# Terraform Google Calendar Provider

This is a [Terraform][terraform] provider for managing meetings on Google
Calendar. It enables you to treat "calendars as code" the same way you already
treat infrastructure as code!

Based off the work of @sethvargo https://github.com/sethvargo/terraform-provider-googlecalendar



## Installation

1. Pull from the module registry:

    ```hcl
    terraform {
      required_providers {
        google-calendar = {
          source = "chewnoill/google-calendar"
          version = "1.0.0"
        }
      }
    }
    ```
1. Configure Provider
    ```hcl
    provider "google-calendar" {
        credentials = "./credentials.json"
        token = "./token.json"
    }
    ```
    two json configuration files are required for this provider to function
    * credentials: these are OAuth application credentials used to access googles APIs. Instructions on creating these credentials is here: https://developers.google.com/workspace/guides/create-credentials
    * token: Once a user has authorized your app to impersonate them, an access token is generated. 

    A script has been provided [scripts/auth.go](./scripts/auth.go) that will walk through the process of using the credentials.json to create a token.json

## Usage

1. Create a google calendar event resource:

    ```hcl
    resource "google-calendar_event" "demo" {
      // Common options
      summary     = "My Demo Terraform Event"
      description = "Long-form description of the event, such as why it's needed"
      location    = "Conference Room B"

      // Start and end times work best if specified as RFC3339.
      start = "2024-03-13T10:00:00-05:00"
      end   = "2024-03-13T11:00:00-05:00"

      // Each attendee is listed separately, and attendees can be marked as
      // optional.
      attendee {
        email = "will@example.com"
      }

      // By default, the user's calendar reminders are used. By setting any
      // reminders, you override all default calendar reminders. The Google API
      // expects calendar  reminder times to be in "minutes", but you can specify
      // them as a Go-style time.Duration value for simplicity here, like "30m" for
      // "30 minutes".
      reminder {
        method = "email"
        before = "2h"
      }

      reminder {
        method = "popup"
        before = "5m"
      }
    }
    ```

1. Run `terraform init` to pull in the provider:

    ```sh
    $ terraform init
    ```

1. Run `terraform plan` and `terraform apply` to create events:

    ```sh
    $ terraform plan

    $ terraform apply
    ```

## Examples

For more examples, please see the [examples][./examples] folder in this
repository.
