# QR Codes
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
- [ ] Scan QR code #1 (+ button in bottom right)
- [ ] Uninstall / reinstall the app
- [ ] Click `customize` to see which options are available for users to change,
      but leave all defaults
- [ ] The app should prompt you to enable backups. Click `Turn on`.
- [ ] Follow the prompts to enable backup to GDrive using your google account
- [ ] Uninstall / reinstall the app
- [ ] Open the app
- [ ] Skip the tour
- [ ] Do not update the app to the latest version
- [ ] Go to `Settings > 2FAS Backup` and toggle on `Google Drive sync`
- [ ] Follow the prompts to recover
  - Go to the main screen
  - Note that the OTPs are being generated for your account, but no password was
    required to backup/recover. This means that all TOTP fields, including the
    secret, were backed up to GDrive in plaintext.
- [ ] Scan QR code #2
- [ ] Go to `Settings > 2FAS Backup > Synchronization settings > Set password`
- [ ] Use backup password `rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr`
- [ ] Uninstall / reinstall the app
- [ ] Open the app
- [ ] Skip the tour
- [ ] Do not update the app to the latest version
- [ ] Go to `Settings > 2FAS Backup` and toggle on `Google Drive sync`
- [ ] Follow the prompts to recover with password `rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr`
  - Go to the main screen to see all OTPs for your accounts. Network traffic
    confirms that the TOTP backup was encrypted before being sent to GDrive.

