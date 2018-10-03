resource "aws_s3_bucket" "private" {
  bucket = "serverlessl-${sha1(data.aws_caller_identity.current.account_id)}"
  acl    = "private"
}

resource "aws_s3_bucket_object" "ca-config" {
  bucket  = "${aws_s3_bucket.private.id}"
  key     = "/${var.ca_name}/ca_config.json"
  content = "${var.profiles}"
}
