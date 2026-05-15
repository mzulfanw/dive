package runtime

// Placeholder for runtime detection and abstraction.
// In the future, this will support Docker, Colima, OrbStack, Podman, etc.

// DetectRuntime returns the current container runtime (docker, colima, etc).
func DetectRuntime() string {
	return "docker" // MVP: always docker
}
