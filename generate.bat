protoc -I=simple/ --go_out=simple simple/simple.proto
protoc -I=enums/ --go_out=enums enums/enum.proto
protoc -I=complex/ --go_out=complex complex/complex.proto
protoc -I=addressbook/ --go_out=addressbook addressbook/addressbook.proto