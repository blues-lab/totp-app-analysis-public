The app encrypts the TOTP backup using a symmetric key derived from a user
provided password. The traffic snippet below highlights that both the ciphertext
and the user password (`Password123!`) are sent to Yandex servers, which
gives Yandex the capability to decrypt the TOTP backup and read its contents,
including the TOTP secret, label, and issuer.

```text
====== BEGIN REQUEST ======
POST registrator.mobile.yandex.net/1/validation/password/?consumer=dev&
Content-Type: application/x-www-form-urlencoded; charset=UTF-8

password=Password123!
====== END REQUEST ======
====== BEGIN REQUEST ======
POST registrator.mobile.yandex.net/1/bundle/yakey_backup/upload/?
Content-Type: application/x-www-form-urlencoded; charset=UTF-8

number=+12029908424
backup=SOhRm917OEvA9L44YDLIKBBIXK8GoMp04Gm2YDNoSiXiP_nR-rsuUi608lt1ZiwUgnqQgKyQHi2CVs4V-wZKkpJCer6puvmGMy89jZIRXgHCH-6_WJPNVyZfWyRoXXXZPNeL_s9HspqATGZiijGaKyEths8FsjPN2WX64k2FaiNq2PbX6x2_rsv5YTvysQM0kGImpus3VQhA6FLlI5nvRTazoZ9j-LhOes69ysZgpyLp4J9AwY_okUsVbuOv1QdKklU
device_name=Google Pixel 3a
====== END REQUEST ======
```
