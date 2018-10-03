resource "aws_lambda_function" "ca" {
  s3_bucket        = "serverlessl"
  s3_key           = "get_ca.zip"
  function_name    = "slssl-${var.ca_name}-ca"
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
