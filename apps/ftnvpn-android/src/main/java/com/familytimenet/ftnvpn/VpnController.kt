package com.familytimenet.ftnvpn

/** Coordinates enrollment/profile lifecycle; platform VPN adapter is intentionally separate. */
class VpnController(private val controlPlane: ControlPlaneClient) {
    var state: VpnState = VpnState.DISCONNECTED
        private set

    suspend fun enroll(deviceId: String): EnrollmentResult = controlPlane.enroll(deviceId)

    suspend fun loadProfile(deviceId: String): TunnelProfile = controlPlane.fetchProfile(deviceId)

    suspend fun reportHealth(deviceId: String, latencyMs: Long?, packetLossPct: Double?) {
        controlPlane.reportHealth(deviceId, state, latencyMs, packetLossPct)
    }

    fun setState(next: VpnState) {
        state = next
    }
}
