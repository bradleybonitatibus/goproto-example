package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "github.com/bradleybonitatibus/goproto-example/complex"
	enumspb "github.com/bradleybonitatibus/goproto-example/enums"
	simplepb "github.com/bradleybonitatibus/goproto-example/simple"
	"google.golang.org/protobuf/proto"
)

func main() {
	sm := doSimple()
	writeToFile("simple.bin", sm)
	sm2 := &simplepb.SimpleMessage{}
	test, err := readFromFile("simple.bin", sm2)
	if err != nil {
		log.Fatalln("Failed to read from file")
	}

	fmt.Println(test)

	doEnum()
	doComplex()

}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1,
		IsSimple:   true,
		Name:       "hello world",
		SampleList: []int32{1, 4, 6, 4},
	}

	return &sm
}

func doEnum() {
	ep := enumspb.EnumMessage{
		Id:           42069,
		DayOfTheWeek: enumspb.DayOfTheWeek_SATURDAY,
	}
	fmt.Println(ep.String())
}

func doComplex() {
	multple := []*complexpb.DummyMessage{
		{
			Id:   23095,
			Name: "HI THERE FIRST DUMMY",
		},
		{
			Id:   390850213,
			Name: "SECOND DUMMY",
		},
	}
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   430,
			Name: "Hello Complex Message",
		},
		MuiltipleDummy: multple,
	}

	fmt.Println(cm.String())
}
func writeToFile(fname string, pb proto.Message) error {
	bytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, bytes, 0644); err != nil {
		log.Fatalln("Can't write to file ", err)
		return err
	}
	return nil
}

func readFromFile(fname string, pb proto.Message) (proto.Message, error) {
	bytes, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Failed to read file")
		return nil, err
	}

	err = proto.Unmarshal(bytes, pb)

	if err != nil {
		log.Fatalln("Failed to deserialize to protocol buffer struct")
		return nil, err
	}
	return pb, nil
}
