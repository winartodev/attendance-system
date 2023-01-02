run:
	go run app/main.go

generate: 
	go generate ./ent

init_ent:
	go run -mod=mod entgo.io/ent/cmd/ent init Employee Attendance User
