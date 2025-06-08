

export interface Playback{
    c2s: Req[]; // 客户端至服务端请求
    s2c: Message[]; // 服务端到客户端请求
}

export interface Config{
    c2sMsgName: string[];
    // c2s: Map[string, Message];
}

export interface Req extends Message{
    params?: Param[]
}

export interface Message{
    time: number;  // 时间秒
    name: string; // 协议名
    data: string; // json字符串
}



export interface Param{
    index: number // 
    srcPath: string
    tarPath: string
    
}

 const playbackMockData: Playback = {
    c2s: [
        {
            time: 1689000000000,
            name: "Request1",
            data: "Data for Request1",
            // params: [
            //     {
            //         index: 1,
            //         srcPath: "/source/path1",
            //         tarPath: "/target/path1"
            //     },
            //     {
            //         index: 2,
            //         srcPath: "/source/path2",
            //         tarPath: "/target/path2"
            //     }
            // ]
        },
        {
            time: 1689000001000,
            name: "Request2",
            data: "Data for Request2",
            // params: [
            //     {
            //         index: 3,
            //         srcPath: "/source/path3",
            //         tarPath: "/target/path3"
            //     }
            // ]
        }
    ],
    s2c: [
        {
            time: 1689000002000,
            name: "Response1",
            data: "Data for Response1"
        },
        {
            time: 1689000003000,
            name: "Response2",
            data: "Data for Response2"
        }
    ]
};
export default playbackMockData;
 