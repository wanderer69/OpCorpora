module arkhangelskiy-dv.ru/OpCorpora

go 1.16

replace arkhangelskiy-dv.ru/OpCorpora => /home/user/Go_projects/OpCorpora

replace arkhangelskiy-dv.ru/SmallDB => /home/user/Go_projects/SmallDB

require (
	arkhangelskiy-dv.ru/SmallDB v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.0.0-20220517181318-183a9ca12b87
	google.golang.org/grpc v1.46.2
	google.golang.org/protobuf v1.28.0
)
