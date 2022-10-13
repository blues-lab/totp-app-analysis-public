![combined QR codes](https://user-images.githubusercontent.com/945571/155418867-b13d4f69-a598-4a5c-8abe-31801aece1f5.png)

- [QR Code 1](https://user-images.githubusercontent.com/945571/155416190-d10440cc-bf4b-4592-952b-ac7aba3b130f.png)
  - otpauth://totp/Blues%20Lab:blues-lab-1@example.com?secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAYQ&issuer=Blues%20Lab&algorithm=SHA1&digits=6&period=30
  - Secret  plaintext: `Hello! Plaintext secret number 1`
- [QR Code 2](https://user-images.githubusercontent.com/945571/155416198-e6fe260a-0305-48da-90e1-137faccdc20c.png)
  - otpauth://totp/Blues%20Lab:blues-lab-2@example.com?secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAZA&issuer=Blues%20Lab&algorithm=SHA1&digits=6&period=30
  - Secret  plaintext: `Hello! Plaintext secret number 2`

# Setup
- [ ] Ensure that you have followed the [instructions
      here](/capture-traffic/README.md) and the Android device and laptop are
      capturing network traffic
- [ ] Make sure the phone has:
  - [ ] Internet access (e.g. WiFi is connected)
- [ ] Install the APK of the app to be analyzed using adb: `$> adb
      install-multiple *.apk`

# Action checklist
- [ ] Optionally, start screen recording with scrcpy
```
scrcpy \
    --stay-awake \
    --bit-rate 4M \
    --max-fps 20 \
    --record android-screen.mp4
```

- [ ] Create a Zoho account in the browser ]here](https://www.zoho.com/signup.html)
- [ ] Open the app
- [ ] Tap 'Sign In'
- [ ] Log into the Zoho account
- [ ] Tap 'Review settings' in the popup
- [ ] Leave defaults, but note the settings
- [ ] Go to home screen
  - [ ] Tap 'Skip' in top right
    - this step is encouraging the user to enable MFA on their Zoho account, which
      is awesome, but not relevant to our current analysis.
  - [ ] Go to 'Authenticator' tab
- - [ ] QR code #1
  - [ ] Go to 'Authenticator' tab
  - [ ] Scan QR code #1
- [ ] Enable passphrase backups
  - [ ] Tap purple 'Backup Secrets' button
  - [ ] Tap 'Set up'
  - [ ] Enter password `123` 
  - [ ] Confirm password `123`
  - [ ] Screenshot (toast message saying it backed up the secrets)
- [ ] Scan QR code #2
- [ ] Uninstall/reinstall the app
- [ ] Open app
- [ ] Log into the Zoho account
- [ ] Tap 'x' to close popup about privacy
- [ ] Tap 'Skip' to close screen about MFA on Zoho account
- [ ] Go to 'Authenticator' tab
- [ ] Enter passphrase `123` and click eye to show the password
- [ ] Tap 'Restore'
- [ ] Delete both TOTP entries
    - [ ] Tap the account
    - [ ] Tap 'Delete'
    - [ ] Tap 'Yes'

