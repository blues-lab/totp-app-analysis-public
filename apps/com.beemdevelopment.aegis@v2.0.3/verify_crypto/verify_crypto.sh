#!/usr/bin/env bash

echo "The decrypt.py script requires the cryptography pip module."
echo "Checking whether it is installed..."
pip_list_output=$(pip3 list | grep "\bcryptography\b")

if [[ -z $pip_list_output ]]; then
  echo "The cryptography pip module WAS NOT found."
  echo "You can install it with the following command: pip3 install cryptography"
else
  echo "The cryptography pip module WAS found."
fi

PWD=$(dirname "$0")

echo "The backup password is the single character 'a'."

echo "$> decrypt.py --input ${PWD}/aegis-export-20220413-125511.json"
"${PWD}"/decrypt.py --input "${PWD}"/aegis-export-20220413-125511.json

echo "$> decrypt.py --input ${PWD}/aegis-backup-20220413-125506.json"
"${PWD}"/decrypt.py --input "${PWD}"/aegis-backup-20220413-125506.json
