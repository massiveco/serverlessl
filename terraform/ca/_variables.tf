variable "ca_name" {
  default = "default"
}

variable "ca_country" {
  default = "Canada"
}

variable "ca_state" {
  default = "Ontario"
}

variable "ca_city" {
  default = "Ottawa"
}

variable "profiles" {
  default = <<PROFILE
{
  "signing": {
    "default": {
      "Expiry": 31536000000000000,
      "usages": ["client auth"]
    },
    "profiles": {
      "client": {
        "usages": ["signing", "key encipherment", "client auth"],
        "Expiry": 31536000000000000
      },
      "server": {
        "usages": ["signing", "key encipherment", "server auth", "client auth"],
        "Expiry": 31536000000000000
      }
    }
  }
}
PROFILE
}
