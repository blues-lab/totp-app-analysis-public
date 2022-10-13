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
Note: this app requires use of an active phone number that can receive SMS.
      During our research, we bought temporary phone numbers from
      [messagebird.com](https://www.messagebird.com/) to protect the privacy of
      our personal information.

- [ ] Optionally, start screen recording with scrcpy
```
scrcpy \
    --stay-awake \
    --bit-rate 2M \
    --max-fps 10 \
    --record android-screen.mp4
```
- [ ] Open the app
  - [ ] Use phone #: `XXX-XXX-XXXX`
  - [ ] Use email: `blueslabtotp+authy@gmail.com`
- [ ] Scan QR code #1
- [ ] Enable the backup functionality in the app
  - [ ] Use password `a` to get error
  - [ ] Use password `123456` to enable backups
- [ ] Scan QR code #2
- [ ] Uninstall/reinstall the app
  - [ ] Use phone #: `XXX-XXX-XXXX`
- [ ] Decrypt the accounts
  - Notice that as soon as the phone number is verified, the issuer and label
    for each account is shown to the user. This highlights that that information
    is not encrypted along with the TOTP secrets.
  - [ ] Use the wrong password to trigger the check for whether the password was
        (in)correct
- [ ] Change backup password to `rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr` (to see if it
      gets sent in the clear)
- [ ] Delete accounts
