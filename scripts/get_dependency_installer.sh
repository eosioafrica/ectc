#!/usr/bin/env bash

# This install script is intended to download and install the latest available
# release of the dep dependency manager for Golang.
#
# It attempts to identify the current platform and an error will be thrown if
# the platform is not supported.
#
# Environment variables:
# - ASSETS_URL (required): map to $1. For logging output
# - INSTALL_LOG_TAG (required): map to $2. For logging output
# - INSTALL_DIRECTORY (required): map to $3
# - ASSETS_DIR (required): map to $3
# - DEP_RELEASE_TAG (optional): defaults to fetching the latest release
# - DEP_OS (optional): use a specific value for OS (mostly for testing)
# - DEP_ARCH (optional): use a specific value for ARCH (mostly for testing)
#
# You can install using this script:
# $ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

set -e

ASSETS_URL="https://raw.githubusercontent.com/khosimorafo/assets/master/install/linux"
INSTALL_LOG_TAG=""
INSTALL_DIRECTORY="."
ASSETS_DIR="."


downloadJSON() {
    url="$2"

    echo "Fetching $url.."
    if test -x "$(command -v curl)"; then
        response=$(curl -s -L -w 'HTTPSTATUS:%{http_code}' -H 'Accept: application/json' "$url")
        body=$(echo "$response" | sed -e 's/HTTPSTATUS\:.*//g')
        code=$(echo "$response" | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')
    elif test -x "$(command -v wget)"; then
        temp=$(mktemp)
        body=$(wget -q --header='Accept: application/json' -O - --server-response "$url" 2> "$temp")
        code=$(awk '/^  HTTP/{print $2}' < "$temp" | tail -1)
        rm "$temp"
    else
        echo "Neither curl nor wget was available to perform http requests."
        exit 1
    fi
    if [ "$code" != 200 ]; then
        echo "Request failed with code $code"
        exit 1
    fi

    eval "$1='$body'"
}

downloadFile() {
    url="$1"
    destination="$2"

    echo "Fetching $url.."
    if test -x "$(command -v curl)"; then
        code=$(curl -s -w '%{http_code}' -L "$url" -o "$destination")
    elif test -x "$(command -v wget)"; then
        code=$(wget -q -O "$destination" --server-response "$url" 2>&1 | awk '/^  HTTP/{print $2}' | tail -1)
    else
        echo "Neither curl nor wget was available to perform http requests."
        exit 1
    fi

    if [ "$code" != 200 ]; then
        echo "Request failed with code $code"
        exit 1
    fi
}

findInstallationDirectory() {
    EFFECTIVE_INSTALL_PATH="$1"
    if [ ! -d "EFFECTIVE_INSTALL_PATH" ]; then
        echo "Installation could not find your install path. Please create and restart installation"
        exit 1
    fi
}

initArch() {
    ARCH=$(uname -m)
    if [ -n "$DEP_ARCH" ]; then
        echo "Using DEP_ARCH"
        ARCH="$DEP_ARCH"
    fi
    case $ARCH in
        amd64) ARCH="amd64";;
        x86_64) ARCH="amd64";;
        *) echo "Architecture ${ARCH} is not supported by this installation script"; exit 1;;
    esac
    echo "ARCH = $ARCH"
}

initOS() {
    OS=$(uname | tr '[:upper:]' '[:lower:]')
    if [ -n "$DEP_OS" ]; then
        echo "Using DEP_OS"
        OS="$DEP_OS"
    fi
    case "$OS" in
        darwin) OS='darwin';;
        linux) OS='linux';;
        *) echo "OS ${OS} is not supported by this installation script"; exit 1;;
    esac
    echo "OS = $OS"
}

# identify platform based on uname output
initArch
initOS

# determine install directory if required
if [ -z "$INSTALL_DIRECTORY" ]; then
    findInstallationDirectory INSTALL_DIRECTORY
fi

echo "Will install into $INSTALL_DIRECTORY"


if grep -i "ubuntu" /etc/os-release; then
    # code if found
    OS_FLAVOR="ubuntu"
    echo $OS_FLAVOR
fi

# determine install directory if required
if [ -z "$OS_FLAVOR" ]; then
    echo 'Error occurred: Could not determine OS distribution type!'
    exit 1
fi
# assemble expected release artifact name
BINARY="install_${OS_FLAVOR}".sh

DEP_ARRAY=( $BINARY )


for (( i=0; i<${#DEP_ARRAY[@]}; i++ ));
do

    DEP=$DEP"${DEP_ARRAY[$i]} "
    echo "Downloading " $DEP

    DOWNLOAD_URL="$ASSETS_URL/$DEP"
    DOWNLOAD_FILE=$(mktemp)

    echo "$DOWNLOAD_URL"

    downloadFile "$DOWNLOAD_URL" "$DOWNLOAD_FILE"

    echo "Setting executable permissions."
    chmod +x "$DOWNLOAD_FILE"

    echo "Moving executable to $INSTALL_DIRECTORY"

    /bin/bash "$DOWNLOAD_FILE"

done
