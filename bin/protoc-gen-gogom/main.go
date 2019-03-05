package main

import (
	"github.com/gogo/protobuf/gogoproto"
	"github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
)

func main() {
	req := command.Read()

	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)

	vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
	vanity.ForEachFile(files, vanity.TurnOnSizerAll)
	vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)

	vanity.ForEachFieldInFilesExcludingExtensions(vanity.OnlyProto2(files), vanity.TurnOffNullableForNativeTypesWithoutDefaultsOnly)
	processFields := func(field *descriptor.FieldDescriptorProto) {
		if field.TypeName != nil && *field.TypeName == ".google.protobuf.Timestamp" {
			vanity.SetBoolFieldOption(gogoproto.E_Stdtime, true)(field)
			vanity.SetBoolFieldOption(gogoproto.E_Nullable, false)(field)
		}
		if field.TypeName != nil && *field.TypeName == ".google.protobuf.Duration" {
			vanity.SetBoolFieldOption(gogoproto.E_Stdduration, true)(field)
			vanity.SetBoolFieldOption(gogoproto.E_Nullable, false)(field)
		}
		if field.IsRepeated() {
			vanity.SetBoolFieldOption(gogoproto.E_Nullable, false)(field)
		}
	}
	vanity.ForEachFieldInFilesExcludingExtensions(files, processFields)
	vanity.ForEachFile(files, vanity.TurnOffGoUnrecognizedAll)

	vanity.ForEachFile(files, vanity.TurnOffGoEnumPrefixAll)
	vanity.ForEachFile(files, vanity.TurnOffGoEnumStringerAll)
	vanity.ForEachFile(files, vanity.TurnOnEnumStringerAll)

	vanity.ForEachFile(files, vanity.TurnOnEqualAll)
	vanity.ForEachFile(files, vanity.TurnOnGoStringAll)
	vanity.ForEachFile(files, vanity.TurnOffGoStringerAll)
	vanity.ForEachFile(files, vanity.TurnOnStringerAll)

	resp := command.Generate(req)
	command.Write(resp)
}
