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
- [ ] Open the app
- [ ] Create an account with a unique password
  - Do not enable Biometric ID / Touch Sensor
- [ ] Scan QR code #1
- [ ] Scan QR code #2
- [ ] Uninstall / reinstall app
- [ ] Sign in again your account same password
- [ ] Change the account password to `123456` and sign in again (to see whether
      the password gets leaked)
- [ ] Delete QR Code #1
- [ ] Delete QR Code #2
