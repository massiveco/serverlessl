resource "aws_s3_bucket" "private" {
  bucket = "serverlessl-${var.ca_name}-${sha1(data.aws_caller_identity.current.account_id)}"
  acl    = "private"
}

resource "aws_s3_bucket_object" "ca-config" {
  bucket  = "${aws_s3_bucket.private.id}"
  key     = "/${var.ca_name}/ca-config.json"
  content = "${var.profiles}"
}
