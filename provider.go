// provider.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
	"os"
	"runtime"
)

func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"credentials": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"google-calendar_event": googleCalendarEvent(),
		},
	}

	provider.ConfigureFunc = func(d *schema.ResourceData) (interface{}, error) {
		terraformVersion := provider.TerraformVersion
		if terraformVersion == "" {
			// Terraform 0.12 introduced this field to the protocol
			// We can therefore assume that if it's missing it's 0.10 or 0.11
			terraformVersion = "0.11+compatible"
		}
		return providerConfigure(d, provider, terraformVersion)
	}
	return provider
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// providerConfigure configures the provider. Normally this would use schema
// data from the provider, but the provider loads all its configuration from the
// environment, so we just tell the config to load.
func providerConfigure(d *schema.ResourceData, p *schema.Provider, terraformVersion string) (interface{}, error) {
	var opts []option.ClientOption

	// Add credential source
	if v := d.Get("credentials").(string); v != "" {
		log.Printf("[TRACE] using supplied credentials")
		b, err := os.ReadFile(v)
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}
		config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
		if err != nil {
			log.Fatalf("Unable to retrieve config: %v", err)
		}
		tok, err := tokenFromFile(d.Get("token").(string))
		if err != nil {
			log.Fatalf("Unable to retrieve token: %v", err)
		}
		client := config.Client(context.Background(), tok)
		opts = append(opts, option.WithHTTPClient(client))
	} else {
		log.Printf("[TRACE] attempting to use default credentials: %#v", d.Get("credentials"))
	}

	// Use a custom user-agent string. This helps google with analytics and it's
	// just a nice thing to do.
	userAgent := fmt.Sprintf("(%s %s) Terraform/%s",
		runtime.GOOS, runtime.GOARCH, terraformVersion)
	opts = append(opts, option.WithUserAgent(userAgent))

	log.Printf("[TRACE] client options: %v", opts)

	// Create the calendar service.
	ctx := context.Background()
	calendarSvc, err := calendar.NewService(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create calendar service: %w", err)
	}
	calendarSvc.UserAgent = userAgent

	return &Config{
		calendar: calendarSvc,
	}, nil
}
