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
- [ ] Download `net.codemonkey.otpgeneratorapp@v6.1` from from APKPure
      [here](https://apkpure.com/code-generator/net.codemonkey.otpgeneratorapp/versions)
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
- [ ] Menu top right > `Scan QR Code` > Scan QR code #1
  - The app seems to have a bug that does not register the QR code the first
    time it is scanned, so repeat this step twice to actually register QR code
    #1
- [ ] Click the menu on the card for TOTP account
  - Note that there are no backup options here
- [ ] Menu top right > `Export` > `Save`
- [ ] Open the saved file using the file explorer app on the phone
  - Note that the backup file contains all TOTP data in plaintext (the TOTP
    secret is displayed in a byte array)
- [ ] Scan QR code #2
- [ ] Menu top right > `Export` > `Save`
- [ ] Open the saved file using the file explorer app on the phone
  - Note that the second backup file contains both TOTP accounts
- [ ] Uninstall / reinstall / open the app
- [ ] Menu top right > `Import` > select backup file
  - Note that the TOTP accounts were restored
- [ ] Done
