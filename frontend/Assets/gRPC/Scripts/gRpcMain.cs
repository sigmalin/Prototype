using System;
using Grpc.Core;
using Protobufs.helloWorld;
using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class gRpcMain : MonoBehaviour
{
    // Start is called before the first frame update
    void Start()
    {
        Channel channel = new Channel("127.0.0.1:5033", ChannelCredentials.Insecure);

        var client = new HelloWorldService.HelloWorldServiceClient(channel);

        var reply = client.Say(new HelloWorldRequest() { Me = "sigma" });
        Debug.Log($"reply = {reply.Msg}");

        channel.ShutdownAsync().Wait();
    }
}
