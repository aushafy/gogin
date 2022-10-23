#!/bin/sh

# --batch to prevent interactive command
# --yes to assume "yes" for questions
gpg --quiet --batch --yes --decrypt --passphrase="$SERVICE_ACCOUNT_PASSPHRASE" \
--output terraform/service-account.json service-account.json.gpg