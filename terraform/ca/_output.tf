output "s3_bucket" {
  value = "${aws_s3_bucket.private.bucket}"
}

output "ca_arn" {
  value = "aws_lambda_function.get_ca.arn"
}
