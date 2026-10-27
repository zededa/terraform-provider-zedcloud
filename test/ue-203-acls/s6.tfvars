# S6: with async_writes = false (v1 routes), c removes a bundle rule.
async_writes = false
instances = {
  a = { project = "on" }
  b = { project = "on", acls = ["allow", "p9022", "p9080"] }
  c = { project = "on", acls = ["allow", "p9080"] }
  d = { project = "off", acls = ["allow", "p7022"] }
}
