package main

import (
	"fmt"
	"github.com/jhump/protoreflect/desc"
	"path/filepath"

	"log"
	"os"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
)

var protoIDNameMap = map[uint16]string{
	40001: "C2L_LoginLobby",
	40002: "L2C_LoginLobby",
	40003: "C2L_CreateNewRoomLobby",
	40004: "L2C_CreateNewRoomLobby",
	40005: "C2L_JoinRoomLobby",
	40006: "L2C_JoinRoomLobby",
	40007: "PlayerInfoRoom",
	40008: "L2C_PlayerTriggerRoom",
	40009: "C2L_LeaveRoomLobby",
	40010: "L2C_LeaveRoomLobby",
	40011: "PlayerBattleInfo",
	40012: "C2L_StartGameLobby",
	40013: "L2C_StartGameLobby",
	40014: "PlayerBattlePoint",
	40015: "L2C_PrepareToEnterBattle",
	40016: "C2L_PreparedToEnterBattle",
	40017: "L2C_AllowToEnterMap",
}

func GetProto(protoDir string) []*desc.FileDescriptor {
	// 获取所有proto文件
	protoFiles, err := collectProtoFiles(protoDir)
	fmt.Println(protoFiles)
	if err != nil {
		log.Fatalf("收集proto失败: %v", err)
	}
	if len(protoFiles) == 0 {
		log.Fatalf("未找到proto文件")
	}

	parser := protoparse.Parser{
		ImportPaths:           []string{protoDir}, // 非常关键！
		InferImportPaths:      true,
		IncludeSourceCodeInfo: true,
	}

	fds, err := parser.ParseFiles(protoFiles...)
	if err != nil {
		log.Fatalf("proto解析失败: %v", err)
	}
	return fds
}
func collectProtoFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(d.Name()) == ".proto" {
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			files = append(files, fmt.Sprintf("%s", rel))
		}
		return nil
	})
	return files, err
}

var (
	opCode = map[uint16]*dynamic.Message{}
)

func main() {
	// 命令行参数
	fds := GetProto("E:\\dev\\game\\NKGMobaBasedOnET\\Proto")

	// 查找目标消息
	var msgDesc *dynamic.Message

	for op := range protoIDNameMap {
		for _, fd := range fds {
			md := fd.FindMessage(protoIDNameMap[op])
			if md != nil {
				msgDesc = dynamic.NewMessage(md)
				opCode[op] = msgDesc
				break
			}
		}
	}
	fmt.Println(opCode)
	//for _, fd := range fds {
	//	md := fd.FindMessage(*messageName)
	//	if md != nil {
	//		msgDesc = dynamic.NewMessage(md)
	//		break
	//	}
	//}
	//if msgDesc == nil {
	//	log.Fatalf("未找到消息类型: %s", *messageName)
	//}
	//
	//// 读取二进制数据
	//data, err := os.ReadFile(*dataFile)
	//if err != nil {
	//	log.Fatalf("读取data文件失败: %v", err)
	//}
	//
	//// 反序列化
	//err = msgDesc.Unmarshal(data)
	//if err != nil {
	//	log.Fatalf("解码失败: %v", err)
	//}
	//
	//// 打印结果（可以输出为 JSON）
	//jsonStr, err := msgDesc.MarshalJSON()
	//if err != nil {
	//	log.Fatalf("转JSON失败: %v", err)
	//}
	//fmt.Println(string(jsonStr))
}
