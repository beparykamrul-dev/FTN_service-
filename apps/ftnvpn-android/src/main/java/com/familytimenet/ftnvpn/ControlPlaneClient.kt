package com.familytimenet.ftnvpn

/** Minimal boundary for the FTN control plane. Real transport is injected separately. */
interface ControlPlaneClient {
    suspend fun enroll(deviceId: String): EnrollmentResult
    suspend fun fetchProfile(deviceId: String): TunnelProfile
    suspend fun reportHealth(deviceId: String, state: VpnState, latencyMs: Long?, packetLossPct: Double?)
}

data class EnrollmentResult(val deviceId: String, val enrolled: Boolean)
data class TunnelProfile(val id: String, val transport: String, val endpoint: String, val allowedRoutes: List<String>)
