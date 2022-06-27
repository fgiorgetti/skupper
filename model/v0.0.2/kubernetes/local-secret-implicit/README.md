This scenario expects that a secret named: skupper-site-ca is created,
before skupper is initialized in the west site, providing the expected
certificates: tls.crt and tls.key.

It also expects that east site has a secret named east-to-west whose
certificate has been signed by the CA defined by skupper-site-ca in the west site.

Services are exposed by selector only.