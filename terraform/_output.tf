output "aws_iam_role_requester" {
  value = "${aws_iam_policy.requester.arn}"
}
