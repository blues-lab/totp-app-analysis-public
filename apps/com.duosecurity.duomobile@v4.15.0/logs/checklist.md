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
- [ ] Click `I have existing accounts` and follow prompts. This shows that there
      are no previous backups in GDrive created by DuoMobile.
- [ ] Click the back arrow to return to the initial main page
- [ ] Click `Set up account`
- [ ] Click `Use a QR code`
- [ ] Scan QR code #1
- [ ] Go to `Settings > General > Duo Restore`
- [ ] Click `Backup accounts with Google Drive`
  -  Note that the DuoMobile app will now backup ONLY the TOTP issuer and label.
     It will NOT backup the TOTP secret
- [ ] Uninstall / reinstall the app
- [ ] Click `I have existing accounts` and follow prompts to recover
  - Note that the app recovered only the TOTP issuer and label for your
    accounts. Without the TOTP secret, you may face account lockout at your
    third party account.
- [ ] Go to `Settings > General > Duo Restore`
- [ ] Click `Automatically reconnect third-party accounts`
  - Note the warning that Duo will not store the password
  - Note the password min/max is 10/128
  - [ ] Use backup password `rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr`
- [ ] Go to home screen and click `+ Add`
- [ ] Click `Use QR code` and scan QR code #2
- [ ] Uninstall / reinstall the app
- [ ] Click `I have existing accounts` and follow prompts to recover with
      password `rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr`
  - Note that the TOTP secret for the second account was backed up
- [ ] Delete TOTP accounts
