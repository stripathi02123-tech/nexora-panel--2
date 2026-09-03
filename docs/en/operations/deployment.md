# Deployment

NEXORA can run directly on the host or behind a reverse proxy. In production, set up access control before exposing it to administrators.

## Service Exposure

The default web port is `8999`:

```text
http://YOUR_SERVER_IP:8999
```

Recommendations:

- Allow only fixed administrator IPs.
- Use a reverse proxy with HTTPS.
- Do not expose the real login URL in public docs or screenshots.

## systemd

Common commands:

```bash
systemctl status nexora
systemctl restart nexora
systemctl enable nexora
journalctl -u nexora -f
```

## Firewall

At minimum, confirm:

- The panel port is open only to trusted sources.
- NAT mapped ports are opened only as needed.
- The SSH management port does not conflict with container mappings.
- IPv6 firewall rules are planned together with IPv4 rules.

## Backups

Back up regularly:

- NEXORA configuration directory.
- SQLite database.
- Container configuration.
- Snapshots or external data backups for important containers.
