# QR Codes
![combined QR codes](https://user-images.githubusercontent.com/945571/155418867-b13d4f69-a598-4a5c-8abe-31801aece1f5.png)

- [QR Code 1](https://user-images.githubusercontent.com/945571/155416190-d10440cc-bf4b-4592-952b-ac7aba3b130f.png)
  - otpauth://totp/Blues%20Lab:blues-lab-1@example.com?secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAYQ&issuer=Blues%20Lab&algorithm=SHA1&digits=6&period=30
  - Secret  plaintext: `Hello! Plaintext secret number 1`
- [QR Code 2](https://user-images.githubusercontent.com/945571/155416198-e6fe260a-0305-48da-90e1-137faccdc20c.png)
  - otpauth://totp/Blues%20Lab:blues-lab-2@example.com?secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAZA&issuer=Blues%20Lab&algorithm=SHA1&digits=6&period=30
  - Secret  plaintext: `Hello! Plaintext secret number 2`

# Setup

- [ ] Enable recording of plaintext network traffic generated by the device
  - Various techniques and tools can be used to record network traffic.
    Instructions for how to use open source tools to record betwork traffic on
    Pixel phones can be found [here](/capture-traffic/README.md).
- [ ] Make sure the phone has Internet access (e.g. WiFi is connected)
- [ ] Download `com.beemdevelopment.aegis@v2.0.3` from from APKPure
      [here](https://apkpure.com/aegis-authenticator-2fa-app/com.beemdevelopment.aegis/versions)
- [ ] Optionally, start screen recording with scrcpy:
```
$> scrcpy \
    --stay-awake \
    --bit-rate 4M \
    --max-fps 20 \
    --record android-screen.mp4
```
- [ ] Install the downloaded APK file(s): `$> adb install-multiple *.apk`

# Action checklist

- [ ] Open the app
- [ ] Click through the Welcome tour
- [ ] Create a password using the single character `a`
- [ ] Change password to `password` and note the password strength meter
- [ ] Change password back to `a`
- [ ] Go to Settings > Backups to see that backups are disabled by default
- [ ] Go back to the home screen
- [ ] Scan QR code #1
- [ ] Enable backups: Settings > Backups > Automatic Backups toggle on
- [ ] Scan QR code #2
- [ ] Create export: Settings > Import & Export > Export
  - [ ] Note the plaintext options (txt and json)
  - [ ] Create an encrypted export
- [ ] Delete TOTP secrets
- [ ] Uninstall/reinstall the app
- [ ] Import from file
