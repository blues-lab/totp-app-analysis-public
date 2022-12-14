<!-- omit in toc -->
# Security and Privacy Failures in Popular 2FA Apps

This repository contains the paper titled `Security and Privacy Failures in
Popular 2FA Apps`, which has been accepted to [USENIX Security
2023](https://www.usenix.org/conference/usenixsecurity23/presentation/gilsenan)
, and its supplemental materials and artifacts.

<!-- omit in toc -->
# Abstract and Paper

The Time-based One-Time Password (TOTP) algorithm is a 2FA method that is widely
deployed because of its relatively low implementation costs and purported
security benefits over SMS 2FA. However, users of TOTP 2FA apps face a critical
usability challenge: maintain access to the secrets stored within the TOTP app,
or risk getting locked out of their accounts. To help users avoid this fate,
popular TOTP apps implement a wide range of backup mechanisms, each with varying
security and privacy implications. In this paper, we define an assessment
methodology for conducting systematic security and privacy analyses of the
backup and recovery functionality of TOTP apps. We identified all general
purpose Android TOTP apps in the Google Play Store with at least 100k installs
that implemented a backup mechanism (n = 22). Our findings show that most backup
strategies end up placing trust in the same technologies that TOTP 2FA is meant
to supersede: passwords, SMS, and email. Many backup implementations shared
personal user information with third parties, had serious cryptographic flaws,
and/or allowed the app developers to access the TOTP secrets in plaintext. We
present our findings and recommend ways to improve the security and privacy of
TOTP 2FA app backup mechanisms.

The repo contains the [full paper
here](/Gilsenan%20et%20al%20-%20Security%20and%20Privacy%20Failures%20in%20Popular%202FA%20Apps.pdf)
and the [bibtex here](/bibtex.txt).

<!-- omit in toc -->
# Artifact Appendix - Verify and Reproduce Findings

The artifacts and related instructions in this repo are intended to allow
researchers to verify and fully reproduce the findings presented in the paper.
The major findings in the paper are documented in Tables 1, 2, and 3:

- **Table 1**: Overview of the backup mechanisms supported in each app
- **Table 2**: Overview of the backup mechanisms that automatically sync data to
               the cloud
- **Table 3**: Cryptographic details of app backup mechanisms that use
               encryption

The artifacts available include:

- [Search Terms](search_terms.txt): The list of search terms that we used to
  identify as many TOTP apps in the Google Play Store as possible (see Section
  4.1 - App Selection).

- **App Checklists**: For each app, the customized checklist that enumerates
  exactly which actions to take within the app and which data to enter while
  recording the network traffic (see Section 4.2.1 - Exploring the App).

- **Decryption Scripts**: For each app that supports encrypted TOTP backups, the
  golang script that implements the decryption process (see Section 4.2.3 -
  Performing Cryptanalysis).

The ordered steps below should be followed individually for each app to verify
and reproduce the paper's findings.

- [Step 1 - Capture and review network traffic for each app (Section 4.2.2)](#step-1---capture-and-review-network-traffic-for-each-app-section-422)
- [Step 2 - Review cryptanalysis for each app that supports encryption (Section 4.2.3)](#step-2---review-cryptanalysis-for-each-app-that-supports-encryption-section-423)
- [Step 3 - Review Android Auto Backup (AAB) usage for each app (Section 5.4)](#step-3---review-android-auto-backup-aab-usage-for-each-app-section-54)

## Step 1 - Capture and review network traffic for each app (Section 4.2.2)

> This step allows you to verify and reproduce:
> - **Table 1**
>   - Backup Mechanisms (except for Android Backup)
> - **Table 2**:
>   - PII to use cloud backups
>   - Backup Location
>   - TOTP Data Leaked
>   - Obtain Backup With...
> - **Table 3**
>   - Password Min Len

For each app, record the network traffic while executing the steps enumerated in
the app checklist (`apps/<id@version>/logs/checklist.md`).

The entries in Table 2 can be *verified* by reviewing the network traffic
snippets that we collected during our research
(`apps/<id@version>/logs/traffic-snippet.txt`). These same Table 2 entries can
be *reproduced* by analyzing your own network traffic recorded while executing
the steps in the checklist for each app.

## Step 2 - Review cryptanalysis for each app that supports encryption (Section 4.2.3)

> This step allows you to verify and reproduce:
> - **Table 3**
>   - Key Source
>   - KDF and Configuration
>   - KDF Salt
>   - Encryption Algorithm
>   - Ciphertext Integrity
>   - Decryption Heuristic

For each app that supported encryption, we implemented the decryption process in
a separate script to verify our observations. Each script has a section near the
top that defines constant variables, including:

1. the ciphertext and IV from the network traffic
2. if applicable, the password that used when enabling the backup mechanism
3. if applicable, any salt passed to a KDF

These scripts are available in the `apps/<id@version>/verify_crypto` directory
and require [Go 1.18 or higher](https://go.dev/doc/install). To run each script:

```bash
$> go mod tidy
$> go run verify_crypto.go
```

<!-- omit in toc -->
### Step 2a - Verify cryptanalysis

This step allows you to *verify* the Table 3 fields listed above.

By default, the values for these constants are the actual values that we
observed in real network traffic generated by each app during our research
(`apps/<id@version>/logs/traffic-snippet.txt`). Executing the script will verify
the Table 3 findings listed above.

<!-- omit in toc -->
### Step 2b - Reproduce cryptanalysis

This step allows you to *reproduce* the Table 3 fields listed above.

To reproduce our cryptanalysis findings, you will need to update the relevant
constants in each script with the values present in your network traffic
captures and/or file exports. Each script is well commented to indicate where
the value can be found in the network traffic generated by the app or exported
file/share created by the app. We also include the relevant network traffic
snippets that we collected from our own network traffic captures during our
research (`apps/<id@version>/logs/traffic-snippet.md`) to assist finding the
correct requests/responses and field values in your traffic captures.

To observe that the `KDF Salt` used by an app is not static, you can record
network traffic while repeating the app's checklist in the previous step; a new
salt value will be used each time. To fully verify that the KDF Salt is using an
appropriate source of randomness requires reviewing the decompiled app code. We
encourage researchers to do this, but explaining that level of detail for each
app is out of the scope of these instructions.

## Step 3 - Review Android Auto Backup (AAB) usage for each app (Section 5.4)

> This step allows you to verify and reproduce:
> - **Table 1**
>   - Android Backup

Android apps are opted into [Android Auto Backup
(AAB)](https://developer.android.com/guide/topics/data/autobackup) by default,
but developers can explicitly opt-out by setting `android:allowBackup="false"`
in the app's manifest file. To verify this setting, you can install and use
[apktool](https://ibotpeaches.github.io/Apktool/install/) to view the app's
`AndroidManifest.xml` file:

```
$> java -jar apktool.jar decode something.apk --output apktool_out
$> cd apktool_out
$> cat AndroidManifest.xml | grep allowBackup
```

The attribute `android:allowBackup="false"` means that the app **does not**
allow data to be backed up via the AAB system. These apps show a `-` in Table 1.

**Note**: Google Authenticator *does* use Android Auto Backup for some
          non-secret fields, but *does not* backup any TOTP data via AAB.
          Therefore, it is marked as `-` in Table 1. See paper for details.

Any value other than `false` means that the app **does** allow data to be backed
up via AAB. For these apps, it is necessary to confirm whether TOTP data is,
in fact, backed up by AAB.

**Note**: We could not get AAB to run without error on the following apps: *Aegis
Authenticator*, *andOTP*, and *FreeOTP Authenticator*. See paper for details.

1. Ensure that a Google account is registered on the Android device
2. If backups are enabled, disable them to erase all backup data
   ([docs](https://support.google.com/android/answer/2819582?hl=en))
   - Device Settings > `Google` > `Backup` > toggle off `Backup by Google One` >
     click `Turn off & delete`
3. Re-enable automatic backups on your Android device
   ([docs](https://support.google.com/android/answer/2819582?hl=en))
   - Device Settings > `Google` > `Backup` > turn on `Backup by Google One`
4. Install the correct version of the APK on the Android phone
   - `$> adb install-multiple *.apk`
   - The checklist for each app contains a link to download the correct version
     of the APK file
5. Scan [QR Code
   1](https://user-images.githubusercontent.com/945571/155416190-d10440cc-bf4b-4592-952b-ac7aba3b130f.png)
   and do not create any backups within the app
6. Uninstall the app
7. Verify that no data was backed up via the Android Auto Backup system
   -  Device Settings > `Google` > `Backup` > verify `Backup by Google One` is
      toggled on and the message `No data backed up` is displayed
8. Reinstall the app
   - `$> adb install-multiple *.apk`
9.  Open the app and confirm that no data was backed up
10. Scan [QR Code
   1](https://user-images.githubusercontent.com/945571/155416190-d10440cc-bf4b-4592-952b-ac7aba3b130f.png)
   again
11. Use `adb` to manually trigger an Android Auto Backup for the given app
   - `$> adb shell bmgr backupnow <apk_id>`
   - Example: `$> adb shell bmgr backupnow com.authenticator.authservice2`
12. Confirm the backup ran successfully
   - The output should include `Package <apk_id> with result: Success`
   - Example: `Package com.authenticator.authservice2 with result: Success`
   - If you encounter an error when triggering AAB for an app, try restarting
     the phone and repeating the previous step to trigger the backup again.
      - Example error: `Transport rejected package because it wasn't able to
        process it at the time`
13. Uninstall / reinstall the app
    - `$> adb install-multiple *.apk`
    - Android should automatically restore the app data when the app is
      reinstalled. If it does not, you may need to manually trigger a recovery
      for the app:
      - `$> adb shell bmgr restore <apk_id>`
      - Example: `$> adb shell bmgr restore com.authenticator.authservice2`
14. Open the app to verify that the backed up data was restored
