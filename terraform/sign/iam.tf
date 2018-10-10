resource "aws_iam_role" "sign" {
  name = "slssl_${var.ca_name}_sign_${var.profile}"
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
  name        = "slssl_${var.ca_name}_sign_${var.profile}"
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
      "Resource": "${var.s3_bucket}/${var.ca_name}*"
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
  name       = "sign-attachment-${var.profile}"
  roles      = ["${aws_iam_role.sign.name}"]
  policy_arn = "${aws_iam_policy.sign.arn}"
}

resource "aws_iam_role" "requester" {
  name = "slssl_${var.ca_name}_requester_${var.profile}"
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
  name        = "slssl_${var.ca_name}_requester_${var.profile}"
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
    },    
    {
      "Effect": "Allow",
      "Action": "lambda:InvokeFunction",
      "Resource": "${var.lambda_ca_arn}"
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "requester" {
  name       = "requester-attachment-${var.profile}"
  roles      = ["${aws_iam_role.requester.name}"]
  policy_arn = "${aws_iam_policy.requester.arn}"

  lifecycle {
    ignore_changes = ["users"]
  }
}
