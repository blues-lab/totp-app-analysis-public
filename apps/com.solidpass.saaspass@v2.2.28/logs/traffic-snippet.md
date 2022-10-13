The traffic snippet below highlights that all TOTP fields, including secret
(`JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAZA`), label (e.g.,
`example-email-1@example.com`), and issuer (`Research Lab`), are sent to the
SAASPASS servers in plaintext using the XMPP protocol.

```xmpp
====== BEGIN REQUEST to 104.154.49.147:0 (backing up) ======
<?xml version="1.0"?>
<message id="6" to="saaspass@saaspass.com" type="chat">
  <body>
  {
    "auth": {
      "clientId": "xyht1h9k0fnu4ajqqr3luk2g6c9nyug5",
      "mac": "ZQAAJyxwbskA1HzF/wR2UEikpaDMpjcI3CZENqOjGHvCg6tFhYHjeeS4jwP+p2544K/9EqCTHRokhebXENL6YCUfrDXy"
    },
    "data": {
      "accountType": "AUTHENTICATOR",
      "appName": "Research Lab",
      "key": "otpauth://totp/Research Lab:example-email-2@example.com?secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAZA&issuer=Research Lab&digits=6&algorithm=SHA1",
      "serviceUrl": "",
      "ssoEnabled": false,
      "username": "example-email-2@example.com"
    },
    "device": {
      "dateTime": 1650552954057,
      "knownRevision": 8,
      "language": "en",
      "spKnownVersion": "2.2.28",
      "spVersion": "2.2.28"
    },
    "requestId": 6,
    "service": "authenticatorCreator",
    "type": "request"
  }
  </body>
</message>
====== END REQUEST ======

====== BEGIN RESPONSE from 104.154.49.147:0 while recovering ======
<?xml version="1.0"?>
<message to="xyht1h9k0fnu4ajqqr3luk2g6c9nyug5@saaspass.com/xyht1h9k0fnu4ajqqr3luk2g6c9nyug5" id="3" type="chat" from="saaspass@saaspass.com/76c908d4">
  <body>
    {
      "type": "response",
      "service": "authenticatorCreator",
      "requestId": 3,
      "result": {
        "success": true,
        "message": null,
        "code": 0
      },
      "data": {
        "revision": 4,
        "revisions": [
          {
            "changeset": {
              "category": "email",
              "action": "create",
              "data": {
                "emailAddress": "example-email-1@example.com",
                "verified": false,
                "id": 6010138944750703000,
                "allowDelete": true
              }
            },
            "revisionId": 3
          },
          {
            "changeset": {
              "category": "authenticator",
              "action": "create",
              "data": {
                "username": "example-email-1@example.com",
                "displayName": null,
                "id": 9218846251831681000,
                "appName": "Research Lab",
                "key": "otpauth://totp/Research Lab:example-email-1@example.com?secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAYQ&amp;issuer=Research Lab&amp;digits=6&amp;algorithm=SHA1",
                "ssoEnabled": false,
                "serviceUrl": null,
                "password": null,
                "accountType": "AUTHENTICATOR",
                "computerName": null,
                "computerType": null,
                "comType": null,
                "computerClientId": null,
                "iconSetVersion": null,
                "iconSet": null,
                "storePasswordOnServer": true,
                "allowAutoLogin": true
              }
            },
            "revisionId": 4
          }
        ],
        "id": 9218846251831681000,
        "username": "example-email-1@example.com",
        "displayName": null,
        "appName": "Research Lab",
        "key": "otpauth://totp/Research Lab:example-email-1@example.com?secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAYQ&amp;issuer=Research Lab&amp;digits=6&amp;algorithm=SHA1",
        "ssoEnabled": false,
        "serviceUrl": null,
        "password": null,
        "accountType": "AUTHENTICATOR",
        "computerName": null,
        "computerType": null,
        "storePasswordOnServer": null,
        "isMerged": null
      }
    }
  </body>
</message>
====== END RESPONSE ======
```
