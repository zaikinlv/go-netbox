package main

import (
	"os"
	"testing"

	"github.com/netbox-community/go-netbox/v4"
)

func HGetClient(t *testing.T) *netbox.APIClient {
	hostname := ""
	token := ""

	if os.Getenv("NETBOX_HOSTNAME") != "" {
		hostname = os.Getenv("NETBOX_HOSTNAME")
	}

	if os.Getenv("NETBOX_TOKEN") != "" {
		token = os.Getenv("NETBOX_TOKEN")
	}

	if hostname == "" || token == "" {
		t.Skip("NETBOX_HOSTNAME and NETBOX_TOKEN must be set")
	}

	client := netbox.NewAPIClientFor(hostname, token)
	return client
}
