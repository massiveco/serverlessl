resource "aws_lambda_function" "sign" {
  s3_bucket        = "serverlessl"
  s3_key           = "get_ca.zip"
  function_name    = "slssl-${var.ca_name}-sign"
  role             = "${aws_iam_role.sign.arn}"
  handler          = "sign"
  runtime          = "go1.x"
  source_code_hash = "e3a07ae170086ac87653204a6a0b21928a384451b323552aa833806897b2ce6d"

  environment {
    variables = {
      CA_LAMBDA        = "slssl-${var.ca_name}-ca"
      SLSSL_S3_BUCKET  = "${var.s3_bucket}"
      SLSSL_S3_PREFIX  = "${var.ca_name}/"
      PROFILE_OVERRIDE = "${var.profile_override}"
    }
  }
}
