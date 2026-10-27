# S2: a changes 8022 to 7022 after create.
instances = {
  a = { project = "on", acls = ["allow", "p7022", "p9080"] }
  b = { project = "on" }
  c = { project = "on", acls = ["allow", "p9022", "p9080"] }
  d = { project = "off", acls = ["allow", "p8022"] }
}
