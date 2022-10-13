The traffic snippet below highlights that all TOTP fields, including secret
(`JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAYQ`), label (e.g.,
`example-email-1@example.com`), and issuer (`Research Lab`), are sent to the
Latch servers in plaintext.

```
====== BEGIN REQUEST ======
POST latch.elevenpaths.com/control/1.8/totp?
Content-Type: application/x-www-form-urlencoded

period=30
accountName=example-email-1@example.com
name=Research Lab
digits=6
secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAYQ
algorithm=SHA1
====== END REQUEST ======

====== BEGIN REQUEST ======
POST latch.elevenpaths.com/control/1.8/totp?
Content-Type: application/x-www-form-urlencoded

period=30
accountName=example-email-2@example.com
name=Research Lab
digits=6
secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAZA
algorithm=SHA1
====== END REQUEST ======
```
