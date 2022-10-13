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

# Action checklist for cloud backups
Note: The cloud backup feature in this app is paid feature. During our research,
      we paid for this feature so that we could test it.

- [ ] Optionally, start screen recording with scrcpy
```
scrcpy \
    --stay-awake \
    --bit-rate 4M \
    --max-fps 20 \
    --record android-screen.mp4
```
- [ ] Open the app
- [ ] Scan QR code #1
- [ ] Enable Cloud Backups
- [ ] Scan QR code #2
- [ ] Click `Backup` in the left-hand pane
- [ ] Uninstall/Reinstall the app
- [ ] Restore the secrets: `Backup/Restore > Restore > Restore from cloud`
- [ ] Delete TOTP accounts

# Action checklist for file and share backups

- [ ] Optionally, start screen recording with scrcpy
```
scrcpy \
    --stay-awake \
    --bit-rate 4M \
    --max-fps 20 \
    --record android-screen.mp4
```
- [ ] Open the app
- [ ] Scan QR code #1
- [ ] Backup to file using password `rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr`
- [ ] Scan QR code #2
- [ ] Backup to file again using password `rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr`
- [ ] Uninstall/Reinstall the app
- [ ] Restore the secrets: `Backup/Restore > Restore > Restore from other location`
- [ ] Delete TOTP accounts
- [ ] Attempt to create a file backup to trigger the error message that there
      are no account to back up