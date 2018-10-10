resource "aws_lambda_function" "ca" {
  s3_bucket        = "serverlessl"
  s3_key           = "ca.zip"
  function_name    = "slssl-${var.ca_name}-ca"
  role             = "${aws_iam_role.ca.arn}"
  handler          = "get_ca"
  source_code_hash = "NWE0MzM5MzM0ZDBhMjBmMGNhMGNmYzRjNDcxMjhiNjhjN2JhMDJiNjhlMTZlMmY5YzI3NmRhNzExNGRiMTU3Mwo="
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
