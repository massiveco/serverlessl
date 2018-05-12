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
      "expiry": "8760h"
    },
    "profiles": {
      "server": {
        "usages": ["signing", "key encipherment", "server auth"],
        "expiry": "8760h"
      },
      "client": {
        "usages": ["signing", "key encipherment", "client auth"],
        "expiry": "8760h"
      },
      "server+client": {
        "usages": ["signing", "key encipherment", "server auth", "client auth"],
        "expiry": "8760h"
      }
    }
  }
}
PROFILE
}
