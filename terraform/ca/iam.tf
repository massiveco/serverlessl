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
