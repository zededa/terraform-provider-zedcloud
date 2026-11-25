# S3: unrelated update of b (cpus), acls still omitted.
instances = {
  a = { project = "on", acls = ["allow", "p7022", "p9080"] }
  b = { project = "on", cpus = 2 }
  c = { project = "on", acls = ["allow", "p9022", "p9080"] }
  d = { project = "off", acls = ["allow", "p8022"] }
}
