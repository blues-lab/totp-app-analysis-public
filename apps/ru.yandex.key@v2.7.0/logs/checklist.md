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
Note: this app requires use of an active phone number that can receive SMS.
      During our research, we bought temporary phone numbers from
      [messagebird.com](https://www.messagebird.com/) to protect the privacy of
      our personal information.

- [ ] Optionally, start screen recording with scrcpy
```
scrcpy \
    --stay-awake \
    --bit-rate 4M \
    --max-fps 20 \
    --record android-screen.mp4
```
- [ ] Click `Add account` to scan QR code #1
- [ ] Create a backup: Gear icon in upper left > `Create backup copy`
   - Phone number used: `XXX-XXX-XXXX`
   - I tried to set the password to `a` and got an error message that password
     must be at least length 6
   - I tried `123456` and got error message "weak password"
   - It accepted `password123!`, `aaaa1!`, `Password123!`
- [ ] Scan QR code #2
- [ ] Uninstall/reinstall the app
- [ ] Restore the backup copy: Gear icon in upper left > `Restore from backup
      copy`
   - Phone number used: `XXX-XXX-XXXX`
    - After confirming the code sent to the phone number, the user is shown a
      warning popup saying that all local accounts in the app will be replaced
      with the ones from the backup. Confirm.
