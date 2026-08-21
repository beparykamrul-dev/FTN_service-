package com.familytimenet.ftnvpn

/** UI/control-plane state; tunnel implementation is supplied by a reviewed adapter. */
enum class VpnState {
    DISCONNECTED,
    CONNECTING,
    CONNECTED,
    DEGRADED,
    RECONNECTING,
    ERROR
}
