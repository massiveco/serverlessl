output "s3_bucket" {
  value = "${aws_s3_bucket.private.bucket}"
}

output "lambda_ca_arn" {
  value = "${aws_lambda_function.ca.arn}"
}
