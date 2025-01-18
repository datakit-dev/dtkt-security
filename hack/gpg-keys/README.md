# GPG Key Management

This directory contains tasks for managing the `security@withdatakit.com` GPG key used for secure communications.

## Prerequisites

- `gpg` installed
- `sops` installed and configured
- `op` (1Password CLI) installed and authenticated
- Access to the DataKit 1Password vault

## Tasks

```shell
# Create/rotate the GPG key (deletes existing, generates new)
task gpg-keys:create

# List current GPG keys
task gpg-keys:list

# Load GPG key from stored files (for new machine setup)
task gpg-keys:load

# Delete all GPG keys for security@withdatakit.com
task gpg-keys:delete-all
```

## Key Rotation Workflow

### 1. Create New Key

```shell
task gpg-keys:create
```

This will:
- Delete any existing `security@withdatakit.com` keys
- Generate a new GPG key with a random passphrase
- Export the public key to `data/.well-known/datakit-security.pub.asc`
- Export the private key to `secrets/restricted/datakit-security.asc`
- Save the passphrase to `secrets/restricted/datakit-security-passphrase.txt`
- Encrypt private key and passphrase via SOPS
- Update 1Password with the new passphrase and key files

### 2. Publish Public Key

Create a branch, commit, and push the updated keys:

```shell
git checkout -b chore/rotate-security-gpg-key
git add ../../data/.well-known/datakit-security.pub.asc
git add secrets/restricted/*.enc.*
git commit -m "chore: rotate security GPG key"
git push -u origin chore/rotate-security-gpg-key
```

Then open a PR for review. The public key will be available at `https://withdatakit.com/.well-known/datakit-security.pub.asc` after merge.

## Setting Up a New Machine

To load the GPG key on a new machine:

```shell
task gpg-keys:load
```

This imports both the public key and the SOPS-decrypted private key into your local GPG keyring.

## Files

| File | Description |
|------|-------------|
| `gpg-key.conf` | GPG key generation configuration |
| `secrets/restricted/datakit-security.asc` | Private key (plaintext, gitignored) |
| `secrets/restricted/datakit-security.enc.asc` | Private key (SOPS encrypted, committed) |
| `secrets/restricted/datakit-security-passphrase.txt` | Passphrase (plaintext, gitignored) |
| `secrets/restricted/datakit-security-passphrase.enc.txt` | Passphrase (SOPS encrypted, committed) |
| `../../data/.well-known/datakit-security.pub.asc` | Public key (published) |
