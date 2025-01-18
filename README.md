<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://assets.datakit.cloud/identity/logo-color-on-blue.svg">
    <img alt="DataKit Logo" src="https://assets.datakit.cloud/identity/logo-color.svg" height="30" />
  </picture>
</p>
<p align="center">Open source & cloud-agnostic data activation tools.</p>

---

# DataKit Security

This repository is a catch all repository for security related documentation, disclosures, advisories, and other security related content not covered by our other existing repositories. If you have an issue specific to one of our other repositories, please report it there.

You can find more information about DataKit and our security practices at [https://withdatakit.com/security](https://withdatakit.com/security).

## Reporting a Vulnerability

We take security issues seriously. If you discover a vulnerability in DataKit, please follow these steps:

### 1. Responsible Disclosure

- Do not publicly disclose security vulnerabilities.
- Report the issue privately to our security team preferably using [Github Advisories](https://github.com/datakit-dev/dtkt-security/security) or by secure email.

  📧 Email: <security@withdatakit.com>

  🔒 GPG Key: [datakit-security.pub.asc](https://withdatakit.com/.well-known/datakit-security.pub.asc)

For more details visit https://withdatakit.com/security/researchers#responsible-disclosure

### 2. What to Include in Your Report

When reporting a vulnerability, please provide:

- A detailed description of the issue.
- Steps to reproduce (if possible).
- The affected version(s).
- Any potential mitigation or patch ideas.

We will acknowledge your report within 48 hours and keep you updated on the progress.

## Security Best Practices

To securely use DataKit, follow these best practices:

### 1. Use Official Channels

Only download DataKit artifacts from official sources like [dtkt.dev](https://dtkt.dev) or from our GitHub Releases.

### 2. Verify Signed Artifacts

DataKit artifacts are signed using Cosign. Before using a downloaded artifact, verify its integrity:

```sh
cosign verify-blob \
  --key https://dtkt.dev/.well-known/cosign.pub \
  --signature artifact.sig \
  artifact
```

### 3. Check for Vulnerabilities

Regularly scan artifacts for vulnerabilities.

For example:

```sh
grype artifact
# or (more efficiently with our pre-generated sbom)
grype sbom:artifact.sbom.spdx.json
```

### 4. Keep Your Versions Updated

Since we currently only support the **latest stable release**, always update artifacts to the latest version.

## Security Tools Used

We use the following tools to ensure the security of our OSS code repositories, artifacts, and deployments. Some are language dependent, while others are general security tools.

### Secrets & Credential Scanning

- **[GitHub Secret Scanning](https://docs.github.com/en/code-security/secret-scanning)** – Detects and prevents secrets from being committed to repositories.

### Code Quality & Linting

- **[Go Report Card](https://github.com/gojp/goreportcard)** – Generates a report on the quality of a Go project.
- **[golangci-lint](https://github.com/golangci/golangci-lint)** – A fast linters runner for Go, catching security and quality issues early.

### Static & Dynamic Code Analysis

- **[GitHub CodeQL](https://github.com/github/codeql)** – Analyzes source code for security vulnerabilities.
- **[Go Fuzzing](https://go.dev/doc/security/fuzz)** – Helps find security vulnerabilities and bugs by fuzz testing Go code.

### Supply Chain Security & Dependency Management

- **[Dependabot](https://github.com/dependabot)** – Automatically detects and updates vulnerable dependencies in our project.

### SBOM & Vulnerability Scanning

- **[syft](https://github.com/anchore/syft)** – Generates a Software Bill of Materials (SBOM) for tracking dependencies.
- **[grype](https://github.com/anchore/grype)** – Scans images, projects, and SBOM data for known vulnerabilities (CVE detection).

### Code Signing & Integrity

- **[cosign](https://github.com/sigstore/cosign)** – Ensures binary authenticity and integrity by signing artifacts.

### Best Practices

- **[OpenSSF Best Practices](https://www.bestpractices.dev/en)** – Provides a set of best practices for secure software development.
- **[OpenSSF Scorecard](https://github.com/ossf/scorecard)** – Evaluates the security posture of our open source project based on best practices.

### Infrastructure & Kubernetes Security

- **[kubescape](https://github.com/kubescape/kubescape)** – Scans Kubernetes manifests and clusters for misconfigurations and vulnerabilities.
- **[kube-score](https://github.com/zegl/kube-score)** – Analyzes Kubernetes object definitions to ensure best practices.
- **[kubesec](https://github.com/controlplaneio/kubesec)** – Scans Kubernetes manifests and performs security risk analysis.

<!-- ### Runtime Security & Monitoring

- **[Falco](https://github.com/falcosecurity/falco)** – Monitors runtime behavior in a containerized/cloud environment. -->

This security stack helps ensure that our OSS code repositories, artifacts, and deployments remains safe, reliable, and compliant with industry best practices.

## Disclosure Timeline

| Timeframe | Action                                |
| --------- | ------------------------------------- |
| 0-2 days  | Acknowledge vulnerability report      |
| 3-7 days  | Investigate and confirm vulnerability |
| 7-14 days | Develop a patch or mitigation         |
| 14+ days  | Release a fix and notify users        |

Critical vulnerabilities may receive expedited patches and releases.

## Hall of Fame (Responsible Disclosures)

We appreciate the efforts of security researchers who help us improve DataKit. If you’d like to be publicly credited for reporting a vulnerability on our [Hall of Fame](https://withdatakit.com/security/researchers/hall-of-fame), let us know! 🎖️

---

**Developed with ❤️ by DataKit**
