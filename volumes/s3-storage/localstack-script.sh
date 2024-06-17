#!/bin/bash

awslocal s3api \
create-bucket --bucket transaction-processor-bucket \
--create-bucket-configuration LocationConstraint=eu-central-1 \
--region us-east-1
