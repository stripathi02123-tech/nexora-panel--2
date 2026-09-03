# Release Process

NEXORA installation and upgrade rely on GitHub Release artifacts. Use semantic version tags such as `v1.1.6`.

## Version Number

Check the version in:

- `backend/internal/version/version.go`
- `frontend/package.json`
- Release tag.

## Release Artifacts

The installer first tries to download the Linux AMD64 or ARM64 archive for the host architecture:

```text
nexora-linux-amd64.tar.gz
nexora-linux-arm64.tar.gz
```

In some cases, it may also try the standalone binary:

```text
nexora-linux-amd64
nexora-linux-arm64
```

## Installer Behavior

- `NEXORA_VERSION=latest`: use GitHub `releases/latest`.
- `NEXORA_VERSION=vX.Y.Z`: download artifacts from the specified release tag.

Example:

```bash
NEXORA_VERSION=v1.1.6 sh install.sh
```

## Post-release Verification

- The installer can download the new version.
- `systemctl status nexora` is healthy.
- `/api/version` returns the new version.
- The web panel can load frontend assets.
- Container list, task queue, and API Key pages open correctly.
