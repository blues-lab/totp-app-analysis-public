The app encrypts the TOTP backup using a symmetric key derived from a user
provided password. The traffic snippet below highlights that both the ciphertext
and the user password (`rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr`) are sent to
Zoho servers, which gives Zoho the capability to decrypt the TOTP backup and
read its contents, including the TOTP secret, label, and issuer. Also sent to
Zoho are the plaintext values of the issuer (`Research Lab`) and label
(e.g., `example-email-1@example.com`).

```
====== BEGIN REQUEST ======
POST accounts.zoho.com/api/v1/account/self/user/self/passphrase?device_token=...&nonce=...&
Content-Type: application/json; charset=UTF-8

{
  "passphrase": {
    "pass_phrase": "rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr",
    "private_key": "...",
    "public_key": "..."
  }
}
====== END REQUEST ======

====== BEGIN REQUEST ======
POST accounts.zoho.com/api/v1/account/self/user/self/tpsecret?device_token=...&nonce=...&
Content-Type: application/json; charset=UTF-8

{
  "tpsecret_all": [
    {
      "app_name": "Research Lab",
      "app_secret": "GLUEKWXJQCWKMIORCAB3HJLGVO4QKD7FCBQRKYKT2UFYWOAZO54DYQPLXGLNL5FVYOBWJOMB7XTVKZLGH2S2IRPDB6YUT6RXDYMNMX6M4ZIGLCMTGL2HDHWL2GWHZM66",
      "label": "example-email-1@example.com"
    },
    {
      "app_name": "Research Lab",
      "app_secret": "GLUEKWXJQCWKMIORCAB3HJLGVO4QKD7FCBQRKYKT2UFYWOAZO54DYQPLXGLNL5FVYOBWJOMB7XTVLU2INMHOYV7LSYJOWPFHE4LPEYOM4ZIGLCMTGL2HDHWL2GWHZM66",
      "label": "example-email-2@example.com"
    }
  ]
}

====== END REQUEST ======
```