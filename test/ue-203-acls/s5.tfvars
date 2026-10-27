# S5: b lists ACLs again after a Cortex-style API edit (run after s4 + api edit).
instances = {
  a = { project = "on" }
  b = { project = "on", acls = ["allow", "p9022", "p9080"] }
  c = { project = "on", acls = ["allow", "p9022", "p9080"] }
  d = { project = "off", acls = ["allow", "p7022"] }
}
