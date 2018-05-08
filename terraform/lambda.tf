resource "aws_lambda_function" "sign" {
  filename         = "../functions/sign/deployment.zip"
  function_name    = "serverlesslSign-${var.ca_name}"
  role             = "${aws_iam_role.sign.arn}"
  handler          = "sign"
  source_code_hash = "${base64sha256(file("../functions/sign/deployment.zip"))}"
  runtime          = "go1.x"

  environment {
    variables = {
      serverlessl_S3_BUCKET = "${aws_s3_bucket.private.bucket}"
      serverlessl_S3_PREFIX = "${var.ca_name}/"
    }
  }
}
