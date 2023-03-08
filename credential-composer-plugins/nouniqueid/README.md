# SPIRE external plugin for customizing Workload SVIDs

This is a `CredentialComposer` plugin removes the `x500UniqueIdentifier` from SVID subject that SPIRE adds by default, e.g.:

```
Subject: C = US, O = SPIRE, x500UniqueIdentifier = a55a66721362bfbb331e2fa21da15354
```

Using this plugin, the subject looks like:
```
Subject: C = US, O = SPIRE
```

## Use this plugin in SPIRE

Note: SPIRE 1.6 is required. 

### Generate plugin binary

```
make build
```

This generates a binary named `credential-composer-no-unique-id`

### Configure SPIRE

Add in `server.conf`:
```
plugins {
    CredentialComposer "nouniqueid" {
        plugin_cmd = "./credential-composer-no-unique-id"
        plugin_data {
        }
     }
}
```