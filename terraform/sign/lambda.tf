resource "aws_lambda_function" "sign" {
  filename         = "../lambda/sign/deployment.zip"
  function_name    = "slssl-${var.ca_name}-sign"
  role             = "${aws_iam_role.sign.arn}"
  handler          = "sign"
  runtime          = "go1.x"
  source_code_hash = "a97d49e9d52452886ff1b36089ca6e607bccc730d843265573c6e44a9a42c9fc"

  environment {
    variables = {
      CA_LAMBDA        = "slssl-${var.ca_name}-ca"
      SLSSL_S3_BUCKET  = "${var.s3_bucket}"
      SLSSL_S3_PREFIX  = "${var.ca_name}/"
      PROFILE_OVERRIDE = "${var.profile_override}"
    }
  }
}
