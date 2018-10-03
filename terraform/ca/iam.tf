resource "aws_iam_role" "ca" {
  name = "slssl_${var.ca_name}_ca"
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

resource "aws_iam_policy" "ca" {
  name        = "slssl_${var.ca_name}_ca"
  description = "A policy for the serverlessl ca functionality"
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

resource "aws_iam_policy_attachment" "ca" {
  name       = "ca-attachment"
  roles      = ["${aws_iam_role.ca.name}"]
  policy_arn = "${aws_iam_policy.ca.arn}"
}
