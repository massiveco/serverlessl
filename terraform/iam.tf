// role for serverlessl caget_ca, sign
// write to s3 bucket, read from s3 bucket
// role for requester
// invoke lambda function

resource "aws_iam_role" "sign" {
  name = "slssl_${var.ca_name}_sign"
  path = "/serverlessl/"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "sign" {
  name        = "slssl_${var.ca_name}_sign"
  description = "A policy for the serverlessl Sign functionality"
  path        = "/serverlessl/"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": "${aws_s3_bucket.private.arn}/${var.ca_name}/*"
    },
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "sign" {
  name       = "sign-attachment"
  roles      = ["${aws_iam_role.sign.name}"]
  policy_arn = "${aws_iam_policy.sign.arn}"
}

resource "aws_iam_role" "get_ca" {
  name = "slssl_${var.ca_name}_get_ca"
  path = "/serverlessl/"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "get_ca" {
  name        = "slssl_${var.ca_name}_get_ca"
  description = "A policy for the serverlessl get_ca functionality"
  path        = "/serverlessl/"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "s3:PutObject"
      ],
      "Effect": "Allow",
      "Resource": "${aws_s3_bucket.private.arn}/${var.ca_name}/*"
    },
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "get_ca" {
  name       = "get_ca-attachment"
  roles      = ["${aws_iam_role.get_ca.name}"]
  policy_arn = "${aws_iam_policy.get_ca.arn}"
}

resource "aws_iam_role" "requester" {
  name = "slssl_${var.ca_name}_requester"
  path = "/serverlessl/"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "requester" {
  name        = "slssl_${var.ca_name}_requester"
  path        = "/serverlessl/"
  description = "Policies for consumers of the serverlessl lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "lambda:InvokeFunction",
      "Resource": "${aws_lambda_function.sign.arn}"
    },    {
      "Effect": "Allow",
      "Action": "lambda:InvokeFunction",
      "Resource": "${aws_lambda_function.get_ca.arn}"
    }
  ]
}
EOF
}


resource "aws_iam_policy_attachment" "requester" {
  name       = "requester-attachment"
  roles      = ["${aws_iam_role.requester.name}"]
  policy_arn = "${aws_iam_policy.requester.arn}"
}
