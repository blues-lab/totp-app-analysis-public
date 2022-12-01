The APK for `com.salesforce.authenticator@v3.8.5` can be downloaded from from
APKPure
[here](https://apkpure.com/salesforce-authenticator/com.salesforce.authenticator/versions).

Unfortunately, it is no longer possible to reproduce our results for this app
because Salesforce has disabled the backup functionality for all users who are
not Salesforce customers.

On November 17, 2022, Salesforce replied to our responsible disclosure report
and stated:

> The backup and restore feature is now for Salesforce customers only. This
> limits our scope to customers who are aware of our privacy policy.
> Additionally, fundamental design changes to the Backup and Restore feature is
> in our roadmap. This will address the privacy concerns stemming from
> Salesforce's technical capability to decrypt TOTP backups.

We tried to replicate the checklist we used during our original research on
v3.8.5, but the backup mechanism would not verify any phone number we entered.
In newer versions of the Salesforce Authenticator app, a prompt is shown when a
backup is attempted stating that backup is not possible because a Salesforce
account is not registered in the app.
