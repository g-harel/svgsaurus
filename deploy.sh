#!/bin/sh

set -e

NAME=svgsaurus
ZIP_PATH=function/function.zip

BUCKET=svgsaurus
OBJECT=function.zip

if [ -z $TOKEN ]; then
    echo
    echo "  Error: TOKEN environement variable is empty (OAuth2)"
    echo "    1. Open Google Cloud Shell in your browser"
    echo "    2. Run command \"gcloud auth application-default print-access-token\""
    echo "    3. Set the token in your current shell using \"export TOKEN=...\""
    echo "    4. Run this script again"
    echo
    echo "   You can unset the variable by running \"unset TOKEN\""
    echo
    exit 1
fi

# check that gcloud cli is configured
gcloud beta functions list

make install

make -C ./function

curl -X POST --data-binary @$ZIP_PATH \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/zip" \
    "https://www.googleapis.com/upload/storage/v1/b/$BUCKET/o?uploadType=media&name=$OBJECT" \
    --fail

if [ $? -ne 0 ]; then
    echo
    echo "  Error: Function archive upload failed, this is likely due to an expired token."
    echo
    exit 1
fi

gcloud beta functions deploy \
    $NAME \
    --region=us-east1 \
    --source=gs://$BUCKET/$OBJECT \
    --memory=128MB \
    --timeout=1s \
    --trigger-http

# check that firebase cli is configured
firebase list

firebase deploy
