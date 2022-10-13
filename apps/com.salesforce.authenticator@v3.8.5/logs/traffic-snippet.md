The app encrypts the TOTP backup using a symmetric key derived from a user
provided password. The traffic snippet below highlights that both the ciphertext
and the user password (`0000`) was sent to Salesforce servers, which gives
Salesforce the capability to decrypt the TOTP backup and read its contents,
including the TOTP secret, label, and issuer.

```
====== BEGIN REQUEST ======
POST authenticator-api.salesforce.com/services/verify/v1/authenticators/ae73c008-3089-42ec-961b-2ebb5e0777de/backup?
Content-Type: application/x-www-form-urlencoded

bundle_secret=4Mkzu8WUgt+N1Y59fGP9BiU4UAdztqW7XTPkRnb0SIc=
passcode_type=numeric_pin
passcode=0000
backup_blob={
  "encrypted_bundle": "kLJYom7Sh73G3as8K/a/xugikHIN1nDrBwW7XML4zH11jHnQdINhAWwNgnAtGgMjux5Og/27DoYVLwPvnFFltLwakb74R1c7qmUGuZgbhyif2v6qaYb71uQvs895krqQD8cBJVTdy18PyU3Dw/7f2670eAT0zywJEUjmbn2/vi+/ZzIDoxuMbsm8A6r/qhblI53gxQrrcz/cr5+5Xoi1cNv5tK0ef9kF2ON2Us4QM7F0PBN5+GFuaf1FHEaIXAiro85QFWzj5qdzLuzgPf2/peZIIGUSlTTnKrzLk8jeGlPRu2tDNF7UJz9AK50p4nevnQVjxv3VYvekags/QGKgXw==",
  "encrypted_keys": [
    "SuE0VfZZt9MwzsOx0CigNLNxnF6dj8eI8TwOZzBsxADm4dAKafbQ4erQGbIjgyD6"
  ],
  "hash_salt": "4+VpmumqWtDmwz1C/RWw/Q==",
  "hash_iterations": 10000,
  "initialization_vector": "Njadu35JsKiP9zxEb9go2g=="
}
====== END REQUEST ======
```
