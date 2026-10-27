"""Print the stored interface ACLs of an edge app instance read from stdin."""
import json
import sys

name = sys.argv[1]
d = json.load(sys.stdin)
if "interfaces" not in d:
    print(name, "->", d.get("httpStatusCode", d))
    sys.exit()
for i in d["interfaces"]:
    rules = []
    for a in i.get("acls") or []:
        m = {x["type"]: x["value"] for x in a.get("matches", [])}
        pm = [x.get("mapparams", {}).get("port") for x in a.get("actions") or [] if x.get("portmap")]
        rule = "%s->%s" % (m.get("lport"), pm[0]) if pm else "allow " + m.get("ip", "?")
        rules.append("#%s %s" % (a.get("id"), rule))
    print("%s %s: %s" % (name, i["intfname"], ", ".join(rules)))
