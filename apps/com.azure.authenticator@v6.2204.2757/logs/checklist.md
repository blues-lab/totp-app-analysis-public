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
- [ ] Download `com.azure.authenticator@v6.2204.2757` from from APKPure
      [here](https://apkpure.com/microsoft-authenticator/com.azure.authenticator/versions)
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
- [ ] Click `I agree` to agree to the privacy terms
- [ ] Click `Skip` in the upper right corner to go to the home screen
- [ ] Click `Add account` > `Other account (Google, Facebook, etc.)`
- [ ] Scan QR code #1
- [ ] Open menu in top right and click `Turn on backup`
- [ ] In the popup that reads "You need a personal Microsoft account to use
      cloud backup.", click `ADD ACCOUNT`
- [ ] Create a Microsoft account
- [ ] Open menu in top right > `Turn on backup` > `ADD ACCOUNT`
- [ ] When asked "Tired of remembering passwords?", click `Or only enable
      two-step verification`
- [ ] Scan QR code #2
- [ ] Uninstall / reinstall app
- [ ] Open app
- [ ] Click `I agree` to agree to the privacy terms
- [ ] Click `Restore from backup` in the bottom center
- [ ] Sign into your Microsoft account
- [ ] Verify that the TOTP data was restored
- [ ] Done
