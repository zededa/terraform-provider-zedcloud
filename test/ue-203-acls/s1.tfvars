# S1: create. a edits a bundle rule, b omits acls, c echoes the bundle,
# d edits a bundle rule on an ordering-disabled project.
instances = {
  a = { project = "on", acls = ["allow", "p8022", "p9080"] }
  b = { project = "on" }
  c = { project = "on", acls = ["allow", "p9022", "p9080"] }
  d = { project = "off", acls = ["allow", "p8022"] }
}
