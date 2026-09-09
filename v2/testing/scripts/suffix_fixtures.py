#!/usr/bin/env python3
"""One-shot rewriter: append __SUFFIX__ to object name/title values in v2/resources/testdata.

Kept in the tree because it is the record of WHICH name/title values are object
names -- rerun it after adding fixtures, and read NEVER/ALLOWED_PATHS before
assuming a `sed` would have done.

Nesting-aware: only attributes on the allowlist of (block-path, attribute) pairs
are rewritten. A blanket `sed` would also hit interface names (eth0/eth1),
manifest resource keys (cpus/memory/resourceType), custom-config keys and
external artifact names that must resolve in the datastore.

The golden .yaml pass then tokenises the same values where the goldens assert
them, so names stay compared rather than being added to an ignore list.

Usage: suffix_fixtures.py [--apply] [ROOT]
"""
import re, sys, pathlib

ATTR = re.compile(
    r'^(?P<pre>\s*(?P<key>name|title|serialno|token)\s*=\s*")(?P<val>[^"]*)(?P<post>"\s*)$')
BLOCK_OPEN = re.compile(r'^\s*(?:(?P<kw>resource|data)\s+"(?P<type>[^"]+)"\s+"(?P<label>[^"]+)"|(?P<blk>[A-Za-z_][A-Za-z0-9_]*)\s*(?:=\s*)?)\s*\{\s*$')

# key -> block contexts where it is an identifier worth suffixing.
# "" == a direct attribute of the resource/data block.
#
# name/title also inside `manifest` (the app manifest carries its own name), but
# NOT in `resources`/`interfaces`/`custom_config`, where `name` is a key the
# controller interprets. serialno and token are device identity, resource-level
# only.
ALLOWED_PATHS = {
    "name":     {"", "manifest"},
    "title":    {"", "manifest"},
    "serialno": {""},
    "token":    {""},
}

# Values that are structural keys or external references, never identifiers.
NEVER = {
    "eth0", "eth1",                                   # kernel-global interface names
    "cpus", "memory", "resourceType", "bootparam",    # manifest resource keys
    "custom_config_name", "var_group", "condition", "indirect",
    "xenial-amd64-docker-20180725", "alpine_24",      # artifacts that must exist in the datastore
    "docker_compose_app_bundle", "docker_runtime_app_bundle",
    "compose_app_bundle", "runtime_app_bundle",
    "title", "test", "test_title", "test-title", "test-name", "required_only-title",
}

# Values where the token cannot go at the end. Each maps a whole fixture value to
# its tokenised form.
#
# iam/user.create.tf: user identity is unique per enterprise, so two runs collide
# on it -- but "user1@example.com__SUFFIX__" is not an address. The token goes in
# the local part, which still collapses to the original when SUFFIX is empty.
MIDSTRING = {
    "user1@example.com": "user1__SUFFIX__@example.com",
    "user2@example.com": "user2__SUFFIX__@example.com",
}
MIDSTRING_ATTR = re.compile(r'^(?P<pre>\s*(?P<key>email|username)\s*=\s*")(?P<val>[^"]*)(?P<post>"\s*)$')


def rewrite(text):
    out, path, changed = [], [], 0
    for line in text.splitlines(keepends=True):
        m = ATTR.match(line)
        mm = MIDSTRING_ATTR.match(line)
        if m and path:                      # inside a resource/data block
            ctx = ".".join(path[1:])        # drop the resource/data frame
            v = m.group("val")
            if (ctx in ALLOWED_PATHS[m.group("key")]
                    and v not in NEVER and "__SUFFIX__" not in v and v != ""):
                line = f'{m.group("pre")}{v}__SUFFIX__{m.group("post")}'
                changed += 1
        elif mm and path and mm.group("val") in MIDSTRING:
            line = f'{mm.group("pre")}{MIDSTRING[mm.group("val")]}{mm.group("post")}'
            changed += 1
        else:
            bo = BLOCK_OPEN.match(line)
            if bo:
                path.append(bo.group("type") if bo.group("kw") else bo.group("blk"))
            elif re.match(r'^\s*\}\s*$', line) and path:
                path.pop()
        out.append(line)
    return "".join(out), changed

# --------------------------------------------------------------------------- #
# Golden .yaml pass
# --------------------------------------------------------------------------- #

# Any scalar key: the goldens spell some fields differently from the fixtures
# (serialno/serialNo, metaData.name), and matching on the VALUE rather than the
# key is both simpler and harder to get wrong -- only values the .tf pass
# actually tokenised are eligible.
YAML_ATTR = re.compile(r'^(?P<pre>\s*(?P<key>[A-Za-z_][A-Za-z0-9_]*):\s*"?)(?P<val>[^"\n]*?)(?P<post>"?\s*)$')

# Golden keys to leave alone even when their value matches a suffixed identifier.
#
# Several fixtures set `description` to the same string as the object's name
# (`description = "test_tf_provider-test_datastore"`). Descriptions are free text
# with no uniqueness constraint, so the .tf pass does not touch them -- and
# tokenising the golden copy would then assert a suffix the provider never sent.
DENY_KEYS = {"description"}

# Golden values the CONTROLLER derives from a name the fixture sets, so the token
# lands mid-string rather than at the end. Each entry is
# derived-value -> (fixture name it is built from, format string).
DERIVED = {
    # project/create.yaml: the edgeview policy's name is "<project name>.edgeviewPolicy"
    "test_tf_provider.edgeviewPolicy": ("test_tf_provider", "{}__SUFFIX__.edgeviewPolicy"),
}


def suffixed_values(root):
    """Identifiers the .tf pass has tokenised, without the token."""
    vals = set()
    for f in root.rglob("*.tf"):
        for line in f.read_text().splitlines():
            m = ATTR.match(line)
            if m and m.group("val").endswith("__SUFFIX__"):
                vals.add(m.group("val")[: -len("__SUFFIX__")])
    return vals


def rewrite_golden(text, vals):
    out, changed = [], 0
    for line in text.splitlines(keepends=True):
        m = YAML_ATTR.match(line)
        if m and m.group("key") not in DENY_KEYS:
            v = m.group("val")
            if v in DERIVED:
                base, fmt = DERIVED[v]
                if base in vals:
                    line = f'{m.group("pre")}{fmt.format(base)}{m.group("post")}'
                    changed += 1
            elif v in MIDSTRING:
                line = f'{m.group("pre")}{MIDSTRING[v]}{m.group("post")}'
                changed += 1
            elif v in vals and "__SUFFIX__" not in v:
                line = f'{m.group("pre")}{v}__SUFFIX__{m.group("post")}'
                changed += 1
        out.append(line)
    return "".join(out), changed


def main():
    args = [a for a in sys.argv[1:] if not a.startswith("--")]
    apply = "--apply" in sys.argv
    root = pathlib.Path(args[0] if args else "v2/resources/testdata")
    verb = "rewrote" if apply else "would rewrite"

    total = files = 0
    for f in sorted(root.rglob("*.tf")):
        new, n = rewrite(f.read_text())
        if n:
            files += 1; total += n
            print(f"{verb} {n:3d} in {f}")
            if apply: f.write_text(new)
    print(f"\nfixtures: {total} lines across {files} files")

    # The goldens follow the fixtures: only names the .tf pass actually
    # tokenised may be tokenised here.
    vals = suffixed_values(root)
    gtotal = gfiles = 0
    for f in sorted(root.rglob("*.yaml")):
        new, n = rewrite_golden(f.read_text(), vals)
        if n:
            gfiles += 1; gtotal += n
            print(f"{verb} {n:3d} in {f}")
            if apply: f.write_text(new)
    print(f"goldens:  {gtotal} lines across {gfiles} files")


if __name__ == "__main__":
    main()
