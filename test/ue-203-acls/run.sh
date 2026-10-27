#!/bin/bash
# Run one test step: plan, apply, plan again, then dump the stored ACLs.
# Usage: ./run.sh s1   (reads s1.tfvars, writes s1.*.log)
s=$1
export TF_CLI_CONFIG_FILE=$PWD/repro.tfrc
tofu plan -no-color -var-file="$s.tfvars" > "$s.plan-before.log" 2>&1
tofu apply -auto-approve -no-color -var-file="$s.tfvars" > "$s.apply.log" 2>&1
echo "apply exit=$?" >> "$s.apply.log"
tofu plan -no-color -detailed-exitcode -var-file="$s.tfvars" > "$s.plan-after.log" 2>&1
echo "plan exit=$?" >> "$s.plan-after.log"
./acls.sh ue203_ai_a ue203_ai_b ue203_ai_c ue203_ai_d > "$s.stored.log" 2>&1
echo "$s done"
