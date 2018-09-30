resource "aws_lambda_function" "sign" {
  filename         = "../lambda/sign/deployment.zip"
  function_name    = "slsslSign-${var.ca_name}"
  role             = "${aws_iam_role.sign.arn}"
  handler          = "sign"
  runtime          = "go1.x"
  source_code_hash = "a97d49e9d52452886ff1b36089ca6e607bccc730d843265573c6e44a9a42c9fc"

  environment {
    variables = {
      SLSSL_S3_BUCKET  = "${aws_s3_bucket.private.bucket}"
      SLSSL_S3_PREFIX  = "${var.ca_name}/"
      PROFILE_OVERRIDE = "${var.profile_override}"
    }
  }
}

resource "aws_lambda_function" "get_ca" {
  s3_bucket        = "serverlessl"
  s3_key           = "get_ca.zip"
  function_name    = "slsslGetCa-${var.ca_name}"
  role             = "${aws_iam_role.get_ca.arn}"
  handler          = "get_ca"
  source_code_hash = "e3a07ae170086ac87653204a6a0b21928a384451b323552aa833806897b2ce6d"
  runtime          = "go1.x"

  environment {
    variables = {
      SLSSL_S3_BUCKET     = "${aws_s3_bucket.private.bucket}"
      SLSSL_S3_PREFIX     = "${var.ca_name}/"
      SLSSL_CA_COMMONNAME = "${var.ca_name}"
      SLSSL_CA_GROUP      = "${var.ca_name}"
      SLSSL_CA_COUNTRY    = "${var.ca_country}"
      SLSSL_CA_STATE      = "${var.ca_state}"
      SLSSL_CA_CITY       = "${var.ca_city}"
    }
  }
}
