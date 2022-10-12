#!/bin/sh

##
# This script downloads (if necessary) and installs (if necessary) the Magisk
# APK and related module that are required to configure the Android device to
# trust all user certificates. It also configures the Android device to proxy
# all HTTP(S) traffic to the laptop on port 8080.
#
# This script is idempotent, so it can be run many times in a row.
##

# exit when any command fails
set -e

VENDOR_DIR="vendor"

MAGISK_VERSION="v24.3"
MAGISK_APK_FILENAME="magisk-${MAGISK_VERSION}.apk"
MAGISK_APK_PATH="${VENDOR_DIR}/${MAGISK_APK_FILENAME}"

MAGISK_MODULE_FILENAME="magisk-module-trust-user-certs.zip"
MAGISK_MODULE_PATH="${VENDOR_DIR}/${MAGISK_MODULE_FILENAME}"

main() {
    mkdir -p ${VENDOR_DIR}

    ##
    # Download and install Magisk APK on the phone if it is not already
    # installed.
    #
    # Docs: https://github.com/topjohnwu/Magisk/blob/master/docs/install.md
    ##

    if [ -z $(adb shell pm list packages | grep magisk) ]
    # If the magisk APK is not already installed, then...
    then
        echo "The Magisk apk IS NOT already installed on the Android device."

        if [ -r ${MAGISK_APK_PATH} ]
        # If the magisk APK exists locally and can be read, then...
        then
            echo "Magisk APK found: ${MAGISK_APK_PATH}"
        else
            echo "Downloading Magisk apk."
            curl \
                --output ${MAGISK_APK_PATH} \
                --location \
                --silent \
                --show-error \
                "https://github.com/topjohnwu/Magisk/releases/download/${MAGISK_VERSION}/Magisk-${MAGISK_VERSION}.apk"
        fi

        echo "Installing the Magisk apk on the Android device."
        adb install ${MAGISK_APK_PATH}
    else
        echo "The Magisk apk IS already installed on the Android device."
    fi

    ##
    # Download the Magisk module and push it to the Android device. The module
    # copies user CA certs to system CA certs so that they are always trusted by
    # apps.
    #
    # Docs: https://github.com/NVISOsecurity/MagiskTrustUserCerts
    ##

    if [ -z $(adb shell "[ -f /sdcard/Download/${MAGISK_MODULE_FILENAME} ] || echo 1") ]
    # If the magisk module zip file already exists on the Android device, then...
    then
        echo "The magisk module zip IS already pushed to the Android device download folder."
    else
        echo "The magisk module zip IS NOT already pushed to the Android device download folder."

        if [ -r ${MAGISK_MODULE_PATH} ]
        # If the magisk module zip file exists locally and can be read, then...
        then
            echo "Magisk module zip found: ${MAGISK_MODULE_PATH}"
        else
            echo "Downloading magisk module zip."
            curl \
                --output ${MAGISK_MODULE_PATH} \
                --location \
                --silent \
                --show-error \
                "https://github.com/NVISOsecurity/MagiskTrustUserCerts/releases/download/v0.4.1/AlwaysTrustUserCerts.zip"
        fi

        echo "Pushing magisk module zip to Android device."
        adb push ${MAGISK_MODULE_PATH} /sdcard/Download
    fi

    ##
    # Configure the device to proxy all HTTP(S) traffic to itself at port 8082
    # Docs: https://mpsocial.com/t/how-to-use-proxy-for-all-apps-in-android-devices/125695
    # Docs: https://developer.android.com/reference/android/provider/Settings.Global#HTTP_PROXY
    ##
    echo "Configuring Android device to proxy all HTTP(S) traffic to itself at port 8082."
    adb shell settings put global http_proxy 127.0.0.1:8082

    ##
    #
    # Configure the device to send all TCP traffic that it receives on port 8082
    # to the laptop on port 8080 (the default port for mitmproxy). Docs:
    # https://handstandsam.com/2016/02/01/network-calls-from-android-device-to-laptop-over-usb-via-adb/
    ##
    echo "Configuring Android device to proxy all TCP traffic on port 8082 to laptop on port 8080 (the default port for mitmproxy)."
    adb reverse tcp:8082 tcp:8080

    echo "Done."
    echo "Note: This script is idempotent, so it can be run many times in a row."
}

main
