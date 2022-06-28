This sample demonstrates two sites, east and west running on systemd sites.
The west site uses interior mode and has a pre-defined CA.
Currently, credentials are to be stored using base64 directly into the yaml.

We still need a generic way to represent a reference that fits both a secret
name or a file / directory that stores all certificates.

Service controller (as well as service sync) are ignored at the moment, considering
that non-k8s sites will simply have a skupper-router instance running.

With that, changes to the yaml will need to be reprocessed by skupper (binary).
