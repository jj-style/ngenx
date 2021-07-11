# ngenx
Auto-generated nginx config for self-hosted services under different sub-domains.

# Usage
- `ngenx -input spec.yaml -output nginx.conf`
- `curl server.com/my-spec.yaml | ngenx -input - -output nginx.conf`
- `ngenx` (read from `spec.yaml` in current directory and output to `stdout`)