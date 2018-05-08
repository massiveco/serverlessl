resource "aws_s3_bucket" "private" {
  bucket = "serverlessl-${sha1(data.aws_caller_identity.current.account_id)}"
  acl    = "private"
}
