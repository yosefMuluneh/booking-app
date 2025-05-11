## GitHub Copilot Chat

- Extension Version: 0.26.7 (prod)
- VS Code: vscode/1.99.3
- OS: Linux

## Network

User Settings:
```json
  "github.copilot.advanced.debug.useElectronFetcher": true,
  "github.copilot.advanced.debug.useNodeFetcher": false,
  "github.copilot.advanced.debug.useNodeFetchFetcher": true
```

Connecting to https://api.github.com:
- DNS ipv4 Lookup: 140.82.121.5 (250 ms)
- DNS ipv6 Lookup: Error (636 ms): getaddrinfo ENOTFOUND api.github.com
- Proxy URL: None (7 ms)
- Electron fetch (configured): HTTP 200 (364 ms)
- Node.js https: HTTP 200 (486 ms)
- Node.js fetch: HTTP 200 (559 ms)
- Helix fetch: HTTP 200 (976 ms)

Connecting to https://api.individual.githubcopilot.com/_ping:
- DNS ipv4 Lookup: 140.82.114.21 (81 ms)
- DNS ipv6 Lookup: Error (82 ms): getaddrinfo ENOTFOUND api.individual.githubcopilot.com
- Proxy URL: None (50 ms)
- Electron fetch (configured): HTTP 200 (655 ms)
- Node.js https: HTTP 200 (772 ms)
- Node.js fetch: HTTP 200 (874 ms)
- Helix fetch: HTTP 200 (860 ms)

## Documentation

In corporate networks: [Troubleshooting firewall settings for GitHub Copilot](https://docs.github.com/en/copilot/troubleshooting-github-copilot/troubleshooting-firewall-settings-for-github-copilot).