The traffic snippet below highlights that values of the issuer (`Research Lab`)
and label (e.g.,`example-email-1@example.com`) TOTP fields are sent to Authy
servers in plaintext.

```
====== BEGIN REQUEST ======
POST api.authy.com/json/users/597149415/authenticator_tokens/update
Content-Type: application/x-www-form-urlencoded

token_id=1652749577
account_type=authenticator
encrypted_seed=xqyRw2btXbd0hAi2YDsQ+4uMbtZoRC4cbvElfl8vKXWOUTubsU3RHM1ZXk6weSzD57YF4ODKXRLVgus7p5cMXg==
digits=6
issuer=Research Lab
name=Research Lab: example-email-1@example.com
salt=heXF25paXzbduhVaY3aZPd10CfYLUdA3
password_timestamp=1652749574
original_name=Research Lab:example-email-1@example.com
logo=Research Lab
====== END REQUEST ======
```
