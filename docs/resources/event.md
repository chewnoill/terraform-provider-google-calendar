---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "google-calendar_event Resource - google-calendar"
subcategory: ""
description: |-
  
---

# google-calendar_event (Resource)



## Example Usage

```terraform
// Create a google calendar event.
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

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `end` (String)
- `start` (String)
- `summary` (String)

### Optional

- `attendee` (Block Set) (see [below for nested schema](#nestedblock--attendee))
- `description` (String)
- `guests_can_invite_others` (Boolean)
- `guests_can_modify` (Boolean)
- `guests_can_see_other_guests` (Boolean)
- `location` (String)
- `reminder` (Block Set) (see [below for nested schema](#nestedblock--reminder))
- `send_notifications` (Boolean)
- `show_as_available` (Boolean)
- `visibility` (String)

### Read-Only

- `event_id` (String)
- `hangout_link` (String)
- `html_link` (String)
- `id` (String) The ID of this resource.

<a id="nestedblock--attendee"></a>
### Nested Schema for `attendee`

Required:

- `email` (String)

Optional:

- `optional` (Boolean)


<a id="nestedblock--reminder"></a>
### Nested Schema for `reminder`

Required:

- `before` (String)
- `method` (String)
