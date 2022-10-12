The Aegis Authenticator app is open source and publishes a python script that
can decrypt TOTP backups created by the app.

The `decrypt.py` script is copy/pasted from the Aegis repo here:
https://github.com/beemdevelopment/Aegis/blob/v2.0.3/docs/decrypt.py.

To decrypt the example TOTP backup files, you can run:

```bash
./verify_crypto.sh
```

You can also update the `verify_crypto.sh` script to point to your own backup
file(s).