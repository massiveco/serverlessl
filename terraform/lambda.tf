resource "aws_lambda_function" "sign" {
  filename         = "../lambda/sign/deployment.zip"
  function_name    = "serverlesslSign-${var.ca_name}"
  role             = "${aws_iam_role.sign.arn}"
  handler          = "sign"
  source_code_hash = "${base64sha256(file("../lambda/sign/deployment.zip"))}"
  runtime          = "go1.x"

  environment {
    variables = {
      serverlessl_S3_BUCKET = "${aws_s3_bucket.private.bucket}"
      serverlessl_S3_PREFIX = "${var.ca_name}/"
    }
  }
}

resource "aws_lambda_function" "init" {
  filename         = "../lambda/init/deployment.zip"
  function_name    = "serverlesslInit-${var.ca_name}"
  role             = "${aws_iam_role.init.arn}"
  handler          = "init"
  source_code_hash = "${base64sha256(file("../lambda/init/deployment.zip"))}"
  runtime          = "go1.x"

  environment {
    variables = {
      serverlessl_S3_BUCKET     = "${aws_s3_bucket.private.bucket}"
      serverlessl_S3_PREFIX     = "${var.ca_name}/"
      serverlessl_CA_COMMONNAME = "${var.ca_name}"
      serverlessl_CA_GROUP      = "${var.ca_name}"
      serverlessl_CA_COUNTRY    = "${var.ca_country}"
      serverlessl_CA_STATE      = "${var.ca_state}"
      serverlessl_CA_CITY       = "${var.ca_city}"
    }
  }
}
