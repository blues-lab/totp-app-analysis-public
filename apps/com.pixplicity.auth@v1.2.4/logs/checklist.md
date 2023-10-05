![combined QR codes](https://user-images.githubusercontent.com/945571/155418867-b13d4f69-a598-4a5c-8abe-31801aece1f5.png)

- [QR Code 1](https://user-images.githubusercontent.com/945571/155416190-d10440cc-bf4b-4592-952b-ac7aba3b130f.png)
  - otpauth://totp/Blues%20Lab:blues-lab-1@example.com?secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAYQ&issuer=Blues%20Lab&algorithm=SHA1&digits=6&period=30
  - Secret  plaintext: `Hello! Plaintext secret number 1`
- [QR Code 2](https://user-images.githubusercontent.com/945571/155416198-e6fe260a-0305-48da-90e1-137faccdc20c.png)
  - otpauth://totp/Blues%20Lab:blues-lab-2@example.com?secret=JBSWY3DPEEQFA3DBNFXHIZLYOQQHGZLDOJSXIIDOOVWWEZLSEAZA&issuer=Blues%20Lab&algorithm=SHA1&digits=6&period=30
  - Secret  plaintext: `Hello! Plaintext secret number 2`

# Setup

> **Note:** Version 1.2.4 of the Pixplicity Authenticator app is included here
>           because the developer implemented some of our recommendations, so we
>           allow people to verify/reproduce those updates.

- [ ] Enable recording of plaintext network traffic generated by the device
  - Various techniques and tools can be used to record network traffic.
    Instructions for how to use open source tools to record betwork traffic on
    Pixel phones can be found [here](/capture-traffic/README.md).
- [ ] Make sure the phone has Internet access (e.g. WiFi is connected)
- [ ] Download `com.pixplicity.auth@v1.2.4` from from APKPure
      [here](https://apkpure.com/authenticator/com.pixplicity.auth/versions)
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
- [ ] Scan QR code #1
- [ ] Click the menu on the card for TOTP account > `Show QR code`
  - Note the new warning screen that requires the user to check a box before the
    app will display the plaintext TOTP data, including the secret
- [ ] Scan QR code #2
- [ ] Menu bottom left > `Back up`
  - Note the minimum password length is 20
- [ ] Create 'Backup' with password `rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr`
- [ ] Click the share icon and email yourself the backup so that you can
      copy/paste the contents into the decryption script
- [ ] Uninstall / reinstall / open the app
- [ ] Menu bottom left > `Restore backup`
  - [ ] Select the backup with whichever option you would like. We copy/pasted
        the backup contents so that the app could read it from the clipboard
- [ ] Enter passphrase `rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr`
- [ ] Done