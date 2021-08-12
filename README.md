# Vault Migrate

This is a simple CLI tool that can facilitate a scrappy Vault migration.

## Usage

1. Set `VAULT_ADDR` and `VAULT_TOKEN` in  your environment. You can run `export VAULT_TOKEN=$(cat ~/.vault-token)` after a successful `vault login`.
2. Run `vault-migrate export {kv-path} {secret-path}`, i.e. `vault-migrate export bridge /`. This can be piped to a file with `> export.json`.
3. Adjust your `VAULT_ADDR` and `VAULT_TOKEN` to point at the destination server.
4. Run `cat export.json | vault-migrate import {kv-path}`.

## Disclaimers

This implementation is far from comprehensive and remains untested. It has worked well enough for our use-case.
