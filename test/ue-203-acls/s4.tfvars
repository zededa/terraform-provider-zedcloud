# S4: a drops its acls block; d edits its ACLs on the ordering-disabled project.
instances = {
  a = { project = "on" }
  b = { project = "on", cpus = 2 }
  c = { project = "on", acls = ["allow", "p9022", "p9080"] }
  d = { project = "off", acls = ["allow", "p7022"] }
}
