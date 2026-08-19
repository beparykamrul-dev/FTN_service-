# FTNVPN Client Platform

FTNVPN is the branded client layer for FTN transport services.

## Targets

- Android phone/tablet
- Android TV / TV-class devices
- Windows PC
- Linux PC
- macOS PC (adapter target)

## Common control plane

All clients use the same authenticated FTN control-plane contract for device registration, tunnel profiles, health, policy, telemetry and revocation. Private keys remain device-local and are never stored in the central database.

## Branding

The supplied FTN orbital emblem is the canonical product mark. The same visual identity is used for:

- launcher/app icon
- login/connect screen
- connection-state screen
- system tray/status surface
- Android TV home surface
- VPN connection widgets
- control-panel background/watermark

The background treatment uses a subtle FTN emblem watermark rather than placing the full logo behind every interactive element, preserving readability.

## Client states

```text
DISCONNECTED
CONNECTING
CONNECTED
DEGRADED
RECONNECTING
BLOCKED_BY_POLICY
REVOKED
```

## Transport selection

The control plane may advertise an approved transport profile. The client only enables profiles supported by the target platform and signed/approved by FTN policy.
