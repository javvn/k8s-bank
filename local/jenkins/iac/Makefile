# Golang parameter(gofmt: code formatter, go vet/golist: lint tools)
TF_CMD=terraform

TF_FMT=$(TF_CMD) fmt

TF_PLAN=$(TF_CMD) plan
TF_APPLY=$(TF_CMD) apply
TF_DESTROY=$(TF_CMD) apply -destroy

TF_APPLY_AUTO=$(TF_CMD) apply -auto-approve
TF_DESTROY_AUTO=$(TF_CMD) apply -destroy -auto-approve

TF_STATE_LIST=$(TF_CMD) state list


fmt:
	$(TF_FMT)

sl:
	$(TF_STATE_LIST)

plan:
	$(TF_PLAN)

destroy:
	$(TF_DESTROY)

destroy-auto:
	$(TF_DESTROY_AUTO)

apply:
	make fmt
	$(TF_APPLY)

apply-auto:
	make fmt
	$(TF_APPLY_AUTO)
