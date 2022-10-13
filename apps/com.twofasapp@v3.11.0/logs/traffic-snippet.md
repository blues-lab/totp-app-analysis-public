The traffic snippet below highlights that the 2FAS app send all TOTP fields,
including secret (e.g., `JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAYQ`),
label (e.g., `example-email-1@example.com`), and issuer (`Research Lab`), that
have been previously registered in the app to Google Drive in plaintext when the
Google Drive backup mechanism is first enabled. After the mechanism is enabled,
the app encrypts all TOTP fields when backing up existing and new accounts. This
UX bug could be easily remedied.

```
====== BEGIN REQUEST ======
PUT www.googleapis.com/upload/drive/v3/files
Content-Type: text/plain

{
  "services": [
    {
      "name": "Research Lab",
      "secret": "JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAYQ",
      "updatedAt": 1653369154872,
      "type": "Unknown",
      "otp": {
        "label": "Research Lab:example-email-1@example.com",
        "account": "example-email-1@example.com",
        "issuer": "Research Lab",
        "digits": 6,
        "period": 30,
        "algorithm": "SHA1",
        "counter": 1,
        "tokenType": "TOTP"
      },
      "order": {
        "position": 0
      }
    }
  ],
  "updatedAt": 1653369154872,
  "schemaVersion": 2,
  "appVersionCode": 3110000,
  "appVersionName": "3.11.0",
  "appOrigin": "android",
  "groups": [],
  "account": "researchlabtotp@gmail.com"
}
====== END REQUEST ======
```
