install:
  just install-air
  just install-ent
  just install-templ

install-air:
  go install github.com/air-verse/air@latest

install-ent:
  go install entgo.io/ent/cmd/ent@latest

install-templ:
  go install github.com/a-h/templ/cmd/templ@latest

gen:
  just gen-ent
  just gen-templ

gen-ent:
  go generate ./ent

gen-templ:
  templ generate

ent-new name:
  ent new {{name}}

dev:
  templ generate --watch & \
  air & \
  wait

test:
  go test -count=2 -p 1 ./...

upgrade:
  go list -u -m all | grep "\[" | grep -v "indirect"
