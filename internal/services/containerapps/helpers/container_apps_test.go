// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package helpers

import (
	"testing"
)

func TestValidateContainerAppRegistry(t *testing.T) {
	cases := []struct {
		Input Registry
		Valid bool
	}{
		{
			Input: Registry{
				Server:            "registry.example.com",
				UserName:          "user",
				PasswordSecretRef: "secretref",
			},
			Valid: true,
		},
		{
			Input: Registry{
				Server:   "registry.example.com",
				Identity: "identity",
			},
			Valid: true,
		},
		{
			Input: Registry{
				Server: "registry.example.com",
			},
			Valid: false,
		},
		{
			Input: Registry{
				Server:            "registry.example.com",
				UserName:          "user",
				PasswordSecretRef: "secretref",
				Identity:          "identity",
			},
			Valid: false,
		},
		{
			Input: Registry{
				Server:            "registry.example.com",
				PasswordSecretRef: "secretref",
			},
			Valid: false,
		},
		{
			Input: Registry{
				Server:   "registry.example.com",
				UserName: "user",
			},
			Valid: false,
		},
	}

	for _, tc := range cases {
		t.Logf("[DEBUG] Testing Value %s", tc.Input)
		err := ValidateContainerAppRegistry(tc.Input)
		valid := err == nil
		if tc.Valid != valid {
			t.Fatalf("Expected %t but got %t for %s", tc.Valid, valid, tc.Input)
		}
	}
}

// Azure may omit the default HTTP scheme on read, which must not cause a diff after refresh.
func TestContainerAppStartupProbeTransport(t *testing.T) {
	for _, transport := range []string{"HTTP", "HTTPS", "TCP"} {
		t.Run(transport, func(t *testing.T) {
			probe := ExpandContainerAppStartupProbe(ContainerAppStartupProbe{
				Transport: transport,
				Port:      8080,
			})
			if transport == "HTTP" {
				probe.HTTPGet.Scheme = nil
			}

			if got := FlattenContainerAppStartupProbe(probe)[0].Transport; got != transport {
				t.Errorf("expected transport %q, got %q", transport, got)
			}
		})
	}
}
