package main

import "testing"

func TestParseSSPorts(t *testing.T) {
	output := `Netid State  Recv-Q Send-Q Local Address:Port Peer Address:PortProcess
tcp   LISTEN 0      4096         0.0.0.0:3000      0.0.0.0:*    users:(("node",pid=45238,fd=13))
udp   UNCONN 0      0                  *:5353            *:*    users:(("Spotify",pid=50348,fd=704))
`

	entries := parseSSPorts(output)
	if len(entries) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(entries))
	}
	if entries[0].Port != "3000" || entries[0].Protocol != "tcp" || entries[0].PID != "45238" {
		t.Fatalf("unexpected tcp entry: %+v", entries[0])
	}
	if entries[0].Address != "All Interfaces" || entries[0].Process != "node" {
		t.Fatalf("unexpected tcp process/address: %+v", entries[0])
	}
	if entries[1].Port != "5353" || entries[1].Protocol != "udp" || entries[1].PID != "50348" {
		t.Fatalf("unexpected udp entry: %+v", entries[1])
	}
}

func TestParseLSOFPorts(t *testing.T) {
	output := `COMMAND     PID  USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
node      45238 pavel   13u  IPv6 0x8da4aa211495fb3d      0t0  TCP *:3000 (LISTEN)
postgres  31485 pavel    8u  IPv6 0xf34f15ccfff19e46      0t0  TCP [::1]:5432 (LISTEN)
Spotify   50348 pavel  704u  IPv4 0xf0ddb8fc3bb7a9a6      0t0  UDP *:5353
Spotify   50375 pavel   32u  IPv6 0x596b91f249541c81      0t0  UDP [::1]:57007->[2600:1901:1:7c5::]:443
identitys   955 pavel    7u  IPv4  0x26cdb681dcdd8f6      0t0  UDP *:*
`

	entries := parseLSOFPorts(output)
	if len(entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(entries))
	}
	if entries[0].Port != "3000" || entries[0].Protocol != "tcp" || entries[0].Address != "All Interfaces" {
		t.Fatalf("unexpected tcp wildcard entry: %+v", entries[0])
	}
	if entries[1].Port != "5432" || entries[1].Address != "::1" {
		t.Fatalf("unexpected ipv6 entry: %+v", entries[1])
	}
	if entries[2].Port != "5353" || entries[2].Protocol != "udp" || entries[2].State != "UDP" {
		t.Fatalf("unexpected udp entry: %+v", entries[2])
	}
}

func TestDedupePorts(t *testing.T) {
	entries := []PortEntry{
		{Port: "3000", Protocol: "tcp", PID: "45238", Process: "node", State: "LISTEN", Address: "All Interfaces"},
		{Port: "3000", Protocol: "tcp", PID: "45238", Process: "node", State: "LISTEN", Address: "All Interfaces"},
		{Port: "5432", Protocol: "tcp", PID: "31485", Process: "postgres", State: "LISTEN", Address: "127.0.0.1"},
	}

	deduped := dedupePorts(entries)
	if len(deduped) != 2 {
		t.Fatalf("expected 2 deduped entries, got %d", len(deduped))
	}
	if deduped[0] != entries[0] || deduped[1] != entries[2] {
		t.Fatalf("unexpected deduped order: %+v", deduped)
	}
}
