# FTNVPN Android

Native Android client foundation for Family Time Network.

## Product scope

- FTN-branded VPN client
- Aether-Core control-plane integration
- WireGuard / AmneziaWG profile support through approved tunnel adapters
- Device registration and revocation
- mTLS-backed control-plane communication
- Connection health and telemetry
- Automatic display/layout adaptation for phone, tablet and Android TV
- Central policy and configuration retrieval
- Local private-key protection; private keys are never stored in the central database

## UI states

Disconnected → Connecting → Connected → Degraded → Reconnecting → Error

The client reports health telemetry to the FTN control plane but does not expose or upload private keys.

## Branding

Use the canonical FTN Orbital brand assets from the repository branding manifest. Keep logo aspect ratio intact and use low-opacity branding only where it does not reduce readability.
