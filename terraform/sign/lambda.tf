resource "aws_lambda_function" "sign" {
  function_name = "slssl-${var.ca_name}-${var.profile}"
  handler       = "sign"
  role          = "${aws_iam_role.sign.arn}"
  runtime       = "go1.x"
  s3_bucket     = "serverlessl"
  s3_key        = "sign.zip"

  environment {
    variables = {
      CA_LAMBDA        = "slssl-${var.ca_name}"
      SLSSL_S3_BUCKET  = "${var.s3_bucket}"
      SLSSL_S3_PREFIX  = "${var.ca_name}/"
      PROFILE_OVERRIDE = "${var.profile}"
    }
  }
}
