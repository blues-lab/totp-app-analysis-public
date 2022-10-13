The traffic snippet below highlights that the Duo Mobile app includes the TOTP
label (e.g., `example-email-1@example.com`) in plaintext in the TOTP backup and
sometimes includes the TOTP issuer (e.g., `GitHub`) in plaintext as well. If the
TOTP issuer is one of the values on a hard coded list, then the app includes the
TOTP issuer in the backup; otherwise, it marks the TOTP issuer as "custom" and
**does not** backup the TOTP issuer.

```
====== BEGIN REQUEST ======
PUT www.googleapis.com/upload/drive/v3/files/1w81f4K6u40lDVB9H1WjmU_a3K7qQrVByivXmA53djAHgXwXO0g?
Content-Type: application/json

[
  {
    "version": 1.0,
    "accountType": "OtpAccount",
    "name": "example-email-1@example.com",
    "logoUri": "file:///data/user/0/com.duosecurity.duomobile/files/duokit/logos/f5bfb98d-12f3-4602-b0a0-42e537deae98.png",
    "pkey": "f5bfb98d-12f3-4602-b0a0-42e537deae98",
    "serviceTypeLabelIsCustom": true
  },
  {
    "version": 1.0,
    "accountType": "OtpAccount",
    "logoUri": "android.resource://com.duosecurity.duomobile/drawable/ic_github",
    "pkey": "42f1e517-f30e-46c8-bbfe-2c6d28cbc67a",
    "serviceTypeLabel": "GitHub",
    "serviceTypeLabelIsCustom": false
  }
]
====== END REQUEST ======
```
