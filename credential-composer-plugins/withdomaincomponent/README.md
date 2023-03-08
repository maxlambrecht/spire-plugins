# SPIRE external plugin for customizing Workload SVIDs

This is a `CredentialComposer` plugin that adds the DNS to the Subject as Domain Component (DC), e.g.:

```
Subject: C = US, O = SPIRE, CN = db-client, DC = db-client
```

The DNS is configured to Workload identity through the registration entry.

## Use this plugin in SPIRE

Note: SPIRE 1.6 is required. 

### Generate plugin binary

```
make build
```

This generates a binary named `credential-composer-add-dc`

### Configure SPIRE

Add in `server.conf`:
```
plugins {
    CredentialComposer "add_dc" {
        plugin_cmd = "./credential-composer-add-dc"
        plugin_data {
        }
     }
}
```