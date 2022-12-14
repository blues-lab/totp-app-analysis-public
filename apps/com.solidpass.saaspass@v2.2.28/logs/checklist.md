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
- [ ] Download `com.solidpass.saaspass@v2.2.28` from from APKPure
      [here](https://apkpure.com/saaspass-authenticator-2fa-app/com.solidpass.saaspass/versions)
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
> **Note**: This app requires use of an active phone number that can receive SMS.
>       During our research, we bought temporary phone numbers from
>       [messagebird.com](https://www.messagebird.com/) to protect the privacy of
>       our personal information.

- [ ] Open the app
  - [ ] Use PIN: `0000`
- [ ] DO NOT enable recovery mechanism when prompted during the onboarding
- [ ] Scan QR code #1
- [ ] Enable the recovery mechanism
  - [ ] Use phone: `XXX-XXX-XXXX`
- [ ] Scan QR code #2
- [ ] Open the recovery settings page
  - Recovery workflow analysis should capture:
    - Ability to force a 20 hour recovery delay
    - Ability to add a recovery question
- [ ] Setup a recovery question
  - [ ] `Hello?/World` to show error message of answer length
  - [ ] `Hello?/World!` to set it up without error
- [ ] Uninstall/reinstall the app
- [ ] Click `Recover your SAASPASS ID`
- [ ] Click `Clone SAASPASS ID` just to get that screen in the video
- [ ] Click `SAASPASS Recovery` and `Continue`
- [ ] Use phone: `XXX-XXX-XXXX`
- [ ] Trigger rate limit by entering incorrect recovery code `10` times
- [ ] Enter correct SMS code
- [ ] Trigger rate limit by answering question wrong `10` times
- [ ] Enter correct SMS code and answer to recovery question
- [ ] Delete accounts
- [ ] In recovery settings, click `Permanently turn off SAASPASS Recovery` to
      see what happens
